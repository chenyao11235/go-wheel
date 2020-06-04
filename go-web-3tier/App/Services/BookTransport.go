package Services

import (
    "github.com/gin-gonic/gin"
    "wheel/go-web-3tier/App"
)

func CreateBookListRequest() App.EncodeRequestFunc {
    return func(context *gin.Context) (i interface{}, e error) {
        bReq := &BookListRequest{}
        err := context.ShouldBindQuery(bReq)
        if err != nil {
            return nil, err
        }
        return bReq, nil
    }
}

func CreateBookDetailRequest() App.EncodeRequestFunc {
    return func(context *gin.Context) (i interface{}, e error) {
        bReq := &BookDetailRequest{}
        err := context.ShouldBindUri(bReq)
        if err != nil {
            return nil, err
        }
        return bReq, nil
    }
}

func CreateBookResponse() App.DecodeResponseFunc {
    return func(context *gin.Context, res interface{}) error {
        context.JSON(200, res)
        return nil
    }
}
