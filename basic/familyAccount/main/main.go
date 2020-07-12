package main

import (
	"LearnGo/basic/familyAccount/utils"
	"fmt"
)

func main() {
	fmt.Println("这是面向对象的方法")
	utils.NewFamilyAccount().MainMenu()
}
