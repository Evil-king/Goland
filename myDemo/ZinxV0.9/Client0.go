package main

import (
	"LearnGo/zinx/znet"
	"fmt"
	io "io"
	"net"
	"time"
)

/*
	模拟客户端
*/

func main() {
	fmt.Println("client start...")

	time.Sleep(1 * time.Second)

	// 直接链接远程服务器 得到一个conn链接
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start err,exit!")
		return
	}

	for {
		//发封包message消息
		dp := znet.NewDataPack()
		binaryMsg, err := dp.Pack(znet.NewMsgPackage(0, []byte("Zinx  Client0 Test Message")))
		//发送二进制消息给服务端
		conn.Write(binaryMsg)
		if err != nil {
			fmt.Println("write error err ", err)
			return
		}

		//先一次性读出流中的head部分
		headBuff := make([]byte, dp.GetHeadLen())
		_, err = io.ReadFull(conn, headBuff)
		if err != nil {
			fmt.Println("read head error")
			break
		}
		//将head进行拆包 得到DataLen|MsgId 的一个Message对象
		msgHead, err := dp.Unpack(headBuff)
		if err != nil {
			fmt.Println("server unpack err:", err)
			return
		}
		//再次从head中读取body部分
		if msgHead.GetDataLen() > 0 {
			//msg 是有data数据的，需要再次读取data数据
			msg := msgHead.(*znet.Message) //将msgHead转为message对象
			msg.Data = make([]byte, msg.GetDataLen())

			//根据dataLen从io中一次性读取字节流
			_, err := io.ReadFull(conn, msg.Data)
			if err != nil {
				fmt.Println("server unpack data err:", err)
				return
			}
			fmt.Println("==> Recv Msg: ID=", msg.Id, ", len=", msg.DataLen, ", data=", string(msg.Data))
		}
		//cpu阻塞
		time.Sleep(1 * time.Second)
	}

}
