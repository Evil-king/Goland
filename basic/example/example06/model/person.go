package model

import "fmt"

type person struct {
	Name   string
	age    int
	salary float64
}

//写一个工厂模式的函数
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

//为了 访问age 和 salary 我们编写一对SetXxx和GetXxx的方法
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年龄不在范围之内")
	}
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSalary(salary float64) {
	if salary >= 3000 && salary <= 30000 {
		p.salary = salary
	} else {
		fmt.Println("薪水不在范围之内")
	}
}

func (p *person) GetSalary() float64 {
	return p.salary
}
