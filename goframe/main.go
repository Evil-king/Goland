package goframe

import (
	"context"
	"fmt"
	redisV8 "github.com/go-redis/redis/v8"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

var (
	config = gredis.Config{
		Address: "r-wz98b1ecdaee7d74.redis.rds.aliyuncs.com:6379",
		Db:      1,
		Pass:    "sRe@2022xkawW",
	}

	redis = redisV8.NewClient(&redisV8.Options{
		Addr:     "r-wz98b1ecdaee7d74.redis.rds.aliyuncs.com:6379",
		Password: "sRe@2022xkawW", // no password set
		DB:       1,               // use default DB
	})
)

type CloudAccount struct {
	AccountId        int    `json:"account_id" orm:"account_id"`
	AccountCloudId   string `json:"account_cloud_id" orm:"account_cloud_id"`
	AccountCloudName string `json:"account_cloud_name" orm:"account_cloud_name"`
	CloudType        string `json:"cloud_type" orm:"cloud_type"`
	Enabled          int    `json:"enabled" orm:"enabled"`
	Deleted          int    `json:"deleted" orm:"deleted"`
	CreatedBy        string `json:"created_by" orm:"created_by"`
	CreatedTime      string `json:"created_time" orm:"created_time"`
	UpdatedBy        string `json:"updated_by" orm:"updated_by"`
	UpdatedTime      string `json:"updated_time" orm:"updated_time"`
	TenantID         int64  `json:"tenant_id" orm:"tenant_id"`
	Owner            string `json:"owner" orm:"owner"`
}

func main() {
	//go count(5, "♈")
	//count(5, "牛")
	//
	//strSlice := make([]string, 0)
	//
	//A(&strSlice)
	//B(&strSlice)
	//
	//for _, value := range strSlice {
	//	fmt.Println(value)
	//}

	gredis.SetConfig(&config)

	//set
	//_, err := g.Redis().Do(context.Background(), "SET", "ccc", "fox", "EX", 60, "NX")

	//setex
	_, err := g.Redis().Do(context.Background(), "SETEX", "ccc", 10, "nima")

	fmt.Println(err)
	//g.Redis().Do(context.Background(), "EXPIRE", "hwq", 10)

	//result, err := g.Redis().Do(context.Background(),"GET","90")
	//accountCloudInfo := new(CloudAccount)
	//if err = result.Struct(accountCloudInfo); err != nil {
	//	panic(err)
	//}
	//fmt.Println(accountCloudInfo.AccountId)
}

func A(s *[]string) {
	*s = append(*s, "1111")
}

func B(s *[]string) {
	*s = append(*s, "2222")
}

func count(n int, animal string) {
	for i := 0; i < n; i++ {
		fmt.Println(i+1, animal)
		time.Sleep(time.Microsecond * 500)
	}
}
