package balance

// the operation which process a number of procedure
// (the procedure could be performed either concurrently or serially)
type batchOperationManager interface {
	createBalances(accountsNum int) error
	getAllBalancesSum(numberOfBalances int) (int64, error)
	addToAllBalances(numberOfBalances int, increment int) error
}
