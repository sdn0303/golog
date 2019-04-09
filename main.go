package golog

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger struct of logger
type Logger struct {
	Logging *zap.Logger
}

var (
	loggerInstance *Logger
	setLogger      sync.Once
)

func logger() *zap.Logger {
	level := zap.NewAtomicLevel()
	level.SetLevel(zapcore.DebugLevel)

	config := zap.Config{
		Level:    level,
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "Time",
			LevelKey:       "Level",
			NameKey:        "Name",
			CallerKey:      "Caller",
			MessageKey:     "Msg",
			StacktraceKey:  "St",
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	return logger
}

//　initializeLogger initialize Logger struct
func initializeLogger() {
	loggerInstance = &Logger{
		Logging: logger(),
	}
}

// GetInstance create & return loggerInstance
func GetInstance() *Logger {
	setLogger.Do(func() {
		initializeLogger()
	})
	return loggerInstance
}
