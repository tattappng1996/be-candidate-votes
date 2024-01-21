package logger

import (
	"context"
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	KeyLogger keyContext = iota
)

var zapLogger *zap.Logger = zap.L()

type keyContext int

func Init() *zap.Logger {
	var err error
	zapLogger, err = newZapConfig().Build()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	zap.ReplaceGlobals(zapLogger)
	defer zapLogger.Sync()

	return zapLogger
}

func Ctx(ctx context.Context) *zap.Logger {
	if ctx != nil {
		if l, ok := ctx.Value(KeyLogger).(*zap.Logger); ok {
			return l
		}
	}
	return zapLogger
}

func newZapConfig() zap.Config {
	var lv = zapcore.InfoLevel
	if envValue, ok := os.LookupEnv("LOG_LEVEL"); ok {
		log.Printf("LOG_LEVEL %s", envValue)
		if err := lv.Set(envValue); err != nil {
			log.Fatalf("Can't parse LOG_LEVEL: %s.", envValue)
		}
	}
	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(lv)
	cfg.EncoderConfig.TimeKey = "timestamp"
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder

	return cfg
}
