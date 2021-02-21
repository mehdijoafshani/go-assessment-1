package amount

type generator interface {
	generateNumber(id int) (int, error)
}
