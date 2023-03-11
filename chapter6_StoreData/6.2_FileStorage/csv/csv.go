package main

// 代码清单6-3	读写CSV文件

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func main() {
	directory := "chapter6_StoreData/6.2_FileStorage/csv/"
	csvFile, err := os.Create(directory + "posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	allPosts := []Post{
		Post{Id: 1, Content: "Hello World!", Author: "Zheng Yu"},
		Post{Id: 2, Content: "Bonjour Monde!", Author: "Cheng Bin"},
		Post{Id: 3, Content: "Hola Mundo!", Author: "Lao Liang"},
		Post{Id: 4, Content: "Greetings Earthlings!", Author: "Zheng Yu"},
	}

	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}

	writer.Flush()

	file, err := os.Open(directory + "posts.csv")
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []Post
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{Id: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}
	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)
}
