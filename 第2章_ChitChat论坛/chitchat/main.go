package main

import (
	"net/http"
)

func main() {

	// 1. 多路复用器
	mux := http.NewServeMux() // 多路复用

	// 2. 服务静态文件
	files := http.FileServer(http.Dir("/public")) // 能为指定目录中的静态文件服务的处理器
	// partten: "/static/", 3.3_处理器和处理器函数 Handler
	mux.Handle("/static/", http.StripPrefix("/static", files)) // 使用StripPrefix函数移除请求URL中的指定前缀

	// 3. 创建处理器函数	处理器 - 3.3_处理器和处理器函数
	// HandleFunc 函数把请求重定向到处理器函数 -> route_main.go

	mux.HandleFunc("", index)
	mux.HandleFunc("/err", err)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)
	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()

	// 4. 使用 4.4_Cookie 进行访问控制

	// 5. 使用模板生成 HTML响应	模板 - template
}
