package main

// 代码清单7-7	使用Marshal函数生成XML文件

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
	Xml     string   `xml:",innerxml"`

	Comments []Comment `xml:"comments>comment"`
}

type Comment struct {
	Id      string `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:"author"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

// struct from xml.go

func main() {
	post := Post{
		Id:      "1",
		Content: "hello world",
		Author: Author{
			Id:   "2",
			Name: "Zheng Yu",
		},
	}

	//output, err := xml.Marshal(&post)
	output, err := xml.MarshalIndent(&post, "", "\t")
	if err != nil {
		fmt.Println("Error marshlling to XML:", err)
		return
	}
	err = os.WriteFile("chapter7_GoWeb/7.4_GoAndXML/marshal/post2.xml", []byte(xml.Header+string(output)), 0644)
	if err != nil {
		fmt.Println("Error writing XML to file:", err)
		return
	}
}
