package log

import (
	"log/slog"

	"os"
)

func CreateLogger() *Logger {
	return &Logger{
		slogTelegramLogger: nil,
		// telegram handler
		slogLogger: slog.New(slog.NewTextHandler(os.Stderr, nil)),
	}
}

// Logger is a main logger struct that internally includes sub-loggers.
type Logger struct {
	slogLogger         *slog.Logger
	slogTelegramLogger *slog.Logger
}

// Info logs at LevelInfo.
func (l *Logger) Info(msg string, args ...any) {
	l.slogLogger.Info(msg, args...)
}

// Warn logs at LevelWarn.
func (l *Logger) Warn(msg string, args ...any) {
	l.slogLogger.Warn(msg, args...)

	if l.slogTelegramLogger != nil {
		l.slogTelegramLogger.Warn(msg, args...)
	}
}

// Error logs at LevelError.
func (l *Logger) Error(msg string, args ...any) {
	l.slogLogger.Error(msg, args...)
	if l.slogTelegramLogger != nil {
		l.slogTelegramLogger.Error(msg, args...)
	}
}

func (l *Logger) With(args ...any) *Logger {
	sloger := l.slogLogger.With(args...)
	if l.slogTelegramLogger != nil {
		return &Logger{slogLogger: sloger, slogTelegramLogger: l.slogTelegramLogger.With(args...)}
	}
	return &Logger{slogLogger: sloger, slogTelegramLogger: nil}
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
		std.slogTelegramLogger = slog.New(slog.NewTextHandler(telegramWriter, nil))
	}
}
