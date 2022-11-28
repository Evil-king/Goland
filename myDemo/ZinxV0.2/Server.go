package main

import "LearnGo/zinx/znet"

/*
	基于Zinx框架来开发的 服务器端应用程序
*/

func main() {
	//创建一个server句柄 使用Zinx的api
	s := znet.NewServer()
	//启动server
	s.Serve()
}
