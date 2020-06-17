package initdata

import (
    "encoding/json"
    "os"
)

type Book struct {
    BookId      int64  `json:"book_id"`
    BookName    string `json:"book_name"`
    BookIntr    string `json:"book_intr"`
    BookPrice1  float32  `json:"book_price1"`
    BookPrice2  float32  `json:"book_price2"`
    BookAuthor  string `json:"book_author"`
    BookPress   string `json:"book_press"`
    BookData    string `json:"book_data"`
    BookKind    int32  `json:"book_kind"`
    BookKindStr string `json:"book_kind_str"`
}

// 单例模式，全局变量
var (
    BookData []*Book
)

const (
    dataFile = "/Users/yaochen/go/src/wheel/grpc-demo/server/initdata/jiong_books.json"
)

func InitData() (err error) {
    var (
        file    *os.File
        decoder *json.Decoder
    )

    if file, err = os.Open(dataFile); err != nil {
        return
    }

    decoder = json.NewDecoder(file)
    if err = decoder.Decode(&BookData); err != nil {
        return
    }
    return
}
