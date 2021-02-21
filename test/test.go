package test

import (
	"github.com/mehdijoafshani/go-assessment-1/internal/config"
	"github.com/mehdijoafshani/go-assessment-1/internal/logger"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"strconv"
)

type Balance struct {
	Id     int
	Amount int
}

// assumptions:
//	- id of a balance equals to its index in Balances slice
var Balances []Balance

func init() {
	Balances = []Balance{
		{Id: 0, Amount: 12_000},
		{Id: 1, Amount: 15_000},
		{Id: 2, Amount: 1_000},
		{Id: 3, Amount: 10},
	}
}

func RewriteTestDataOnFiles() {
	RemoveAllTestFiles()
	dir := config.Data.TestAccountsDir

	for _, b := range Balances {
		data := strconv.Itoa(b.Amount)

		f, err := os.Create(filepath.Join(dir, strconv.Itoa(b.Id)+".txt"))
		if err != nil {
			panic("failed to create test data")
		}

		_, err = f.WriteString(data)
		if err != nil {
			panic("failed to write test data")
		}

		err = f.Close()
		if err != nil {
			logger.Zap().Error("failed to close test file", zap.Error(err))
		}
	}
}

func RemoveAllTestFiles() {
	dir := config.Data.TestAccountsDir

	d, err := os.Open(dir)
	if err != nil {
		panic("failed to open testdata dir")
	}

	defer d.Close()

	names, err := d.Readdirnames(-1)
	if err != nil {
		panic("failed to read testdata files")
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			panic("failed to remove testdata files")
		}
	}
}
