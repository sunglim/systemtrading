package log

import "testing"

func TestLogger(t *testing.T) {
	logger := CreateLogger()
	logger.Print("nike")
	logger.MyTest()
}
