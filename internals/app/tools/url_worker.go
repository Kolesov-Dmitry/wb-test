package tools

import (
	"sync"

	"dima.org/wb-test/internals/app/common"
)

// URLWorker interface
type URLWorker interface {
	common.Cloneable

	// Process starts the worker's main loop
	Process(wg *sync.WaitGroup, urlsChan <-chan string, resultsChan chan<- interface{})
}
