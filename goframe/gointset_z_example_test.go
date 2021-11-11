package main

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/container/gset"
	"github.com/gogf/gf/frame/g"
	"testing"
)

func TestNewIntSet(t *testing.T) {
	intSet := gset.NewIntSet(true)
	intSet.Add([]int{1, 2, 3}...)
	fmt.Println(intSet.Slice())
}

func TestNewIntSetFrom(t *testing.T) {
	strSet := gset.NewIntSetFrom([]int{1, 2, 3}, true)
	fmt.Println(strSet.Slice())

	// Mya Output:

}

func TestNewIntAdd(t *testing.T) {
	intSet := gset.NewIntSetFrom([]int{1, 2, 3}, true)
	intSet.Add(1)
	fmt.Println(intSet.Slice())
	fmt.Println(intSet.AddIfNotExist(1))
}

func TestNewIntAddIfNotExist(t *testing.T) {
	intSet := gset.NewIntSetFrom([]int{1, 2, 3}, true)
	intSet.Add(1)
	fmt.Println(intSet.Slice())
	fmt.Println(intSet.AddIfNotExist(1))
}

func TestNewIntAddIfNotExistFunc(t *testing.T) {
	intSet := gset.NewIntSetFrom([]int{1, 2, 3}, true)
	intSet.Add(1)
	fmt.Println(intSet.Slice())
	fmt.Println(intSet.AddIfNotExistFunc(5, func() bool {
		return true
	}))
}

func TestNewIntAddIfNotExistFuncLock(t *testing.T) {
	strSet := gset.NewIntSetFrom([]int{1, 2, 3}, true)
	strSet.Add(1)
	fmt.Println(strSet.Slice())
	fmt.Println(strSet.AddIfNotExistFuncLock(4, func() bool {
		return true
	}))
}

func TestNewIntClear(t *testing.T) {
	intSet := gset.NewIntSetFrom([]int{1, 2, 3}, true)
	fmt.Println(intSet.Size())
	intSet.Clear()
	fmt.Println(intSet.Size())
}

func TestNewIntComplement(t *testing.T) {
	intSet := gset.NewIntSetFrom([]int{1, 2, 3, 4, 5}, true)
	s := gset.NewIntSetFrom([]int{1, 2, 3}, true)
	fmt.Println(s.Complement(intSet).Slice())
}

func TestNewIntContains(t *testing.T) {
	var set1 gset.IntSet
	set1.Add(1, 4, 5, 6, 7)
	fmt.Println(set1.Contains(1))

	var set2 gset.IntSet
	set2.Add(1, 4, 5, 6, 7)
	fmt.Println(set2.Contains(8))

}

func TestNewIntContainsI(t *testing.T) {
	var set gset.StrSet
	set.Add("a")
	fmt.Println(set.ContainsI("a"))
	fmt.Println(set.ContainsI("A"))
	fmt.Println(set.ContainsI("A"))
}

func TestNewIntDiff(t *testing.T) {
	s1 := gset.NewIntSetFrom([]int{1, 2, 3}, true)
	s2 := gset.NewIntSetFrom([]int{1, 2, 3, 4}, true)
	fmt.Println(s2.Diff(s1).Slice())
}

func TestNewIntEqual(t *testing.T) {

	s1 := gset.NewIntSetFrom([]int{1, 2, 3}, true)
	s2 := gset.NewIntSetFrom([]int{1, 2, 3, 4}, true)
	fmt.Println(s2.Equal(s1))

	s3 := gset.NewIntSetFrom([]int{1, 2, 3}, true)
	s4 := gset.NewIntSetFrom([]int{1, 2, 3}, true)
	fmt.Println(s3.Equal(s4))
}

func TestNewIntIntersect(t *testing.T) {
	s1 := gset.NewIntSet(true)
	s1.Add([]int{1, 2, 3}...)
	var s2 gset.IntSet
	s2.Add([]int{1, 2, 3, 4}...)
	fmt.Println(s2.Intersect(s1).Slice())
}

func TestNewIntIsSubsetOf(t *testing.T) {
	s1 := gset.NewIntSet(true)
	s1.Add([]int{1, 2, 3, 4}...)
	var s2 gset.IntSet
	s2.Add([]int{1, 2, 4}...)
	fmt.Println(s2.IsSubsetOf(s1))
}

func TestNewIntIterator(t *testing.T) {
	s1 := gset.NewIntSet(true)
	s1.Add([]int{1, 2, 3, 4}...)
	s1.Iterator(func(v int) bool {
		fmt.Println("Iterator", v)
		return true
	})
}

func TestNewIntJoin(t *testing.T) {
	s1 := gset.NewIntSet(true)
	s1.Add([]int{1, 2, 3, 4}...)
	fmt.Println(s1.Join(","))
}

func TestNewIntLockFunc(t *testing.T) {
	s1 := gset.NewIntSet(true)
	s1.Add([]int{1, 2}...)
	s1.LockFunc(func(m map[int]struct{}) {
		m[3] = struct{}{}
	})
	fmt.Println(s1.Slice())
}

func TestNewIntRLockFunc(t *testing.T) {
	s1 := gset.NewIntSet(true)
	s1.Add([]int{1, 2, 3, 4}...)
	s1.RLockFunc(func(m map[int]struct{}) {
		fmt.Println(m)
	})
}

func TestNewIntMarshalJSON(t *testing.T) {
	type Student struct {
		Id     int
		Name   string
		Scores *gset.IntSet
	}
	s := Student{
		Id:     1,
		Name:   "john",
		Scores: gset.NewIntSetFrom([]int{100, 99, 98}, true),
	}
	b, _ := json.Marshal(s)
	fmt.Println(string(b))
}

func TestNewIntMerge(t *testing.T) {
	s1 := gset.NewIntSet(true)
	s1.Add([]int{1, 2, 3, 4}...)

	s2 := gset.NewIntSet(true)
	fmt.Println(s1.Merge(s2).Slice())
}

func TestNewIntPop(t *testing.T) {
	s1 := gset.NewIntSet(true)
	s1.Add([]int{1, 2, 3, 4}...)

	fmt.Println(s1.Pop())
}

func TestNewIntPops(t *testing.T) {
	s1 := gset.NewIntSet(true)
	s1.Add([]int{1, 2, 3, 4}...)
	for _, v := range s1.Pops(2) {
		fmt.Println(v)
	}
}

func TestNewIntRemove(t *testing.T) {
	s1 := gset.NewIntSet(true)
	s1.Add([]int{1, 2, 3, 4}...)
	s1.Remove(1)
	fmt.Println(s1.Slice())
}

func TestNewIntSlice(t *testing.T) {
	s1 := gset.NewIntSet(true)
	s1.Add([]int{1, 2, 3, 4}...)
	fmt.Println(s1.Slice())
}

func TestNewIntString(t *testing.T) {
	s1 := gset.NewIntSet(true)
	s1.Add([]int{1, 2, 3, 4}...)
	fmt.Println(s1.String())
}

func TestNewIntSum(t *testing.T) {
	s1 := gset.NewIntSet(true)
	s1.Add([]int{1, 2, 3, 4}...)
	fmt.Println(s1.Sum())

}

func TestNewIntUnion(t *testing.T) {
	s1 := gset.NewIntSet(true)
	s1.Add([]int{1, 2, 3, 4}...)
	s2 := gset.NewIntSet(true)
	s2.Add([]int{1, 2, 4}...)
	fmt.Println(s1.Union(s2).Slice())
}

func TestNewIntUnmarshalValue(t *testing.T) {
	//s := gset.NewIntSetFrom([]string{"a"}, true)
	//s.UnmarshalValue([]string{"b", "c"})
	//fmt.Println(s.Slice())
}

func TestNewIntUnmarshalJSON(t *testing.T) {
	b := []byte(`{"Id":1,"Name":"john","Scores":[100,99,98]}`)
	type Student struct {
		Id     int
		Name   string
		Scores *gset.IntSet
	}
	s := Student{}
	json.Unmarshal(b, &s)
	fmt.Println(s)
}

func TestWalk(t *testing.T) {
	var (
		set    gset.IntSet
		names  = g.SliceInt{2, 50}
		prefix = 10
	)
	set.Add(names...)
	// Add prefix for given table names.
	set.Walk(func(item int) int {
		return prefix + item
	})
	fmt.Println(set.Slice())
}
