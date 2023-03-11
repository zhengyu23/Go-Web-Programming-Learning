package main

// 代码清单4-9	通过 WriterHeader 方法将状态码写入到响应当中

import (
	"fmt"
	"net/http"
)

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service, try next door")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/writeheader", writeHeaderExample)
	server.ListenAndServe()
}
