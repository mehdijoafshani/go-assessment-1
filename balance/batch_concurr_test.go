package balance

import (
	"github.com/mehdijoafshani/go-assessment-1/internal/amount"
	"github.com/mehdijoafshani/go-assessment-1/internal/concurrency"
	"github.com/mehdijoafshani/go-assessment-1/internal/logger"
	"github.com/mehdijoafshani/go-assessment-1/internal/storage"
	"github.com/mehdijoafshani/go-assessment-1/test"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func concurrentBatchOperation() concurrentBatch {
	storageMng := storage.CreateManager()
	amountMng := amount.CreateAmountManager()
	singleOpMng := createSingleOperationTask(storageMng, amountMng)
	concurrencyMng := concurrency.CreateManager()

	return createConcurrentBatch(storageMng, amountMng, concurrencyMng, singleOpMng)
}

func TestConcurrentBatchGetAllBalancesSum(t *testing.T) {
	test.RewriteTestDataOnFilesByNumberOfBalances(1000)

	concurrBatchOpMng := concurrentBatchOperation()

	numberOfBalances := len(test.Balances)
	expectedSum := int64(0)
	for i := 0; i < numberOfBalances; i++ {
		expectedSum += int64(test.Balances[i].Amount)
	}

	result, err := concurrBatchOpMng.getAllBalancesSum(numberOfBalances)
	if err != nil {
		logger.Zap().Error("TestConcurrentBatchGetAllBalancesSum", zap.Error(err))
	}

	assert.Nil(t, err, "the method should return no error")
	assert.Equal(t, expectedSum, result, "the calculated sum of files should match the actual one")
}
