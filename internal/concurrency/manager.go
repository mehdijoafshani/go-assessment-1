package concurrency

import (
	"github.com/mehdijoafshani/go-assessment-1/internal/config"
	"github.com/mehdijoafshani/go-assessment-1/internal/logger"
	"go.uber.org/zap"
	"sync/atomic"
)

type Manager struct {
	pattern pattern
}

func (m Manager) ScheduleReadAllBalancesSum(balancesNum int, worker func(ids <-chan int, results chan<- int, error chan<- error)) (int64, error) {
	balancesSum := int64(0)

	err := m.pattern.start(balancesNum,
		config.Data.MaxConcurrentGoroutines,
		worker,
		//result handler:
		func(result int) {
			atomic.AddInt64(&balancesSum, int64(result))
		})
	if err != nil {
		logger.Zap().Error("failed to run read all balances concurrently", zap.Error(err))
		return 0, err
	}

	logger.Zap().Info("concurrency.manager read all done !")
	return balancesSum, nil
}

func (m Manager) ScheduleCreateBalances(balancesNum int, worker func(ids <-chan int, results chan<- int, error chan<- error)) error {
	err := m.pattern.start(balancesNum,
		config.Data.MaxConcurrentGoroutines,
		worker,
		func(result int) {
			// the result is only a simple signal here
		})
	if err != nil {
		logger.Zap().Error("failed to run create balances concurrently", zap.Error(err))
		return err
	}

	logger.Zap().Info("concurrency.manager create done !")
	return nil
}

func (m Manager) ScheduleUpdateBalances(balancesNum int, worker func(ids <-chan int, results chan<- int, error chan<- error)) error {
	err := m.pattern.start(balancesNum,
		config.Data.MaxConcurrentGoroutines,
		worker,
		func(result int) {
			// the result is only a simple signal here
		})
	if err != nil {
		logger.Zap().Error("failed to run update balances concurrently", zap.Error(err))
		return err
	}

	logger.Zap().Info("concurrency.manager update done !")
	return nil
}

func CreateManager() Manager {
	return Manager{
		pattern: createPoolingWaitForAllPattern(),
	}
}
