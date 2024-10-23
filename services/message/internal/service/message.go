package service

import (
	"context"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	pb "microGo/api/proto/message/v1"
	"microGo/pkg/kafka"
	"microGo/pkg/kafka/events"
	"net/http"
	"time"
)

type MessageService struct {
	pb.UnimplementedMessageServiceServer
	producer *kafka.Producer
	upgrader websocket.Upgrader
	clients  map[string]*websocket.Conn
}

func (s *MessageService) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.Message, error) {
	if err := s.checkFriendship(req.FromUser, req.ToUser); err != nil {
		return nil, err
	}

	message := &pb.Message{
		Id:        generateID(),
		FromUser:  req.FromUser,
		ToUser:    req.ToUser,
		Content:   req.Content,
		CreatedAt: timestamppb.Now(),
	}

	event := &events.MessageSentEvent{
		ID:        message.Id,
		FromUser:  message.FromUser,
		ToUser:    message.ToUser,
		Content:   message.Content,
		CreatedAt: time.Now(),
	}

	if err := s.producer.Publish(events.TopicMessageSent, message.Id, event); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to publish message event: %v", err)
	}

	if conn, ok := s.clients[message.ToUser]; ok {
		if err := conn.WriteJSON(message); err != nil {
			log.Printf("failed to send websocket message: %v", err)
		}
	}

	return message, nil
}

func (s *MessageService) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}

	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("failed to upgrade connection: %v", err)
		return
	}

	s.clients[userID] = conn
	defer func() {
		conn.Close()
		delete(s.clients, userID)
	}()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return
		}
		if messageType == websocket.TextMessage {
			var msg pb.SendMessageRequest
			if err := json.Unmarshal(p, &msg); err != nil {
				continue
			}

			s.SendMessage(context.Background(), &msg)
		}
	}
}
