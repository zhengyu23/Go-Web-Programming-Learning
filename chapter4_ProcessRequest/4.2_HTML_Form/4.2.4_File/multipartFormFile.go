package main

// 代码清单4-6	通过MultipartForm字段接收用户上传的文件

import (
	"fmt"
	"io"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)
	fileHeader := r.MultipartForm.File["uploaded"][0]
	file, err := fileHeader.Open()

	if err == nil {
		data, err := io.ReadAll(file)
		//fmt.Fprintln(w, string(data), err)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
