package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"golang.org/x/net/context"

	pb "github.com/walkerrandolphsmith/grpc-gateway-basics/echo"
)

type server struct {
}

func (s *server) Echo(ctx context.Context, msg *pb.EchoMessage) (*pb.EchoMessage, error) {
	log.Println(msg, ctx)
	grpc.SendHeader(ctx, metadata.New(map[string]string{
		"foo": "foo1",
		"bar": "bar1",
	}))
	grpc.SetTrailer(ctx, metadata.New(map[string]string{
		"foo": "foo2",
		"bar": "bar2",
	}))
	return msg, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterEchoServer(grpcServer, &server{})
	log.Println("listening to port *:9090")
	grpcServer.Serve(lis)
}