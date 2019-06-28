package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/kim-nam-hoon/grpchello/proto"
	"google.golang.org/grpc"
)

type server struct {
}

func (*server) Greeting(ctx context.Context, req *proto.SayRequest) (*proto.SayResponse, error) {
	return &proto.SayResponse{Say: "Hello"}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalf("Failed net Linsten : %v", err)
	}

	s := grpc.NewServer()

	proto.RegisterSayServiceServer(s, &server{})

	go func() {
		fmt.Println("GRPC SERVER RUNNING...")

		err := s.Serve(lis)
		if err != nil {
			log.Fatalf("Failed GRPC SERVER Serve : %v ", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	fmt.Println("Stoping the server")
	s.GracefulStop()

	fmt.Println("Closing the listener")
	lis.Close()

}
