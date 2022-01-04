package main

import (
	"fmt"
	"time"
)

func main() {
	//go count(5,"♈")
	//count(5,"牛")

	//strSlice := make([]string, 0)
	//
	//A(&strSlice)
	//B(&strSlice)
	//
	//for _, value := range strSlice {
	//	fmt.Println(value)
	//}
}

func A(s *[]string) {
	*s = append(*s, "1111")
}

func B(s *[]string) {
	*s = append(*s, "2222")
}

func count(n int, animal string) {
	for i := 0; i < n; i++ {
		fmt.Println(i+1, animal)
		time.Sleep(time.Microsecond * 500)
	}
}
