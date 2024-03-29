package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	var reply string
	err = client.Call("HelloService.Hello", "Hwq", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
