package concurrency

import (
	"github.com/mehdijoafshani/go-assessment-1/internal/config"
	"github.com/mehdijoafshani/go-assessment-1/internal/logger"
	"path/filepath"
)

func init() {
	config.SetupTestViper(filepath.Join("..", ".."))
	logger.SetupZap()
}
