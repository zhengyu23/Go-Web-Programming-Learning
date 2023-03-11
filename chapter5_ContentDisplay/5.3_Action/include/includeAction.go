package main

// 代码清单5-11	调用嵌套模板

import (
	"html/template"
	"net/http"
)

func includeProcess(w http.ResponseWriter, r *http.Request) {
	directory := "chapter5_ContentDisplay/5.3_Action/include/"
	t1 := directory + "includeTmpl.html"
	t2 := directory + "includeTmpl2.html"
	t, _ := template.ParseFiles(t1, t2)
	t.Execute(w, "Hello World!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/includeProcess", includeProcess)
	server.ListenAndServe()
}
