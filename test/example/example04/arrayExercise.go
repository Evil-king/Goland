package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//创建一个byte类型的26个元素的数组，分别放置'A'-'Z'
	//使用for循环访问所有元素并打印出来
	//思路
	//1、声明一个数组var myChars [26]byte
	//2、使用for循环，利用 字符可以进行运算的特点来赋值 'A'+1->'B'
	//3、使用for打印即可
	var myChar [26]byte
	for i := 0; i < 26; i++ {
		myChar[i] = byte('A' + i)
	}
	for i := 0; i < 26; i++ {
		fmt.Printf("%c ", myChar[i])
	}

	fmt.Println()
	//请求出一个数组的最大值，并得到对应的下标
	//思路：
	//1、声明一个数组 var intArr[5] = [...]int {1,-1,9,90,11}
	//2、假定第一个元素就是最大值，下标为0
	//3、然后从第二个元素开始循环比较，如果发现有更大的，则交换
	var intArr [5]int = [...]int{1, -1, 9, 90, 11}
	maxVal := intArr[0]
	maxIndex := 0
	for i := 1; i < len(intArr); i++ {
		if maxVal < intArr[i] {
			maxVal = intArr[i]
			maxIndex = i
		}
	}
	fmt.Printf("maxVal=%v maxIndex=%v", maxVal, maxIndex)

	fmt.Println()
	//请求出一个数组的和以及平均值，for-range
	//思路
	//先声明一个数组
	//求出和
	//求出平均值
	var intArr2 [5]int = [...]int{1, -1, 9, 90, 12}
	sum := 0
	for _, val := range intArr2 {
		sum += val
	}
	average := float64(sum) / float64(len(intArr2))
	fmt.Printf("sum=%v,average=%v", sum, average)

	fmt.Println()

	//要求：随机生成五个数，并将其反转打印
	var intArr3 [5]int
	length := len(intArr3)
	//为了每次生成的随机数不一样，我们需要给一个seed值
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<length ;i++  {
		intArr3[i] = rand.Intn(100)
	}
	fmt.Println(intArr3)
	temp := 0
	for i:=0;i<length/2 ;i++  {
		temp = intArr3[length-1 - i]
		intArr3[length-1 - i] = intArr3[i]
		intArr3[i] = temp
	}
	fmt.Println(intArr3)
}
