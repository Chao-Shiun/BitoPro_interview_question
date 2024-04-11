package logger

import (
	"errors"
	"log"

	"go.uber.org/zap"

	"BitoPro_interview_question/config"
)

type (
	LogInfo interface {
		Debug(args ...interface{})
		Info(args ...interface{})
		Warn(args ...interface{})
		Error(args ...interface{})
		Panic(args ...interface{})
		Fatal(args ...interface{})
	}

	LogFormat interface {
		Debugf(template string, args ...interface{})
		Infof(template string, args ...interface{})
		Warnf(template string, args ...interface{})
		Errorf(template string, args ...interface{})
		Panicf(template string, args ...interface{})
		Fatalf(template string, args ...interface{})
	}

	LogInfoFormat interface {
		LogInfo
		LogFormat
	}
)

type Logger struct {
	zapSugarLogger *zap.SugaredLogger
}

func NewLogger(c *config.Config) (LogInfoFormat, error) {
	if c.Logger.Use == "zapLogger" {
		z, er := NewZapLogger(c)
		if er != nil {
			log.Fatalf("can't initialize zap logger: %v", er)
			return nil, er
		}
		return &Logger{zapSugarLogger: z}, nil

	}
	return nil, errors.New("logger not supported : " + c.Logger.Use)
}

func (l *Logger) Debug(args ...interface{}) {
	l.zapSugarLogger.Debug(args)
}

func (l *Logger) Info(args ...interface{}) {
	l.zapSugarLogger.Info(args)
}

func (l *Logger) Warn(args ...interface{}) {
	l.zapSugarLogger.Warn(args)
}

func (l *Logger) Error(args ...interface{}) {
	l.zapSugarLogger.Error(args)
}

func (l *Logger) Panic(args ...interface{}) {
	l.zapSugarLogger.Panic(args)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.zapSugarLogger.Fatal(args)
}

func (l *Logger) Debugf(template string, args ...interface{}) {
	l.zapSugarLogger.Debugw(template, args)
}

func (l *Logger) Infof(template string, args ...interface{}) {
	l.zapSugarLogger.Infow(template, args...)
}

func (l *Logger) Warnf(template string, args ...interface{}) {
	l.zapSugarLogger.Warnw(template, args)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	l.zapSugarLogger.Errorw(template, args)
}

func (l *Logger) Panicf(template string, args ...interface{}) {
	l.zapSugarLogger.Panicw(template, args)
}

func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.zapSugarLogger.Fatalw(template, args)
}
