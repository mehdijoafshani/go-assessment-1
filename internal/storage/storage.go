package storage

import (
	"github.com/mehdijoafshani/go-assessment-1/balance"
	"github.com/mehdijoafshani/go-assessment-1/internal/config"
)

func Get() balance.Storage {
	url := config.Data.AccountsDir
	return createFileStorage(url)
}
