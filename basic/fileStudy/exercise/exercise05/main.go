package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/**
将一张图片拷贝到另一个文件夹中
*/
//自己编写一个函数，接受两个文件路径 srcFileName(源文件) dstFileName(目标文件)
func CopeFile(srcFileName string, dstFileName string) (Writer int64, err error) {
	//srcFileName 需要是一个reader类型的
	file, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("Open file is err=", err)
	}
	reader := bufio.NewReader(file)

	defer file.Close()

	//dstFileName 需要是一个writer类型的
	dsFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666) //只读如果文件不存在这创建
	if err != nil {
		fmt.Println("OpenFile file is err=", err)
	}
	writer := bufio.NewWriter(dsFile)
	defer dsFile.Close()

	return io.Copy(writer, reader)
}

func main() {
	srcFileName := "/Users/fox/Desktop/map.jpg"
	dstFileName := "/Users/fox/project/map.jpg"
	_, err := CopeFile(srcFileName, dstFileName)
	if err == nil {
		fmt.Printf("拷贝完成")
	} else {
		fmt.Println("拷贝失败")
	}

}
