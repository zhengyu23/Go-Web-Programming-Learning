package main

// 代码清单6-6	使用Go对Postgre执行CRUD操作
// 代码清单6-7	用于创建数据库句柄的函数
// 代码清单6-8	创建一篇帖子
// 代码清单6-9	获取一篇帖子
// 代码清单6-10	更新一篇帖子
// 代码清单6-11	删除一篇帖子
// 代码清单6-12	一次获取多篇帖子

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres",
		"user=gwp dbname=gwp password=741213 sslmode=disable")
	if err != nil {
		panic(err)
	}
}

// Posts 获得指定数量的帖子
func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("select id, content, author form posts limit $1", limit)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

// GetPost 获得单独一片帖子
func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where "+
		"id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

// 创建一篇新帖子
func (post *Post) Create() (err error) {
	statement := "insert into posts (content, author) values ($1,$2) returning id"
	// Prepare creates a prepared statement for later queries or executions.
	stmt, err := Db.Prepare(statement)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

// Update 更新帖子
func (post *Post) Update() (err error) {
	_, err = Db.Exec("update posts set content = $2, author = $3 where id = $1",
		post.Id, post.Content, post.Author)
	return
}

// Delete 删除帖子
func (post *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}

func main() {
	post := Post{Content: "Hello World", Author: "Zheng Yu"}

	fmt.Println(post)
	post.Create()
	fmt.Println(post)

	readPost, _ := GetPost(post.Id)
	fmt.Println(readPost)

	readPost.Content = "Bonjour Monde!"
	readPost.Author = "Pierre"
	readPost.Update()

	posts, _ := Posts(1)
	fmt.Println(posts)

	readPost.Delete()
}
