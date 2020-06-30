package main

import (
    "github.com/gin-gonic/gin"
    "wheel/go-web-3tier/App"
    "wheel/go-web-3tier/App/Services"
    . "wheel/go-web-3tier/AppInit"
)

func main() {
    router := gin.Default()
    v1 := router.Group("v1")
    {
        bookListServiceEndpoint := Services.BookListEndPoint(&Services.BookService{}) //图书endpoint
        bookListHandler := App.RegisterHandler(
            bookListServiceEndpoint,          //业务最终函数
            Services.CreateBookListRequest(), // 怎么取请求参数
            Services.CreateBookResponse(),    //怎么处理响应
        )
        bookDetailServiceEndpoint := Services.BookDetailEndPoint(&Services.BookService{}) //图书endpoint
        bookDetailHandler := App.RegisterHandler(
            bookDetailServiceEndpoint,          //业务最终函数
            Services.CreateBookDetailRequest(), // 怎么取请求参数
            Services.CreateBookResponse(),      //怎么处理响应
        )

        v1.Handle(HTTP_METHOD_GET, "/prods", bookListHandler)
        v1.Handle(HTTP_METHOD_GET, "/prods/:id", bookDetailHandler)
    }
    router.Run(SERVER_ADDRESS)
}
