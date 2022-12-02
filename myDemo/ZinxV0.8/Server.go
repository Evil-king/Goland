package main

import (
	"LearnGo/zinx/ziface"
	"LearnGo/zinx/znet"
	"fmt"
)

/*
	基于Zinx框架来开发的 服务器端应用程序
*/

//ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

//Test Handle
func (pr *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingZinxRouter Handle...")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId=", request.GetMsgId(), ", data=", string(request.GetData()))

	//回写数据
	err := request.GetConnection().SendMsg(200, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}

//hell Zinx test 自定义路由
type HelloRouter struct {
	znet.BaseRouter
}

//Test Handle
func (pr *HelloRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call HelloZinxRouter Handle...")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId=", request.GetMsgId(), ", data=", string(request.GetData()))

	//回写数据
	err := request.GetConnection().SendMsg(201, []byte("Hello Welcome Zinx!!!"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	//1、创建一个server句柄 使用Zinx的api
	s := znet.NewServer()

	//2、给当前zinx框架添加自定义的router
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})

	//3、启动server
	s.Serve()
}
