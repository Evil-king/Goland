package main

import (
	"LearnGo/test/customerManager/service"
	"fmt"
)

type customerView struct {
	key             string
	loop            bool
	customerService *service.CustomerService
}

//显示所有的客户信息
func (this *customerView) List() {
	//获取当前所有客户信息 客户信息都在切片中
	customers := this.customerService.List()
	//显示
	fmt.Println("---------------------------客户列表---------------------------")
	fmt.Println("编号\t 姓名\t 性别\t 年龄\t 电话\t 邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Printf("\n-------------------------客户列表完成-------------------------\n\n")
}

//显示主菜单
func (this *customerView) MainMenu() {
	for {
		fmt.Println("-----------------客户信息管理软件-----------------")
		fmt.Println("						1添 加客 户")
		fmt.Println("						2修 改客 户")
		fmt.Println("						3删 除客 户")
		fmt.Println("						4客 户列 表")
		fmt.Println("						5退     出")
		fmt.Print("请选择(1-5):")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			fmt.Println("添 加 客 户")
		case "2":
			fmt.Println("修 改 客 户")
		case "3":
			fmt.Println("删 除 客 户")
		case "4":
			this.List()
		case "5":
			this.loop = false
		default:
			fmt.Println("你的输入有误，请重新输入...")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("你退出了客户关系管理系统...")
}

func main() {
	//在main函数中，创建一个customerView,并运行显示主菜单...
	customerView := customerView{
		key:  "",
		loop: true,
	}
	//初始化 customerService字段
	customerView.customerService = service.NewService()
	//显示主菜单
	customerView.MainMenu()
}
