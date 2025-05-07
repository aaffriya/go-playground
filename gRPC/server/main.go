package main

import (
    "context"
    "log"
    "net"

    pb "github.com/aaffriya/go-playground/gRPC/proto"
    "google.golang.org/grpc"
)

type greeterServer struct {
    pb.UnimplementedGreeterServer
}

func (s *greeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
    return &pb.HelloReply{Message: "Hello " + req.Name}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterGreeterServer(grpcServer, &greeterServer{})

    log.Println("Server listening on port 50051")
    grpcServer.Serve(lis)
}
