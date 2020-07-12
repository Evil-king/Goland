package main

import "fmt"

func main() {

	fmt.Println("请输入年龄:")
	//fmt.Scanln(&age)
	//goland支持在if中，直接定义一个变量，比如下面

	if age := 20; age > 18 {
		fmt.Println("你的年龄大于18，要对自己的行为负责")
	}
	//if age > 18 {
	//	fmt.Println("你的年龄大于18，要对自己的行为负责")
	//}

}
