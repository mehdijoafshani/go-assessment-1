package test

import (
	"github.com/mehdijoafshani/go-assessment-1/internal/balance"
)

type Account struct {
	Id      int
	Balance int
}

// assumptions:
//	- id of a balance equals to its index in Accounts slice
var Accounts []Account

func init() {
	Accounts = []Account{
		{Id: 0, Balance: 12_000},
		{Id: 1, Balance: 15_000},
		{Id: 2, Balance: 1_000},
	}
}

func ChangeNumberOfAccounts(accounts int) {
	Accounts = make([]Account, 0, accounts)
	for i := 0; i < accounts; i++ {
		balance, err := balance.CreateAmountManager().GenerateBalance(i)
		if err != nil {
			panic("failed to generate test data, failed to generate amount")
		}
		Accounts = append(Accounts, Account{
			Id:      i,
			Balance: balance,
		})
	}
}
