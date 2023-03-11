package main

import (
	"fmt"
	"net/http"
)

func process_(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024) // 从 multipart 编码的表单内取多少字节的数据
	fmt.Fprintln(w, r.MultipartForm)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process_)
	server.ListenAndServe()
}
