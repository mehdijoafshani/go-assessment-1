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
	totalBalance := int64(0)

	numberOfBalances, err := sb.storageMng.NumberOfBalances()
	if err != nil {
		logger.Zap().Error("failed to get the number of balances", zap.Error(err))
		return 0, err
	}

	for i := 0; i < numberOfBalances; i++ {
		id := i

		amount, err := sb.storageMng.GetBalance(id)
		if err != nil {
			logger.Zap().Error("failed to get balance", zap.Int("id", id), zap.Error(err))
			return 0, err
		}

		totalBalance += int64(amount)
	}

	return totalBalance, nil
}

func (sb serialBatch) addToAll(increment int) error {
	numberOfBalances, err := sb.storageMng.NumberOfBalances()
	if err != nil {
		logger.Zap().Error("failed to get the number of balances", zap.Error(err))
		return err
	}

	for i := 0; i < numberOfBalances; i++ {
		id := i

		err := sb.storageMng.IncreaseBalance(id, increment)
		if err != nil {
			logger.Zap().Error("failed to update balance", zap.Int("id", id), zap.Error(err))
			return err
		}
	}

	return nil
}

func createSerialBatch(storageMng StorageManager, amountManager AmountManager) serialBatch {
	return serialBatch{
		storageMng:    storageMng,
		amountManager: amountManager,
	}
}
