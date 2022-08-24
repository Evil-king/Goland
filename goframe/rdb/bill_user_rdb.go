package rdb

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"goframe/usertype"
)

func InsertUser(tx *gdb.TX, data usertype.User) (sql.Result, error) {
	return tx.Model("user").Insert(data)
}

func InsertUserDetails(tx *gdb.TX, data usertype.UserDetails) error {
	_, err := tx.Model("user_details").Insert(data)
	return err
}
