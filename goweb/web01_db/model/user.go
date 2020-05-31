package model

import (
	"LearnGo/goweb/web01_db/utils"
	"fmt"
)

//User 结构体
type User struct {
	ID       int
	Username string
	Password string
	Email    string
}

func (user *User) AddUser() error {
	//写sqlyuju
	sqlStr := "insert into tbl_user(username,password,email) values(?,?,?)"
	//执行
	_,err :=utils.Db.Exec(sqlStr,"hwq","111111","hwq_8910&163.com")
	if err != nil {
		fmt.Println("执行出现异常:",err)
		return err
	}
	return nil
}

func (user *User) AddUser2() error {
	//写sqlyuju
	sqlStr := "insert into tbl_user(username,password,email) values(?,?,?)"
	//预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常:",err)
		return err
	}
	//执行
	_,err2 :=inStmt.Exec("admin","123456","hwq@sina.com")
	if err2 != nil {
		fmt.Println("执行出现异常:",err2)
		return err2
	}
	return nil
}
