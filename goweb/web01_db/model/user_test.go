package model

import (
	"fmt"
	"testing"
)

func TestAddUser(t *testing.T) {
	fmt.Println("测试添加用户方法1")
	user := &User{}
	user.AddUser()
	user.AddUser2()
}
