package Strategy

import "testing"

func TestLogging(t *testing.T) {
	fileLog := &FileLogging{}
	fileLog.Info()
	fileLog.Error()

	db := &DbLogging{}
	db.Info()
	db.Error()
}
