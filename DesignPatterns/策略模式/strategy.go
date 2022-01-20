package strategy

import "fmt"

// 策略模式 定义一系列算法，让这些算法在运行时可以互换，使得分离算法，符号开闭原则。

type Payment struct {
	context  *PaymentContext //支付内容属性
	strategy PaymentStrategy //支付的策略
}

func (p *Payment) Pay() {
	p.strategy.Pay(p.context) // payment的方法通过策略去调用自己的pay方法
}

type PaymentContext struct {
	Name, CardID string
	Money        int
}

type PaymentStrategy interface {
	Pay(*PaymentContext)
}

func NewPayment(name, cardId string, money int, strategy PaymentStrategy) *Payment {
	return &Payment{
		context: &PaymentContext{
			Name:   name,
			CardID: cardId,
			Money:  money,
		},
		strategy: strategy,
	}
}

type Cash struct{}

func (*Cash) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by cash", ctx.Money, ctx.Name)
}

type Bank struct{}

func (*Bank) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by bank account %s", ctx.Money, ctx.Name, ctx.CardID)

}
