package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = "" // if you want stack trace to be logger, then remove this line.
	config.EncoderConfig = encoderConfig
	log, err = config.Build(zap.AddCallerSkip(1))

	//log, err = zap.NewProduction(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}
}

//Info level logger
func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

//Debug level logger
func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

//Error level logger
func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}
