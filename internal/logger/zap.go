package logger

import (
	"github.com/mehdijoafshani/go-assessment-1/internal/config"
	"go.uber.org/zap"
	"log"
	"sync"
)

var zapLogger *zap.Logger

// to make sure zap would be setup only once
var setupOnce sync.Once

func SetupZap() {
	setupOnce.Do(func() {
		var cfg zap.Config
		if config.Data.IsProduction {
			cfg = zap.NewProductionConfig()
		} else {
			cfg = zap.NewDevelopmentConfig()
		}

		cfg.OutputPaths = []string{
			config.Data.LogsFile,
		}

		//explicitly define var err, to avoid confusion in the next line
		var err error
		zapLogger, err = cfg.Build()
		if err != nil {
			log.Panic("failed to build zap logger", err)
		}
	})
}

func Zap() *zap.Logger {
	return zapLogger
}
