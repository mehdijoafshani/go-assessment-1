package storage

import (
	"github.com/mehdijoafshani/go-assessment-1/internal/config"
	"github.com/mehdijoafshani/go-assessment-1/internal/logger"
	"go.uber.org/zap"
)

// TODO rename
type Manager struct {
	file file
}

func (m Manager) CreateBalance(id int, content int) error {
	err := m.file.createInt(id, content)
	if err != nil {
		logger.Zap().Error("failed to create balance in file", zap.Error(err))
		return err
	}

	return nil
}

func (m Manager) GetBalance(id int) (int, error) {
	balance, err := m.file.getInt(id)
	if err != nil {
		logger.Zap().Error("failed to get balance from file", zap.Error(err), zap.Int("id", id))
		return 0, err
	}

	return balance, nil
}

func (m Manager) IncreaseBalance(id int, newContent int) error {
	balance, err := m.GetBalance(id)
	if err != nil {
		logger.Zap().Error("failed to get balance from file to update", zap.Error(err), zap.Int("id", id))
		return err
	}

	err = m.file.updateInt(id, newContent+balance)
	if err != nil {
		logger.Zap().Error("failed to update balance in file", zap.Error(err), zap.Int("id", id))
		return err
	}

	return nil
}

func Get() Manager {
	url := config.Data.AccountsDir

	return Manager{
		file: createDistributedFileStorage(url),
	}
}
