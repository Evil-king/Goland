package main

import (
	"fmt"
	"testing"
)

func TestMyDemo(t *testing.T) {
	arr := []int{1, 2, 3, 2, 4, 3, 5, 6, 5}

	// 统计元素出现的次数
	countMap := make(map[int]int)
	for _, v := range arr {
		countMap[v]++
	}

	// 创建新的切片存储不重复的元素
	newArr := make([]int, 0)
	for _, v := range arr {
		if countMap[v] == 1 {
			newArr = append(newArr, v)
		}
	}

	// 将重复的数据添加到新切片末尾
	for _, v := range arr {
		if countMap[v] > 1 {
			newArr = append(newArr, v)
		}
	}
	for _, v := range newArr {
		fmt.Printf("%d", v)
	}
}

func TwoSum(nums []int, target int) []int {

	//for i := 0; i < len(nums); i++ {
	//	for j := i + 1; j < len(nums); j++ {
	//		if nums[i]+nums[j] == target {
	//			return []int{i, j}
	//		}
	//	}
	//}
	//return []int{}

	hashTable := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if j, ok := hashTable[target-nums[i]]; ok {
			return []int{j, i}
		}
		hashTable[nums[i]] = i
	}
	return []int{}
}
func TestName(t *testing.T) {
	//result := TwoSum([]int{2, 7, 11, 15}, 17)
	//g.Dump(result)
}
