package main

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestSetStr(t *testing.T) {
	config = gredis.Config{
		Address: "r-wz98b1ecdaee7d74.redis.rds.aliyuncs.com:6379",
		Db:      1,
		Pass:    "sRe@2022xkawW",
	}
	gredis.SetConfig(&config)
	_, err := g.Redis().Do(context.Background(), "SET", "hwq", "boy")
	fmt.Println(err)
}

func TestSetInterfaces(t *testing.T) {
	type Students struct {
		Name string `json:"name"`
		Age  string `json:"age"`
	}
	config = gredis.Config{
		Address: "r-wz98b1ecdaee7d74.redis.rds.aliyuncs.com:6379",
		Db:      1,
		Pass:    "sRe@2022xkawW",
	}
	gredis.SetConfig(&config)
	stu := Students{
		Name: "hwq",
		Age:  "25",
	}
	_, err := g.Redis().Do(context.Background(), "SET", "stuInfo", stu)
	fmt.Println(err)
}

func TestSetKeyValueAndTimeOut(t *testing.T) {
	config = gredis.Config{
		Address: "r-wz98b1ecdaee7d74.redis.rds.aliyuncs.com:6379",
		Db:      1,
		Pass:    "sRe@2022xkawW",
	}
	gredis.SetConfig(&config)
	// SET key value [EX seconds] [PX milliseconds] [NX|XX]
	//EX seconds – 设置键key的过期时间，单位时秒
	//PX milliseconds – 设置键key的过期时间，单位时毫秒
	//NX – 只有键key不存在的时候才会设置key的值
	//XX – 只有键key存在的时候才会设置key的值
	_, err := g.Redis().Do(context.Background(), "SET", "hwq", "boy", "EX", 10)
	fmt.Println(err)
}

func TestGet(t *testing.T) {
	config = gredis.Config{
		Address: "r-wz98b1ecdaee7d74.redis.rds.aliyuncs.com:6379",
		Db:      1,
		Pass:    "sRe@2022xkawW",
	}
	gredis.SetConfig(&config)
	_, err := g.Redis().Do(context.Background(), "GET", "hwq")

	fmt.Println(err)
}

func TestGetInterfaces(t *testing.T) {
	type Students struct {
		Name string `json:"name"`
		Age  string `json:"age"`
	}
	config = gredis.Config{
		Address: "r-wz98b1ecdaee7d74.redis.rds.aliyuncs.com:6379",
		Db:      1,
		Pass:    "sRe@2022xkawW",
	}
	gredis.SetConfig(&config)
	data, err := g.Redis().Do(context.Background(), "GET", "stuInfo")
	if err = data.Struct(new(Students)); err != nil {
		panic(err)
	} else {
		fmt.Printf(`%+v`, data)
	}
}
