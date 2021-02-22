package balance

import (
	"github.com/mehdijoafshani/go-assessment-1/internal/amount"
	"github.com/mehdijoafshani/go-assessment-1/internal/logger"
	"github.com/mehdijoafshani/go-assessment-1/internal/storage"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Manager struct {
	batch       batchOperationManager
	storageMng  StorageManager
	singleOpMng singleOperationManager
}

func (m Manager) Create(accountsNum int) error {
	// rule: balances should be created only once
	areBalancesCreated, err := m.storageMng.AreBalancesCreated()
	if err != nil {
		logger.Zap().Error("failed to check if balances are created before", zap.Error(err))

	}
	if areBalancesCreated {
		logger.Zap().Error("failed to create balances, they are created before")
		return errors.New("balances are created before, it is allowed to be created only once")
	}

	err = m.batch.createBalances(accountsNum)
	if err != nil {
		logger.Zap().Error("failed to create accounts in batchOperationManager", zap.Error(err))
	}

	// define truncErr to avoid hierarchical code
	var truncErr error
	if err != nil {
		truncErr = m.storageMng.Truncate()
	}
	if truncErr != nil {
		logger.Zap().Fatal("failed to truncate the storage after failed creation. The system data is in invalid state", zap.Error(truncErr))
	}
	if err != nil {
		return err
	}

	return nil
}

func (m Manager) GetAll() (int64, error) {
	numberOfBalances, err := m.storageMng.NumberOfBalances()
	if err != nil {
		logger.Zap().Error("failed to getBalance the number of balances", zap.Error(err))
		return 0, err
	}

	balance, err := m.batch.getAllBalancesSum(numberOfBalances)
	if err != nil {
		logger.Zap().Error("failed to getBalance all balances in batchOperationManager", zap.Error(err))
		return 0, err
	}

	return balance, nil
}

func (m Manager) Get(id int) (int, error) {
	balance, err := m.singleOpMng.getBalance(id)
	if err != nil {
		logger.Zap().Error("failed to read the balance from the storageMng", zap.Error(err))
		return 0, err
	}

	return balance, nil
}

func (m Manager) AddToAll(increment int) error {
	numberOfBalances, err := m.storageMng.NumberOfBalances()
	if err != nil {
		logger.Zap().Error("failed to getBalance the number of balances", zap.Error(err))
		return err
	}

	err = m.batch.addToAllBalances(numberOfBalances, increment)
	if err != nil {
		logger.Zap().Error("failed to addBalance extra balance to all accounts", zap.Error(err))
		// TODO (IMPORTANT): rollback made changes !
		return err
	}

	return nil
}

func (m Manager) Add(increment int, id int) error {
	// TODO: is zero balance allowed (negative increment)?
	err := m.singleOpMng.addBalance(id, increment)
	if err != nil {
		logger.Zap().Error("failed to update the balance in the storageMng", zap.Error(err))
		return err
	}

	return nil
}

func CreateManager(isConcurrent bool) Manager {
	storageMng := storage.CreateManager()
	amountMng := amount.CreateAmountManager()
	singleOpMng := createSingleOperationTask(storageMng, amountMng)
	// TODO: replace the nil with a factory function call
	var concurrencyMng ConcurrencyManager = nil

	mng := Manager{
		storageMng: storageMng,
	}

	if isConcurrent {
		mng.batch = createConcurrentBatch(storageMng, amountMng, concurrencyMng, singleOpMng)
	} else {
		mng.batch = createSerialBatch(storageMng, amountMng, singleOpMng)
	}

	return mng
}
