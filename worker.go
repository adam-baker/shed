package main

import (
	"context"
	"fmt"
	"shed/workloads"
	"time"
)

type Worker struct {
	ID        int
	jobChan   chan *Job
	scheduler *Scheduler
}

func NewWorker(id int, scheduler *Scheduler) *Worker {
	return &Worker{
		ID:        id,
		jobChan:   make(chan *Job),
		scheduler: scheduler,
	}
}

func (w *Worker) Start() {
	go w.heartbeat()
	go w.run()
}

func (w *Worker) heartbeat() {
	for {
		w.scheduler.mu.Lock()
		w.scheduler.workerStatus[w.ID] = time.Now()
		w.scheduler.mu.Unlock()
		time.Sleep(1 * time.Second)
	}
}

func (w *Worker) run() {
	for job := range w.jobChan {
		fmt.Printf("Worker %d starting job %d\n", w.ID, job.ID)

		var workload workloads.Workload

		switch job.JobType {
		case ShortJob:
			workload = workloads.Short{}
		case MediumJob:
			workload = workloads.Medium{}
		case LongJob:
			workload = workloads.Long{}
		default:
			fmt.Printf("Unknown job type %s\n", job.JobType)
			continue
		}

		// Execute workload logic
		err := workload.Execute(context.Background(), job.Payload)
		if err != nil {
			fmt.Printf("Worker %d failed job %d: %v\n", w.ID, job.ID, err)
			continue
		}

		w.scheduler.mu.Lock()
		job.Status = Completed
		fmt.Printf("Worker %d completed job %d\n", w.ID, job.ID)
		w.scheduler.mu.Unlock()
	}
}
