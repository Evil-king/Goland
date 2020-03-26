package main

import "fmt"

type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

func (account *Account) Deposit(money float64, pwd string) {
	//输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("输入的密码不正确")
		return
	}
	//看看存款金额是否正确
	if money <= 0 {
		fmt.Println("输入的金额不正确")
		return
	}
	account.Balance += money
	fmt.Println("存款成功")
}

func (account *Account) Withdraw(money float64, pwd string) {
	//输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("输入的密码不正确")
		return
	}
	//看看存款金额是否正确
	if money <= 0 || money > account.Balance {
		fmt.Println("输入的金额不正确")
		return
	}
	account.Balance -= money
	fmt.Println("存款成功")
}

func (account *Account) queryBalance(pwd string) {
	//输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("输入的密码不正确")
		return
	}
	fmt.Printf("你的账号为=%v,你的余额为=%v", account.AccountNo, account.Balance)
}

func main() {
	account := Account{
		AccountNo: "工商银行",
		Pwd:       "666666",
		Balance:   100000,
	}

	//account.Deposit(float64(99999),"666666")
	account.Withdraw(float64(10000),"666666")
	account.queryBalance("666666")
}
