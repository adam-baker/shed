package main

import (
	"fmt"
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
		time.Sleep(2 * time.Second)
		w.scheduler.mu.Lock()
		job.Status = Completed
		fmt.Printf("Worker %d completed job %d\n", w.ID, job.ID)
		w.scheduler.mu.Unlock()
	}
}
