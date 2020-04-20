package main

import "fmt"

func main() {
	//声明一个变量，保存接受用户输入的选项
	key := ""
	//声明一个变量，控制是否退出for
	loop := true

	//定义账户的余额[]
	balance := 10000.0
	//每次收支的金额
	money := 0.0
	//每次收支的说明
	note := ""
	//定一个变量，记录是否有收支记录的行为
	flag := false
	//收支的详情使用字符串记录 当有收支时，需要对details 进行拼接处理即可
	details := "收支\t账户金额\t收支金额\t说   明"

	//显示主菜单
	for {
		fmt.Println("------------------------家庭收支记账软件------------------------")
		fmt.Println("                        1 收支明细")
		fmt.Println("						 2 登记收入")
		fmt.Println("						 3 登记支出")
		fmt.Println("						 4 退出软件")
		fmt.Println("请选择(1-4): ")
		fmt.Scan(&key)

		switch key {
		case "1":
			fmt.Println("当前收支明细")
			if flag{
				fmt.Println(details)
			} else {
				fmt.Println("当前没有收支明细...来一笔吧！")
			}
		case "2":
			fmt.Println("本次收入金额：")
			fmt.Scan(&money)
			balance += money
			fmt.Println("本次收入说明：")
			fmt.Scan(&note)
			flag = true
			details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
		case "3":
			fmt.Println("本次支出金额：")
			fmt.Scan(&money)
			if money > balance{
				fmt.Println("余额的金额不足")
				break
			}
			balance -= money
			fmt.Println("本次支出说明：")
			fmt.Scan(&note)
			flag = true
			details += fmt.Sprintf("\n支出\t%v\t%v\t%v", balance, money, note)
		case "4":
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
				loop = false
			}
		default:
			fmt.Println("请输入正确的选项..")
		}

		if !loop {
			break
		}
	}
	fmt.Println("你退出家庭记账软件的使用...")
}
