package balance

type concurrentBatch struct {
}

func (cb concurrentBatch) create(accountsNum int) error {
	//TODO impl
	return nil
}

func (cb concurrentBatch) getAll() (int64, error) {
	//TODO impl
	return 0, nil
}

func (cb concurrentBatch) addToAll(increment int) error {
	//TODO impl
	return nil
}

func createConcurrentBatch() concurrentBatch {
	return concurrentBatch{}
}
