package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"BitoPro_interview_question/config"
)

func NewZapLogger(config *config.Config) (*zap.SugaredLogger, error) {

	var logger *zap.Logger
	var cfg zap.Config
	var err error
	if strings.ToLower(config.Logger.Environment) == "dev" {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewProductionConfig()
	}

	cfg.DisableStacktrace = true
	cfg.DisableCaller = true
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	cfg.Encoding = "json"
	cfg.EncoderConfig.TimeKey = "log_time"
	cfg.EncoderConfig.LevelKey = "log_level"
	cfg.EncoderConfig.MessageKey = "log_msg"
	atom := zap.NewAtomicLevel()
	logger = zap.New(zapcore.NewCore(zapcore.NewJSONEncoder(cfg.EncoderConfig), zapcore.Lock(os.Stdout), atom))
	logger, err = cfg.Build()

	if err != nil {
		panic(err)
	}

	defer logger.Sync()

	return logger.Sugar(), nil
}
