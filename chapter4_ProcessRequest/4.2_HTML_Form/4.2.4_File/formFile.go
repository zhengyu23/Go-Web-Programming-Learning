package main

// 代码清单4-7	使用 FormFile 方法获取被上传的文件

import (
	"fmt"
	"io"
	"net/http"
)

func process_(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("uploaded")
	if err == nil {
		data, err := io.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process_)
	server.ListenAndServe()
}
