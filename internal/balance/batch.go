package balance

// the operation which process a number of procedure
// (the procedure could be performed either concurrently or serially)
type batch interface {
	create(accountsNum int) error
	getAll() (int64, error)
	addToAll(increment int) error
}
