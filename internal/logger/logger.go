package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	logger *zap.SugaredLogger
}

func NewLogger(level string) *Logger {
	encoderCfg := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}
	encoder := zapcore.NewJSONEncoder(encoderCfg)
	l := encodeLevel(level)
	core := zapcore.NewCore(encoder, os.Stdout, l)
	logger := zap.New(core)

	return &Logger{logger: logger.Sugar()}
}

func (l *Logger) Debug(msg string) {
	l.logger.Debug(msg)
}

func (l *Logger) Info(msg string) {
	l.logger.Info(msg)
}

func (l *Logger) Infof(msg string, args ...interface{}) {
	l.logger.Infof(msg, args)
}

func (l *Logger) Infow(msg string, w ...interface{}) {
	l.logger.Infow(msg, w...)
}

func (l *Logger) Warn(msg string) {
	l.logger.Warn(msg)
}

func (l *Logger) Error(msg string) {
	l.logger.Error(msg)
}

func (l *Logger) Errorf(msg string, args ...interface{}) {
	l.logger.Errorf(msg, args)
}

func (l *Logger) Errorw(msg string, w ...interface{}) {
	l.logger.Errorw(msg, w...)
}

func encodeLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}
