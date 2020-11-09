package main

import (
	"fmt"
	"net/http"
)
//自定义多路复用器
func handler(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintln(w,"通过自己创建的多路复用器来处理请求")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/myMux",handler)

	http.ListenAndServe(":9999",mux)
}
