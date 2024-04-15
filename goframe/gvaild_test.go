package goframe

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestValidStruct(t *testing.T) {
	type BizReq struct {
		ID     uint   `v:"required" dc:"Your ID"`
		Name   string `v:"required #名字不能为空"`
		Gender uint   `v:"in:0,1,2" dc:"0:Secret;1:Male;2:Female"`
		//WifeName    string `v:"required-if:gender,1"`
		//HusbandName string `v:"required-if:gender,2"`
	}
	var (
		ctx = context.Background()
		req = BizReq{
			ID:     1,
			Name:   "",
			Gender: 1,
		}
	)
	if err := g.Validator().Data(req).Run(ctx); err != nil {
		g.Log().Error(ctx, err)
	}
}
