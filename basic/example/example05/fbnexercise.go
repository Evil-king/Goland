package main

import "fmt"

func fbn(n int) []uint64 {
	//声明一个切片，切片大小 n
	fbnSlice := make([]uint64, n)
	//第1个和第2个数的斐波那契 为1
	fbnSlice[0] = 1
	fbnSlice[1] = 1
	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	return fbnSlice
}

func main() {
	//测试一把看看是否好用
	fbnSlice := fbn(20)
	fmt.Println("fbnSlice=", fbnSlice)
}
