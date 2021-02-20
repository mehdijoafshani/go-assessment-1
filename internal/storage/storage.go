package storage

import "github.com/mehdijoafshani/go-assessment-1/internal/config"

func Get() {
	url := config.Data.AccountsDir
	createFileStorage(url)
}
