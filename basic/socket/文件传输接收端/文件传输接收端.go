package main

import (
	"fmt"
	"net"
	"os"
)

func revFile(conn net.Conn, fileName string) {
	//按文件名创建文件
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create err", err)
		return
	}
	defer f.Close()

	//从网络读取文件 写入到本地
	buf := make([]byte, 4096)
	for {
		n, _ := conn.Read(buf)
		if n == 0 {
			fmt.Println("文件读取成功")
			return
		}
		//读多少，写多少
		conn.Write(buf[:n])
	}
}

func main() {
	//创建连接用于监听socket
	listener, err := net.Listen("tcp", "217.0.0.1:8008")
	if err != nil {
		fmt.Println("net.Listen err", err)
		return
	}
	defer listener.Close()
	//阻塞监听
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept err", err)
		return
	}
	defer conn.Close()

	//读取文件名,保存
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err", err)
		return
	}

	fileName := string(buf[:n])
	//回写给服务端OK
	conn.Write([]byte("ok"))

	//接收文件内容
	revFile(conn, fileName)
}
