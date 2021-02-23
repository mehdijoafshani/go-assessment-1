package concurrency

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"sync/atomic"
	"testing"
)

func TestPoolingWaitForAllStart(t *testing.T) {
	poolWait := createPoolingWaitForAllPattern()

	jobs := 1000
	maxGoRoutines := 50

	var processedJobs int32 = 0

	err := poolWait.start(jobs, maxGoRoutines, func(ids <-chan int, results chan<- int, error chan<- error) {
		for id := range ids {
			results <- id
		}
	}, func(result int) {
		atomic.AddInt32(&processedJobs, 1)
	})

	assert.Nil(t, err, "no error signal had been sent")
	assert.Equal(t, processedJobs, int32(jobs))
}

func TestPoolingWaitForAllStartErrorHappens(t *testing.T) {
	poolWait := createPoolingWaitForAllPattern()

	jobs := 1000
	maxGoRoutines := 50

	var processedJobs int32 = 0

	err := poolWait.start(jobs, maxGoRoutines, func(ids <-chan int, results chan<- int, error chan<- error) {
		for id := range ids {
			error <- errors.New("some error")
			results <- id
		}
	}, func(result int) {
		atomic.AddInt32(&processedJobs, 1)
	})

	assert.NotNil(t, err, "an error was sent to the channel")
}
