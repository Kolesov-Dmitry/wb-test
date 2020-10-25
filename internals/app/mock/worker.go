package mock

import (
	"sync"

	"dima.org/wb-test/internals/app/common"
)

// Worker is a URLWorker mock object
type Worker struct {
	result int
}

// NewWorker makes new CountWorker instance
func NewWorker() *Worker {
	return &Worker{
		result: 0,
	}
}

// Process ...
func (w *Worker) Process(wg *sync.WaitGroup, urlsChan <-chan string, resultsChan chan<- interface{}) {
	total := 0
	for {
		select {
		case <-urlsChan:
			total++

		default:
			resultsChan <- total
			wg.Done()
			return
		}
	}
}

// Clone clones the worker
func (w *Worker) Clone() common.Cloneable {
	return &Worker{
		result: 0,
	}
}
