package main

import (
	"chitchat/data"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	if err == nil {
		_, err := session(w, r) // 判断用户是否登录
		publicTmplFiles := []string{"templates/layout.html",
			"templates/navbar.html",
			"templates/index.html"}
		privateTmplFiles := []string{"templates/layout.html",
			"templates/navbar.html",
			"templates/index.html"}
		var templates *template.Template

		// 未登录使用public导航条，已登录使用private导航条
		if err != nil {
			// template.ParseFiles(): 对模板文件进行语法分析,并创建出相应的模板
			templates = template.Must(template.ParseFiles(privateTmplFiles...))
		} else {
			templates = template.Must(template.ParseFiles(publicTmplFiles...))
		}
		templates.ExecuteTemplate(w, "layout", threads) // layout 模板包含 content 和 navbar, 所以会一并执行
	}
}

func indexOld(w http.ResponseWriter, r *http.Request) {
	files := []string{"templates/layout.html",
		"templates/navbar.html",
		"templates/index.html"}

	templates := template.Must(template.ParseFiles(files...))
	threads, err := data.Threads() // Threads 从数据库里取出所有帖子并返回给index处理器函数
	if err != nil {
		templates.ExecuteTemplate(w, "layout", threads)
	}
}

func err(w http.ResponseWriter, r *http.Request) {}

func login(w http.ResponseWriter, r *http.Request) {}

func logout(w http.ResponseWriter, r *http.Request) {}

func signup(w http.ResponseWriter, r *http.Request) {}

func signupAccount(w http.ResponseWriter, r *http.Request) {}

// 对用户的身份进行验证,并存放到 http_cookie
func authenticate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 从浏览器取得
	// data.UserByEmail 查找
	// data.Encrypt 加密
	user, _ := data.UserByEmail(r.PostFormValue("email")) //
	if user.Password == data.Encrypt(r.PostFormValue("password")) {
		session := user.CreateSession()
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)

	} else {
		http.Redirect(w, r, "login", 302) // 假如未通过验证,重定向到登录页面
	}
}

func newThread(w http.ResponseWriter, r *http.Request) {}

func createThread(w http.ResponseWriter, r *http.Request) {}

func postThread(w http.ResponseWriter, r *http.Request) {}

func readThread(w http.ResponseWriter, r *http.Request) {}
