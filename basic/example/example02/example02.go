package main

import "fmt"

func main() {

	/**
	打印矩形
	***
	***
	***
	*/
	//for i:=1;i<=3;i++{
	//	for j:=1;j<=3;j++{
	//		fmt.Print("*")
	//	}
	//	fmt.Println()
	//}

	/**
	打印直角三角形
	*
	**
	***
	*/
	//for i:=1;i<=3;i++{
	//	for j:=1;j<=i;j++{
	//		fmt.Print("*")
	//	}
	//	fmt.Println()
	//}

	/**
		打印金字塔
	         *    第一层一个 2*层数-1  空格数=总层数-当前层数
		    ***   第二层三个 2*层数-1  空格数=总层数-当前层数
	       *****  第三层五个 2*层数-1  空格数=总层数-当前层数
	*/
	//控制行数
	for i := 1; i <= 3; i++ {
		//空格数
		for k := 1; k <= 3-i; k++ {
			fmt.Print(" ")
		}
		//控制每行打印多少个*
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
