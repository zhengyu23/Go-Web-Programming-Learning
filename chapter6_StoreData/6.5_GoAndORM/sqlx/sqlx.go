package main

// 代码清单6-17	使用Sqlx重新实现论坛程序

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Post struct {
	Id         int
	Content    string
	AuthorName string `db:author`
}

var Db *sqlx.DB

func init() {
	var err error
	Db, err = sqlx.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func GetPost(id int) (err error) {
	post := Post{}
	err = Db.QueryRowx("select id, content, author form posts where id = $1", id).StructScan(&post)
	if err != nil {
		return
	}
	return
}

func (post *Post) Create() (err error) {
	err = Db.QueryRow("insert into posts (content, author) values($1, $2) returning id",
		post.Content, post.AuthorName).Scan(&post.Id)
	return
}

func main() {
	post := Post{Content: "Hello World", AuthorName: "Zheng Yu"}
	post.Create()
	fmt.Println(post)
}
