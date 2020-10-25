package net

import (
	"sync"
	"testing"

	"dima.org/wb-test/internals/app/mock"
)

func TestProcess(t *testing.T) {
	urlsChan := make(chan string, 5)
	urlsChan <- "Url 1"
	urlsChan <- "Url 2"
	urlsChan <- "Url 3"
	urlsChan <- "Url 4"
	urlsChan <- "Url 5"

	resultsChan := make(chan interface{}, 1)

	var wg sync.WaitGroup
	wg.Add(1)

	worker := NewCountWorker(RequesterFunc(mock.RequestURL))
	worker.Process(&wg, urlsChan, resultsChan)

	total := <-resultsChan
	if total != mock.TotalGoAppearingNumber {
		t.Fail()
	}
}
