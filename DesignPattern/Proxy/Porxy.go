package main

import "fmt"

//抽象主题
type BeautyWoman interface {
	//对男人抛媚眼
	MakeEyesWithMan()
	//和男人浪漫的约会
	HappyWithMan()
}

//具体主题
type PanJinLian struct{}

func (p *PanJinLian) MakeEyesWithMan() {
	fmt.Println("潘金莲对本官抛了个媚眼")
}

func (p *PanJinLian) HappyWithMan() {
	fmt.Println("潘金莲和本官共度了浪漫的约会。")
}

//代理人 王婆
type WangPo struct {
	woman BeautyWoman
}

//初始化王婆
func NewProxy(woman BeautyWoman) BeautyWoman {
	return &WangPo{woman}
}

func (p *WangPo) MakeEyesWithMan() {
	p.woman.MakeEyesWithMan()
}

func (p *WangPo) HappyWithMan() {
	p.woman.HappyWithMan()
}

//西门庆
func main() {
	//通过王婆找潘金莲
	wangpo := NewProxy(new(PanJinLian))
	wangpo.MakeEyesWithMan()
	wangpo.HappyWithMan()
}
