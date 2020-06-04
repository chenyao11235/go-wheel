package Services

import (
    "context"
    "wheel/go-web-3tier/App"
)

// 图书列表信息
type BookListRequest struct {
    Size int `form:"size"`
}

//  图书详细信息
type BookDetailRequest struct {
    BookId int `uri:"id" binding:"required,gt=0,max=70000"`
}

type BookResponse struct {
    Result interface{} `json:"result"`
}

func BookListEndPoint(book *BookService) App.Endpoint {
    return func(ctx context.Context, request interface{}) (response interface{}, err error) {
        req := request.(*BookListRequest)
        return &BookResponse{Result: book.LoadBookList(req)}, nil
    }
}

func BookDetailEndPoint(book *BookService) App.Endpoint {
    return func(ctx context.Context, request interface{}) (response interface{}, err error) {
        req := request.(*BookDetailRequest)
        return &BookResponse{Result: book.LoadBookDetail(req)}, nil
    }
}
