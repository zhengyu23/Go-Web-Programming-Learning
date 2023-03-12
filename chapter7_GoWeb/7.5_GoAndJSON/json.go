package main

// 代码清单7-10	JSON分析程序	-> main1()
// 代码清单7-11	使用Decoder对JSON进行语言分析 -> main2()
// 代码清单7-13	使用Encoder把结构编码为JSON

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	//read()
	//decode()
	encode()
}

func read() {
	filename := "chapter7_GoWeb/7.5_GoAndJSON/post.json"
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}
	var post Post
	json.Unmarshal(jsonData, &post)
	fmt.Println(post)
}

func decode() {
	filename := "chapter7_GoWeb/7.5_GoAndJSON/post.json"
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	for {
		var post Post
		err := decoder.Decode(&post)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}
		fmt.Println(post)
	}
}

func encode() {
	post := Post{
		Id:      1,
		Content: "abc",
		Author: Author{
			Id:   2,
			Name: "LLL",
		},
		Comments: []Comment{
			{
				Id:      3,
				Content: "#3",
				Author:  "#3",
			},
			{
				Id:      4,
				Content: "4",
				Author:  "4",
			},
		},
	}
	filename := "chapter7_GoWeb/7.5_GoAndJSON/post2.json"
	jsonFIle, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating JSON file:", err)
		return
	}
	encoder := json.NewEncoder(jsonFIle)
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encoding JSON to file:", err)
		return
	}
}
