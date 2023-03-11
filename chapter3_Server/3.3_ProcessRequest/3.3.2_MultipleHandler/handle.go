package main

// 代码清单3-7	使用多个处理器对请求进行处理

import (
	"fmt"
	"net/http"
)

// 通过 http.Handle 函数将处理器绑定至 DefaultServeMux

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

type WorldHandler struct{}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "World!")
}

func main() {
	hello := HelloHandler{}
	world := WorldHandler{}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	// 通过 http.Handle 函数将处理器绑定至 DefaultServeMux
	http.Handle("/hello", &hello) // 参数: pattern, 处理器
	http.Handle("/world", &world)
	server.ListenAndServe()
}
