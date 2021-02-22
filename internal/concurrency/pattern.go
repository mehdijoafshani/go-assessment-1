package concurrency

type pattern interface {
	start(jobsNum int, maxGoroutines int, worker func(ids <-chan int, results chan<- int, error chan<- error), resultHandler func(result int)) error
}
