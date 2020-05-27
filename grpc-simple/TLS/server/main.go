package main

import (
    "context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials"
    "log"
    "net"
    pb "wheel/grpc-simple/TLS/proto"
)

type SearchService struct{}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
    return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

const PORT = "9001"

func main() {
    c, err := credentials.NewServerTLSFromFile("../conf/server.pem", "../conf/server.key")
    if err != nil {
        log.Fatalf("credentials.NewServerTLSFromFile err: %v", err)
    }

    server := grpc.NewServer(grpc.Creds(c))
    pb.RegisterSearchServiceServer(server, &SearchService{})

    lis, err := net.Listen("tcp", ":"+PORT)
    if err != nil {
        log.Fatal("net.Listen err: %v", err)
    }

    server.Serve(lis)
}
