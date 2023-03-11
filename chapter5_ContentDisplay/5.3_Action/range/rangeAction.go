package main

// 代码清单5-5	迭代动作演示

import (
	"html/template"
	"net/http"
)

func rangeProcess(w http.ResponseWriter, h *http.Request) {
	filename := "chapter5_ContentDisplay/5.3_Action/range/rangeTmpl.html"
	t, _ := template.ParseFiles(filename)
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	t.Execute(w, daysOfWeek)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/rangeProcess", rangeProcess)
	server.ListenAndServe()
}
