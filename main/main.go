package main

import (
	"github.com/mehdijoafshani/go-assessment-1/api"
	"github.com/mehdijoafshani/go-assessment-1/internal/config"
	"github.com/mehdijoafshani/go-assessment-1/internal/logger"
	"go.uber.org/zap"
)

func main() {
	// config should come before logger
	config.SetupViper()
	logger.SetupZap()

	err := api.StartRestServer()
	if err != nil {
		logger.Zap().Fatal("failed to start REST service", zap.Error(err))
	}

	// TODO: close possible opened connections
}
