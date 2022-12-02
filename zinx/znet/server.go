package znet

import (
	"LearnGo/zinx/utils"
	"LearnGo/zinx/ziface"
	"fmt"
	"net"
)

// Server iServer的接口实现，定义一个Server的服务器模块
type Server struct {
	//服务器的名称
	Name string
	//服务器绑定的ip版本
	IPVersion string
	//服务器监听的IP
	Ip string
	//服务器监听的端口
	Port int
	//当前Server的消息管理模块，用来绑定MsgId和对应的处理方法
	MsgHandler ziface.IMsgHandler
	//当前Server的链接管理器
	ConnManager ziface.IConnManager
	// =======================
	//新增两个hook函数原型

	//该Server创建链接之后自动调用Hook函数-OnConnStart
	OnConnStart func(connection ziface.IConnection)
	//该Server销毁链接之前自动调用Hook函数-OnConnStop
	OnConnStop func(connection ziface.IConnection)
}

// NewServer 初始化Server
func NewServer() ziface.IServer {
	utils.GlobalObject.Reload()
	return &Server{
		Name:        utils.GlobalObject.Name,
		IPVersion:   "tcp4",
		Ip:          utils.GlobalObject.Host,
		Port:        utils.GlobalObject.TcpPort,
		MsgHandler:  NewMsgHandler(),
		ConnManager: NewConnManager(),
	}
}

// Start 启动服务器
func (s *Server) Start() {
	fmt.Printf("[Zinx] Server Listener at IP :%s,Port :%d,is starting\n",
		utils.GlobalObject.Host,
		utils.GlobalObject.TcpPort)
	fmt.Printf("[Zinx] Version: %s, MaxConn: %d,  MaxPacketSize: %d\n",
		utils.GlobalObject.Version,
		utils.GlobalObject.MaxConn,
		utils.GlobalObject.MaxPacketSize)

	go func() {
		//启动WorkerPool
		s.MsgHandler.StartWorkerPool()

		//1、获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.Ip, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err:", err)
			return
		}
		//2、监听服务器的地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
			return
		}

		var cid uint32
		cid = 0

		//3、阻塞的等待客户端链接，处理客户端链接也为(读写)
		for {
			//3.1 如果有客户端链接过来，阻塞会返回
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("AcceptTCP err", err)
				continue
			}

			//3.2 设置服务器最大连接控制,如果超过最大连接，那么则关闭此新的连接(防止服务器因为连接I/O过多而出现问题)
			if s.ConnManager.Len() >= utils.GlobalObject.MaxConn {
				fmt.Println("当前连接数已达到服务器最大上限.....")
				conn.Close()
				continue
			}

			//3.3 将处理新链接的业务方法 和conn 进行绑定 得到我们的链接模块
			dealConn := NewConnection(s, conn, cid, s.MsgHandler)
			cid++

			//3.4 启动当前的链接业务处理
			go dealConn.Start()
		}
	}()

}

// Stop 停止服务器
func (s *Server) Stop() {
	// 将一些服务器的资源、状态或者一些已经开辟的链接信息 进行停止或者回收
	fmt.Println("[STOP] Zinx server , name ", s.Name)

	//将其他需要清理的连接信息或者其他信息 也要一并停止或者清理
	s.ConnManager.ClearConn()
}

// Serve 运行服务器
func (s *Server) Serve() {
	//启动server的服务功能
	s.Start()

	//TODO 做一些启动服务器之后的额外业务

	//阻塞状态
	select {}
}

// AddRouter 路由功能，给当前的服务注册一个路由方法，供客户端的链接处理试用
func (s *Server) AddRouter(msgId uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouter(msgId, router)
	fmt.Println("Add Router Succ!!")
}

// GetConnMgr 得到链接管理
func (s *Server) GetConnMgr() ziface.IConnManager {
	return s.ConnManager
}

// SetOnConnStart 设置该Server的连接创建时Hook函数
func (s *Server) SetOnConnStart(hookFunc func(ziface.IConnection)) {
	s.OnConnStart = hookFunc
}

// SetOnConnStop 设置该Server的连接断开时的Hook函数
func (s *Server) SetOnConnStop(hookFunc func(ziface.IConnection)) {
	s.OnConnStop = hookFunc
}

// CallOnConnStart 调用连接OnConnStart Hook函数
func (s *Server) CallOnConnStart(conn ziface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("---> CallOnConnStart....")
		s.OnConnStart(conn)
	}
}

// CallOnConnStop 调用连接OnConnStop Hook函数
func (s *Server) CallOnConnStop(conn ziface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("---> CallOnConnStop....")
		s.OnConnStop(conn)
	}
}
