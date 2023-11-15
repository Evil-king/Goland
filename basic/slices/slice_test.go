package main

import (
	"fmt"
	"testing"
)

func TestOne(t *testing.T) {
	s := []int{1, 2}
	s = append(s, 3, 4, 5)
	fmt.Println(len(s), cap(s))
}
