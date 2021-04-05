package main

import (
	"fmt"
	"strings"
)

func wordCountFun(str string) map[string]int {
	s := strings.Fields(str)
	m := make(map[string]int)

	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			m[s[i]] ++
		} else {
			m[s[i]] = 1
		}
	}
	return m
}

func main() {
	//resp, err := http.Get("http://www.zhenai.com/zhenghun")
	////resp, err := http.Get("https://www.auluckylottery.com/results/lucky-ball-10")
	//if err !=nil{
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//
	//if resp.StatusCode != http.StatusOK {
	//	fmt.Println("Error:status code",resp.StatusCode)
	//	return
	//}
	//	all,err:=ioutil.ReadAll(resp.Body)
	//	if err !=nil{
	//		panic(err)
	//	}
	//	fmt.Printf("%s\n",all)

	str := "I Love my work I I I I Love Love my family too"
	mRet :=wordCountFun(str)

	for k,v := range mRet{
		fmt.Printf("%s:%d\n",k,v)
	}
}
