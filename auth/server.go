package main

import (
    "log"
    "net"

    "google.golang.org/grpc"
    pb "github.com/abdulrahmanalotaibi/macro/protos/auth"
)

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterHealthServer(s, &server{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}