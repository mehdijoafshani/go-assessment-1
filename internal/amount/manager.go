package amount

import (
	"github.com/mehdijoafshani/go-assessment-1/balance"
	"github.com/mehdijoafshani/go-assessment-1/internal/config"
	"github.com/mehdijoafshani/go-assessment-1/internal/logger"
	"go.uber.org/zap"
)

type Manager struct {
	generator generator
}

func (m Manager) GenerateBalanceAmount(id int) (int, error){
	balanceAmount, err := m.generator.generateNumber(id)
	if err != nil {
		logger.Zap().Error("failed to generate number", zap.Int("id", id), zap.Error(err))
		return 0, err
	}

	return balanceAmount, nil
}

func CreateAmountManager() balance.AmountManager {
	return Manager{
		generator: createRandomAmountManager(config.Data.RandomBalanceMinRange, config.Data.RandomBalanceMaxRange),
	}
}
