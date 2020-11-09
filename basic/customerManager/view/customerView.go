package main

import (
	"LearnGo/basic/customerManager/model"
	"LearnGo/basic/customerManager/service"
	"fmt"
)

type customerView struct {
	key             string
	loop            bool
	customerService *service.CustomerService
}

//显示所有的客户信息
func (this *customerView) list() {
	//获取当前所有客户信息 客户信息都在切片中
	customers := this.customerService.List()
	//显示
	fmt.Println("---------------------------客户列表---------------------------")
	fmt.Println("编号\t 姓名\t 性别\t 年龄\t 电话\t 邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Printf("\n\n-------------------------客户列表完成-------------------------\n\n")
}

func (this *customerView) delete() {
	fmt.Println("---------------------------删除客户---------------------------")
	fmt.Println("请选择待删除客户编号(-1退出):")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("确认是否删除(Y/N):")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "Y" || choice == "y" {
		if this.customerService.Delete(id) {
			fmt.Println("删除完成")
		} else {
			fmt.Println("删除失败")
		}
	}
}

//添加客户的方法
func (this *customerView) add() {
	fmt.Println("---------------------添加客户---------------------")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("电邮:")
	email := ""
	fmt.Scanln(&email)
	//组装切片 枸酱一个新的customer实体
	customers := model.NewCustomerFactory2(name, gender, age, phone, email)
	//调用customerService的方法
	if this.customerService.Add(customers) {
		fmt.Println("添加完成")
	} else {
		fmt.Println("添加失败")
	}

}

//添加客户的方法
func (this *customerView) update() {
	fmt.Println("---------------------修改客户---------------------")
	fmt.Println("id:")
	id := 0
	fmt.Scanln(&id)
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("电邮:")
	email := ""
	fmt.Scanln(&email)
	//组装切片 枸酱一个新的customer实体
	customers := model.NewCustomerFactory(id, name, gender, age, phone, email)
	//调用customerService的方法
	if this.customerService.Update(customers) {
		fmt.Println("修改完成")
	} else {
		fmt.Println("修改失败")
	}

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
			this.add()
		case "2":
			this.update()
		case "3":
			this.delete()
		case "4":
			this.list()
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
