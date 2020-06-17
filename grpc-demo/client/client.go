package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"
	"wheel/grpc-demo/client/my_resolver"
	pb "wheel/grpc-demo/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/status"
)

func GetBook(c pb.BookServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := c.GetBook(ctx, &pb.BookRequest{BookId: 2000})
	if err != nil {
		log.Println(status.Code(err))
		return
	}
	fmt.Println(resp.Name)
}

func SearchBookByPrice(c pb.BookServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	// 流模式设置超时时间的意思是：如果超过规定时间，流数据还没有发送完毕就视为超时，并返回异常
	stream, err := c.SearchBookByPrice(ctx, &pb.GetBooksByPrice{Min: 50, Max: 120})
	if err != nil {
		log.Println(status.Code(err))
		return
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("数据传完了")
			return
		}
		if err != nil {
			log.Println(status.Code(err))
			return
		}
		for _, book := range resp.GetBooks() {
			fmt.Println(book.Id, book.Price)
		}
	}
}

func SearchBookByIds(c pb.BookServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	stream, err := c.SearchBookByIds(ctx)
	if err != nil {
		log.Println(status.Code(err))
		return
	}

	books := []*pb.BookRequest{
		&pb.BookRequest{
			BookId: 5009,
		},
		&pb.BookRequest{
			BookId: 5009,
		},
	}
	_ = stream.Send(&pb.BooksRequest{Books: books})

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Println(status.Code(err))
		return
	}
	fmt.Println(resp)

}

func SearchBookByKind(c pb.BookServiceClient) {

}

var (
	serviceConfig = `{
	"loadBalancingPolicy": "round_robin",
	"healthCheckConfig": {
		"serviceName": ""
	}
    }`
	service_name = "/service/book/"
)

func main() {
	// 使用默认的 my_resolver
	//r := my_resolver.NewDefaultResolver()
	// 使用基于etcd的resolver
	r := my_resolver.NewDynamicResolverBuilder()
	resolver.Register(r)
	address := fmt.Sprintf("%s:///%s", r.Scheme(), service_name)
	// 与服务端建立连接
	options := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultServiceConfig(serviceConfig),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, address, options...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	prodClient := pb.NewBookServiceClient(conn)

	for range time.Tick(time.Second) {
		GetBook(prodClient)
	}

	//SearchBookByPrice(prodClient)

	for range time.Tick(time.Second) {
	}
}
