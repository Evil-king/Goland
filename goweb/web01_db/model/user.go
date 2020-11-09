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

//查询语句 直接查询
func (user *User) AddUser2() error {
	//写sqlyuju
	sqlStr := "insert into tbl_user(username,password,email) values(?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, "hwq", "111111", "hwq_8910&163.com")
	if err != nil {
		fmt.Println("执行出现异常:", err)
		return err
	}
	return nil
}

//查询语句 预编译
func (user *User) AddUser() error {
	//写sqlyuju
	sqlStr := "insert into tbl_user(username,password,email) values(?,?,?)"
	//预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常:", err)
		return err
	}
	//执行
	_, err2 := inStmt.Exec("admin", "123456", "hwq@sina.com")
	if err2 != nil {
		fmt.Println("执行出现异常:", err2)
		return err2
	}
	return nil
}

//通过入参user 查询一条记录
func (user *User) GetUserById() (*User, error) {
	sqlStr := "select id,username,password,email from user where id=?"
	row := utils.Db.QueryRow(sqlStr,user.ID)
	var id int
	var username string
	var password string
	var email string
	err := row.Scan(&id,&username,&password,&email)
	if err != nil{
		return nil,err
	}
	u := &User{
		ID: id,
		Username: username,
		Password: password,
		Email: email,
	}
	return u,err
}

//查询所有数据 并且返回一个切片
func (user *User) GetUsers() ([]*User,error)  {
	sqlStr := "select id,username,password,email from user"
	rows,err := utils.Db.Query(sqlStr)
	if err != nil{
		return nil, err
	}
	//创建User切片
	var users []*User
	for rows.Next(){
		var id int
		var username string
		var password string
		var email string
		err := rows.Scan(&id,&username,&password,&email)
		if err != nil{
			return nil,err
		}
		u := &User{
			ID: id,
			Username: username,
			Password: password,
			Email: email,
		}
		users = append(users,u)
	}
	return users,nil
}