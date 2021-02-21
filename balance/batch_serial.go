package balance

type serialBatch struct {
	storageMng StorageManager
}

func (sb serialBatch) create(accountsNum int) error {
	//TODO impl
	return nil
}

func (sb serialBatch) getAll() (int64, error) {
	//TODO impl
	return 0, nil
}

func (sb serialBatch) addToAll(increment int) error {
	//TODO impl
	return nil
}

func createSerialBatch(storageMng StorageManager) serialBatch {
	return serialBatch{
		storageMng: storageMng,
	}
}
