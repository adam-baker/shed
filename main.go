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

	// Add jobs
	for j := 1; j <= 10; j++ {
		job := &Job{ID: j, Payload: fmt.Sprintf("Payload %d", j)}
		scheduler.AddJob(job)
		fmt.Printf("Job %d added\n", j)
	}

	select {}
}
