package main

import (
	"fmt"
	"net/http"
)

func process__(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "(1)", r.FormValue("hello"))     // url
	fmt.Fprintln(w, "(2)", r.PostFormValue("hello")) // 表单
	fmt.Fprintln(w, "(3)", r.PostForm)               // 表单
	fmt.Fprintln(w, "(4)", r.MultipartForm)          // 表单
	fmt.Fprintln(w, "(5)", r.ParseForm())            // 无 只支持 application/x-www-form-urlencoded
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process__)
	server.ListenAndServe()
}
