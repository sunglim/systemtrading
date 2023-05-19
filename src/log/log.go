package log

import (
	"bytes"
	"fmt"
	"log"
)

func CreateLogger() *Logger {
	var buf bytes.Buffer
	return &Logger{nativeLogger: log.New(&buf, "logger", log.Lshortfile)}
}

type Logger struct {
	nativeLogger *log.Logger
}

func (l Logger) Print(message string) {
	l.nativeLogger.Print(message)
}

func (l Logger) MyTest() {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "logger: ", log.Lshortfile)
	)

	logger.Print("Hello, log file!")

	fmt.Print(&buf)
}
