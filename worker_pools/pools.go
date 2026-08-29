package main

import (
	"fmt"
	"sync"
	"sync/atomic"

	"golang.org/x/sync/errgroup"
)

type Task struct {
	f func() error
}

func NewTask(funcArg func() error) Task {
	return Task{
		f: funcArg,
	}
}

type WorkerPools struct {
	runningWorkers int64
	capacity       int64
	JobChan        chan Task
	mutex          sync.Mutex
	errGroup       errgroup.Group
}

func NewWorkerPools(taskNum, cap int64) WorkerPools {
	return WorkerPools{
		errGroup: errgroup.Group{},
		capacity: cap,
		JobChan:  make(chan Task, taskNum),
	}
}

func (w *WorkerPools) GetCap() int64 {
	return w.capacity
}

func (w *WorkerPools) GetRunningWorkers() int64 {
	return atomic.LoadInt64(&w.runningWorkers)
}

func (w *WorkerPools) InCrRunningWorkers() {
	atomic.AddInt64(&w.runningWorkers, 1)
}

func (w *WorkerPools) DeCrRunningWorkers() {
	atomic.AddInt64(&w.runningWorkers, -1)
}

func (w *WorkerPools) run() {
	w.InCrRunningWorkers()
	w.errGroup.Go(func() error {
		defer w.DeCrRunningWorkers()
		for job := range w.JobChan {
			err := job.f()
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (w *WorkerPools) AddTasks(task Task) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	if w.GetCap() > w.GetRunningWorkers() {
		w.run()
	}

	w.JobChan <- task
}

func main() {
	workers := NewWorkerPools(10, 3)
	for i := 0; i < 20; i++ {
		task := NewTask(func() error {
			fmt.Println("job : i ", i, "started")
			return nil
		})
		workers.AddTasks(task)
	}
	close(workers.JobChan)
	err := workers.errGroup.Wait()
	if err != nil {
		return
	}

}
