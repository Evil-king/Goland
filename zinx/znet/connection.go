package znet

import (
	"LearnGo/zinx/ziface"
	"errors"
	"fmt"
	"io"
	"net"
)

/*
	链接模块
*/

type Connection struct {
	//当前链接的socket TCP套接字
	Conn *net.TCPConn

	//链接的ID
	ConnId uint32

	//当前的链接状态
	isClosed bool

	//告知当前链接已经退出的/停止 channel
	ExitChan chan bool

	//消息管理MsgId和对应处理方法的消息管理模块
	MsgHandler ziface.IMsgHandler
}

// NewConnection 初始化链接模块的方法
func NewConnection(conn *net.TCPConn, connId uint32, msgHandler ziface.IMsgHandler) *Connection {
	return &Connection{
		Conn:       conn,
		ConnId:     connId,
		MsgHandler: msgHandler,
		isClosed:   false,
		ExitChan:   make(chan bool, 1),
	}
}

// StartReader 连接的读业务方法
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Println("connId=", c.ConnId, "Reader is exit,remote addr is ", c.RemoteAddr().String())
	defer c.Stop()

	for {
		//读取客户端的数据到buf中
		//buf := make([]byte, utils.GlobalObject.MaxPacketSize)
		//_, err := c.Conn.Read(buf)
		//if err != nil {
		//	fmt.Println("recv buf err", err)
		//	continue
		//}

		//创建拆包对象
		dp := NewDataPack()

		//读取客户端的Msg head
		headBuff := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(c.GetTCPConnection(), headBuff)
		if err != nil {
			fmt.Println("read msg head error ", err)
			break
		}

		//拆包 得到msgID和dataLen 放在msg中
		msg, err := dp.Unpack(headBuff)
		if err != nil {
			fmt.Println("unpack error ", err)
			break
		}

		//根据dataLen 读取 data
		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())
			_, err := io.ReadFull(c.GetTCPConnection(), data)
			if err != nil {
				fmt.Println("read msg data error ", err)
				break
			}
		}

		//放在msg.data中
		msg.SetData(data)

		//得到当前conn数据的Request请求数据
		req := Request{
			conn:    c,
			message: msg,
		}

		//创建消息管理对象
		//根据绑定好的MsgId 找到对应处理的api业务
		go c.MsgHandler.DoMsgHandler(&req)

	}
}

// Start 启动链接 让当前的链接准备开始工作
func (c *Connection) Start() {
	fmt.Println("Conn Start().... ConnID=", c.ConnId)
	//启动从当前链接的读数据的业务
	go c.StartReader()
	//TODO 启动从当前链接写数据的业务

}

// Stop 停止链接 结束当前链接的工作
func (c *Connection) Stop() {
	fmt.Println("Conn Stop()....ConnId = ", c.ConnId)
	if c.isClosed {
		return
	}
	c.isClosed = true

	//关闭stock链接
	c.Conn.Close()

	//回收资源
	close(c.ExitChan)
}

// GetTCPConnection 获取当前链接绑定的socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// GetConnId 获取当前链接模块的链接ID
func (c *Connection) GetConnId() uint32 {
	return c.ConnId
}

// RemoteAddr 获取远程客户端的 TCP专题 IP port
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// SendMsg 发送数据 将数据发送给远程的客户端
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed == true {
		return errors.New("Connection closed when send msg")
	}
	//将data封包
	dp := NewDataPack()

	binaryMsg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		fmt.Println("Pack error msg id = ", msgId)
		return errors.New("Pack error msg ")
	}
	//写回给客户端
	if _, err := c.Conn.Write(binaryMsg); err != nil {
		fmt.Println("Write msg id ", msgId, " error ")
		return errors.New("conn Write error")
	}
	return nil
}
