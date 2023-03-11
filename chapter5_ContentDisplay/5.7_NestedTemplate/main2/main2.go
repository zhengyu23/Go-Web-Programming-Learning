package main

// 代码清单5-28	处理器使用在不同模板文件中定义的同名模板
import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func process2(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	var t *template.Template
	directory := "chapter5_ContentDisplay/5.7_NestedTemplate/main2/"
	if rand.Intn(10) > 5 {
		t, _ = template.ParseFiles(directory+"layout2.html", directory+"red_hello.html")
	} else {
		t, _ = template.ParseFiles(directory+"layout2.html", directory+"blue_hello.html")
	}
	t.ExecuteTemplate(w, "layout2", "")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process2", process2)
	server.ListenAndServe()
}
