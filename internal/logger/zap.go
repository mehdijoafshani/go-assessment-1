package logger

import (
	"go.uber.org/zap"
)

var logger *zap.Logger

// No need to "sync.once", as the logger is going to be created only once
// No need to be thread safe
func init() {
	logger = zap.NewExample()
}

func Zap() *zap.Logger {
	return logger
}
