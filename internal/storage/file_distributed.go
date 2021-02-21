package storage

import (
	"github.com/mehdijoafshani/go-assessment-1/internal/logger"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type fileStorage struct {
	url string
}

func (fs fileStorage) createInt(id int, content int) error {
	data := strconv.Itoa(content)

	f, err := os.Create(fileName(fs.url, id))
	if err != nil {
		logger.Zap().Error("failed to create file (probably already exists)", zap.Error(err))
		// JFYI despite the fact that creation should be done only once, I will not call return in this case
		// as it's business layer's responsibility to apply these kind of rules
	}

	_, err = f.WriteString(data)
	if err != nil {
		logger.Zap().Error("failed to write into the file", zap.Error(err), zap.String("file", fileName(fs.url, id)))
	}

	err = f.Close()
	if err != nil {
		logger.Zap().Error("failed to close the file", zap.Error(err))
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
	data := strconv.Itoa(newContent)

	f, err := os.OpenFile(fileName(fs.url, id), os.O_WRONLY|os.O_TRUNC, os.ModeAppend)
	if err != nil {
		logger.Zap().Error("failed to open the file", zap.Error(err))
		return err
	}

	_, err = f.WriteString(data)
	if err != nil {
		logger.Zap().Error("failed to write into the file", zap.Error(err), zap.String("file", fileName(fs.url, id)))
		return err
	}

	err = f.Close()
	if err != nil {
		logger.Zap().Error("failed to close the file", zap.Error(err))
	}

	return nil
}

func (fs fileStorage) isDirEmpty() (bool, error) {
	dir, err := os.Open(fs.url)
	if err != nil {
		logger.Zap().Error("failed to open the dir", zap.Error(err), zap.String("dir", fs.url))
		return false, err
	}

	defer func() {
		err := dir.Close()
		if err != nil {
			logger.Zap().Error("failed to close the dir", zap.Error(err), zap.String("dir", fs.url))
		}
	}()

	dirContents, err := dir.Readdirnames(-1)
	if err != nil {
		logger.Zap().Error("failed to check the number of files within the directory", zap.Error(err), zap.String("dir", fs.url))
		return false, err
	}

	isDirEmpty := len(dirContents) == 0
	return isDirEmpty, nil
}

func (fs fileStorage) truncateDir() error {
	dir, err := os.Open(fs.url)
	if err != nil {
		logger.Zap().Error("failed to open the directory", zap.Error(err), zap.String("dir", fs.url))
		return err
	}

	defer func() {
		err := dir.Close()
		if err != nil {
			logger.Zap().Error("failed to close the directory", zap.Error(err), zap.String("dir", fs.url))
		}
	}()

	names, err := dir.Readdirnames(-1)
	if err != nil {
		logger.Zap().Error("failed to read the directory's content", zap.Error(err), zap.String("dir", fs.url))
		return err
	}
	for _, name := range names {
		filePath := filepath.Join(fs.url, name)
		err = os.RemoveAll(filePath)
		if err != nil {
			logger.Zap().Error("failed to remove the file", zap.Error(err), zap.String("file", filePath))
			return err
		}
	}

	return nil
}

func (fs fileStorage) dirFilesNumber(fileExtension string) (int, error) {
	dir, err := os.Open(fs.url)
	if err != nil {
		logger.Zap().Error("failed to open the directory", zap.Error(err), zap.String("dir", fs.url))
		return 0, err
	}

	defer func() {
		err := dir.Close()
		if err != nil {
			logger.Zap().Error("failed to close the directory", zap.Error(err), zap.String("dir", fs.url))
		}
	}()

	names, err := dir.Readdirnames(-1)
	if err != nil {
		logger.Zap().Error("failed to read the directory's content", zap.Error(err), zap.String("dir", fs.url))
		return 0, err
	}

	totalBalances := 0
	for _, name := range names {
		if strings.HasSuffix(name, fileExtension) {
			totalBalances++
		}
	}

	return totalBalances, nil
}

func createDistributedFileStorage(url string) fileStorage {
	return fileStorage{
		url: url,
	}
}
