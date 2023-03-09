package main

import "net/http"

func main() {
	HttpsWeb()
}

// 最简单的web服务
func BestEasyWeb() {
	http.ListenAndServe("", nil) // 默认8080
}

// 带有附加配置的Web服务器
func ConfigWeb() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}
	server.ListenAndServe()
}

// 通过HTTPS提供服务
func HttpsWeb() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}
	//	cert.pem: SSL证书 -> SSL(Secure Socket Layer,安全套接层)
	//		是一种通过公钥基础设施(Public Key Infrastructure, PKI)
	//		为通信双方提供数据加密和身份验证的协议. 其中通信双方通常是客户端和服务器.

	//	key.pem: 服务器私钥
	server.ListenAndServeTLS("cert.pem", "key.pem")
}
