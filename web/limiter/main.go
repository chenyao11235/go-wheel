package limiter

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

// 该包展示了如何进行速率限制，比如 同一时刻最多有多少个请求,
//原理： 创建一个有缓存长度chan，来一个请求就往chan中添加一个元素，当一个请求完成之后从chan中释放一个元素

type middlewareHandler struct {
	r *httprouter.Router
	l *ConnLimiter
}

func (m middlewareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 判断是否超过当前的并发阈值
	if !m.l.GetConn() {
		w.WriteHeader(http.StatusTooManyRequests)
		io.WriteString(w, "Too many requests")
		return
	}
	m.r.ServeHTTP(w, r)
}

func NewMiddlewareHandler(r *httprouter.Router) http.Handler {
	m := middlewareHandler{}
	m.r = r

	m.l = NewConnLimiter(100)
	return m
}

func main() {
	router := httprouter.New()
	router.POST("/login", Login)
	mh := NewMiddlewareHandler(router)
	http.ListenAndServe(":8080", mh)
}
