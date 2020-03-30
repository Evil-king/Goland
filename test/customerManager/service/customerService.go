package service

import "LearnGo/test/customerManager/model"

//该service，完成对customer的操作
type CustomerService struct {

	 //定义一个切片
	 customers []model.Customer
	 //声明一个字段，表示当前切片含有多少个客户 该字段还可以作为新客户的id
	 customerNum int
}


