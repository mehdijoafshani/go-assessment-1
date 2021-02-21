package balance

// the operation which process a number of procedure
// (the procedure could be performed either concurrently or serially)
type batch interface {
	create(accountsNum int) error
	getAll(numberOfBalances int) (int64, error)
	addToAll(numberOfBalances int, increment int) error
}
