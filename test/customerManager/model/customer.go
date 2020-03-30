package model

type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

//利用工厂模式返回一个Customer对象实利
func NewCustomerFactory(name string, gender string, age int, phone string, email string) *Customer {
	return &Customer{
		Id:     0,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}
