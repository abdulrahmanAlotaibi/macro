package macro

import (
    "context"
    "log"
    "time"

    "google.golang.org/grpc"
    pb "github.com/abdulrahmanalotaibi/macro/protos/auth/authh"
)

func f() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewHealthClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    r, err := c.Check(ctx, &pb.CheckRequest{})
    if err != nil {
        log.Fatalf("could not check health: %v", err)
    }
    log.Printf("Health status: %s", r.GetStatus())
}