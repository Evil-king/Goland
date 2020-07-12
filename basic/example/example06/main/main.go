package main

import (
	"LearnGo/basic/example/example06/model"
	"fmt"
)

func main() {
	p := model.NewPerson("Fox")
	p.SetAge(10)
	p.SetSalary(float64(5000))

	fmt.Printf("员工=%v,年龄=%v,薪水=%v", p.Name, p.GetAge(), p.GetSalary())

}
