package log

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger  *zap.Logger
	logFile *os.File
)

func InitLogger() error {
	logDir := "logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return err
	}

	// ファイルを開く
	var err error
	logFile, err = os.OpenFile(
		filepath.Join(logDir, "app.log"),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		return err
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	logLevel := zapcore.DebugLevel

	fileCore := zapcore.NewCore(fileEncoder, zapcore.AddSync(logFile), logLevel)
	consoleCore := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), logLevel)

	core := zapcore.NewTee(fileCore, consoleCore)

	Logger = zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)

	return nil
}

func Info(msg string, fields ...zap.Field)  { Logger.Info(msg, fields...) }
func Error(msg string, fields ...zap.Field) { Logger.Error(msg, fields...) }
func Warn(msg string, fields ...zap.Field)  { Logger.Warn(msg, fields...) }
func Debug(msg string, fields ...zap.Field) { Logger.Debug(msg, fields...) }
func Fatal(msg string, fields ...zap.Field) { Logger.Fatal(msg, fields...) }

func Sync() {
	if Logger != nil {
		_ = Logger.Sync()
	}
}

func Close() {
	if logFile != nil {
		_ = logFile.Close()
	}
}
