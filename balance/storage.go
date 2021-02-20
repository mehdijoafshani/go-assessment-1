package balance

type Storage interface {
	Create(id int, content string) error
	Read(id int) (int, error)
	Update(id int, newContent int) error
}
