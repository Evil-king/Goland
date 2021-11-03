package gostrset

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/container/gset"
	"strconv"
	"testing"
	"time"
)

func TestNewStrSet(t *testing.T) {
	var strSet gset.StrSet
	strSet.Add([]string{"str1", "str2", "str3"}...)

	strSet.Iterator(func(v string) bool {
		fmt.Println("Iterator ", v)
		return true
	})
}

func TestNewStrSetFrom(t *testing.T) {
	strSet := gset.NewStrSetFrom([]string{"str1", "str2", "str3"}, true)
	fmt.Println(strSet.Slice())

	// Mya Output:

}

func TestAdd(t *testing.T) {
	strSet := gset.NewStrSetFrom([]string{"str1", "str2", "str3"}, true)
	strSet.Add("str")
	fmt.Println(strSet.Slice())
	fmt.Println(strSet.AddIfNotExist("str"))
}

func TestAddIfNotExist(t *testing.T) {
	var strSet gset.StrSet
	fmt.Println(strSet.AddIfNotExist("str"))
}

func TestAddIfNotExistFunc(t *testing.T) {
	strSet := gset.NewStrSetFrom([]string{"str1", "str2", "str3"}, true)
	strSet.Add("str")
	fmt.Println(strSet.Slice())
	fmt.Println(strSet.AddIfNotExistFunc("str5", func() bool {
		return true
	}))
}

func TestAddIfNotExistFuncLock(t *testing.T) {
	strSet := gset.NewStrSetFrom([]string{"str1", "str2", "str3"}, true)
	strSet.Add("str")
	fmt.Println(strSet.Slice())
	fmt.Println(strSet.AddIfNotExistFuncLock("str4", func() bool {
		return true
	}))
}

func TestClear(t *testing.T) {
	var strSet gset.StrSet
	strSet.Add([]string{"str1", "str2", "str3"}...)
	fmt.Println(strSet.Size())
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

	s1 := gset.NewStrSetFrom([]string{"a", "b", "c"}, true)
	s2 := gset.NewStrSetFrom([]string{"a", "b", "c", "d"}, true)
	fmt.Println(s2.Equal(s1))

	s3 := gset.NewStrSetFrom([]string{"a", "b", "c"}, true)
	s4 := gset.NewStrSetFrom([]string{"a", "b", "c"}, true)
	fmt.Println(s3.Equal(s4))
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
	s1 := gset.NewStrSet(true)
	num := 1
	for i := 0; i < 5; i++ {
		go func() {
			s1.LockFunc(func(m map[string]struct{}) {
				m[strconv.Itoa(num)] = struct{}{}
				num++
			})
		}()
		go func() {
			s1.LockFunc(func(m map[string]struct{}) {
				m[strconv.Itoa(num)] = struct{}{}
				num++
			})
		}()
	}

	//time.Sleep(time.Duration(1) * time.Second)
	//go func() {
	//	for i := 0; i < 5; i++ {
	//		s1.LockFunc(func(m map[string]struct{}) {
	//			m[strconv.Itoa(num)] = struct{}{}
	//			num ++
	//		})
	//	}
	//}()

	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println(s1.Slice())

}

func TestRLockFunc(t *testing.T) {
	s1 := gset.NewStrSet(true)
	s1.Add([]string{"a", "b", "c", "d"}...)
	s1.RLockFunc(func(m map[string]struct{}) {
		fmt.Println(m)
	})
}

func TestMarshalJSON(t *testing.T) {
	type Student struct {
		Id     int
		Name   string
		Scores *gset.IntSet
	}
	s := Student{
		Id:     1,
		Name:   "john",
		Scores: gset.NewIntSetFrom([]int{100, 99, 98}),
	}
	b, _ := json.Marshal(s)
	fmt.Println(string(b))
}

func TestMerge(t *testing.T) {
	s1 := gset.NewStrSet(true)
	s1.Add([]string{"a", "b", "c", "d"}...)

	s2 := gset.NewStrSet(true)
	fmt.Println(s1.Merge(s2).Slice())
}

func TestPop(t *testing.T) {
	s1 := gset.NewStrSet(true)
	s1.Add([]string{"a", "b", "c", "d"}...)

	fmt.Println(s1.Pop())
}

func TestPops(t *testing.T) {
	s1 := gset.NewStrSet(true)
	s1.Add([]string{"a", "b", "c", "d"}...)
	for _, v := range s1.Pops(2) {
		fmt.Println(v)
	}
}

func TestRemove(t *testing.T) {
	s1 := gset.NewStrSet(true)
	s1.Add([]string{"a", "b", "c", "d"}...)
	s1.Remove("a")
	fmt.Println(s1.Slice())
}

func TestSlice(t *testing.T) {
	s1 := gset.NewStrSet(true)
	s1.Add([]string{"a", "b", "c", "d"}...)
	fmt.Println(s1.Slice())
}

func TestString(t *testing.T) {
	s1 := gset.NewStrSet(true)
	s1.Add([]string{"a", "b", "c", "d"}...)
	fmt.Println(s1.String())
}

func TestSum(t *testing.T) {
	s1 := gset.NewStrSet(true)
	s1.Add([]string{"a", "b", "c", "d"}...)
	fmt.Println(s1.Sum())
}

func TestUnion(t *testing.T) {
	s1 := gset.NewStrSet(true)
	s1.Add([]string{"a", "b", "c", "d"}...)
	s2 := gset.NewStrSet(true)
	s2.Add([]string{"a", "b", "c", "d"}...)
	fmt.Println(s1.Union(s2).Slice())
}

func TestUnmarshalValue(t *testing.T) {
	s := gset.NewStrSetFrom([]string{"a"}, true)
	s.UnmarshalValue([]string{"b", "c"})
	fmt.Println(s.Slice())
}

func TestUnmarshalJSON(t *testing.T) {
	b := []byte(`{"Id":1,"Name":"john","Scores":["100","99","98"]}`)
	type Student struct {
		Id     int
		Name   string
		Scores *gset.StrSet
	}
	s := Student{}
	json.Unmarshal(b, &s)
	fmt.Println(s)
}
