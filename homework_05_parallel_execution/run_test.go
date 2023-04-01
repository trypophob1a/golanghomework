package homework05parallelexecution

import (
	"errors"
	"fmt"
	"math/rand"
	"runtime"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

func TestRun(t *testing.T) {
	defer goleak.VerifyNone(t)

	t.Run("if were errors in first M tasks, than finished not more N+M tasks", func(t *testing.T) {
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32

		for i := 0; i < tasksCount; i++ {
			err := fmt.Errorf("error from task %d", i)
			tasks = append(tasks, func() error {
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
				atomic.AddInt32(&runTasksCount, 1)
				return err
			})
		}

		workersCount := 10
		maxErrorsCount := 23
		err := Run(tasks, workersCount, maxErrorsCount)
		fmt.Printf("\n\n>>>>>>> !%v! <<<<<<<<<<\n\n", err)
		require.Truef(t, errors.Is(err, ErrErrorsLimitExceeded), "actual err - %v", err)
		require.LessOrEqual(t, runTasksCount, int32(workersCount+maxErrorsCount), "extra tasks were started")
	})

	t.Run("tasks without errors", func(t *testing.T) {
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32
		var sumTime time.Duration

		for i := 0; i < tasksCount; i++ {
			taskSleep := time.Millisecond * time.Duration(rand.Intn(100))
			sumTime += taskSleep

			tasks = append(tasks, func() error {
				time.Sleep(taskSleep)
				atomic.AddInt32(&runTasksCount, 1)
				return nil
			})
		}

		workersCount := 5
		maxErrorsCount := 1

		start := time.Now()
		err := Run(tasks, workersCount, maxErrorsCount)
		elapsedTime := time.Since(start)
		require.NoError(t, err)

		require.Equal(t, runTasksCount, int32(tasksCount), "not all tasks were completed")
		require.LessOrEqual(t, int64(elapsedTime), int64(sumTime/2), "tasks were run sequentially?")
	})

	t.Run("empty slice tasks", func(t *testing.T) {
		tasks := make([]Task, 0)
		result := Run(tasks, 0, 0)
		require.Nil(t, result)
	})

	t.Run("zero count task", func(t *testing.T) {
		tasksCount := 0
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32
		atomic.AddInt32(&runTasksCount, 0)

		workersCount := 100
		maxErrorsCount := 200
		result := Run(tasks, workersCount, maxErrorsCount)
		require.Nil(t, result)

		require.Equal(t, runTasksCount, int32(tasksCount), "not all tasks were completed")
	})

	t.Run("test empty", func(t *testing.T) {
		Empty()
	})

	t.Run("test without error with require.Eventually", func(t *testing.T) {
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32

		for i := 0; i < tasksCount; i++ {
			tasks = append(tasks, func() error {
				atomic.AddInt32(&runTasksCount, 1)
				return nil
			})
		}

		workersCount := 5
		maxErrorsCount := 1

		start := time.Now()
		err := Run(tasks, workersCount, maxErrorsCount)
		elapsedTime := time.Since(start)
		require.NoError(t, err)

		require.Eventually(t, func() bool {
			return atomic.LoadInt32(&runTasksCount) == int32(tasksCount)
		}, 5*time.Second, 100*time.Millisecond, "not all tasks were completed")

		require.LessOrEqual(t, int64(elapsedTime), int64(time.Duration(10*tasksCount)*time.Millisecond), "tasks were run sequentially?")
	})
}

func TestRunMultipleGoroutines(t *testing.T) {
	defer goleak.VerifyNone(t)

	tasks := []Task{
		func() error {
			return nil
		},
		func() error {
			return nil
		},
	}

	n := 1
	m := 0

	err := Run(tasks, n, m)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if runtime.NumGoroutine() != n+1 {
		t.Errorf("Expected %d goroutines, but got %d",
			n+1,
			runtime.NumGoroutine(),
		)
	}
}
