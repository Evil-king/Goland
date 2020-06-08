package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n",
		s, len(s), cap(s))
}

func sliceOps() {
	fmt.Println("Creating Slice")
	var s []int // Zero value for slice is nil

	for i := 0; i < 100; i++ {
		//printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 16)
	s3 := make([]int, 10, 32)
	printSlice(s2) //[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0], len=16, cap=16
	printSlice(s3) //[0 0 0 0 0 0 0 0 0 0], len=10, cap=32

	fmt.Println("Copying slice")
	copy(s2, s1)   //这里的copy是将s1复制到s2后面
	printSlice(s2) //[2 4 6 8 0 0 0 0 0 0 0 0 0 0 0 0], len=16, cap=16

	fmt.Println("Deleting elements from slice") //go中对于切片没有删除的操作 只能通过append函数来操作
	//比如要删除s2中的8这个元素 就需要先取到s2[:3]=[2,4,6] 再取到s2[4:] = [0 0 0 0 0 0 0 0 0 0 0 0] 就相当于把第三个元素跳掉
	//ss := s2[:3]
	//sf := s2[4:]
	//printSlice(ss)
	//printSlice(sf)
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)


	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]

	fmt.Println(front)
	printSlice(s2)

	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]

	fmt.Println(tail)
	printSlice(s2)
}
