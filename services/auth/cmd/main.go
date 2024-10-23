package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"microGo/pkg/auth"
	pb "microservices/api/proto/auth/v1"
)

type AuthService struct {
	pb.UnimplementedAuthServiceServer
	tokenManager *auth.TokenManager
}

func NewAuthService(tokenManager *auth.TokenManager) *AuthService {
	return &AuthService{
		tokenManager: tokenManager,
	}
}

func (s *AuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	token, err := s.tokenManager.GenerateToken(req.Email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate token: %v", err)
	}

	refreshToken, err := s.tokenManager.GenerateRefreshToken(req.Email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate refresh token: %v", err)
	}

	return &pb.AuthResponse{
		AccessToken:  token,
		RefreshToken: refreshToken,
		User: &pb.UserInfo{
			Id:       "dummy-id",
			Email:    req.Email,
			Nickname: req.Nickname,
		},
	}, nil
}

func (s *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	token, err := s.tokenManager.GenerateToken(req.Email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate token: %v", err)
	}

	refreshToken, err := s.tokenManager.GenerateRefreshToken(req.Email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate refresh token: %v", err)
	}

	return &pb.AuthResponse{
		AccessToken:  token,
		RefreshToken: refreshToken,
		User: &pb.UserInfo{
			Id:       "dummy-id",
			Email:    req.Email,
			Nickname: "dummy-nickname",
		},
	}, nil
}

func (s *AuthService) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	userID, err := s.tokenManager.ValidateToken(req.Token)
	if err != nil {
		return &pb.ValidateTokenResponse{
			Valid: false,
		}, nil
	}

	return &pb.ValidateTokenResponse{
		Valid:  true,
		UserId: userID,
	}, nil
}
