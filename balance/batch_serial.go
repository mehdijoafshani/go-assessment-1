package balance

import (
	"github.com/mehdijoafshani/go-assessment-1/internal/logger"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type serialBatch struct {
	storageMng    StorageManager
	amountManager amountManager
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
		amount, err := sb.amountManager.generateBalance(id)
		if err != nil {
			logger.Zap().Error("failed to generate balance amount", zap.Int("id", id), zap.Error(err))
			return err
		}

		err = sb.storageMng.CreateBalance(id, amount)
		if err != nil {
			logger.Zap().Error("failed to create balance", zap.Int("id", id), zap.Error(err))
			// TODO: truncate storage
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

func createSerialBatch(storageMng StorageManager, amountManager amountManager) serialBatch {
	return serialBatch{
		storageMng:    storageMng,
		amountManager: amountManager,
	}
}
