package main

// 代码清单4-14	使用 SetCookie 方法设置 4.4_Cookie

import (
	"net/http"
)

func setCookieFast(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Go Web Programming",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "Manning Publication Co",
		HttpOnly: true,
	}
	http.SetCookie(w, &c1) // 更快地设置浏览器cookie
	http.SetCookie(w, &c2)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/set_cookie", setCookieFast)
	server.ListenAndServe()
}
