package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
)

// 代码清单6-4	使用gob包读写二进制数据

type Post struct {
	Id      int
	Content string
	Author  string
}

func store(data interface{}, filename string) {
	buffer := new(bytes.Buffer) // 数据存在buffer结构中
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}

func load(data interface{}, filename string) {
	raw, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}

func main() {
	directory := "chapter6_StoreData/6.2_FileStorage/gob/"
	post := Post{Id: 1, Content: "Hello World", Author: "Zheng Yu"}
	store(post, directory+"post1")
	var postRead Post
	load(&postRead, directory+"post1")
	fmt.Println(postRead)
}
