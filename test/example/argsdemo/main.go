package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("命令行的参数有", len(os.Args))
	//便利os.Aargs切片，就可以得到所有命令行输入数值
	for i, v := range os.Args {
		fmt.Printf("args[%v]=%v\n", i, v)
	}
}
