package balance

type ConcurrencyManager interface {
	ScheduleReadAllBalancesSum(balancesNum int,
		worker func(ids <-chan int, results chan<- int, errors chan<- error)) (int64, error)
	ScheduleCreateBalances(balancesNum int,
		worker func(ids <-chan int, errors chan<- error)) error
	// map[int]int is a map of id -> balance
	ScheduleUpdateBalances(balancesNum int,
		worker func(ids <-chan int, errors chan<- error)) error
}
