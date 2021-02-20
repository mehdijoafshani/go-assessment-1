package balance

import (
	"github.com/mehdijoafshani/go-assessment-1/internal/logger"
	"go.uber.org/zap"
)

type Manager struct {
	batch batch
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
	//TODO impl
	return 0, nil
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
	//TODO impl
	return nil
}

func CreateManager(isConcurrent bool) Manager {
	mng := Manager{}

	if isConcurrent {
		mng.batch = createSerialBatch()
	} else {
		mng.batch = createConcurrentBatch()
	}

	return mng
}
