package main

// 代码清单5-25 使用显式定义的模板

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	filename := "chapter5_ContentDisplay/5.7_NestedTemplate/layout.html"
	t, _ := template.ParseFiles(filename)
	t.ExecuteTemplate(w, "layout", "")

}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
