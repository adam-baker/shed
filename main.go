package main

import (
	"fmt"
)

func main() {
	scheduler := NewScheduler()

	// Create workers
	for i := 1; i <= 3; i++ {
		worker := NewWorker(i, scheduler)
		scheduler.mu.Lock()
		scheduler.workers[i] = worker
		scheduler.mu.Unlock()
		worker.Start()
		fmt.Printf("Worker %d started\n", i)
	}

	// Scheduler routines
	go scheduler.AssignJobs()
	go scheduler.MonitorWorkers()

	jobs := []*Job{
		{ID: 1, Payload: "Job 1 payload", JobType: ShortJob},
		{ID: 2, Payload: "Job 2 payload", JobType: MediumJob},
		{ID: 3, Payload: "Job 3 payload", JobType: LongJob},
		{ID: 4, Payload: "Job 4 payload", JobType: ShortJob},
		{ID: 5, Payload: "Job 5 payload", JobType: MediumJob},
	}

	// Add jobs
	for _, job := range jobs {
		scheduler.AddJob(job)
		fmt.Printf("Job %d added\n", job.ID)
	}

	select {}
}
