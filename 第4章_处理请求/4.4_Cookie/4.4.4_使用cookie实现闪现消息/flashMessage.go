package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

// 代码清单4-17	使用Go的 4.4_Cookie 实现闪现消息
// 代码清单4-18	设置消息 -> setMessage

func setMessage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Hello World!")
	c := http.Cookie{
		Name:  "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, &c)
}
func showMessage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("flash")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "No message found")
		}
	} else {
		rc := http.Cookie{
			Name:    "flash",         // 设置同名为"flash"的cookie
			MaxAge:  -1,              // MaxAge 过期时间设置为负数
			Expires: time.Unix(1, 0), // Expires 设置为一个已经过去的时间
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(val))
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/set_message", setMessage)
	http.HandleFunc("/show_message", showMessage)
	server.ListenAndServe()
}
