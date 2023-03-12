package main

// 代码清单7-6	使用Decoder分析XML

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

// struct from xml.go

func main() {
	filename := "chapter7_GoWeb/7.4_GoAndXML/decoder/post.xml"
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()

	// 流方式传输的XML文件
	decoder := xml.NewDecoder(xmlFile) // 根据给定的xml数据，生成相应的解码器
	for {                              // 没迭代一次解码器中所有XML数据
		t, err := decoder.Token() // 就从解码器中获取一个token
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding XML into tokens:", err)
		}

		switch se := t.(type) { // 检查token的类型
		case xml.StartElement:
			if se.Name.Local == "comment" {
				var comment Comment
				decoder.DecodeElement(&comment, &se) // 将XML数据解码至结构
			}
		}
	}

}
