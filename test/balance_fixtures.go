package test

import (
	"github.com/mehdijoafshani/go-assessment-1/internal/amount"
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

func ChangeNumberOfBalances(balances int) {
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
}
