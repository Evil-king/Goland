package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//将一个文件倒入到另一个文件中去

	//首先将a文件读取到内存中
	filePath := "/Users/fox/Desktop/a.txt"

	filePath2 := "/Users/fox/Desktop/b.txt"

	//先自动写入一些数据到a.txt中
	file, err := os.OpenFile(filePath,os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("writer a.txt file is err", err)
		return
	}
	defer file.Close()
	str := "This is a ignore day\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 1; i++ {
		writer.WriteString(str)
	}
	writer.Flush()

	//打开文件路径并获取到data切片
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Open file is err", err)
		return
	}

	//将读取到的内存写入到b文件中
	err = ioutil.WriteFile(filePath2, data, 0)
	if err != nil {
		fmt.Printf("writer file err=%v,", err)
	}

}
