package main

import (
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	//这里我们循环的接受客户端发送的链接
	defer conn.Close() //关闭conn

	for {
		//创建一个新的切面
		buf := make([]byte, 1024)
		//conn.Read(buf) 等待客户端通过conn发送信息 如果客户端没有write 那么协程就阻塞在这里
		fmt.Printf("服务器在等待客户端%s 发送信息\n" , conn.RemoteAddr().String())
		n, err := conn.Read(buf) //从conn读取
		if err != io.EOF {
			fmt.Println("客户端退出")
			return
		}
		//显示客户端发送的内容到服务器的终端
		fmt.Print(string(buf[:n]))
	}

}

func main() {
	fmt.Println("服务器哦开始监听......")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	defer listen.Close() //延时关闭

	//循环等待客户端来连接
	for {
		//等待客户端链接
		fmt.Println("等待客户端来链接.....")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
		} else {
			fmt.Printf("Accept() suc=%v\n，客户端ip=%v\n", conn, conn.RemoteAddr())
		}
		//这里准备一个协程，为客户端服务
		go process(conn)
	}

	fmt.Printf("listen suc=%v\n", listen)
}
