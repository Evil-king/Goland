package main

import "fmt"

//抽象武器策略
type WeaponStrategy interface {
	UseWeapon() //使用武器的方法
}

//具体的策略
type Ak47 struct{}

func (a *Ak47) UseWeapon() {
	fmt.Println("使用Ak47 去战斗")
}

//具体的策略
type Knife struct{}

func (k *Knife) UseWeapon() {
	fmt.Println("使用匕首 去战斗")
}

//实施策略的环境
type Hero struct {
	strategy WeaponStrategy // 调用一个抽象的策略
}

//设置一个策略
func (h *Hero) SetWeaponStrategy(s WeaponStrategy) {
	h.strategy = s
}

func (h *Hero) Fight() {
	h.strategy.UseWeapon() //调用策略
}

func main() {
	hero := Hero{}
	//更换策略1
	hero.SetWeaponStrategy(new(Ak47))
	hero.Fight()

	hero.SetWeaponStrategy(new(Knife))
	hero.Fight()
}
