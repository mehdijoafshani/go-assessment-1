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

func (cb concurrentBatch) createBalances(accountsNum int) error {
	err := cb.concurrencyMng.ScheduleCreateBalances(accountsNum, func(ids <-chan int, errors chan<- error) {
		for id := range ids {
			err := cb.singleOpMng.createBalance(id)
			if err != nil {
				logger.Zap().Error("failed to create balance", zap.Int("id", id), zap.Error(err))
			}

			errors <- err
		}
	})

	if err != nil {
		return err
	}

	return nil
}

func (cb concurrentBatch) getAllBalancesSum(numberOfBalances int) (int64, error) {
	sum, err := cb.concurrencyMng.ScheduleReadAllBalancesSum(numberOfBalances, func(ids <-chan int, results chan<- int, errors chan<- error) {
		for id := range ids {
			balance, err := cb.singleOpMng.getBalance(id)
			errors <- err
			results <- balance
		}
	})

	if err != nil {
		logger.Zap().Error("failed to getBalance sum of all balances", zap.Error(err))
		return 0, err
	}

	return sum, nil
}

func (cb concurrentBatch) addToAllBalances(numberOfBalances int, increment int) error {
	err := cb.concurrencyMng.ScheduleUpdateBalances(numberOfBalances, func(ids <-chan int, errors chan<- error) {
		for id := range ids {
			err := cb.singleOpMng.addBalance(id, increment)
			logger.Zap().Error("failed to increase balance", zap.Int("id", id), zap.Error(err))
			errors <- err
		}
	})

	if err != nil {
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
