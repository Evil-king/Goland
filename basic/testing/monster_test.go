package monster

import (
	"testing"
)

//测试用例 用于测试Store方法
func TestMonster_Store(t *testing.T) {
	//先创建一个Monster实例
	monster := Monster{
		Name:  "红孩儿",
		Age:   30,
		Skill: "三味真火",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("monster.Strot() 错误，希望为=%v 实际为=%v", true, res)
	}
	t.Logf("monster.Strot() 测试成功")
}

func TestMonster_ReStore(t *testing.T) {
	//先创建一个空的Monster实例
	var monster = Monster{}
	res := monster.ReStore()
	if !res {
		t.Fatalf("monster.ReStore() 错误，希望为=%v 实际为=%v", true, res)
	}
	//进一步判断
	if monster.Name != "红孩儿" {
		t.Fatalf("monster.ReStore() 错误，希望为=%v 实际为=%v", "红孩儿", monster.Name)
	}
	t.Logf("monster.Strot() 测试成功")
}
