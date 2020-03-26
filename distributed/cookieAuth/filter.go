package main

import (
	"net/http"
	"strings"
)

// 过滤器
// 对某些url需要进行某些操作，比如说身份验证

type FilterHandler func(w http.ResponseWriter, r *http.Request) error

type Filter struct {
	filterMap map[string]FilterHandler
}

func NewFilter() *Filter {
	return &Filter{
		filterMap: make(map[string]FilterHandler),
	}
}

// 注册拦截器
func (f *Filter) RegisterFilterUri(uri string, handler FilterHandler) {
	f.filterMap[uri] = handler
}

//
func (f *Filter) GetUriHandler(uri string) FilterHandler {
	return f.filterMap[uri]
}

type WebHandler func(w http.ResponseWriter, r *http.Request)

// 执行拦截器
func (f *Filter) Handle(webHandle WebHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		for uri, handle := range f.filterMap {
			if strings.Contains(r.RequestURI, uri) {
				if err := handle(w, r); err != nil {
					_, _ = w.Write([]byte(err.Error()))
					return
				}
				break
			}
		}
		webHandle(w, r)
	}
}
