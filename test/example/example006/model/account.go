package model

import "fmt"

type account struct {
	Name string
	accountNo string
	balance float64
	pwd string
}

//编写一个工厂函数
func NewAccount(accountNo string,balance float64,pwd string) *account {
	return &account{
		accountNo:accountNo,
		balance:balance,
		pwd:pwd,
	}
}


func NewAccountName(name string) *account {
	return &account{
		Name:name,
	}
}

//提供相对应的setXxx和getXxx方法
func (a *account) SetAccountNo(no string)  {
	if len(no) > 6 && len(no) <10 {
		a.accountNo = no
	} else {
		fmt.Println("账号长度有问题，必须在6～10之间")
	}
}

func (a *account) GetAccountNo() string {
	return a.accountNo
}

func (a *account) SetBalance(b float64)  {
	if b > 20 {
		a.balance = b
	} else {
		fmt.Println("余额必须大余20")
	}
}

func (a *account) GetBalance() float64 {
	return a.balance
}

func (a *account) SetPwd(pwd string)  {
	if len(pwd) >= 6{
		a.pwd = pwd
	} else {
		fmt.Println("密码必须大余6位")
	}
}

func (a *account) GetPwd() string {
	return a.pwd
}