package balance

import (
	"github.com/mehdijoafshani/go-assessment-1/internal/logger"
	"go.uber.org/zap"
)

type serialBatch struct {
	storageMng      StorageManager
	amountManager   AmountManager
	singleOperation singleOperationManager
}

func (sb serialBatch) createBalances(accountsNum int) error {
	for i := 0; i < accountsNum; i++ {
		id := i
		err := sb.singleOperation.createBalance(id)
		if err != nil {
			logger.Zap().Error("failed to create balance", zap.Int("id", id), zap.Error(err))
			return err
		}
	}

	return nil
}

func (sb serialBatch) getAllBalancesSum(numberOfBalances int) (int64, error) {
	totalBalance := int64(0)

	for i := 0; i < numberOfBalances; i++ {
		id := i

		balance, err := sb.singleOperation.getBalance(id)
		if err != nil {
			logger.Zap().Error("failed to getBalance balance", zap.Int("id", id), zap.Error(err))
			return 0, err
		}

		totalBalance += int64(balance)
	}

	return totalBalance, nil
}

func (sb serialBatch) addToAllBalances(numberOfBalances int, increment int) error {
	for i := 0; i < numberOfBalances; i++ {
		id := i

		err := sb.singleOperation.addBalance(id, increment)
		if err != nil {
			logger.Zap().Error("failed to increase balance", zap.Int("id", id), zap.Error(err))
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
