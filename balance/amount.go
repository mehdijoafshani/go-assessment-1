package balance

import "github.com/mehdijoafshani/go-assessment-1/internal/config"

type amountManager interface {
	// the current implementation generates random balances, the id parameter here is for the possible
	//future cases in which the id of the balance would be needed to generate the amount of the balance
	generateBalance(id int) (int, error)
}

func createAmountManager() amountManager {
	return createRandomAmountManager(config.Data.RandomBalanceMinRange, config.Data.RandomBalanceMaxRange)
}
