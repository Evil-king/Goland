package monster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

//给Monster绑定方法store 可以将一个Monster·变量序列化保存到文件
func (this *Monster) Store() bool {
	//先序列化
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println("Marshal err=", err)
		return false
	}
	//保存文件
	filePath := "d:/monster.ser"
	err = ioutil.WriteFile(filePath,data,0666)
	if err != nil{
		fmt.Println("writer file err=", err)
		return false
	}
	return true
}

//反序列化
func (this *Monster) ReStore() bool {
	//先从文件中读取序列化的字符串
	filePath := "d:/monster.ser"
	data,err:=ioutil.ReadFile(filePath)
	if err != nil{
		fmt.Println("read file err=", err)
		return false
	}
	//使用读取到的data []byte 进行反序列化
	err = json.Unmarshal(data,this)
	if err != nil{
		fmt.Println("Unmarshal  err=", err)
		return false
	}
	return true
}
