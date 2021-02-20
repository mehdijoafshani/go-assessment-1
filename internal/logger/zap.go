package logger

import (
	"github.com/mehdijoafshani/go-assessment-1/internal/config"
	"go.uber.org/zap"
	"log"
)

var logger *zap.Logger

// No need to "sync.once", as the logger is going to be created only once
// No need to be thread safe
func init() {
	var cfg zap.Config
	if config.Data.Production {
		cfg = zap.NewProductionConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
	}

	cfg.OutputPaths = []string{
		config.Data.LogsDir,
	}

	//explicitly define var err, to avoid confusion in the next line
	var err error
	logger, err = cfg.Build()
	if err != nil {
		log.Panic("failed to build zap logger", err)
	}
}

func Zap() *zap.Logger {
	return logger
}
