package main

import (
	"LearnGo/basic/example/example006/model"
	"fmt"
)

func main() {

	//account := model.NewAccount("1213123",40,"1111111")
	account1 := model.NewAccountName("胡文卿")
	account1.SetAccountNo("1213123")
	account1.SetBalance(10)
	account1.SetPwd("111111")
	fmt.Printf("姓名=%v,账号=%v,余额=%v,密码=%v", account1.Name, account1.GetAccountNo(), account1.GetBalance(), account1.GetPwd())

}
