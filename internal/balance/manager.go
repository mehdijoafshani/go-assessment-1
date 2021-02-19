package balance

type Manager struct {
}

func (m Manager) Create(accountsNum int) error {
	//TODO impl
	return nil
}

func (m Manager) GetAll() (int64, error) {
	//TODO impl
	return 0, nil
}

func (m Manager) Get(id int) (int, error) {
	//TODO impl
	return 0, nil
}

func (m Manager) AddToAll(increment int) error {
	//TODO impl
	return nil
}

func (m Manager) Add(increment int, id int) error {
	//TODO impl
	return nil
}

func CreateManager(isConcurrent bool) Manager {
	return Manager{}
}
