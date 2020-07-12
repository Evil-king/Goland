package main

import "fmt"

/**
Goland中的继承体现
*/

type Student struct {
	Name  string
	Age   int
	Score int
}

//显示他的成绩
func (s *Student) ShowInfo() {
	fmt.Printf("学生的姓名=%v 年龄=%v 成绩=%v", s.Name, s.Age, s.Score)
}

//设置Score
func (s *Student) SetScore(score int) {
	s.Score = score
}

//小学生
type Pupil struct {
	Student
}

//大学生
type Collage struct {
	Student
}

//打印xx学生正在考试
func (p *Pupil) testing() {
	fmt.Println("小学生正在考试.........")
}

func (c *Collage) testing() {
	fmt.Println("大学生正在考试.........")
}

func main() {
	pupil := &Pupil{}
	pupil.Student.Name = "Fox"
	pupil.Student.Age = 10
	pupil.Student.Score = 100
	pupil.ShowInfo()
	pupil.testing()

	collage := &Collage{}
	collage.Student.Name = "Jim"
	collage.Student.Age = 24
	collage.Student.Score = 100
	collage.ShowInfo()
	collage.testing()
}
