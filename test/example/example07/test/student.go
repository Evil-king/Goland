package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//声明一个学生结构体
type Student struct {
	Name  string
	Age   int
	Score float64
}

//定一个学生类型的切片
type student []Student

func (stu student) Len() int {
	return len(stu)
}

func (stu student) Less(i, j int) bool {
	return stu[i].Score > stu[j].Score
}

func (stu student) Swap(i, j int) {
	stu[i], stu[j] = stu[j], stu[i]
}

func main() {
	var stu student
	for i := 0; i < 10; i++ {
		stus := Student{
			Name:  fmt.Sprintf("学生～%d", rand.Intn(100)),
			Age:   rand.Intn(100),
			Score: rand.Float64(),
		}
		stu = append(stu, stus)
	}

	//调用sort.Sort
	sort.Sort(stu)

	for _, v := range stu {
		fmt.Println(v)
	}
}
