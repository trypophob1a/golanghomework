package homework05parallelexecution

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Empty() {
	tasks := make([]Task, 0)
	err := Run(tasks, 0, 0)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n int, m int) error {

	queueTask := make(chan Task, n)
	var errCounter uint32
	errorLimit := uint32(m)
	wg := &sync.WaitGroup{}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(counter *uint32) {
			defer wg.Done()

			for task := range queueTask {
				if atomic.LoadUint32(counter) >= errorLimit {
					return
				}

				err := task()
				if err != nil {
					atomic.AddUint32(counter, 1)
				}
			}
		}(&errCounter)
	}

	for _, task := range tasks {
		queueTask <- task
	}
	close(queueTask)
	wg.Wait()

	if errCounter >= errorLimit && errorLimit != 0 {
		return ErrErrorsLimitExceeded
	}

	return nil
}
