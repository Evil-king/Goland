package main

import "fmt"

func main() {
	//使用for-range遍历map
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "天津"
	for k, v := range cities {
		fmt.Printf("k=%v,v=%v\n", k, v)
	}

	studentMap := make(map[string]map[string]string)

	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["name"] = "tom"
	studentMap["stu01"]["sex"] = "男"
	studentMap["stu01"]["address"] = "北京长安街"

	studentMap["stu02"] = make(map[string]string, 3)
	studentMap["stu02"]["name"] = "mary"
	studentMap["stu02"]["sex"] = "女"
	studentMap["stu02"]["address"] = "上海黄浦江"

	for k, v := range cities {
		fmt.Printf("k=%v,v=%v\n", k, v)
	}

}
