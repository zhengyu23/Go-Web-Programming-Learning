package main

// 代码清单6-2	对文件进行读写

import (
	"fmt"
	"os"
)

func main() {
	directory := "chapter6_StoreData/6.2_FileStorage/csv/"
	data := []byte("Hello World!\n")

	err := os.WriteFile(directory+"data1", data, 0644)
	if err != nil {
		panic(err)
	}
	read1, _ := os.ReadFile(directory + "data1")
	fmt.Println(string(read1))

	file1, _ := os.Create(directory + "data2")
	defer file1.Close()
	bytes, _ := file1.Write(data) // 返回写入字节数 bytes
	fmt.Printf("Wrote %d bytes to file\n", bytes)

	file2, _ := os.Open(directory + "data2")
	defer file2.Close()
	read2 := make([]byte, len(data))
	bytes, _ = file2.Read(read2)
	fmt.Printf("Read %d bytes from file\n", bytes)
	fmt.Println(string(data))

}
