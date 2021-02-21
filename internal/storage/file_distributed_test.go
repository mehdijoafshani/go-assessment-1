package storage

import (
	"github.com/mehdijoafshani/go-assessment-1/internal/config"
	"github.com/mehdijoafshani/go-assessment-1/test"
	"testing"
)

func TestGetInt(t *testing.T) {
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

func TestGetIntNoFile(t *testing.T) {
	initTestEnv()
	test.RemoveAllTestFiles()

	file := createDistributedFileStorage(config.Data.TestAccountsDir)
	_, err := file.getInt(test.Balances[0].Id)
	if err == nil {
		t.Error(err)
	}
}
