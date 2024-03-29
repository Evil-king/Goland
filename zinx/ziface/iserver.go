package ziface

// IServer 定义一个服务器接口
type IServer interface {
	// Start 启动服务器
	Start()
	// Stop 停止服务器
	Stop()
	// Serve 运行服务器
	Serve()
	// AddRouter 路由功能，给当前的服务注册一个路由方法，供客户端的链接处理试用
	AddRouter(msgId uint32, router IRouter)
	// GetConnMgr 得到链接管理
	GetConnMgr() IConnManager
	// SetOnConnStart 设置该Server的连接创建时Hook函数
	SetOnConnStart(func(IConnection))
	// SetOnConnStop 设置该Server的连接断开时的Hook函数
	SetOnConnStop(func(IConnection))
	// CallOnConnStart 调用连接OnConnStart Hook函数
	CallOnConnStart(conn IConnection)
	// CallOnConnStop 调用连接OnConnStop Hook函数
	CallOnConnStop(conn IConnection)
}
