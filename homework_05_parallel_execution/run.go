package homework05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type worker struct {
	errorLimit   uint32
	errorCounter uint32
	wg           *sync.WaitGroup
}

func newWorker(errorLimit uint32) *worker {
	wg := &sync.WaitGroup{}
	return &worker{errorLimit: errorLimit, wg: wg}
}

func process(Task Task, queue <-chan struct{}, worker *worker) {
	defer func() {
		<-queue
		worker.wg.Done()
	}()
	if atomic.LoadUint32(&worker.errorCounter) >= worker.errorLimit {
		return
	}
	err := Task()
	if err != nil {
		atomic.AddUint32(&worker.errorCounter, 1)
		return
	}
}

func Run(tasks []Task, n int, m int) error {
	queue := make(chan struct{}, n)
	defer close(queue)
	errLimit := uint32(m)
	worker := newWorker(errLimit)

	for _, task := range tasks {
		if atomic.LoadUint32(&worker.errorCounter) >= errLimit {
			return ErrErrorsLimitExceeded
		}
		queue <- struct{}{}
		worker.wg.Add(1)
		go process(task, queue, worker)
	}

	worker.wg.Wait()
	return nil
}
