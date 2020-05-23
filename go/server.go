package main

import (
	"context"
	"log"
	"net"

	"github.com/U-taro-ogw/grpc-training-ruby-go/go/lib"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
		return
	}
	grpcServer := grpc.NewServer()
	training.RegisterTrainingServer(grpcServer, &server{})
	log.Printf("==========>>>>>>>>>> training server is running!")
	grpcServer.Serve(listener)
}

type server struct {}

func (s *server) Hoge(ctx context.Context, req *training.Empty) (*training.HogeResponse, error) {
	hogeResponse := &training.HogeResponse{
		Language: "Golang",
	}
	return hogeResponse, nil
}
