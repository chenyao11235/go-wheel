package main

import (
    "context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials"
    "log"
    pb "wheel/grpc-simple/TLS/proto"
)

const PORT = "9001"

func main() {
    c, err := credentials.NewClientTLSFromFile("../conf/server.pem", "go-grpc-simple")
    if err != nil {
        log.Fatalf("credentials.NewClientTLSFromFile err: %v", err)
    }

    conn, err := grpc.Dial(":"+PORT, grpc.WithTransportCredentials(c))
    if err != nil {
        log.Fatal("grpc.Dial err: ", err)
    }

    defer conn.Close()
    client := pb.NewSearchServiceClient(conn)
    resp, err := client.Search(context.Background(), &pb.SearchRequest{Request: "gRPC"})
    if err != nil {
        log.Fatal("client.search err:%v ", err)
    }
    log.Printf("resp: %s", resp.GetResponse())
}
