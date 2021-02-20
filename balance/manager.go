package balance

import (
	"github.com/mehdijoafshani/go-assessment-1/internal/logger"
	"github.com/mehdijoafshani/go-assessment-1/internal/storage"
	"go.uber.org/zap"
)

type Manager struct {
	batch   batch
	storage Storage
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
	balance, err := m.storage.GetBalance(id)
	if err != nil {
		logger.Zap().Error("failed to read the balance from the storage", zap.Error(err))
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
	err := m.storage.IncreaseBalance(id, increment)
	if err != nil {
		logger.Zap().Error("failed to update the balance in the storage", zap.Error(err))
		return err
	}

	return nil
}

func CreateManager(isConcurrent bool) Manager {
	mng := Manager{
		storage: storage.Get(),
	}

	if isConcurrent {
		mng.batch = createSerialBatch()
	} else {
		mng.batch = createConcurrentBatch()
	}

	return mng
}
