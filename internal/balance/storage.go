package balance

type Storage interface {
	Create(content string) error
	Read(name string) (string, error)
	Update(name string, newContent string) error
}
