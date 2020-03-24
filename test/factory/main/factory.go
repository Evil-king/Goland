package main

import "LearnGo/test/factory/model"
import "fmt"

func main() {

	var stu = model.NewStudent("tom~",90)
	fmt.Println(stu.Score)
	fmt.Println(stu.Name)
}
