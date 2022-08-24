package model

import (
	"fmt"
	"goframe/pdb"
	"goframe/usertype"
	"testing"
)

func TestUserCreate(t *testing.T) {
	pdb.SetDB(&pdb.DBConfig{
		Host:     "127.0.0.1",
		Port:     "3306",
		User:     "root",
		Pass:     "12345678",
		Database: "db_test",
		Debug:    true,
	})
	err := UserCreate(nil, usertype.UserRequest{
		UserPhone: "18888888",
		UserName:  "Fox",
	})
	fmt.Println(err)
}
