package log

import (
	"log"
	"log/slog"

	"os"
)

func CreateLogger() *Logger {
	return &Logger{
		// By default, set as nil.
		telegramLogger: nil,
		nativeLogger:   log.New(os.Stdout, "tradingbot", log.Lshortfile),
		slogLogger:     slog.New(slog.NewTextHandler(os.Stderr, nil)),
	}
}

// Logger is a main logger struct that internally includes sub-loggers.
type Logger struct {
	slogLogger     *slog.Logger
	telegramLogger *log.Logger
	nativeLogger   *log.Logger
}

// Info logs at LevelInfo.
func (l *Logger) Info(msg string, args ...any) {
	l.slogLogger.Info(msg, args...)
}

// Warn logs at LevelWarn.
func (l *Logger) Warn(msg string, args ...any) {
	l.slogLogger.Warn(msg, args...)
}

// Error logs at LevelError.
func (l *Logger) Error(msg string, args ...any) {
	l.slogLogger.Error(msg, args...)
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

func SetTelegramLoggerByToken(telegramToken string, telegramChatId int64) {
	if telegramToken == "" {
		return
	}

	telegramWriter := CreateTelegramWriter(telegramToken, telegramChatId)
	if telegramWriter != nil {
		SetTelegramLogger(log.New(telegramWriter, "", log.Ldate|log.Ltime))
	}
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
