package model

import (
	"LearnGo/goweb/goweb_01/utils"
	"fmt"
)

type User struct {
	ID int
	username string
	password string
	email string
}
//不带预编译
func (user *User) AddUser() error  {
	sqlStr := "insert into user(username,password,email) values(?,?,?)"
	_,err:=utils.Db.Exec(sqlStr,"admin","123456","111@163.com")
	if err != nil{
		fmt.Println("插入数据有误")
		return err
	}
	return nil
}
//带预编译
func (user *User) AddUser2() error  {
	sqlStr := "insert into user(username,password,email) values(?,?,?)"
	stmt,err:=utils.Db.Prepare(sqlStr)
	if err != nil{
		fmt.Println("预编译出现异常",err)
		return err
	}
	_,err = stmt.Exec("hwq","111111","23232@163.com")
	if err != nil {
		fmt.Println("执行出现异常:", err)
		return err }
	return nil
}

//查询一行记录
func (user *User) GetUserById() (*User,error)  {
	sqlStr := "select id,username,password,email form user where id = ?"
	row := utils.Db.QueryRow(sqlStr,user.ID)
	var username string
	var password string
	var email string
	err:=row.Scan(&username,&password,&email)
	if err != nil{
		return nil, err
	}
	//赋值给User
	users:=&User{
		username: username,
		password: password,
		email: email,
	}
	return users,err
}