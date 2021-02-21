package balance

type concurrentBatch struct {
	storageMng    StorageManager
	amountManager AmountManager
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

func createConcurrentBatch(storageMng StorageManager, amountManager AmountManager) concurrentBatch {
	return concurrentBatch{
		storageMng:    storageMng,
		amountManager: amountManager,
	}
}
