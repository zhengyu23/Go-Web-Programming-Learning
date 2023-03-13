package main

// 代码清单8-3	对main.go进行测试的main_test.go文件
// go test -v -cover

import (
	"testing"
)

func TestDecode(t *testing.T) {
	post, err := decode("post.json") // go test 的根目录是当前目录
	if err != nil {
		t.Error(err)
	}
	if post.Id != 1 {
		t.Error("Wrong id, was expecting 1 but got ", post.Id)
	}
	if post.Content != "Hello World" {
		t.Error("Wrong content, was expecting 'Hello World' but got ", post.Content)
	}
}

func TestEncode(t *testing.T) {
	t.Skip("Skipping encoding for now")
}
