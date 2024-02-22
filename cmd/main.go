package main

import (
	"context"
	"fmt"
	desc "github.com/FreylGit/chat-server/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

type server struct {
	desc.UnimplementedChatV1Server
}

const grpcPort = 50052

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("Failed to listen:%d", grpcPort)
	}
	log.Println("Start run server")
	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server")
	}

}

func (s *server) Create(ctx context.Context, request *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Create user:%s\n", request.Usernames)
	return &desc.CreateResponse{Id: int64(len(request.Usernames))}, nil
}

func (s *server) Delete(ctx context.Context, request *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Delete user by id:%d", request.GetId())
	return nil, nil
}

func (s *server) SendMessage(ctx context.Context, request *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("Text message:%s", request.GetText())
	return nil, nil
}
