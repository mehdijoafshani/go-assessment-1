package storage

import (
	"github.com/mehdijoafshani/go-assessment-1/balance"
	"github.com/mehdijoafshani/go-assessment-1/internal/config"
	"github.com/mehdijoafshani/go-assessment-1/internal/logger"
	"go.uber.org/zap"
	"io/ioutil"
	"strconv"
)

const (
	fileExtension = ".txt"
)

type fileStorage struct {
	url string
}

func (fs fileStorage) Create(id int, content string) error {
	// TODO implement
	return nil
}

func (fs fileStorage) Read(id int) (int, error) {
	fileName := config.Data.AccountsDir + strconv.Itoa(id) + fileExtension

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		logger.Zap().Error("file not found", zap.Error(err))
		return 0, err
	}

	balance, err := strconv.Atoi(string(data))
	if err != nil {
		logger.Zap().Error("file content was not numeric", zap.Error(err))
		return 0, err
	}

	return balance, nil
}

func (fs fileStorage) Update(id int, newContent int) error {
	// TODO implement
	return nil
}

func createFileStorage(url string) balance.Storage {
	return fileStorage{
		url: url,
	}
}
