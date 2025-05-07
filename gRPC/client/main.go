package main

import (
    "context"
    "log"
    "time"

    pb "github.com/aaffriya/go-playground/gRPC/proto"
    "google.golang.org/grpc"
)

func main() {
    conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    c := pb.NewGreeterClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "Go Developer"})
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }

    log.Printf("Greeting: %s", r.Message)
}
