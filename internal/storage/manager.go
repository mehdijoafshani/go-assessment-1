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

func (m Manager) AreAccountsCreated() (bool, error) {
	isEmpty, err := m.file.isDirEmpty()
	if err != nil {
		logger.Zap().Error("failed to check if any account is created in the storage", zap.Error(err))
		return false, err
	}

	areAccountsCreated := !isEmpty
	return areAccountsCreated, nil
}

func (m Manager) CreateAccount(id int, balance int) error {
	err := m.file.createInt(id, balance)
	if err != nil {
		logger.Zap().Error("failed to create account in the storage", zap.Error(err), zap.Int("id", id))
		return err
	}

	return nil
}

func (m Manager) GetBalance(id int) (int, error) {
	balance, err := m.file.getInt(id)
	if err != nil {
		logger.Zap().Error("failed to get balance from the storage", zap.Error(err), zap.Int("id", id))
		return 0, err
	}

	return balance, nil
}

func (m Manager) IncreaseBalance(id int, newBalance int) error {
	balance, err := m.GetBalance(id)
	if err != nil {
		logger.Zap().Error("failed to get balance from file to update", zap.Error(err), zap.Int("id", id))
		return err
	}

	err = m.file.updateInt(id, newBalance+balance)
	if err != nil {
		logger.Zap().Error("failed to update balance in file", zap.Error(err), zap.Int("id", id))
		return err
	}

	return nil
}

func (m Manager) Truncate() error {
	err := m.file.truncateDir()
	if err != nil {
		logger.Zap().Error("failed to truncate the storage", zap.Error(err))
		return err
	}

	return nil
}

func (m Manager) NumberOfAccounts() (int, error) {
	numbers, err := m.file.dirFilesNumber(config.Data.AccountFileExtension)
	if err != nil {
		logger.Zap().Error("failed to get the number of accounts in the storage", zap.Error(err))
		return 0, err
	}

	return numbers, nil
}

func CreateManager() Manager {
	url := config.Data.AccountsDir

	return Manager{
		file: createDistributedFileStorage(url),
	}
}
