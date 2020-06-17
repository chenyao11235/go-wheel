package service

import (
    "context"
    pb "wheel/grpc-demo/proto"
    "wheel/grpc-demo/server/initdata"
)

type BookService struct {
    pb.UnimplementedBookServiceServer
}

func NewBookService() *BookService {
    return &BookService{}
}

func (b BookService) GetBook(ctx context.Context, request *pb.BookRequest) (*pb.BookResponse, error) {
    for _, v := range initdata.BookData {
        if request.BookId == v.BookId {
            return &pb.BookResponse{
                Id:          v.BookId,
                Name:        v.BookName,
                Author:      v.BookAuthor,
                Price:       v.BookPrice1,
                Intro:       v.BookIntr,
                Press:       v.BookPress,
                PublishDate: v.BookData,
                Kind:        v.BookKind,
            }, nil
        }
    }

    return nil, nil
}

func (b BookService) SearchBookByPrice(request *pb.GetBooksByPrice, server pb.BookService_SearchBookByPriceServer) error {
    books := make([]*pb.BookResponse, 10)
    for _, v := range initdata.BookData {
        if v.BookPrice1 > request.Min && v.BookPrice1 < request.Max {
            resp := &pb.BookResponse{
                Id:          v.BookId,
                Name:        v.BookName,
                Author:      v.BookName,
                Price:       v.BookPrice1,
                Intro:       v.BookIntr,
                Press:       v.BookPress,
                PublishDate: v.BookData,
                Kind:        v.BookKind,
            }
            books = append(books, resp)
            // 每三book个组成一个response返回
            if len(books)%3 == 0 {
                _ = server.Send(&pb.BookListResponse{Books: books})
                books = books[:0]
            }
        }
    }
    return nil
}

func (b BookService) SearchBookByIds(server pb.BookService_SearchBookByIdsServer) error {
    panic("implement me")
}

func (b BookService) SearchBookByKind(server pb.BookService_SearchBookByKindServer) error {
    panic("implement me")
}
