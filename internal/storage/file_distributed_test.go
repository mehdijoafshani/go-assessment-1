package storage

import (
	"github.com/mehdijoafshani/go-assessment-1/internal/config"
	"github.com/mehdijoafshani/go-assessment-1/test"
	"strconv"
	"testing"
)

func TestFileDistributedGetInt(t *testing.T) {
	initTestEnv()
	test.RewriteTestDataOnFiles()

	file := createDistributedFileStorage(config.Data.TestAccountsDir)
	balance, err := file.getInt(test.Balances[0].Id)
	if err != nil {
		t.Error(err)
	}

	if balance != test.Balances[0].Amount {
		t.Error()
	}
}

func TestFileDistributedGetIntInvalidId(t *testing.T) {
	initTestEnv()
	test.RewriteTestDataOnFiles()

	file := createDistributedFileStorage(config.Data.TestAccountsDir)
	_, err := file.getInt(len(test.Balances))
	if err == nil {
		t.Error(err)
	}
}

func TestFileDistributedGetIntNoFile(t *testing.T) {
	initTestEnv()
	test.RemoveAllTestFiles()

	file := createDistributedFileStorage(config.Data.TestAccountsDir)
	_, err := file.getInt(test.Balances[0].Id)
	if err == nil {
		t.Error(err)
	}
}

func TestFileDistributedCreateInt(t *testing.T) {
	initTestEnv()
	test.RemoveAllTestFiles()

	id := test.Balances[0].Id
	amount := test.Balances[0].Amount

	file := createDistributedFileStorage(config.Data.TestAccountsDir)
	err := file.createInt(id, amount)
	if err != nil {
		t.Error(err)
	}

	balance, err := strconv.Atoi(test.ReadTestDataContentFromTestFile(id))
	if err != nil {
		t.Error(err, "balance amount save on the file was not numeric")
	}

	if balance != amount {
		t.Error(err, "balance amount save on the file was incorrect")
	}
}

func TestFileDistributedCreateIntOnExistingFile(t *testing.T) {
	initTestEnv()
	test.RewriteTestDataOnFiles()

	id := test.Balances[0].Id
	amount := test.Balances[0].Amount

	file := createDistributedFileStorage(config.Data.TestAccountsDir)
	err := file.createInt(id, amount)
	if err != nil {
		t.Error(err)
	}

	balance, err := strconv.Atoi(test.ReadTestDataContentFromTestFile(id))
	if err != nil {
		t.Error(err, "balance amount save on the file was not numeric")
	}

	if balance != amount {
		t.Error(err, "balance amount save on the file was incorrect")
	}
}

func TestFileDistributedUpdateInt(t *testing.T) {
	initTestEnv()
	test.RewriteTestDataOnFiles()

	id := test.Balances[0].Id
	newAmount := test.Balances[0].Amount + 1000

	file := createDistributedFileStorage(config.Data.TestAccountsDir)
	err := file.updateInt(id, newAmount)
	if err != nil {
		t.Error(err)
	}

	balance, err := strconv.Atoi(test.ReadTestDataContentFromTestFile(id))
	if err != nil {
		t.Error(err, "balance amount save on the file was not numeric")
	}

	if balance != newAmount {
		t.Error(err, "balance amount save on the file was incorrect")
	}
}

func TestFileDistributedUpdateIntMissingFile(t *testing.T) {
	initTestEnv()
	test.RemoveAllTestFiles()

	id := test.Balances[0].Id
	newAmount := test.Balances[0].Amount + 1000

	file := createDistributedFileStorage(config.Data.TestAccountsDir)
	err := file.updateInt(id, newAmount)
	if err == nil {
		t.Error(err)
	}
}
