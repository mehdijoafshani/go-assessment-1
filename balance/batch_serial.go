package balance

import (
	"github.com/mehdijoafshani/go-assessment-1/internal/logger"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type serialBatch struct {
	storageMng    StorageManager
	amountManager AmountManager
}

func (sb serialBatch) create(accountsNum int) error {
	// rule: balances should be created only once
	areBalancesCreated, err := sb.storageMng.AreBalancesCreated()
	if err != nil {
		logger.Zap().Error("failed to check if balances are created before", zap.Error(err))

	}
	if areBalancesCreated {
		logger.Zap().Error("failed to create balances, they are created before")
		return errors.New("balances are created before, it is allowed to be created only once")
	}

	for i := 0; i < accountsNum; i++ {
		id := i
		amount, err := sb.amountManager.GenerateBalanceAmount(id)
		if err != nil {
			logger.Zap().Error("failed to generate balance amount", zap.Int("id", id), zap.Error(err))
			return err
		}

		err = sb.storageMng.CreateBalance(id, amount)
		// define truncErr to avoid hierarchical code
		var truncErr error
		if err != nil {
			logger.Zap().Error("failed to create balance", zap.Int("id", id), zap.Error(err))
			truncErr = sb.storageMng.Truncate()
		}
		if truncErr != nil {
			logger.Zap().Fatal("failed to truncate the storage after failed creation. The system data is in invalid state", zap.Error(truncErr))
		}
		if err != nil {
			return err
		}
	}

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

func createSerialBatch(storageMng StorageManager, amountManager AmountManager) serialBatch {
	return serialBatch{
		storageMng:    storageMng,
		amountManager: amountManager,
	}
}
