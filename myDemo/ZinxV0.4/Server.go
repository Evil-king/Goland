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

//Test PreRouter
func (pr *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router PreHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping...\n"))
	if err != nil {
		fmt.Println("call back before ping err")
	}
}

//Test Handle
func (pr *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping\n"))
	if err != nil {
		fmt.Println("call back ping...ping...ping err")
	}
}

//Test PostRouter
func (pr *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router PostHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping...\n"))
	if err != nil {
		fmt.Println("call back after ping err")
	}
}

func main() {
	//创建一个server句柄 使用Zinx的api
	s := znet.NewServer()
	//给当前zinx框架添加一个自定义的router
	s.AddRouter(0, &PingRouter{})
	//启动server
	s.Serve()
}
