package storage

import (
	"github.com/mehdijoafshani/go-assessment-1/internal/config"
	"path/filepath"
	"strconv"
)

type file interface {
	getInt(id int) (int, error)
	createInt(id int, i int) error
	updateInt(id int, i int) error
	isDirEmpty() (bool, error)
	truncateDir() error
	dirFilesNumber(fileExtension string) (int, error)
}

func fileName(dir string, id int) string {
	return filepath.Join(dir, strconv.Itoa(id)+config.Data.BalanceFileExtension)
}
