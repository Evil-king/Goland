package model

//这里是因为Student 首字母大写 别的包可以引用的到
type student struct {
	Name  string
	Score float64
}

func NewStudent(n string, s float64) *student {
	return &student{
		Name:  n,
		Score: s,
	}
}
