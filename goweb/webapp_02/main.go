package main

import (
	"fmt"
	"net/http"
	"time"
)

type MyHandler struct {}

//自己创建的处理器
func (m *MyHandler) ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintln(w,"通过自己创建的处理器处理请求！")
}

func main() {
	myHandle := MyHandler{}

	//http.Handle("/myHandler",&myHandle)
	//
	//http.ListenAndServe(":9999",nil)

	server := http.Server{
		Addr: ":9999",
		Handler: &myHandle,
		ReadHeaderTimeout: 2*time.Second,
	}

	server.ListenAndServe()
}
