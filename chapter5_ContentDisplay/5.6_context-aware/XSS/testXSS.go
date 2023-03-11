package main

// 代码清单5-20	测试XSS攻击

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	filename := "chapter5_ContentDisplay/5.6_context-aware/XSS/invalidTmpl.html"
	t, _ := template.ParseFiles(filename)
	t.Execute(w, r.FormValue("comment"))
}

func form(w http.ResponseWriter, r *http.Request) {
	filename := "chapter5_ContentDisplay/5.6_context-aware/XSS/form.html"
	t, _ := template.ParseFiles(filename)
	t.Execute(w, nil)
	// 输入<script>alert('Pwnd!'); </script>
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/form", form)
	server.ListenAndServe()
}
