package main

import (
    "context"
    "google.golang.org/grpc"
    "log"
    "net"
    pb "wheel/go-grpc-example/proto"
)

type SearchService struct{}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
    return &pb.SearchResponse{
        Response: r.GetRequest() + "server"}, nil
}

const PORT = "9001"

func main() {
    server := grpc.NewServer()
    pb.RegisterSearchServiceServer(server, &SearchService{})

    lis, err := net.Listen("tcp", ":"+PORT)
    if err != nil {
        log.Fatal("net.listen err: %v", err)
    }
    server.Serve(lis)
}
