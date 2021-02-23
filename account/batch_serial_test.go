package account

import (
	"github.com/mehdijoafshani/go-assessment-1/internal/balance"
	"github.com/mehdijoafshani/go-assessment-1/internal/storage"
	"github.com/mehdijoafshani/go-assessment-1/test"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func serialBatchOperation() serialBatch {
	storageMng := storage.CreateManager()
	amountMng := balance.CreateAmountManager()
	singleOpMng := createSingleOperationTask(storageMng, amountMng)

	return createSerialBatch(storageMng, amountMng, singleOpMng)
}

func TestSerialBatchGetAllBalancesSum(t *testing.T) {
	test.ChangeNumberOfAccounts(1000)
	serialBatchOpMng := serialBatchOperation()
	numberOfAccounts := len(test.Accounts)

	test.RewriteTestDataOnFiles()

	expectedSum := int64(0)
	for i := 0; i < numberOfAccounts; i++ {
		expectedSum += int64(test.Accounts[i].Balance)
	}

	result, err := serialBatchOpMng.getAllBalancesSum(numberOfAccounts)
	assert.Nil(t, err, "the method should return no error")
	assert.Equal(t, expectedSum, result, "the calculated sum of files should match the actual one")
}

func TestSerialBatchCreateBalances(t *testing.T) {
	test.ChangeNumberOfAccounts(1000)
	serialBatchOpMng := concurrentBatchOperation()
	numberOfAccounts := len(test.Accounts)

	test.RemoveAllTestFiles()

	err := serialBatchOpMng.createAccounts(numberOfAccounts)
	assert.Nil(t, err, "the method should return no error")
	assert.Equal(t, numberOfAccounts, test.NumberOfFiles(), "number of created files should match with the number of balances in the request")

	for id := 0; id < numberOfAccounts; id++ {
		actualBalanceAmountStr := test.ReadTestDataContentFromTestFile(id)

		_, err := strconv.Atoi(actualBalanceAmountStr)
		assert.Nil(t, err, "content of created files should be numeric")
	}
}

func TestSerialBatchAddToAllBalances(t *testing.T) {
	test.ChangeNumberOfAccounts(1000)
	serialBatchOpMng := concurrentBatchOperation()
	numberOfAccounts := len(test.Accounts)
	increment := 1000

	test.RewriteTestDataOnFiles()

	err := serialBatchOpMng.addToAllBalances(numberOfAccounts, increment)
	assert.Nil(t, err, "the method should return no error")

	for id := 0; id < numberOfAccounts; id++ {
		actualBalanceAmountStr := test.ReadTestDataContentFromTestFile(id)
		expectedBalanceAmount := test.Accounts[id].Balance + increment

		actualBalanceAmount, err := strconv.Atoi(actualBalanceAmountStr)
		assert.Nil(t, err, "content of created files should be numeric")
		assert.Equal(t, expectedBalanceAmount, actualBalanceAmount, "content of created files should be numeric")
	}
}
