package net

import (
	"log"
	"strings"
	"sync"

	"dima.org/wb-test/internals/app/common"
)

// CountWorker implements tools.URLWorker interface
type CountWorker struct {
	requester Requester
	result    int
}

// NewCountWorker makes new CountWorker instance
func NewCountWorker(r Requester) *CountWorker {
	return &CountWorker{
		requester: r,
		result:    0,
	}
}

// Process ...
func (w *CountWorker) Process(wg *sync.WaitGroup, urlsChan <-chan string, resultsChan chan<- interface{}) {
	total := 0
	for {
		select {
		case url := <-urlsChan:
			if body, err := w.requester.RequestURL(url); err == nil {
				count := strings.Count(body, "Go")
				log.Println(url, " - ", count)
				total += count
			} else {
				log.Println(url, " - ", err)
			}

		default:
			resultsChan <- total
			wg.Done()
			return
		}
	}
}

// Clone clones the worker
func (w *CountWorker) Clone() common.Cloneable {
	return &CountWorker{
		requester: w.requester,
		result:    0,
	}
}
