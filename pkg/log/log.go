package log

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggingMode string

const (
	Debug   = "debug"
	Release = "release"
)

func NewLoggingMode(ms string) (LoggingMode, bool) {
	switch ms {
	case Debug:
		return Debug, true
	case Release:
		return Release, true
	}
	return "", false
}

func NewLogger(mode LoggingMode, path string) *zap.Logger {
	switch mode {
	case Debug:
		f := zapcore.AddSync(&lumberjack.Logger{
			Filename:   path,
			MaxSize:    100,
			MaxBackups: 3,
			MaxAge:     14,
		})
		msync := zapcore.NewMultiWriteSyncer(f, os.Stdout)

		conf := zap.NewDevelopmentConfig()
		core := zapcore.NewCore(
			zapcore.NewConsoleEncoder(conf.EncoderConfig),
			msync,
			zap.DebugLevel,
		)
		return zap.New(core)
	default:
		f := zapcore.AddSync(&lumberjack.Logger{
			Filename:   path,
			MaxSize:    100,
			MaxBackups: 3,
			MaxAge:     14,
		})
		msync := zapcore.NewMultiWriteSyncer(f, os.Stdout)

		conf := zap.NewProductionConfig()
		core := zapcore.NewCore(
			zapcore.NewConsoleEncoder(conf.EncoderConfig),
			msync,
			zap.ErrorLevel,
		)
		return zap.New(core)
	}
}
