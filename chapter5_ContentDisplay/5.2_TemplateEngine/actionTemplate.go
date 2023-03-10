package main

import (
	"html/template"
	"net/http"
)

// 代码清单5-2	在处理器函数中触发模板引擎

func process(w http.ResponseWriter, r *http.Request) {
	directory := "./chapter5_ContentDisplay/5.2_TemplateEngine"
	t, _ := template.ParseFiles(directory + "/simpleTemplate.html")
	t.Execute(w, "Hello World!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
