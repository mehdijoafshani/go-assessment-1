package account

type StorageManager interface {
	AreAccountsCreated() (bool, error)
	CreateAccount(id int, amount int) error
	GetBalance(id int) (int, error)
	IncreaseBalance(id int, newBalance int) error
	Truncate() error
	NumberOfAccounts() (int, error)
}
