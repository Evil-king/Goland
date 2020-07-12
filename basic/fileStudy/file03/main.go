package main

import (
	"fmt"
	"io/ioutil"
)

/**
  这种方式是一次性全部读取文件中的内容 不适合文件内容过大
*/
func main() {
	file := "/Users/fox/Desktop/temp.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("read file err=%v", err)
	}
	fmt.Printf("%v", string(content)) //content 是以byte[]的形式显示的 需要用string转一下
	//ioutil.ReadFile不需要打开文件和关闭文件 文件的操作都封装在ReadFile函数内部
}
