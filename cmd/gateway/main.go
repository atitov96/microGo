package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptrace"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	pb "microGo/gen/go/auth/v1"

	"github.com/swaggo/http-swagger"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(
		runtime.WithErrorHandler(CustomErrorHandler),
		runtime.WithForwardResponseOption(httptrace.ForwardResponseOption),
	)

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := pb.RegisterAuthServiceHandlerFromEndpoint(
		ctx,
		mux,
		"localhost:50051",
		opts,
	)
	if err != nil {
		log.Fatalf("Failed to register gateway: %v", err)
	}

	swaggerHandler := httpSwagger.Handler(
		httpSwagger.URL("/swagger/auth.swagger.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Content-Type", "Content-Length", "Authorization"},
	})

	router := http.NewServeMux()
	router.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.Dir("./gen/openapiv2"))))
	router.Handle("/swagger-ui/", swaggerHandler)
	router.Handle("/", mux)

	handler := corsMiddleware.Handler(router)

	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint
		if err := server.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
	}()

	log.Printf("Server listening on :8080")
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
}

func CustomErrorHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	httpStatus := runtime.HTTPStatusFromCode(status.Code(err))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	errResponse := struct {
		Error string `json:"error"`
		Code  int    `json:"code"`
	}{
		Error: status.Convert(err).Message(),
		Code:  httpStatus,
	}

	json.NewEncoder(w).Encode(errResponse)
}
