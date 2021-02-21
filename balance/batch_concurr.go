package balance

import "github.com/mehdijoafshani/go-assessment-1/internal/amount"

type concurrentBatch struct {
	storageMng    StorageManager
	amountManager amount.amountManager
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

func createConcurrentBatch(storageMng StorageManager, amountManager amount.amountManager) concurrentBatch {
	return concurrentBatch{
		storageMng:    storageMng,
		amountManager: amountManager,
	}
}
