package main

import (
	pb "GRPC_server/chat"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)


type Server struct {
	pb.UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	log.Printf("RECEIVED MESSAGE BODY FROM CLIENT: %s", in.Body)
	return &pb.Message{Body: "The current version of this server is: BLUE"}, nil
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Starting GRPC server")

	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := Server{}

	grpcServer := grpc.NewServer()

	pb.RegisterChatServiceServer(grpcServer, &s)

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
