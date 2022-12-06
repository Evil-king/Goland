package main

import "fmt"

//题目一
//func deferFuncParameter(){
//	var aInt = 1
//	 defer fmt.Println(aInt)
//
//	aInt = 2
//	return
//}

//题目二
func printArray(array *[3]int) {
	for i := range array {
		fmt.Println(array[i])
	}
}

func deferFuncParameter() {
	var aArray = [3]int{1, 2, 3}

	defer printArray(&aArray)

	aArray[0] = 10
	return
}

//题目三
func deferFuncReturn() (result int) {
	i := 1

	defer func() {
		result++
	}()

	return i
}

func main() {
	fmt.Println(deferFuncReturn())
}
