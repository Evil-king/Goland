package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//打开已经存在的文件 输入并且追加到原来的后面
	filePath := "/Users/fox/Desktop/writer.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Open file err", err)
		return
	}

	defer file.Close()

	//开始写入
	str := "你很的很棒\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
