package balance

type ConcurrencyManager interface {
	ScheduleReadAllBalancesSum(balancesNum int, worker func(ids <-chan int, results chan<- int, error chan<- error)) (int64, error)
	ScheduleCreateBalances(balancesNum int, worker func(ids <-chan int, results chan<- int, error chan<- error)) error
	ScheduleUpdateBalances(balancesNum int, worker func(ids <-chan int, results chan<- int, error chan<- error)) error
}
