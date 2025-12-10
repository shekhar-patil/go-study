package main

import (
	"fmt"
	"sync"
)

type Job struct {
	Id    int
	Value int
}

type Result struct {
	WorkerId int
	Result   int
	JobValue int
}

type WorkerPool struct {
	numJobs    int
	numWorkers int
	jobCh      chan Job
	resultCh   chan Result
	wg         sync.WaitGroup
}

func (wp *WorkerPool) createJobs() {
	defer close(wp.jobCh)

	for i := 1; i <= wp.numJobs; i++ {
		wp.jobCh <- Job{Id: i, Value: i}
	}
}

func (wp *WorkerPool) StartWorkers(workerId int) {
	defer wp.wg.Done()

	for job := range wp.jobCh {
		wp.resultCh <- Result{JobValue: job.Value, WorkerId: workerId, Result: job.Value * job.Value}
	}
}

func (wp *WorkerPool) CreateWorkers() {
	for i := 1; i <= wp.numWorkers; i++ {
		wp.wg.Add(1)
		go wp.StartWorkers(i)
	}
}

func main() {
	numJobs := 30
	numWorkers := 3
	var wg sync.WaitGroup

	wp := WorkerPool{
		numJobs:    numJobs,
		numWorkers: numWorkers,
		jobCh:      make(chan Job),
		resultCh:   make(chan Result),
	}

	wp.CreateWorkers()

	go func() {
		wp.createJobs()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for result := range wp.resultCh {
			fmt.Println("Worker Id", result.WorkerId, "Job Value", result.JobValue, "Result: ", result.Result)
		}
	}()

	wp.wg.Wait()
	close(wp.resultCh)
	wg.Wait()
}
