package main

import "fmt"

//医生-命令接受者
type Doctor struct{}

func (d *Doctor) treatEye() {
	fmt.Println("医生治疗眼睛")
}

func (d *Doctor) treatNose() {
	fmt.Println("医生治疗鼻子")
}

//抽象的命令
type Command interface {
	Treat()
}

//治疗眼睛的病单
type CommandTreatEye struct {
	doctor *Doctor
}

func (cmd *CommandTreatEye) Treat() {
	cmd.doctor.treatEye()
}

//治疗鼻子的病单
type CommandTreatNose struct {
	doctor *Doctor
}

func (cmd *CommandTreatNose) Treat() {
	cmd.doctor.treatNose()
}

func main() {
	//创建医生
	doctor := new(Doctor)
	//看眼睛
	cmdEye := CommandTreatEye{doctor}
	cmdEye.Treat()
	//看鼻子
	cmdNose := CommandTreatNose{doctor}
	cmdNose.Treat()
}
