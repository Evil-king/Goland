package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("/Users/fox/Desktop/temp.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}
	//当函数退出时，要及时的关闭file
	defer file.Close() //记得及时关闭 否则会内存泄露

	//创建一个*Reader  是带缓冲的 默认的大小是4096
	reader := bufio.NewReader(file)
	//循环读取文件内容
	for {
		str, _, err := reader.ReadLine()
		if err == io.EOF { //io.EOF表示文件末尾
			break
		}
		//输出内容
		fmt.Println(string(str))
	}
	fmt.Println("读取文件结束...")
}
