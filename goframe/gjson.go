package main

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"testing"
)

func TestJsonToStruct(t *testing.T) {
	data :=
		`
{
    "count" : 1,
    "array" : ["John", "Ming"]
}`
	type Users struct {
		Count int
		Array []string
	}

	if j, err := gjson.DecodeToJson(data); err != nil {
		panic(err)
	} else {
		users := new(Users)
		if err := j.Struct(users); err != nil {
			panic(err)
		}
		fmt.Printf(`%+v`, users)
	}
}

func TestJsonToStruct1(t *testing.T) {
	data :=
		`{
    "users" : {
        "count" : 1,
        "array" : ["John", "Ming"]
    },
	"students" : {
        "name" : "fox",
        "age" : "10"
    }
}`
	if j, err := gjson.DecodeToJson(data); err != nil {
		panic(err)
	} else {
		type Users struct {
			Count int
			Array []string
		}
		type Students struct {
			Name string
			Age  string
		}
		//users := new(Users)
		students := new(Students)
		if err := j.GetStruct("students", students); err != nil {
			panic(err)
		}
		fmt.Printf(`%+v`, students)
	}
}

func TestGoroutine(t *testing.T) {

}
