package account

// #SOLID: O
// Instead of relying on a concrete struct and implement createBalances, ... methods, an interface has been defined
// This will give us the power of changing the functionality with bringing new solutions (extending the usage of this interface instead of changing single implementation)
// This indicates that Open/Closed principle has been applied here
// We have to pay attention that critical changes should be applied and extending is not for all cases

// the operation which process a number of procedure
// (the procedure could be performed either concurrently or serially)
type batchOperationManager interface {
	createBalances(accountsNum int) error
	getAllBalancesSum(numberOfBalances int) (int64, error)
	addToAllBalances(numberOfBalances int, increment int) error
}
