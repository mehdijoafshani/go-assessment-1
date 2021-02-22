package concurrency

import (
	"github.com/mehdijoafshani/go-assessment-1/internal/logger"
	"go.uber.org/zap"
)

type pooling struct {
}

func (p pooling) start(jobsNum int, maxGoroutines int, worker func(ids <-chan int, results chan<- int, error chan<- error), resultHandler func(result int)) error {
	jobs := make(chan int, jobsNum)
	results := make(chan int, jobsNum)
	errorCh := make(chan error, jobsNum)

	// create workers
	for workerIndex := 0; workerIndex < maxGoroutines; workerIndex++ {
		go worker(jobs, results, errorCh)
	}

	// start jobs
	for jobIndex := 0; jobIndex < jobsNum; jobIndex++ {
		jobs <- jobIndex
		logger.Zap().Info("pooling, job sent", zap.Int("job id", jobIndex))
	}
	close(jobs)

	for len(results) > 0 {
		select {
		case err := <-errorCh:
			logger.Zap().Info("pooling, error received", zap.Error(err))
			return err
		case result := <-results:
			logger.Zap().Info("pooling, result received", zap.Int("result", result))
			resultHandler(result)
		}
	}

	return nil
}

func createPoolingPattern() pooling {
	return pooling{}
}
