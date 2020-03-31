package service

import "LearnGo/test/customerManager/model"

//该service，完成对customer的操作
type CustomerService struct {

	 //定义一个切片
	 customers []model.Customer
	 //声明一个字段，表示当前切片含有多少个客户 该字段还可以作为新客户的id
	 customerNum int
}


//编写一个方法，可以返回 *CustomerService
func NewService() *CustomerService  {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	 customer := model.NewCustomerFactory(1,"张三","男",20,"112","zhangsan@163,com")
	customerService.customers = append(customerService.customers,customer)
	return customerService
}

//返回客户切片
func (this *CustomerService) List() []model.Customer {
	return this.customers
}

//添加客户
func (this *CustomerService) Add(customer model.Customer) bool {
	//自动分配id的规则，就是添加的顺序
	this.customerNum++
	customer.Id = this.customerNum
	//将创建好的customer加入到切片中
	this.customers = append(this.customers,customer)
	return true
}