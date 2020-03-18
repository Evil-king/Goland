package main

import "fmt"

/**
第10天只有一个桃子
第9天有几个桃子 = (第10天桃子数量 +1)*2
规律：第n天的桃子数量 peach(n) = (peach(n+1)+1)*2
 */
func peach(n int) int {
	if n > 10 || n < 1{
		fmt.Println("输入的天数不对")
		return 0
	}
	if n == 10{
		return 1
	} else {
		return (peach(n+1)+1)*2
	}
}

func main() {
	fmt.Println("第9天桃子的数量是=",peach(9))
}
