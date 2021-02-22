package account

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

const jobDone = 1

func (cb concurrentBatch) createAccounts(accountsNum int) error {
	err := cb.concurrencyMng.ScheduleCreateBalances(accountsNum, func(ids <-chan int, results chan<- int, errorCh chan<- error) {
		for id := range ids {
			err := cb.singleOpMng.createAccount(id)
			if err != nil {
				logger.Zap().Error("failed to create balance", zap.Int("id", id), zap.Error(err))
				errorCh <- err
				return
			}
			results <- jobDone
		}
	})

	if err != nil {
		return err
	}

	return nil
}

func (cb concurrentBatch) getAllBalancesSum(accountsNum int) (int64, error) {
	sum, err := cb.concurrencyMng.ScheduleReadAllBalancesSum(accountsNum, func(ids <-chan int, results chan<- int, errorCh chan<- error) {
		for id := range ids {
			balance, err := cb.singleOpMng.getBalance(id)
			if err != nil {
				logger.Zap().Error("failed to getBalance sum of balance", zap.Error(err), zap.Int("id", id))
				errorCh <- err
				return
			}
			results <- balance
		}
	})

	if err != nil {
		logger.Zap().Error("failed to getBalance sum of all balances", zap.Error(err))
		return 0, err
	}

	return sum, nil
}

func (cb concurrentBatch) addToAllBalances(numberOfAccounts int, increment int) error {
	err := cb.concurrencyMng.ScheduleUpdateBalances(numberOfAccounts, func(ids <-chan int, results chan<- int, errorCh chan<- error) {
		for id := range ids {
			err := cb.singleOpMng.addBalance(id, increment)
			if err != nil {
				logger.Zap().Error("failed to increase balance", zap.Int("id", id), zap.Error(err))
				errorCh <- err
				return
			}
			results <- jobDone
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
