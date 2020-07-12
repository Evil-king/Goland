package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {

	var allChan chan interface{}
	allChan = make(chan interface{}, 3)

	allChan <- 10
	allChan <- "jack"
	cat1 := Cat{Name: "小花猫", Age: 18}
	allChan <- cat1

	//取出
	<-allChan
	<-allChan
	newCat := <-allChan
	fmt.Printf("newCat=%T,newCat=%v", newCat, newCat)
	a := newCat.(Cat) //类型断言
	fmt.Printf("newCat.Name=%v", a.Name)
}
