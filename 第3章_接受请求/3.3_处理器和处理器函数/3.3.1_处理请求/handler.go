package main

// 代码清单3-6	处理请求

import (
	"fmt"
	"net/http"
)

// MyHandler 处理器: 处理全部请求
type MyHandler struct{}

func (s *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}
	server.ListenAndServe()
}
