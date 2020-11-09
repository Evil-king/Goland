package Strategy

import "fmt"

/**
策略模式
*/

//实现一个日志记录器
type LogManager struct {
	Logging
}

func NewLogManager(logManager LogManager) *LogManager {
	return &LogManager{logManager}
}

//抽象日志
type Logging interface {
	Info()
	Error()
}

//文件-具体日志实现
type FileLogging struct {
}

func (fl *FileLogging) Info() {
	fmt.Println("文件记录Info日志")
}

func (fl *FileLogging) Error() {
	fmt.Println("文件记录Error日志")
}

//数据库-具体日志实现
type DbLogging struct {
}

func (db *DbLogging) Info() {
	fmt.Println("数据库记录Info日志")
}

func (db *DbLogging) Error() {
	fmt.Println("数据库记录Error日志")
}
