package main

// 代码清单9-1	goroutine使用示例

import (
	"fmt"
	"time"
)

func printNumbers1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
	}
}
func printLetter1() {
	for i := 'A'; i < 'A'+10; i++ {
		fmt.Printf("%d", i)
	}
}

func print1() {
	printNumbers1()
	printLetter1()
}

func main() {
	go print1()
	time.Sleep(1 * time.Millisecond)
}
