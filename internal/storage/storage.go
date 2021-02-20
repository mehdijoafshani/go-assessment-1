package storage

import "github.com/mehdijoafshani/go-assessment-1/internal/config"

type storage interface {
	create(content string) error
	read(name string) (string, error)
	update(name string, newContent string) error
}

func Get() {
	url := config.Data.AccountsDir
	createFileStorage(url)
}
