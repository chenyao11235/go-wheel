package main

import (
    "context"
    "google.golang.org/grpc"
    "log"
    pb "wheel/grpc-simple/simple/proto"
)

const PORT = "9001"

func main() {
    conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
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
