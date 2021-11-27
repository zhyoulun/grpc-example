package main

import (
	"context"
	pb "github.com/zhyoulun/grpc-example/proto/sdk/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedHelloServer
}

func (s *server) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
	log.Printf("receive request: %+v", in.GetName())
	return &pb.PingResponse{
		Content: "pong " + in.GetName(),
	}, nil
}

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:12345")
	if err != nil {
		log.Fatalf("net Listen fail: %+v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &server{})
	log.Printf("server listening at %+v", ln.Addr())
	if err := s.Serve(ln); err != nil {
		log.Fatalf("failed to serve: %+v", err)
	}
}
