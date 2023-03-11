package main

// 代码清单4-3	读取请求主题中的数据

import (
	"fmt"
	"net/http"
)

func body(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	//fmt.Fprintln(w, string(body))
	//a, _ := r.Body.Read(body)
	r.Body.Read(body)
	//fmt.Fprintln(w, a)
	fmt.Fprintln(w, string(body))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/body", body)
	server.ListenAndServe()
}
