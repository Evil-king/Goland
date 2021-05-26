package factorymethod

import "testing"

func computer(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}

func TestOperator(t *testing.T) {
	var (
		factory OperatorFactory
	)

	factory = PlusOperatorFactory{}
	if computer(factory, 1, 2) != 3 {
		t.Fatal("error with factory method pattern")
	}
	factory = MinusOperatorFactory{}
	if computer(factory, 4, 2) != 2 {
		t.Fatal("error with factory method pattern")
	}
}
