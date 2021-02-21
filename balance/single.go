package balance

import (
	"github.com/mehdijoafshani/go-assessment-1/internal/logger"
	"go.uber.org/zap"
)

type singleOperationManager interface {
	create(id int) error
	get(id int) (int, error)
	add(id int, increment int) error
}

type singleOperationImpl struct {
	storageMng StorageManager
	amountMng  AmountManager
}

func (s singleOperationImpl) create(id int) error {
	amount, err := s.amountMng.GenerateBalanceAmount(id)
	if err != nil {
		logger.Zap().Error("failed to generate balance amount", zap.Int("id", id), zap.Error(err))
		return err
	}

	err = s.storageMng.CreateBalance(id, amount)
	if err != nil {
		logger.Zap().Error("failed to create balance", zap.Int("id", id), zap.Error(err))
		return err
	}

	return nil
}

func (s singleOperationImpl) get(id int) (int, error) {
	balance, err := s.storageMng.GetBalance(id)
	if err != nil {
		logger.Zap().Error("failed to fetch balance", zap.Int("id", id), zap.Error(err))
		return 0, err
	}

	return balance, nil
}

func (s singleOperationImpl) add(id int, increment int) error {
	err := s.storageMng.IncreaseBalance(id, increment)
	if err != nil {
		logger.Zap().Error("failed to increase balance", zap.Int("id", id), zap.Error(err))
		return err
	}

	return nil
}

func createSingleOperationTask(storageMng StorageManager, amountMng AmountManager) singleOperationImpl {
	return singleOperationImpl{
		storageMng: storageMng,
		amountMng:  amountMng,
	}
}