package main

// 代码清单5-7	对点进行设置

import (
	"html/template"
	"net/http"
)

func rangeProcess(w http.ResponseWriter, h *http.Request) {
	filename := "chapter5_ContentDisplay/5.3_Action/withArg/withArgTmpl.html"
	t, _ := template.ParseFiles(filename)
	t.Execute(w, "hello")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/rangeProcess", rangeProcess)
	server.ListenAndServe()
}
