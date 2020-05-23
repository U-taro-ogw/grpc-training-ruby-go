package main

import (
	"context"
	"log"
	"net"

	pd "github.com/U-taro-ogw/grpc-training-ruby-go/go/lib"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
		return
	}
	grpcServer := grpc.NewServer()
	pd.RegisterTrainingServer(grpcServer, &server{})
	log.Printf("==========>>>>>>>>>> training server is running!")
	grpcServer.Serve(listener)
}

type server struct {}

func (s *server) Hoge(ctx context.Context, req *pd.Empty) (*pd.HogeResponse, error) {
	hogeResponse := &pd.HogeResponse{
		Language: "Golang",
	}
	return hogeResponse, nil
}
