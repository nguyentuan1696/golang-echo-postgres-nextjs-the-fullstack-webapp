package logger

import (
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"sync"
	"thichlab-backend-docs/constant"
)

var lock = &sync.Mutex{}
var mLog *log.Logger

func NewLogger(logPath string, logPrefix string) *log.Logger {
	if mLog == nil {
		lock.Lock()
		defer lock.Unlock()
		if mLog == nil {
			mLog = &log.Logger{}
		}
	}

	// Set config in debug mode
	if viper.GetBool("Debug") {
		mLog = log.New()
		mLog.Formatter = &log.TextFormatter{TimestampFormat: constant.TimeFormatDefault}
	} else {
		mLog = log.New()
		mLog.Out = io.Discard
		mLog.Hooks.Add(lfshook.NewHook(
			lfshook.PathMap{
				log.InfoLevel:  logPath + "/" + logPrefix + "_info.log",
				log.TraceLevel: logPath + "/" + logPrefix + "_trace.log",
				log.WarnLevel:  logPath + "/" + logPrefix + "_warn.log",
				log.DebugLevel: logPath + "/" + logPrefix + "_debug.log",
				log.ErrorLevel: logPath + "/" + logPrefix + "_error.log",
				log.FatalLevel: logPath + "/" + logPrefix + "_fatal.log",
				log.PanicLevel: logPath + "/" + logPrefix + "_panic.log",
			},
			&log.JSONFormatter{
				TimestampFormat: constant.TimeFormatDefault,
			},
		))
	}

	return mLog
}

// Debug logs a debug message.
func Debug(format string, v ...any) {
	mLog.Debugf(constant.LogDebugPrefix+format, v...)
}

// Info logs an info message.
func Info(format string, v ...any) {
	mLog.Infof(constant.LogInfoPrefix+format, v...)
}

// Warn logs a warning message.
func Warn(format string, v ...any) {
	mLog.Warnf(constant.LogWarnPrefix+format, v...)
}

// Error logs an error message.
func Error(format string, v ...any) {
	mLog.Errorf(constant.LogErrorPrefix+format, v...)
}

// Fatal logs a fatal message and exits the program.
func Fatal(format string, v ...any) {
	mLog.Fatalf(constant.LogFatalPrefix+format, v...)
}

// Panic logs a panic message and panics.
func Panic(format string, v ...any) {
	mLog.Panicf(constant.LogPanicPrefix+format, v...)
}
