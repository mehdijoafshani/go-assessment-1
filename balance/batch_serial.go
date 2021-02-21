package balance

import (
	"github.com/mehdijoafshani/go-assessment-1/internal/logger"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type serialBatch struct {
	storageMng      StorageManager
	amountManager   AmountManager
	singleOperation singleOperationManager
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
		err = sb.singleOperation.create(id)
		if err != nil {
			logger.Zap().Error("failed to create balance", zap.Int("id", id), zap.Error(err))
			break
		}
	}

	// define truncErr to avoid hierarchical code
	var truncErr error
	if err != nil {
		truncErr = sb.storageMng.Truncate()
	}
	if truncErr != nil {
		logger.Zap().Fatal("failed to truncate the storage after failed creation. The system data is in invalid state", zap.Error(truncErr))
	}
	if err != nil {
		return err
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

		balance, err := sb.singleOperation.get(id)
		if err != nil {
			logger.Zap().Error("failed to get balance", zap.Int("id", id), zap.Error(err))
			return 0, err
		}

		totalBalance += int64(balance)
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

		err := sb.singleOperation.add(id, increment)
		if err != nil {
			logger.Zap().Error("failed to update balance", zap.Int("id", id), zap.Error(err))
			// TODO rollback changes
			return err
		}
	}

	return nil
}

func createSerialBatch(storageMng StorageManager, amountManager AmountManager, singleOperation singleOperationManager) serialBatch {
	return serialBatch{
		storageMng:      storageMng,
		amountManager:   amountManager,
		singleOperation: singleOperation,
	}
}
