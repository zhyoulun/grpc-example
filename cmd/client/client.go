package main

import (
	"context"
	pb "github.com/zhyoulun/grpc-example/proto/sdk/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:12345", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("grpc Dial fail: %+v", err)
	}
	defer conn.Close()

	c := pb.NewHelloClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Ping(ctx, &pb.PingRequest{
		Name: "haha",
	})
	if err != nil {
		log.Fatalf("c Ping fail: %+v", err)
	}
	log.Printf("ping response: %+v", r.GetContent())
}
