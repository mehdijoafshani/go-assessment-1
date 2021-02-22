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

	logger.Zap().Info("concurrency.manager done !")
	return balancesSum, nil
}

func (m Manager) ScheduleCreateBalances(balancesNum int, worker func(ids <-chan int, results chan<- int, error chan<- error)) error {
	panic("not implemented")
}

func (m Manager) ScheduleUpdateBalances(balancesNum int, worker func(ids <-chan int, results chan<- int, error chan<- error)) error {
	panic("not implemented")
}

func CreateManager() Manager {
	return Manager{
		pattern: createPoolingPattern(),
	}
}
