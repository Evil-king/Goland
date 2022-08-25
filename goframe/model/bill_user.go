package model

import (
	"context"
	"github.com/gogf/gf/database/gdb"
	"goframe/rdb"
	"goframe/usertype"
)

func UserCreate(ctx context.Context, data usertype.UserRequest) error {
	user := usertype.User{}
	user.Name = data.UserName
	rdb.TransCtx(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 添加user
		dbResult, err := rdb.InsertUser(tx, user)
		if err != nil {
			return err
		}
		userId, _ := dbResult.LastInsertId()
		panic(0 / 1)
		// 添加userDetails
		userDetail := usertype.UserDetails{}
		userDetail.UserId = userId
		userDetail.UserName = data.UserName
		userDetail.UserPhone = data.UserPhone
		userDetail.UserAddress = "广州市天河区"
		err = rdb.InsertUserDetails(tx, userDetail)
		if err != nil {
			return err
		}
		return err
	})

	return nil
}
