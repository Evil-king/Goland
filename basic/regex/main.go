package main

import (
	"fmt"
	"regexp"
)

//传入[]byte，返回[]byte
func findTest() {
	str := "ab001234hah120210a880218end"
	reg := regexp.MustCompile("\\d{6}") //六位连续的数字
	fmt.Println("------Find------")
	//返回str中第一个匹配reg的字符串
	data := reg.Find([]byte(str))
	fmt.Println(string(data))

	fmt.Println("------FindAll------")
	//返回str中所有匹配reg的字符串
	//第二个参数表示最多返回的个数，传-1表示返回所有结果
	dataSlice := reg.FindAll([]byte(str), -1)
	for _, v := range dataSlice {
		fmt.Println(string(v))
	}
}

//传入[]byte，返回首末位置索引
func findIndexTest() {
	fmt.Println("------FindIndex------")
	//返回第一个匹配的字符串的首末位置
	reg2 := regexp.MustCompile("start\\d*end") //start开始，end结束，中间全是数字
	str2 := "00start123endhahastart120PSend09start10000end"
	//index[0]表示开始位置，index[1]表示结束位置
	index := reg2.FindIndex([]byte(str2))
	fmt.Println("start:", index[0], ",end:", index[1], str2[index[0]:index[1]])

	fmt.Println("------FindAllIndex------")
	//返回所有匹配的字符串首末位置
	indexSlice := reg2.FindAllIndex([]byte(str2), -1)
	for _, v := range indexSlice {
		fmt.Println("start:", v[0], ",end:", v[1], str2[v[0]:v[1]])
	}
}

//传入string，返回string（更加方便）
func findStringTest() {
	fmt.Println("------FindString------")

	str := "ab001234hah120210a880218end"
	reg := regexp.MustCompile("\\d{6}") //六位连续的数字
	fmt.Println(reg.FindString(str))
	fmt.Println(reg.FindAllString(str, -1))
	//以下两个方法是类似的
	fmt.Println(reg.FindStringIndex(str))
	fmt.Println(reg.FindIndex([]byte(str)))
}

//查找汉字
func findChinesString() {
	str := "hello中国hello世界和平hi好"
	reg := regexp.MustCompile("[\\p{Han}]+")
	fmt.Println(reg.FindAllString(str, -1))
	//[中国 世界和平 好]
}

//查找数字或小写字母
func findNumOrLowerLetter() {
	str := "HAHA00azBAPabc09FGabHY99"
	reg := regexp.MustCompile("[\\d|a-z]+")
	fmt.Println(reg.FindAllString(str, -1))
	//[00az abc09 ab 99]
}

//查找并替换
func findAndReplace() {
	str := "Welcome for Beijing-Tianjin CRH train."
	reg := regexp.MustCompile(" ")
	fmt.Println(reg.ReplaceAllString(str, "@")) //将空格替换为@字符

	//Welcome@for@Beijing-Tianjin@CRH@train.
}

func main() {
	text := `Hello 世界！123 Go.`
	reg := regexp.MustCompile(`[a-z]+`)             // 查找连续的小写字母
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // 输出结果["ello" "o"]

	reg = regexp.MustCompile(`[^a-z]+`)             // 查找连续的非小写字母
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["H" " 世界！123 G" "."]

	reg = regexp.MustCompile(`[\w]+`)               // 查找连续的单词字母
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["Hello" "123" "Go"]

	reg = regexp.MustCompile(`[^\w\s]+`)            // 查找连续的非单词字母、非空白字符
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["世界！" "."]

	reg = regexp.MustCompile(`[[:upper:]]+`)        // 查找连续的大写字母
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["H" "G"]

	reg = regexp.MustCompile(`[[:^ascii:]]+`)       // 查找连续的非 ASCII 字符
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["世界！"]

	reg = regexp.MustCompile(`[\pP]+`)              // 查找连续的标点符号
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["！" "."]

	reg = regexp.MustCompile(`[\PP]+`)              // 查找连续的非标点符号字符
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["Hello 世界" "123 Go"]

	reg = regexp.MustCompile(`[\p{Han}]+`)          // 查找连续的汉字
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["世界"]

	reg = regexp.MustCompile(`[\P{Han}]+`)          // 查找连续的非汉字字符
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["Hello " "！123 Go."]

	reg = regexp.MustCompile(`Hello|Go`)            // 查找 Hello 或 Go
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["Hello" "Go"]

	reg = regexp.MustCompile(`^H.*\s`)              // 查找行首以 H 开头，以空格结尾的字符串
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["Hello 世界！123 "]

	reg = regexp.MustCompile(`(?U)^H.*\s`)          // 查找行首以 H 开头，以空白结尾的字符串（非贪婪模式）
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["Hello "]

	reg = regexp.MustCompile(`(?i:^hello).*Go`)     // 查找以 hello 开头（忽略大小写），以 Go 结尾的字符串
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["Hello 世界！123 Go"]

	reg = regexp.MustCompile(`\QGo.\E`)             // 查找 Go.
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["Go."]

	reg = regexp.MustCompile(`(?U)^.* `)            // 查找从行首开始，以空格结尾的字符串（非贪婪模式）
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["Hello "]

	reg = regexp.MustCompile(` [^ ]*$`)             // 查找以空格开头，到行尾结束，中间不包含空格字符串
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // [" Go."]

	reg = regexp.MustCompile(`(?U)\b.+\b`)          // 查找“单词边界”之间的字符串
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["Hello" " 世界！" "123" " " "Go"]

	reg = regexp.MustCompile(`[^ ]{1,4}o`)          // 查找连续 1 次到 4 次的非空格字符，并以 o 结尾的字符串
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["Hello" "Go"]

	reg = regexp.MustCompile(`(?:Hell|G)o`)         // 查找 Hello 或 Go
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["Hello" "Go"]

	//reg = regexp.MustCompile(`(?PHell|G)o`)                   // 查找 Hello 或 Go，替换为 Hellooo、Gooo
	//fmt.Printf("%q\n", reg.ReplaceAllString(text, "${n}ooo")) // "Hellooo 世界！123 Gooo."

	reg = regexp.MustCompile(`(Hello)(.*)(Go)`)              // 交换 Hello 和 Go
	fmt.Printf("%q\n", reg.ReplaceAllString(text, "$3$2$1")) // "Go 世界！123 Hello."

	//reg = regexp.MustCompile(`[\f\t\n\r\v\123\x7F\x{10FFFF}\\\^\$\.\*\+\?\{\}\(\)\[\]\|]`)
	//fmt.Printf("%q\n", reg.ReplaceAllString("\f\t\n\r\v\123\x7F\U0010FFFF\\^$.*+?{}()[]|", "-"))
	// "----------------------"

}
