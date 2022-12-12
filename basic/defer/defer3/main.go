package main

//知识点4: 有名函数返回值遇见defer情况
//
//	在没有defer的情况下，其实函数的返回就是与return一致的，但是有了defer就不一样了。
//
//  我们通过知识点2得知，先return，再defer，所以在执行完return之后，还要再执行defer里的语句，依然可以修改本应该返回的结果。

import "fmt"

func returnButDefer() (t int) { //t初始化0， 并且作用域为该函数全域

	defer func() {
		t = t * 10
	}()

	return 1
}

func main() {
	fmt.Println(returnButDefer())
}
