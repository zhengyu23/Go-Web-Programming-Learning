package main

// 代码清单3-11	串联多个处理器

import (
	"fmt"
	"net/http"
)

type HelloHandler_ struct{}

func (h HelloHandler_) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello!")
}

func log_(h http.Handler) http.Handler {
	// 使用HandlerFunc直接将匿名函数转换成一个Handler
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Handler called - %T\n", h)
		h.ServeHTTP(w, r)
	})
}

func protect(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Handler called - %T\n", h)
		h.ServeHTTP(w, r)
	})
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	hello := HelloHandler_{}
	http.Handle("/hello", protect(log_(hello)))
	server.ListenAndServe()
}
