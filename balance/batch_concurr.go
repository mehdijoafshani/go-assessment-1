package balance

import (
	"github.com/mehdijoafshani/go-assessment-1/internal/logger"
	"go.uber.org/zap"
)

type concurrentBatch struct {
	storageMng     StorageManager
	amountMng      AmountManager
	concurrencyMng ConcurrencyManager
	singleOpMng    singleOperationManager
}

func (cb concurrentBatch) create(accountsNum int) error {
	err := cb.concurrencyMng.ScheduleCreateBalances(accountsNum, func(ids <-chan int, errors chan<- error) {
		for id := range ids {
			err := cb.singleOpMng.create(id)
			errors <- err
		}
	})

	var truncateErr error
	if err != nil {
		logger.Zap().Error("failed to create balances", zap.Error(err))
		truncateErr = cb.storageMng.Truncate()
	}
	if truncateErr != nil {
		logger.Zap().Fatal("failed to remove not completed balances from the storage, system data is in invalid state", zap.Error(truncateErr))
	}
	if err != nil {
		return err
	}

	return nil
}

func (cb concurrentBatch) getAll() (int64, error) {
	numberOfBalances, err := cb.storageMng.NumberOfBalances()
	if err != nil {
		logger.Zap().Error("failed to get the number of balances", zap.Error(err))
		return 0, err
	}

	sum, err := cb.concurrencyMng.ScheduleReadAllBalancesSum(numberOfBalances, func(ids <-chan int, results chan<- int, errors chan<- error) {
		for id := range ids {
			balance, err := cb.singleOpMng.get(id)
			errors <- err
			results <- balance
		}
	})

	if err != nil {
		logger.Zap().Error("failed to get sum of all balances", zap.Error(err))
		return 0, err
	}

	return sum, nil
}

func (cb concurrentBatch) addToAll(increment int) error {
	numberOfBalances, err := cb.storageMng.NumberOfBalances()
	if err != nil {
		logger.Zap().Error("failed to get the number of balances", zap.Error(err))
		return err
	}

	err = cb.concurrencyMng.ScheduleUpdateBalances(numberOfBalances, func(ids <-chan int, errors chan<- error) {
		for id := range ids {
			err := cb.singleOpMng.add(id, increment)
			errors <- err
		}
	})

	if err != nil {
		logger.Zap().Error("failed to increase all balances", zap.Error(err))
		//TODO rollback changes
		return err
	}

	return nil
}

func createConcurrentBatch(storageMng StorageManager, amountMng AmountManager, concurrencyMng ConcurrencyManager, singleOpMng singleOperationManager) concurrentBatch {
	return concurrentBatch{
		storageMng:     storageMng,
		amountMng:      amountMng,
		concurrencyMng: concurrencyMng,
		singleOpMng:    singleOpMng,
	}
}
