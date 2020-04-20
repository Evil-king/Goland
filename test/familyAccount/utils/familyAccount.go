package utils

import "fmt"

type FamilyAccount struct {
	//声明一个变量，保存接受用户输入的选项
	key string
	//声明一个变量，控制是否退出for
	loop bool
	//定义账户的余额[]
	balance float64
	//每次收支的金额
	money float64
	//每次收支的说明
	note string
	//定一个变量，记录是否有收支记录的行为
	flag bool
	//收支的详情使用字符串记录 当有收支时，需要对details 进行拼接处理即可
	//details := "收支\t账户金额\t收支金额\t说   明"
	details string
}

//编写一个工厂模式的构造，返回一个*FamilyAccount的实例
func NewFamilyAccount() *FamilyAccount  {
	return &FamilyAccount{
		key:     "",
		loop:    false,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "收支\t账户金额\t收支金额\t说   明",
	}
}

//给该结构体绑定相应的方法
//显示主菜单
func (this *FamilyAccount) MainMenu() {
	for {
		fmt.Println("------------------------家庭收支记账软件------------------------")
		fmt.Println("                        1 收支明细")
		fmt.Println("						 2 登记收入")
		fmt.Println("						 3 登记支出")
		fmt.Println("						 4 退出软件")
		fmt.Println("请选择(1-4): ")
		fmt.Scan(&this.key)

		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.exit()
		default:
			fmt.Println("请输入正确的选项..")
		}
		if !this.loop {
			break
		}
	}
}

//将显示明细写成一个方法
func (this *FamilyAccount) showDetails()  {
	fmt.Println("当前收支明细")
	if this.flag{
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有收支明细...来一笔吧！")
	}
}

//将收入金额写成一个方法
func (this *FamilyAccount) income() {
	fmt.Println("本次收入金额：")
	fmt.Scan(&this.money)
	this.balance += this.money
	fmt.Println("本次收入说明：")
	fmt.Scan(&this.note)
	this.flag = true
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, this.money, this.note)
}

//将支出金额写成一个方法
func (this *FamilyAccount) pay()  {
	fmt.Println("本次支出金额：")
	fmt.Scan(&this.money)
	if this.money > this.balance{
		fmt.Println("余额的金额不足")
	}
	this.balance -= this.money
	fmt.Println("本次支出说明：")
	fmt.Scan(&this.note)
	this.flag = true
	this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", this.balance, this.money, this.note)
}

//将退出写成一个方法
func (this *FamilyAccount) exit()  {
	fmt.Println("你确定要退出吗？y/n")
	chioce := ""
	for {
		fmt.Scan(&chioce)
		if chioce == "y" || chioce == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入 y/n")
	}
	if chioce == "y"{
		this.loop = false
	}
}