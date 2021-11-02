package gostrset

import (
	"fmt"
	"github.com/gogf/gf/container/gset"
	"testing"
)

func TestNewStrSet(t *testing.T) {
	var strSet gset.StrSet
	strSet.Add([]string{"str1", "str2", "str3"}...)

	strSet.Iterator(func(v string) bool {
		fmt.Println("Iterator ", v)
		return true
	})
}

func TestAddIfNotExist(t *testing.T) {
	var strSet gset.StrSet
	fmt.Println(strSet.AddIfNotExist("str"))
}

func TestAddIfNotExistFunc(t *testing.T) {
	var strSet gset.StrSet
	fmt.Println(strSet.AddIfNotExistFunc("str", func() bool {
		return true
	}))
}

func TestAddIfNotExistFuncLock(t *testing.T) {
	var strSet gset.StrSet
	fmt.Println(strSet.AddIfNotExistFuncLock("str", func() bool {
		return true
	}))
}

func TestClear(t *testing.T) {
	var strSet gset.StrSet
	strSet.Add([]string{"str1", "str2", "str3"}...)
	strSet.Clear()

	fmt.Println(strSet.Size())
}

func TestComplement(t *testing.T) {
	strSet := gset.NewStrSet(true)
	strSet.Add([]string{"str1", "str2", "str3", "str4", "str5"}...)
	s := gset.NewStrSet(true)
	s.Add([]string{"str1", "str2", "str3"}...)

	fmt.Println(s.Complement(strSet).Slice())
}

func TestContains(t *testing.T) {
	var set gset.StrSet
	set.Add("a")
	fmt.Println(set.Contains("a"))
	fmt.Println(set.Contains("A"))
	fmt.Println(set.ContainsI("A"))
}

func TestContainsI(t *testing.T) {
	var set gset.StrSet
	set.Add("a")
	fmt.Println(set.ContainsI("a"))
	fmt.Println(set.ContainsI("A"))
	fmt.Println(set.ContainsI("A"))
}

func TestDiff(t *testing.T) {
	s1 := gset.NewStrSet(true)
	s1.Add([]string{"a", "b", "c"}...)
	var s2 gset.StrSet
	s2.Add([]string{"a", "b", "c", "d"}...)
	// 差集
	fmt.Println(s2.Diff(s1).Slice())
}

func TestEqual(t *testing.T) {
	s1 := gset.NewStrSet(true)
	s1.Add([]string{"a", "b", "c"}...)
	var s2 gset.StrSet
	s2.Add([]string{"a", "b", "c", "d"}...)
	fmt.Println(s2.Equal(s1))
}

func TestIntersect(t *testing.T) {
	s1 := gset.NewStrSet(true)
	s1.Add([]string{"a", "b", "c"}...)
	var s2 gset.StrSet
	s2.Add([]string{"a", "b", "c", "d"}...)
	// 交集
	fmt.Println(s2.Intersect(s1).Slice())
}

func TestIsSubsetOf(t *testing.T) {
	s1 := gset.NewStrSet(true)
	s1.Add([]string{"a", "b", "c", "d"}...)
	var s2 gset.StrSet
	s2.Add([]string{"a", "b", "d"}...)
	fmt.Println(s2.IsSubsetOf(s1))
}

func TestIterator(t *testing.T) {
	s1 := gset.NewStrSet(true)
	s1.Add([]string{"a", "b", "c", "d"}...)
	s1.Iterator(func(v string) bool {
		fmt.Println("Iterator", v)
		return true
	})
}

func TestJoin(t *testing.T) {
	s1 := gset.NewStrSet(true)
	s1.Add([]string{"a", "b", "c", "d"}...)
	fmt.Println(s1.Join(","))
}

func TestLockFunc(t *testing.T) {
	//s1 := gset.NewStrSet(true)
	//s1.Add([]string{"a", "b", "c", "d"}...)
	//s1.LockFunc(func(m map[string]struct{}) {
	//	fmt.Println(m["a"])
	//})
}
