package storage

import (
	"github.com/mehdijoafshani/go-assessment-1/internal/config"
	"strconv"
)

const (
	fileExtension = ".txt"
)

type file interface {
	getInt(id int) (int, error)
	createInt(id int, i int) error
	updateInt(id int, i int) error
}

func fileName(id int) string {
	return config.Data.AccountsDir + strconv.Itoa(id) + fileExtension
}
