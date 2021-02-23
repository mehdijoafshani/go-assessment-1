package balance

type generator interface {
	generateNumber(id int) (int, error)
}
