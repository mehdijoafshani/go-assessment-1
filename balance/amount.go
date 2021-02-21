package balance

type AmountManager interface {
	// the current implementation generates random balances, the id parameter here is for the possible
	//future cases in which the id of the balance would be needed to generate the amount of the balance
	GenerateBalanceAmount(id int) (int, error)
}
