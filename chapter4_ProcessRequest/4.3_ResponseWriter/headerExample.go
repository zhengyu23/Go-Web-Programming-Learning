package main

// 代码清单4-10	通过编写首部实现客户端重定向

import (
	"net/http"
)

func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://gooogle.com")
	w.WriteHeader(302)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/redirect", headerExample)
	server.ListenAndServe()
}
