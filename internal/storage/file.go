package storage

type fileStorage struct {
	url string
}

func (fs fileStorage) create(content string) error {
	// TODO implement
	return nil
}

func (fs fileStorage) read(name string) (string, error) {
	// TODO implement
	return "", nil
}

func (fs fileStorage) update(name string, newContent string) error {
	// TODO implement
	return nil
}

func createFileStorage(url string) storage {
	return fileStorage{
		url: url,
	}
}
