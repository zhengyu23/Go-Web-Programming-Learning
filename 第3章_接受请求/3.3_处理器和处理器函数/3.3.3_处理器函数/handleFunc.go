package main

// 代码清单3-8	使用处理器函数处理请求

import (
	"fmt"
	"net/http"
)

// HandlerFunc函数类型，可以把一个带有正确签名的函数f转换成一个带有方法f的处理器Handler.
// mux.Handle(pattern, HandlerFunc(3.3_处理器和处理器函数))

//	签名 - signature

// 处理器函数： 与 ServeHTTP 方法拥有相同签名的函数, 一种创建处理器的便捷方式
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "hello!")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "world!")
}

/*
func (h * HelloHandler)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}
*/

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	// 比 Handle 方便
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)
	server.ListenAndServe()
}
