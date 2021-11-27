package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/zhyoulun/grpc-example/proto/sdk/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"sync"
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
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go runGRPCServer(wg)
	go runGateway(wg)
	wg.Wait()
}

func runGateway(wg *sync.WaitGroup) {
	defer wg.Done()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterHelloHandlerFromEndpoint(ctx, mux, "127.0.0.1:12345", opts)
	if err != nil {
		log.Fatalf("pb RegisterHelloHandlerFromEndpoint fail: %+v", err)
	}

	log.Printf("12345 for grpc, 12346 for grpc gateway")

	err = http.ListenAndServe("127.0.0.1:12346", mux)
	if err != nil {
		log.Fatalf("http ListenAndServe fail: %+v", err)
	}
}

func runGRPCServer(wg *sync.WaitGroup) {
	defer wg.Done()

	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &server{})

	ln, err := net.Listen("tcp", "127.0.0.1:12345")
	if err != nil {
		log.Fatalf("net Listen fail: %+v", err)
	}
	log.Printf("server listening at %+v", ln.Addr())
	if err := s.Serve(ln); err != nil {
		log.Fatalf("failed to serve: %+v", err)
	}
}
