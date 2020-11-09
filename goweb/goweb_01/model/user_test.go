package model

import (
	"testing"
)

func TestUser_AddUser(t *testing.T) {
	user :=&User{}
	//user.AddUser()
	user.AddUser2()
}

func TestUser_GetUserById(t *testing.T) {
	user:=&User{
		ID: 1,
	}
	u,_ :=user.GetUserById()
}