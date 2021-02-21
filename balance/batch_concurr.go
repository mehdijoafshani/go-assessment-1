package balance

type concurrentBatch struct {
	storageMng StorageManager
}

func (cb concurrentBatch) create(accountsNum int) error {
	//TODO impl
	return nil
}

func (cb concurrentBatch) getAll() (int64, error) {
	//TODO impl
	return 0, nil
}

func (cb concurrentBatch) addToAll(increment int) error {
	//TODO impl
	return nil
}

func createConcurrentBatch(storageMng StorageManager) concurrentBatch {
	return concurrentBatch{
		storageMng: storageMng,
	}
}
