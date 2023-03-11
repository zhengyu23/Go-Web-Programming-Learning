package main

// 代码清单3-4	通过 HTTPS 提供服务

import "net/http"

// HttpsWeb 通过HTTPS提供服务

func main() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}
	//	cert.pem: SSL证书 -> 通过HTTPS提供服务(Secure Socket Layer,安全套接层)
	//		是一种通过公钥基础设施(Public Key Infrastructure, PKI)
	//		为通信双方提供数据加密和身份验证的协议. 其中通信双方通常是客户端和服务器.

	//	key.pem: 服务器私钥
	server.ListenAndServeTLS("cert.pem", "key.pem")
}
