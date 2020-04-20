package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//打开已经存在的文件 将原来的内容读出显示在终端 并且追加五句
	filePath := "/Users/fox/Desktop/writer.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Open file err", err)
		return
	}

	defer file.Close()

	//读出显示在终端
	reader := bufio.NewReader(file)
	for {
		str,err := reader.ReadString('\n')
		if err == io.EOF { //io.EOF表示文件末尾
			break
		}
		//输出内容
		fmt.Print(string(str))
	}

	//开始写入
	str := "1111\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
