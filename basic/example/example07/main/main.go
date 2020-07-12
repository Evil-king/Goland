package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//声明一个Hero结构体
type Hero struct {
	Name string
	Age  int
}

//声明一个Hero的切片类型
type HeroSlice []Hero

//现实Interface接口 返回长度
func (hs HeroSlice) Len() int {
	return len(hs)
}

//决定使用什么标准进行排序
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}

//进行交换
func (hs HeroSlice) Swap(i, j int) {
	//temp := hs[i]
	//hs[i] = hs[j]
	//hs[j] = temp
	//上面三句话可以写成下面一句话
	hs[i], hs[j] = hs[j], hs[i]
}

func main() {
	var heros HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄～%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heros = append(heros, hero)
	}

	//调用sort.Sort
	sort.Sort(heros)

	for _, v := range heros {
		fmt.Println(v)
	}
}
