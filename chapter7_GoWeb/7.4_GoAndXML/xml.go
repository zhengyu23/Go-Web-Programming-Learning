package main

// 代码清单7-2	对XML进行分析
// 代码清单7-5	带有comments结构字段的Post结构

import (
	"encoding/xml"
	"fmt"
	"io"
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

func main() {
	filename := "chapter7_GoWeb/7.4_GoAndXML/post.xml"
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()
	xmlData, err := io.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading XML file:", err)
	}
	var post Post
	xml.Unmarshal(xmlData, &post)
	fmt.Println(post)
}
