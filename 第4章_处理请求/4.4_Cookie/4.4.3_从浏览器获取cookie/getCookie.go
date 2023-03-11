package main

// 代码清单4-15	从请求的首部获取 4.4_Cookie
// 代码清单4-16	使用 Cookie 方法	-> 取得单独的键值对格式的 4.4_Cookie

import (
	"fmt"
	"net/http"
)

func getCookie(w http.ResponseWriter, r *http.Request) {
	h := r.Header["Cookie"]
	fmt.Fprintln(w, h) // [first_cookie="Go Web Programming"; second_cookie="Manning Publication Co"]

	// 取得单独的键值对格式的 4.4_Cookie
	c1, err := r.Cookie("first_cookie")
	if err != nil {
		fmt.Fprintln(w, "Cannot get the first 4.4_Cookie!")
	}
	cs := r.Cookies()
	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, cs)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/get_cookie", getCookie)
	server.ListenAndServe()
}
