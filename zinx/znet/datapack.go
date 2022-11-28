package znet

import (
	"LearnGo/zinx/utils"
	"LearnGo/zinx/ziface"
	"bytes"
	"encoding/binary"
	"errors"
)

//封包拆包类实例，暂时不需要成员
type DataPack struct{}

//封包拆包实例初始化方法
func NewDataPack() *DataPack {
	return &DataPack{}
}

//获取包头长度方法
func (d *DataPack) GetHeadLen() uint32 {
	//Id uint32(4字节) + DataLen uint32(4字节)
	return 8
}

//封包方法(压缩数据)
func (d *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	//封装一个存放bytes字节缓冲
	dataBuff := bytes.NewBuffer([]byte{})

	//写dataLen
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}

	//写msgID
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}

	//写data数据
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

//拆包方法(解压数据)
func (d *DataPack) Unpack(data []byte) (ziface.IMessage, error) {
	//创建一个输入二进制数据的ioReader
	databuff := bytes.NewReader(data)

	//只解压head的信息，得到dataLen和msgID
	msg := &Message{}

	//读dataLen
	if err := binary.Read(databuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	//读msgID
	if err := binary.Read(databuff, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}
	//判断dataLen的长度是否超出我们允许的最大包长度
	if utils.GlobalObject.MaxPacketSize > 0 && msg.DataLen > utils.GlobalObject.MaxPacketSize {
		return nil, errors.New("Too large msg data recieved")
	}
	return msg, nil
}
