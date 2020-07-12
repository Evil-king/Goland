package main

import (
	"fmt"
	"runtime"
)

//0、1、1、2、3、5、8、13、.....
//  a  b  a+b
//     a  b  a+b
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	cupNums := runtime.NumCPU()
	fmt.Println(cupNums)
	//f := fibonacci()
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
}
