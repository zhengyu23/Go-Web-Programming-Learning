package main

// 代码清单5-3	在处理器里面生成一个随机数

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func ifProcess(w http.ResponseWriter, h *http.Request) {
	filename := "chapter5_ContentDisplay/5.3_Action/if/ifTmpl.html"
	t, _ := template.ParseFiles(filename)
	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 5)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/ifProcess", ifProcess)
	server.ListenAndServe()
}
