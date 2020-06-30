package Services

import (
    "wheel/go-web-3tier/AppInit"
    "wheel/go-web-3tier/Models"
)

type BookService struct {
}

func (this *BookService) LoadBookList(req *BookListRequest) *Models.BookList {
    prods := &Models.BookList{}
    AppInit.GetDB().Limit(req.Size).Order("book_id desc").Find(prods)
    return prods
}

func (this *BookService) LoadBookDetail(req *BookDetailRequest) *Models.Books {
    books := &Models.Books{}
    if AppInit.GetDB().Find(books, req.BookId).RowsAffected != 1 {
        return nil
    }
    return books
}
