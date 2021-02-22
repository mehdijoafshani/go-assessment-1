package concurrency

import (
	"github.com/mehdijoafshani/go-assessment-1/internal/logger"
	"go.uber.org/zap"
)

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

	for {
		select {
		case err := <-errorCh:
			logger.Zap().Info("pooling, error received", zap.Error(err))
			break
		case result := <-results:
			logger.Zap().Info("pooling, result received", zap.Int("result", result))
			remainingJobs--
			resultHandler(result)
		}

		if remainingJobs == 0 {
			break
		}
	}
	close(jobs)

	return nil
}

func createPoolingWaitForAllPattern() poolingWaitForAll {
	return poolingWaitForAll{}
}
