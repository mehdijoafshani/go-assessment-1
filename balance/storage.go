package balance

type StorageManager interface {
	AreBalancesCreated() (bool, error)
	CreateBalance(id int, amount int) error
	GetBalance(id int) (int, error)
	IncreaseBalance(id int, newContent int) error
	Truncate() error
}
