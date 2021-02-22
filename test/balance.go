package test

import (
	"github.com/mehdijoafshani/go-assessment-1/internal/amount"
	"github.com/mehdijoafshani/go-assessment-1/internal/config"
	"github.com/mehdijoafshani/go-assessment-1/internal/logger"
	"go.uber.org/zap"
	"io/ioutil"
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
	}
}

func RewriteTestDataOnFilesByNumberOfBalances(balances int) {
	Balances = make([]Balance, 0, balances)
	for i := 0; i < balances; i++ {
		balanceAmount, err := amount.CreateAmountManager().GenerateBalanceAmount(i)
		if err != nil {
			panic("failed to generate test data, failed to generate amount")
		}
		Balances = append(Balances, Balance{
			Id:     i,
			Amount: balanceAmount,
		})
	}

	RewriteTestDataOnFiles()
}

func RewriteTestDataOnFiles() {
	RemoveAllTestFiles()
	dir := config.Data.AccountsDir

	for _, b := range Balances {
		data := strconv.Itoa(b.Amount)

		f, err := os.Create(filepath.Join(dir, strconv.Itoa(b.Id)+config.Data.BalanceFileExtension))
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
	dir := config.Data.AccountsDir

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

func ReadTestDataContentFromTestFile(id int) string {
	dir := config.Data.AccountsDir
	fileName := strconv.Itoa(id) + config.Data.BalanceFileExtension

	data, err := ioutil.ReadFile(filepath.Join(dir, fileName))
	if err != nil {
		panic("test file not found")
	}

	return string(data)
}
