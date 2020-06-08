package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println("arr[2:6] = ", arr[2:6]) //从2开始到不包含6的位置 2，3，4，5
	fmt.Println("arr[:6] = ", arr[:6])   //从开始到不包含6之前的所有元素 0, 1, 2, 3, 4, 5

	s1 := arr[2:]
	fmt.Println("s1 = ", s1) //从2开始到最后
	s2 := arr[:]
	fmt.Println("s2 = ", s2) //全部的元素

	fmt.Println("After updateSlices(s1)")
	updateSlice(s1)  //修改切片s1的元素
	fmt.Println(s1)  //[100 3 4 5 6 7]
	fmt.Println(arr) //打印数组 [0 1 100 3 4 5 6 7]

	fmt.Println("After updateSlices(s2)")
	updateSlice(s2)  //修改切片s2的元素
	fmt.Println(s2)  //[100 1 100 3 4 5 6 7]
	fmt.Println(arr) //打印数组 [100 1 100 3 4 5 6 7]

	fmt.Println("Reslice")
	fmt.Println(s2) //[100 1 100 3 4 5 6 7]
	s2 = s2[:5]
	fmt.Println(s2) //[100 1 100 3 4]
	s2 = s2[:2]
	fmt.Println(s2) ////[100 1 ]

	fmt.Println("Extending slice")
	arr[0], arr[2] = 0, 2
	fmt.Println("arr = ", arr) //[0 1 2 3 4 5 6 7]
	s1 = arr[2:6]
	s2 = arr[3:5]
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n",
		s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n",
		s2, len(s2), cap(s2))

	s3 := append(s2,10)
	s4 := append(s3,11)
	s5 := append(s4,12)
	fmt.Println("s3, s4, s5 =", s3, s4, s5)
	// s4 and s5 no longer view arr.
	fmt.Println("arr =", arr)

	fmt.Println("Uncomment to see sliceOps demo")
	sliceOps()
}
