package storage

import (
	"github.com/mehdijoafshani/go-assessment-1/internal/config"
	"github.com/mehdijoafshani/go-assessment-1/test"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestFileDistributedGetInt(t *testing.T) {
	test.RewriteTestDataOnFiles()

	file := createDistributedFileStorage(config.Data.AccountsDir)
	balance, err := file.getInt(test.Balances[0].Id)
	assert.Nil(t, err, "getInt should return no error")

	assert.Equal(t, test.Balances[0].Amount, balance, "returned balance should match the test data")
}

func TestFileDistributedGetIntInvalidId(t *testing.T) {
	test.RewriteTestDataOnFiles()

	file := createDistributedFileStorage(config.Data.AccountsDir)
	_, err := file.getInt(len(test.Balances))
	assert.Error(t, err, "getInt should return an error when the given id does not exist")
}

func TestFileDistributedGetIntNoFile(t *testing.T) {
	test.RemoveAllTestFiles()

	file := createDistributedFileStorage(config.Data.AccountsDir)
	_, err := file.getInt(test.Balances[0].Id)
	assert.Error(t, err, "getInt should return an error when there is no balance file")

}

func TestFileDistributedCreateInt(t *testing.T) {
	test.RemoveAllTestFiles()

	id := test.Balances[0].Id
	expectedAmount := test.Balances[0].Amount

	file := createDistributedFileStorage(config.Data.AccountsDir)
	err := file.createInt(id, expectedAmount)
	assert.Nil(t, err, "should return no error")

	actualBalance, err := strconv.Atoi(test.ReadTestDataContentFromTestFile(id))
	assert.Nil(t, err, "the balance in the file ought to be numeric")
	assert.Equal(t, expectedAmount, actualBalance, "the testdata amount should match the fetched one")
}

func TestFileDistributedCreateIntOnExistingFile(t *testing.T) {
	test.RewriteTestDataOnFiles()

	id := test.Balances[0].Id
	expectedAmount := test.Balances[0].Amount

	file := createDistributedFileStorage(config.Data.AccountsDir)
	err := file.createInt(id, expectedAmount)
	assert.Nil(t, err, "should return no error")

	actualBalance, err := strconv.Atoi(test.ReadTestDataContentFromTestFile(id))
	assert.Nil(t, err, "saved balance amount should be numeric")
	assert.Equal(t, expectedAmount, actualBalance, "the testdata amount should match the created one")
}

func TestFileDistributedUpdateInt(t *testing.T) {
	test.RewriteTestDataOnFiles()

	id := test.Balances[0].Id
	newAmount := test.Balances[0].Amount + 1000

	file := createDistributedFileStorage(config.Data.AccountsDir)
	err := file.updateInt(id, newAmount)
	assert.Nil(t, err, "should return no error")

	actualBalance, err := strconv.Atoi(test.ReadTestDataContentFromTestFile(id))
	assert.Nil(t, err, "saved balance amount should be numeric")
	assert.Equal(t, newAmount, actualBalance, "the testdata amount should match the updated one")
}

func TestFileDistributedUpdateIntMissingFile(t *testing.T) {
	test.RemoveAllTestFiles()

	id := test.Balances[0].Id
	newAmount := test.Balances[0].Amount + 1000

	file := createDistributedFileStorage(config.Data.AccountsDir)
	err := file.updateInt(id, newAmount)
	assert.Error(t, err, "should return an error when the balance file is missing")
}
