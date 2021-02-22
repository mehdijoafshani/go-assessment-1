package concurrency

import (
	"github.com/mehdijoafshani/go-assessment-1/internal/logger"
	"go.uber.org/zap"
)

// #SOLID: L
// Liskov Substitution Principle is related to inheritance, while in Golang composition is used
// Also embedding a struct into another one only copies the BEHAVIOR of a struct to another one, it does not mean the embedding struct is extending the embedded struct
// There is no inheritance in Go. However, I can replace this start() method with any other possible implementation of concurrency.pattern interface
// And it will not break, this is the closest definition of Liskov Principle I can present here.
type poolingWaitForAll struct {
}

func (p poolingWaitForAll) start(jobsNum int, maxGoroutines int, worker func(ids <-chan int, results chan<- int, error chan<- error), resultHandler func(result int)) error {
	jobs := make(chan int, jobsNum)
	results := make(chan int, jobsNum)
	errorCh := make(chan error, jobsNum)

	remainingJobs := jobsNum

	// create workers
	for workerIndex := 0; workerIndex < maxGoroutines; workerIndex++ {
		go worker(jobs, results, errorCh)
	}

	// start jobs
	for jobIndex := 0; jobIndex < jobsNum; jobIndex++ {
		jobs <- jobIndex
		logger.Zap().Info("pooling, job sent", zap.Int("job id", jobIndex))
	}

	var err error
infiniteLoop:
	for {
		select {
		case err = <-errorCh:
			logger.Zap().Info("pooling, error received", zap.Error(err))
			break infiniteLoop
		case result := <-results:
			logger.Zap().Info("pooling, result received", zap.Int("result", result))
			remainingJobs--
			// TODO possible alternative: include resultHandler in worker and use waitGroup here
			resultHandler(result)
		}

		if remainingJobs == 0 {
			break
		}
	}
	close(jobs)

	return err
}

func createPoolingWaitForAllPattern() poolingWaitForAll {
	return poolingWaitForAll{}
}
