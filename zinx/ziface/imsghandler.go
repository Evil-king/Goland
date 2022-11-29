package ziface

/*
	消息管理的抽象层
*/

type IMsgHandler interface {
	DoMsgHandler(request IRequest)
	AddRouter(msgId uint32, router IRouter)
}
