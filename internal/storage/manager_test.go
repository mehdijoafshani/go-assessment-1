package storage

import (
	"github.com/mehdijoafshani/go-assessment-1/internal/config"
	"github.com/mehdijoafshani/go-assessment-1/internal/logger"
	"path/filepath"
)

func initTestEnv() {
	config.SetupViper(filepath.Join("..", ".."))
	logger.SetupZap()
}
