package main

import (
    "context"
    "log"

    "google.golang.org/grpc"
    pb "protos/auth"
)

type server struct {
    pb.UnimplementedHealthServer
}

func (s *server) Check(ctx context.Context, in *pb.CheckRequest) (*pb.CheckResponse, error) {
    log.Printf("Received health check request")
    return &pb.CheckResponse{Status: "OK"}, nil
}