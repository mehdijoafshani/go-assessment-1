package storage

import (
	"github.com/mehdijoafshani/go-assessment-1/internal/logger"
	"go.uber.org/zap"
	"io/ioutil"
	"strconv"
)

type fileStorage struct {
	url string
}

func (fs fileStorage) createInt(id int, content int) error {
	data := []byte(strconv.Itoa(content))

	// TODO remove files permission hardcode
	err := ioutil.WriteFile(fileName(fs.url, id), data, 644)
	if err != nil {
		logger.Zap().Error("failed to create file", zap.Error(err), zap.Int("id", id))
		return err
	}

	return nil
}

func (fs fileStorage) getInt(id int) (int, error) {
	data, err := ioutil.ReadFile(fileName(fs.url, id))
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

func (fs fileStorage) updateInt(id int, newContent int) error {
	data := []byte(strconv.Itoa(newContent))

	// TODO remove files permission hardcode
	err := ioutil.WriteFile(fileName(fs.url, id), data, 0666)
	if err != nil {
		logger.Zap().Error("failed to update file", zap.Error(err), zap.Int("id", id))
		return err
	}

	return nil
}

func createDistributedFileStorage(url string) fileStorage {
	return fileStorage{
		url: url,
	}
}
