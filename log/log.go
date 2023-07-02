package log

import (
	"log"
	"os"
)

func CreateLogger() *Logger {
	return &Logger{
		// By default, set as nil.
		telegramLogger: nil,
		nativeLogger:   log.New(os.Stdout, "tradingbot", log.Lshortfile)}
}

type Logger struct {
	telegramLogger *log.Logger
	nativeLogger   *log.Logger
}

func (l Logger) Printf(format string, v ...any) {
	l.nativeLogger.Printf(format, v...)
	if l.telegramLogger != nil {
		l.telegramLogger.Printf(format, v...)
	}
}

func (l Logger) Println(v ...any) {
	if l.telegramLogger != nil {
		// Add a new line for beaitfify.
		v = append([]any{"\n"}, v)
		std.telegramLogger.Println(v...)
	}
	l.nativeLogger.Println(v...)
}

func (l Logger) SetPrefix(prefix string) {
	if l.telegramLogger != nil {
		l.telegramLogger.SetPrefix(prefix)
	}
	l.nativeLogger.SetPrefix(prefix)
}

var std = CreateLogger()

func Default() *Logger {
	return std
}

func SetTelegramLogger(logger *log.Logger) {
	std.telegramLogger = logger
}

func Printf(format string, v ...any) {
	if std.telegramLogger != nil {
		std.telegramLogger.Printf(format, v...)
	}
	std.nativeLogger.Printf(format, v...)
}

func Println(v ...any) {
	if std.telegramLogger != nil {
		// Add a new line for beaitfify.
		v = append([]any{"\n"}, v)
		std.telegramLogger.Println(v...)
	}
	std.nativeLogger.Println(v...)
}
