package znet

import "LearnGo/zinx/ziface"

type Request struct {
	//已经和客户端建立好的链接
	conn ziface.IConnection
	//客户端请求的数据
	message ziface.IMessage
}

//得到当前链接
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

//得到请求的消息数据
func (r *Request) GetData() []byte {
	return r.message.GetData()
}

//得到请求的消息数据ID
func (r *Request) GetMsgId() uint32 {
	return r.message.GetMsgId()
}
