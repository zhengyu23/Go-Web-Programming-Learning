package main

// 代码清单5-17	为了展示模板中的上下文感知特性而设置的处理器
import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	filename := "chapter5_ContentDisplay/5.6_context-aware/invalidTmpl.html"
	t, _ := template.ParseFiles(filename)
	content := `I asked: <i>"What's up?"</i>`
	t.Execute(w, content)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
