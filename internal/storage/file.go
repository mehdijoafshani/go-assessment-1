package storage

import "github.com/mehdijoafshani/go-assessment-1/internal/balance"

type fileStorage struct {
	url string
}

func (fs fileStorage) Create(content string) error {
	// TODO implement
	return nil
}

func (fs fileStorage) Read(name string) (string, error) {
	// TODO implement
	return "", nil
}

func (fs fileStorage) Update(name string, newContent string) error {
	// TODO implement
	return nil
}

func createFileStorage(url string) balance.Storage {
	return fileStorage{
		url: url,
	}
}
