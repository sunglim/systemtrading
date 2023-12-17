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
		slogLogger:     slog.New(slog.NewTextHandler(os.Stderr, nil)),
	}
}

// Logger is a main logger struct that internally includes sub-loggers.
type Logger struct {
	slogLogger     *slog.Logger
	telegramLogger *log.Logger
}

// Info logs at LevelInfo.
func (l *Logger) Info(msg string, args ...any) {
	l.slogLogger.Info(msg, args...)
}

// Warn logs at LevelWarn.
func (l *Logger) Warn(msg string, args ...any) {
	l.slogLogger.Warn(msg, args...)
	if l.telegramLogger != nil {
		l.telegramLogger.Println(msg)
	}
}

// Error logs at LevelError.
func (l *Logger) Error(msg string, args ...any) {
	l.slogLogger.Error(msg, args...)
	if l.telegramLogger != nil {
		l.telegramLogger.Println(msg)
	}
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
