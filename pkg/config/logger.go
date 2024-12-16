package config

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func SetupLogger() (*zap.Logger, error) {
	// Configuração do logger
	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel), // Nível de log
		Development: true,
		Encoding:    "console", // Formato legível (console)
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "timestamp",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder, // Níveis coloridos
			EncodeTime:     zapcore.ISO8601TimeEncoder,       // Tempo ISO8601
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"}, // Logs no console
		ErrorOutputPaths: []string{"stderr"}, // Logs de erro no console
	}
	logger, err := config.Build()
	if err != nil {
		log.Fatal("Error creating logger")
		return nil, err
	}
	defer logger.Sync()
	Logger = logger
	return logger, nil
}
