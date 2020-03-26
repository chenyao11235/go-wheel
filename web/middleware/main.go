//改包展示了如何给web服务器添加一个中间件，使得每个请求都要经过这些中间件

//1. 使用httprouter作为server处理器，处理器实现了Handler接口，该接口需要实现httpServe方法
//2. 修改默认的处理器可以在创建Serve的时候通过Handler字段指定，也可以在http.ListenAndServe方法中通过参数指定
package middleware

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
}

// ServeHTTP是Handler接口需要实现的方法
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 在这里可以添加中间操作
	validateUserSession(r)
	m.r.ServeHTTP(w, r)
}

// 创建中间件Handler
func NewMiddlerWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func main() {
	router := httprouter.New()
	router.POST("/login", Login)
	mh := NewMiddlerWareHandler(router)
	http.ListenAndServe(":8080", mh)
}
