package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

func TestDataPack(t *testing.T) {
	/*
		模拟服务器
	*/
	//1、创建stockTCP
	listen, err := net.Listen("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("server Listen err", err)
		return
	}

	//创建一个go 承载 负责从客户端处理业务
	go func() {
		for {
			//2、从客户端读取数据，拆包处理
			conn, err := listen.Accept()
			if err != nil {
				fmt.Println("server accept err", err)
			}
			go func(conn net.Conn) {
				//处理客户端的请求
				//---------> 拆包的过程 <--------
				//定义一个拆包对象dp
				dp := NewDataPack()
				for {
					//第一次从conn读 把包的head读出来
					headData := make([]byte, dp.GetHeadLen())
					_, err := io.ReadFull(conn, headData)
					if err != nil {
						fmt.Println("read head error")
						break
					}
					msgHead, err := dp.Unpack(headData)
					if err != nil {
						fmt.Println("server Unpack error", err)
						return
					}
					if msgHead.GetDataLen() > 0 {
						//第二次从conn读 根据head中的dataLen 再读取data内容
						msg := msgHead.(*Message)
						msg.Data = make([]byte, msg.GetDataLen())

						//根据datalen的长度再次从io流中读取
						_, err := io.ReadFull(conn, msg.Data)
						if err != nil {
							fmt.Println("server Unpack data error", err)
							return
						}
						fmt.Println("==> Recv Msg: ID=", msg.Id, ", len=", msg.DataLen, ", data=", string(msg.Data))
					}

				}
			}(conn)
		}
	}()

	/*
		模拟客户端
	*/
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client dial error", err)
		return
	}

	//创建一个封包对象
	dp := NewDataPack()

	//模拟粘包过程 封装两个msg一同发送
	//封装第一个msg1包
	msg1 := &Message{
		Id:      1,
		DataLen: 5,
		Data:    []byte{'h', 'e', 'l', 'l', 'o'},
	}
	sendData1, err := dp.Pack(msg1)
	if err != nil {
		return
	}
	//封装第一个msg2包
	msg2 := &Message{
		Id:      2,
		DataLen: 7,
		Data:    []byte{'w', 'o', 'r', 'l', 'd', '!', '!'},
	}
	sendData2, err := dp.Pack(msg2)
	if err != nil {
		return
	}
	//将两个包粘在一起
	sendData1 = append(sendData1, sendData2...)
	//一次性发生给服务端
	conn.Write(sendData1)

	//客户端阻塞
	select {}
}
