package main

// 代码清单5-14	创建模板自定义函数

import (
	"html/template"
	"net/http"
	"time"
)

func formatData(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func process(w http.ResponseWriter, h *http.Request) {
	funcMap := template.FuncMap{"fdate": formatData}
	t := template.New("invalidTmpl.html").Funcs(funcMap)
	t, _ = t.ParseFiles("chapter5_ContentDisplay/5.5_function/invalidTmpl.html")
	t.Execute(w, time.Now())
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
