package main

import (
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
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
		if err := j.Scan(users); err != nil {
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
		type Students struct {
			Name string `json:"name"`
			Age  string `json:"age"`
		}
		//users := new(Users)
		students := new(Students)
		if err := j.Get("students", students); err != nil {
			panic(err)
		} else {
			fmt.Printf(`%+v`, students)
		}
	}
}

func TestGoroutine(t *testing.T) {

}
