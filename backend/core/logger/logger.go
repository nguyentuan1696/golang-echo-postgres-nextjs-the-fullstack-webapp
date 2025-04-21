package logger

import (
	"log/slog"
	"os"
	"path/filepath"
)

type LogLevel string

const (
	LogLevelDebug LogLevel = "DEBUG"
	LogLevelInfo  LogLevel = "INFO"
	LogLevelWarn  LogLevel = "WARN"
	LogLevelError LogLevel = "ERROR"
)

type LogConfig struct {
	Level      LogLevel
	JSONFormat bool
	FilePath   string
}

type Logger struct {
	*slog.Logger
}

var defaultLogger *Logger

func NewLogger(config LogConfig) (*Logger, error) {
	var level slog.Level
	switch config.Level {
	case LogLevelDebug:
		level = slog.LevelDebug
	case LogLevelInfo:
		level = slog.LevelInfo
	case LogLevelWarn:
		level = slog.LevelWarn
	case LogLevelError:
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level: level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{
					Key:   "timestamp",
					Value: a.Value,
				}
			}
			return a
		},
	}

	var handler slog.Handler
	writer := os.Stdout

	if config.FilePath != "" {
		// Create logs directory if it doesn't exist
		logDir := filepath.Join("logs")
		if err := os.MkdirAll(logDir, 0755); err != nil {
			return nil, err
		}

		// Open log file with full path
		logPath := filepath.Join(logDir, config.FilePath)
		file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return nil, err
		}
		writer = file
	}

	handler = slog.NewJSONHandler(writer, opts)
	logger := slog.New(handler)
	return &Logger{logger}, nil
}

// Initialize the default logger
func Init(config LogConfig) error {
	logger, err := NewLogger(config)
	if err != nil {
		return err
	}
	defaultLogger = logger
	return nil
}

// Helper methods for different log levels
func (l *Logger) Debug(msg string, args ...any) {
	l.Logger.Debug(msg, args...)
}

func (l *Logger) Info(msg string, args ...any) {
	l.Logger.Info(msg, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	l.Logger.Warn(msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	l.Logger.Error(msg, args...)
}

func (l *Logger) With(args ...any) *Logger {
	return &Logger{l.Logger.With(args...)}
}


// Global logger functions
func Error(msg string, args ...any) {
	if defaultLogger == nil {
		panic("Logger not initialized. Call Init() first")
	}
	defaultLogger.Error(msg, args...)
}

func Debug(msg string, args ...any) {
	if defaultLogger == nil {
		panic("Logger not initialized. Call Init() first")
	}
	defaultLogger.Debug(msg, args...)
}

func Info(msg string, args ...any) {
	if defaultLogger == nil {
		panic("Logger not initialized. Call Init() first")
	}
	defaultLogger.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	if defaultLogger == nil {
		panic("Logger not initialized. Call Init() first")
	}
	defaultLogger.Warn(msg, args...)
}

func With(args ...any) *Logger {
	if defaultLogger == nil {
		panic("Logger not initialized. Call Init() first")
	}
	return defaultLogger.With(args...)
}
