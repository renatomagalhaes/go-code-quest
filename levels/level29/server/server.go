package main

import (
    "context"
    "log"
    "net"

    "go-code-quest/levels/level29/proto"

    "google.golang.org/grpc"
)

type server struct {
    proto.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
    return &proto.HelloReply{Message: "Hello, " + req.Name}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    proto.RegisterGreeterServer(s, &server{})

    log.Println("Server is running on port 50051...")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
