package balance

type Storage interface {
	CreateBalance(id int, content int) error
	GetBalance(id int) (int, error)
	IncreaseBalance(id int, newContent int) error
}
