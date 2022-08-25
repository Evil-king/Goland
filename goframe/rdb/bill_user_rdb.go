package rdb

import (
	"context"
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"goframe/pdb"
	"goframe/usertype"
)

func InsertUser(tx *gdb.TX, data usertype.User) (sql.Result, error) {
	return tx.Model("user").Insert(data)
}

func InsertUserDetails(tx *gdb.TX, data usertype.UserDetails) error {
	_, err := tx.Model("user_details").Insert(data)
	return err
}

func TransCtx(ctx context.Context, fn func(ctx context.Context, tx *gdb.TX) error) error {
	orm, err := pdb.GetDBInstance()
	if err != nil {
		panic(err)
	}
	return orm.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		return fn(ctx, tx)
	})
}

//func SelectById(id string) usertype.User {
//	orm, err := pdb.GetDBInstance()
//	if err != nil {
//		panic(err)
//	}
//	var user = new(usertype.User)
//	err := orm.Model("user").Where("id",id).Scan(&user)
//	if err != nil {
//		panic(err)
//	}
//	if err == sql.ErrNoRows {
//
//	}
//	return err
//}
