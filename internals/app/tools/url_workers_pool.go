package tools

import (
	"bufio"
	"io"
	"sync"
)

// URLWorkersPool is a pool of workers
type URLWorkersPool struct {
	maxWorkersAmount int         // the maximum amount of workers
	workersAmount    int         // current amount of workers
	inputStream      io.Reader   // the input stream to obtain urls from
	result           interface{} // the final result
	worker           URLWorker   //
}

// NewURLWorkersPool make new URLWorkersPool instance
// Inputs:
// - maximum amount of workers value
// - input stream to obtain urls
func NewURLWorkersPool(max int, w URLWorker, stream io.Reader) *URLWorkersPool {
	return &URLWorkersPool{
		maxWorkersAmount: max,
		workersAmount:    0,
		inputStream:      stream,
		result:           nil,
		worker:           w,
	}
}

// Start starts the URLs handling process
// Note: bloking until all URLs are handled
func (p *URLWorkersPool) Start() {
	var wg sync.WaitGroup

	urlsChan := make(chan string, p.maxWorkersAmount)
	resultsChan := make(chan interface{}, p.maxWorkersAmount)

	// read the input stream
	sc := bufio.NewScanner(p.inputStream)
	for sc.Scan() {
		urlsChan <- sc.Text()
		p.runWorker(&wg, urlsChan, resultsChan)
	}

	wg.Wait()

	close(urlsChan)
	close(resultsChan)

	p.collectResults(resultsChan)
}

// Result returns the total result
func (p *URLWorkersPool) Result() interface{} {
	return p.result
}

func (p *URLWorkersPool) collectResults(resultsChan <-chan interface{}) {
	result := 0
	for res := range resultsChan {
		result += res.(int)
	}

	p.result = result
}

func (p *URLWorkersPool) runWorker(wg *sync.WaitGroup, urlsChan <-chan string, resultsChan chan<- interface{}) {
	if p.workersAmount < p.maxWorkersAmount {
		p.workersAmount++
		worker := p.worker.Clone().(URLWorker)
		go worker.Process(wg, urlsChan, resultsChan)
		wg.Add(1)
	}
}
