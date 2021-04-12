package main

import (
	"fmt"
	"net"
)

//创建用户结构体类型
type Client struct {
	C    chan string
	Name string
	Addr string
}

//创建全局map，存储在线用户
var onlineMap map[string]Client

//创建全局channel 传递用户消息
var message = make(chan string)

func WriteMessageToClient(clnt Client, conn net.Conn) {
	//监听 用户自带channel是否自带消息
	for msg := range clnt.C {
		conn.Write([]byte(msg + "\n"))
	}
}

func MakeMsg(clnt Client, msg string) (buf string) {
	buf = "[" + clnt.Addr + "]" + clnt.Name + ": " + msg
	return buf
}

func HandlerConnect(conn net.Conn) {
	defer conn.Close()

	//获取用户 网络地址 ID+pory
	netAddr := conn.RemoteAddr().String()
	//创建新连接用户的 结构体 默认用户是 IP+port
	clnt := Client{make(chan string), netAddr, netAddr}

	//将新连接用户添加到 在线map中 key:IP+port value client
	onlineMap[netAddr] = clnt

	//创建专门用来给当前用户发送消息的go程
	go WriteMessageToClient(clnt, conn)

	//发送用户上线到全局channel中
	//message <- "[" + netAddr + "]" + clnt.Name + "login"
	message <- MakeMsg(clnt, "login")

	//创建一个匿名 go 程，专门处理用户发送的消息
	go func() {
		for {
			buf := make([]byte, 4096)
			n, err := conn.Read(buf)
			if n == 0 {
				fmt.Printf("检测到客户端:%s退出\n", clnt.Name)
				return
			}
			if err != nil {
				fmt.Println("conn.Read err", err)
				return
			}
			//将读到的用户消息 保存到msg中 string类型
			msg := string(buf[:n])
			//将读到的用户消息，写入到message中
			message <- MakeMsg(clnt, msg)
		}
	}()

	for {

	}
}

func Manager() {
	//初始化onlineMap
	onlineMap := make(map[string]Client)
	//监听全局channel是否有 数据 有数据存储到msg 无数据阻塞
	for {
		msg := <-message
		//循环发送消息给在线用户 要想执行 必须msg := <-message 执行完 接触阻塞
		for _, clnt := range onlineMap {
			clnt.C <- msg
		}
	}

}

func main() {
	//创建监听套接字
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("net.Listen err", err)
		return
	}
	defer listener.Close()

	//创建管理者go程 管理map和全局channel
	go Manager()

	//循环监听客户端连接请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err", err)
			return
		}
		//启动go程处理客户端数据请求
		go HandlerConnect(conn)
	}
}