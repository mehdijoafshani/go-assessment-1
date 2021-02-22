package account

import (
	"github.com/mehdijoafshani/go-assessment-1/internal/amount"
	"github.com/mehdijoafshani/go-assessment-1/internal/concurrency"
	"github.com/mehdijoafshani/go-assessment-1/internal/storage"
	"github.com/mehdijoafshani/go-assessment-1/test"
	"github.com/stretchr/testify/assert"
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
	test.ChangeNumberOfBalances(1000)
	concurrBatchOpMng := concurrentBatchOperation()
	numberOfBalances := len(test.Balances)

	for i := 0; i < 100; i++ {
		test.RewriteTestDataOnFiles()

		expectedSum := int64(0)
		for i := 0; i < numberOfBalances; i++ {
			expectedSum += int64(test.Balances[i].Amount)
		}

		result, err := concurrBatchOpMng.getAllBalancesSum(numberOfBalances)
		assert.Nil(t, err, "the method should return no error")
		assert.Equal(t, expectedSum, result, "the calculated sum of files should match the actual one")
	}
}

func TestConcurrentBatchCreateBalances(t *testing.T) {
	test.ChangeNumberOfBalances(1000)
	concurrBatchOpMng := concurrentBatchOperation()
	numberOfBalances := len(test.Balances)

	for i := 0; i < 100; i++ {
		test.RemoveAllTestFiles()

		err := concurrBatchOpMng.createBalances(numberOfBalances)
		assert.Nil(t, err, "the method should return no error")
		assert.Equal(t, numberOfBalances, test.NumberOfFiles(), "number of created files should match with the number of balances in the request")

		for id := 0; id < numberOfBalances; id++ {
			actualBalanceAmountStr := test.ReadTestDataContentFromTestFile(id)

			_, err := strconv.Atoi(actualBalanceAmountStr)
			assert.Nil(t, err, "content of created files should be numeric")
		}
	}
}

func TestConcurrentBatchAddToAllBalances(t *testing.T) {
	test.ChangeNumberOfBalances(1000)
	concurrBatchOpMng := concurrentBatchOperation()
	numberOfBalances := len(test.Balances)
	increment := 1000

	for i := 0; i < 100; i++ {
		test.RewriteTestDataOnFiles()

		err := concurrBatchOpMng.addToAllBalances(numberOfBalances, increment)
		assert.Nil(t, err, "the method should return no error")

		for id := 0; id < numberOfBalances; id++ {
			actualBalanceAmountStr := test.ReadTestDataContentFromTestFile(id)
			expectedBalanceAmount := test.Balances[id].Amount + increment

			actualBalanceAmount, err := strconv.Atoi(actualBalanceAmountStr)
			assert.Nil(t, err, "content of created files should be numeric")
			assert.Equal(t, expectedBalanceAmount, actualBalanceAmount, "content of created files should be numeric")
		}
	}
}
