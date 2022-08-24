package utils

import (
	"fmt"
	"github.com/gohouse/converter"
	"testing"
)

func TestMysql(t *testing.T) {
	t2t := converter.NewTable2Struct()
	t2t.Config(&converter.T2tConfig{
		//StructNameToHump: true,   // 这里配置一下即可**
	})
	err := t2t.SavePath("../usertype/billusertype/user_details_type.go").
		Dsn("root:12345678@tcp(127.0.0.1:3306)/db_test?charset=utf8").
		TagKey("orm").
		EnableJsonTag(true).
		Table("user_details").
		Run()
	fmt.Println(err)
}
