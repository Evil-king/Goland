package znet

import (
	"LearnGo/zinx/ziface"
	"errors"
	"fmt"
	"sync"
)

type ConnManager struct {
	connections map[uint32]ziface.IConnection //管理的连接信息
	connLock    sync.Locker                   //读写连接的读写锁
}

// NewConnManager 初始化方法
func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]ziface.IConnection),
	}
}

// Add 添加链接
func (cm *ConnManager) Add(conn ziface.IConnection) {
	//保护共享资源Map 加写锁
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	//将conn连接添加到ConnManager中
	cm.connections[conn.GetConnId()] = conn

	fmt.Println("connection add to ConnManager successfully: conn num = ", cm.Len())

}

// Remove 删除连接
func (cm *ConnManager) Remove(conn ziface.IConnection) {
	//保护共享资源Map 加写锁
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	//删除连接信息
	delete(cm.connections, conn.GetConnId())

	fmt.Println("connection Remove ConnID=", conn.GetConnId(), " successfully: conn num = ", cm.Len())
}

// Get 利用ConnID获取链接
func (cm *ConnManager) Get(connID uint32) (ziface.IConnection, error) {
	//保护共享资源Map 加写锁
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	if conn, ok := cm.connections[connID]; ok {
		return conn, nil
	} else {
		return nil, errors.New("connection not found")
	}
}

// Len 获取当前连接
func (cm *ConnManager) Len() int {
	return len(cm.connections)
}

func (cm *ConnManager) ClearConn() {
	//保护共享资源Map 加写锁
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	//停止并删除全部的连接信息
	for connID, conn := range cm.connections {
		//停止
		conn.Stop()
		//删除
		delete(cm.connections, connID)
	}
	fmt.Println("Clear All Connections successfully: conn num = ", cm.Len())
}
