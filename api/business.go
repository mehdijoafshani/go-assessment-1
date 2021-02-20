package api

type balanceManager interface {
	Create(accountsNum int) error
	GetAll() (int64, error)
	Get(accId int) (int, error)
	// TODO rename add with another name
	AddToAll(increment int) error
	Add(increment int, accId int) error
}
