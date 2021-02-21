package storage

import (
	"path/filepath"
	"strconv"
)

const (
	fileExtension = ".txt"
)

type file interface {
	getInt(id int) (int, error)
	createInt(id int, i int) error
	updateInt(id int, i int) error
	isDirEmpty() (bool, error)
	truncateDir() error
	dirFilesNumber(fileExtension string) (int, error)
}

func fileName(url string, id int) string {
	return filepath.Join(url, strconv.Itoa(id)+fileExtension)
}
