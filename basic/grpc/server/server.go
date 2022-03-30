package main

import (
	"LearnGo/basic/grpc"
	"log"
	"net"
	"net/rpc"
)

func main() {
	rpc.RegisterName("HelloService", new(grpc.HelloService))

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}
	rpc.ServeConn(conn)
}
