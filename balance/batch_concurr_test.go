package balance

import (
	"github.com/mehdijoafshani/go-assessment-1/internal/amount"
	"github.com/mehdijoafshani/go-assessment-1/internal/concurrency"
	"github.com/mehdijoafshani/go-assessment-1/internal/logger"
	"github.com/mehdijoafshani/go-assessment-1/internal/storage"
	"github.com/mehdijoafshani/go-assessment-1/test"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"strconv"
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
	for i := 0; i < 100; i++ {
		test.ChangeNumberOfBalances(1000)
		test.RewriteTestDataOnFiles()

		concurrBatchOpMng := concurrentBatchOperation()

		numberOfBalances := len(test.Balances)
		expectedSum := int64(0)
		for i := 0; i < numberOfBalances; i++ {
			expectedSum += int64(test.Balances[i].Amount)
		}

		result, err := concurrBatchOpMng.getAllBalancesSum(numberOfBalances)
		if err != nil {
			logger.Zap().Error("error on getAllBalancesSum", zap.Error(err))
		}

		assert.Nil(t, err, "the method should return no error")
		assert.Equal(t, expectedSum, result, "the calculated sum of files should match the actual one")
	}
}

func TestConcurrentBatchCreateBalances(t *testing.T) {
	for i := 0; i < 100; i++ {
		test.ChangeNumberOfBalances(1000)
		test.RemoveAllTestFiles()

		concurrBatchOpMng := concurrentBatchOperation()

		numberOfBalances := len(test.Balances)

		err := concurrBatchOpMng.createBalances(numberOfBalances)
		if err != nil {
			logger.Zap().Error("error on createBalances", zap.Error(err))
		}

		assert.Nil(t, err, "the method should return no error")
		assert.Equal(t, numberOfBalances, test.NumberOfFiles(), "number of created files should match with the number of balances in the request")

		for id := 0; id < numberOfBalances; id++ {
			actualBalanceAmountStr := test.ReadTestDataContentFromTestFile(id)

			_, err := strconv.Atoi(actualBalanceAmountStr)
			assert.Nil(t, err, "content of created files should be numeric")
		}
	}
}
