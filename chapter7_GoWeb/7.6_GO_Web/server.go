package main

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"
)

// 代码清单7-15	定义在server.go文件内的Go Web服务

// curl -i -X POST -H "Content-Type:application/json" -d '{"content":"My first post","author":"Zheng yu"}' http://127.0.0.1:8080/post/
// psql -U gwp -d gwp -c "select * from posts;"

type Post struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/post/", handleRequest)
	server.ListenAndServe()
}

// handleRequest 多路复用器负责将请求转发给正确的处理器函数
func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handleUpdate(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	default:
		http.Error(w, "invalid request", http.StatusInternalServerError)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// handleGet 获取指定的帖子
func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	// ① 从url中获取id
	// ② 调用retrieve函数获取post
	// ③ 将post结构体marshal为JSON
	// ④ 浏览器数据:json
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := retrieve(id)
	if err != nil {
		return
	}
	output, err := json.MarshalIndent(post, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// handlePost 创建新的帖子
func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	// ① 读取Body
	// ② 将Body内容Unmarshal为post结构体
	// ③ post调用create方法创建新的帖子
	// ④ 浏览器数据:200
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var post Post
	json.Unmarshal(body, &post)
	err = post.create()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

// handlePut 更新指定的帖子
func handleUpdate(w http.ResponseWriter, r *http.Request) (err error) {
	// ① 从URL中获取id
	// ② 调用retrieve函数获取post
	// ③ 读取Body
	// ④ 将Body内容Unmarshal为post结构体
	// ⑤ post调用update方法
	// ⑥ 浏览器数据:200
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	post, err := retrieve(id)

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	json.Unmarshal(body, &post)

	err = post.update()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

// handleDelete 删除指定的帖子
func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	// ① 从url中获取id
	// ② 调用retrieve函数获取post
	// ③ post调用delete方法删除自身
	// ④ 浏览器数据:200
	id, err := strconv.Atoi(path.Base(r.URL.Path))

	post, err := retrieve(id)

	err = post.delete()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}
