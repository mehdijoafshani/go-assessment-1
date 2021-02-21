package balance

import (
	"github.com/mehdijoafshani/go-assessment-1/internal/amount"
	"github.com/mehdijoafshani/go-assessment-1/internal/logger"
	"github.com/mehdijoafshani/go-assessment-1/internal/storage"
	"go.uber.org/zap"
)

type Manager struct {
	batch      batch
	storageMng StorageManager
}

func (m Manager) Create(accountsNum int) error {
	err := m.batch.create(accountsNum)
	if err != nil {
		logger.Zap().Error("failed to create accounts in batch", zap.Error(err))
		return err
	}
	return nil
}

func (m Manager) GetAll() (int64, error) {
	balance, err := m.batch.getAll()
	if err != nil {
		logger.Zap().Error("failed to get all balances in batch", zap.Error(err))
		return 0, err
	}

	return balance, nil
}

func (m Manager) Get(id int) (int, error) {
	balance, err := m.storageMng.GetBalance(id)
	if err != nil {
		logger.Zap().Error("failed to read the balance from the storageMng", zap.Error(err))
		return 0, err
	}

	return balance, nil
}

func (m Manager) AddToAll(increment int) error {
	err := m.batch.addToAll(increment)
	if err != nil {
		logger.Zap().Error("failed to add extra balance to all accounts", zap.Error(err))
		return err
	}

	return nil
}

func (m Manager) Add(increment int, id int) error {
	// TODO: is zero balance allowed (negative increment)?
	err := m.storageMng.IncreaseBalance(id, increment)
	if err != nil {
		logger.Zap().Error("failed to update the balance in the storageMng", zap.Error(err))
		return err
	}

	return nil
}

func CreateManager(isConcurrent bool) Manager {
	mng := Manager{
		storageMng: storage.CreateManager(),
	}

	if isConcurrent {
		mng.batch = createConcurrentBatch(mng.storageMng, amount.CreateAmountManager())
	} else {
		mng.batch = createSerialBatch(mng.storageMng, amount.CreateAmountManager())
	}

	return mng
}
