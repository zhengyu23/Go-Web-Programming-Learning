package main

// 代码清单6-14	使用Go语言实现一对多以及多对多关系
// 代码清单6-15	创建评论，并建立评论与帖子之间的关系
// 代码清单6-16	获取帖子及其评论

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

type Post struct {
	Id       int
	Content  string
	Author   string
	Comments []Comment
}

type Comment struct {
	Id      int
	Content string
	Author  string
	Post    *Post
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

// Create 创建一条评论
func (comment *Comment) Create() (err error) {
	if comment.Post == nil {
		err = errors.New("post not found")
		return
	}
	err = Db.QueryRow("insert into comments (content, author, post_id) values ($1, $2, $3) returning id",
		comment.Content, comment.Author, comment.Post.Id).Scan(&comment.Id)
	return
}

// GetPost 获得单独一篇帖子
func GetPost(id int) (post Post, err error) {
	post = Post{}
	post.Comments = []Comment{}

	err = Db.QueryRow("select id, content, author from posts where id = $1",
		id).Scan(&post.Id, &post.Content, &post.Author)

	rows, err := Db.Query("select id, content, author from comments")
	if err != nil {
		return
	}

	for rows.Next() {
		comment := Comment{Post: &post}
		err = rows.Scan(&comment.Id, &comment.Content, &comment.Author)
		if err != nil {
			return
		}
		post.Comments = append(post.Comments, comment)
	}
	rows.Close()
	return
}

// Create 创建一篇新帖子
func (post *Post) Create() (err error) {
	err = Db.QueryRow("insert into posts (content, author) values ($1, $2) returning id",
		post.Content, post.Author).Scan(&post.Id)
	return
}

func main() {
	post := Post{Content: "Hello World", Author: "Zheng Yu"}
	post.Create()

	comment := Comment{Content: "Good post!", Author: "Cai", Post: &post}
	comment.Create()
	readPost, _ := GetPost(post.Id)
	fmt.Println(readPost)
	fmt.Println(readPost.Comments)
	fmt.Println(readPost.Comments[0].Post)
}
