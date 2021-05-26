package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		//colly.AllowedDomains("news.baidu.com"),
		colly.UserAgent("Opera/9.80 (Windows NT 6.1; U; zh-cn) Presto/2.9.168 Version/11.50"))

	c.OnHTML("div[LotteryIndex lotteryHall]", func(e *colly.HTMLElement) {
		// <div class="hotnews" alog-group="focustop-hotnews"> 下所有的a解析
		// <div class="lotteryTitle amlhc"> 下的nextTime解析
		e.ForEach("div[mainLottery lotteryTitle amlhc]", func(i int, el *colly.HTMLElement) {
			context := el.Text
			fmt.Println("context:", context)
		})
		////<div class="lotteryNum"> lotteryNum下的 lastTitle 解析
		//e.ForEach(".lotteryNum lastTitle  lotteryIssue", func(i int, element *colly.HTMLElement) {
		//	periodNum := element.Text
		//	fmt.Printf("期号 %s",periodNum)
		//})
		//
		////<div class=""> lotteryNum下的 num 解析
		//e.ForEach(".lotteryNum num span", func(i int, element *colly.HTMLElement) {
		//	outNumber := element.Text
		//	fmt.Printf("开奖号码 %s",outNumber)
		//})
	})
	c.Visit("https://www.macau-jc.com/pc/#/")

}
