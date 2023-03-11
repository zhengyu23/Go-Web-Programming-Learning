package main

// 代码清单6-1	在内存里面存储数据

import (
	"fmt"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

var PostById map[int]*Post
var PostsByAuthor map[string][]*Post

func store(post Post) {
	PostById[post.Id] = &post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
}

func main() {
	PostById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	post1 := Post{Id: 1, Content: "Hello World!", Author: "Zheng Yu"}
	post2 := Post{Id: 2, Content: "Bonjour Monde!", Author: "Cheng Bin"}
	post3 := Post{Id: 3, Content: "Hola Mundo!", Author: "Lao Liang"}
	post4 := Post{Id: 4, Content: "Greetings Earthlings!", Author: "Zheng Yu"}

	store(post1)
	store(post2)
	store(post3)
	store(post4)

	for _, post := range PostsByAuthor["Zheng Yu"] {
		fmt.Println(post)
	}
	for _, post := range PostsByAuthor["Lao Liang"] {
		fmt.Println(post)
	}
}
