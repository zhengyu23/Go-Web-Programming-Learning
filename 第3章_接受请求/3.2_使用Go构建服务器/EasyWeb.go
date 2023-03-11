package main

// 代码清单3-1	最简单的 Web 服务器
// 代码清单3-2	带有附加配置的 Web 服务器
// 代码清单3-3	Server 结构的配置选项

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
}

// BestEasyWeb 最简单的web服务
func BestEasyWeb() {
	http.ListenAndServe("", nil) // 默认8080
}

// ConfigWeb 带有附加配置的Web服务器
func ConfigWeb() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}
	server.ListenAndServe()
}

// Server 结构的配置选项
type Server struct {
	Addr           string
	Handler        http.Handler
	ReadTimeout    time.Duration
	MaxHeaderBytes int
	TLSConfig      *tls.Config
	TLSNextProto   map[string]func(*http.Server, *tls.Conn, http.Handler)
	ConnState      func(conn net.Conn, state http.ConnState)
	ErrorLog       *log.Logger
}
