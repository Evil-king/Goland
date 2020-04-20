package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
写入数据到文件
 */
func main() {
	filePath := "/Users/fox/Desktop/writer.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("Open file err=%v", err)
		return
	}

	//及时关闭file句柄
	defer file.Close()

	str := "Hello,Gardon\n"
	//写入时带缓存的 *Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	//因为writer是带缓存的 所有str是需要Flush 否则都是写在缓存中的
	writer.Flush()
	fmt.Println("文件写入完成.....")
}
