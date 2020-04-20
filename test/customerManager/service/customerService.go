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

//根据id查找客户在切片中对应的下标，如果没有该客户返回-1
func (this *CustomerService) FindById(id int) int  {
	index := -1
	//变量切片
	for i:=0;i< len(this.customers);i++{
		if this.customers[i].Id == id{
			//说明找到了
			index = i
		}
	}
	return index
}

//根据id删除客户(从切片中删除)
func (this *CustomerService) Delete(id int) bool  {
	index := this.FindById(id)
	//说明没有这个客户
	if index == -1{
		return false
	}
	//如何从切片中删除
	this.customers = append(this.customers[:index],this.customers[index+1:]...)
	return true
}

//根据id更新客户信息
func (this *CustomerService) Update(customer model.Customer) bool  {
	index := this.FindById(customer.Id)
	//说明没有这个客户
	if index == -1{
		return false
	}
	//遍历切片中去找到对应的customer对象 然后进行更新
	for i,_ := range this.customers{
		if i == index {
			this.customers[i] = customer
		}
	}
	return true
}


