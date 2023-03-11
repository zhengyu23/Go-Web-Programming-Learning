package main

// 代码清单3-10	串联两个处理器函数

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func hello_(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

// HandlerFunc函数类型，可以把一个带有正确签名的函数f转换成一个带有方法f的Handler.
// mux.Handle(pattern string, handler Handler)
func log(h http.HandlerFunc) http.HandlerFunc {
	// 返回处理器函数
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h(w, r)
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/hello", log(hello_)) // 参数: pattern, 处理器函数
	server.ListenAndServe()

}
