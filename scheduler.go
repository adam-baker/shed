package main

import (
	"fmt"
	"sync"
	"time"
)

type Scheduler struct {
	mu           sync.Mutex
	jobs         map[int]*Job
	jobQueue     chan *Job
	workers      map[int]*Worker
	workerStatus map[int]time.Time
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		jobs:         make(map[int]*Job),
		jobQueue:     make(chan *Job, 100),
		workers:      make(map[int]*Worker),
		workerStatus: make(map[int]time.Time),
	}
}

func (s *Scheduler) AddJob(job *Job) {
	s.mu.Lock()
	job.Status = Queued
	s.jobs[job.ID] = job
	s.mu.Unlock()

	s.jobQueue <- job
}

func (s *Scheduler) AssignJobs() {
	for job := range s.jobQueue {
		go func(j *Job) {
			for {
				s.mu.Lock()
				worker := s.selectHealthyWorker()
				if worker != nil {
					j.Status = InProgress
					j.WorkerID = worker.ID
					j.AssignedAt = time.Now()
					s.mu.Unlock()

					worker.jobChan <- j
					break
				}
				s.mu.Unlock()
				time.Sleep(time.Second)
			}
		}(job)
	}
}

func (s *Scheduler) selectHealthyWorker() *Worker {
	now := time.Now()
	for id, lastBeat := range s.workerStatus {
		if now.Sub(lastBeat) < 3*time.Second {
			return s.workers[id]
		}
	}
	return nil
}

func (s *Scheduler) MonitorWorkers() {
	for {
		time.Sleep(2 * time.Second)
		now := time.Now()

		s.mu.Lock()
		for id, lastBeat := range s.workerStatus {
			if now.Sub(lastBeat) > 5*time.Second {
				fmt.Printf("Worker %d failed!\n", id)
				for _, job := range s.jobs {
					if job.WorkerID == id && job.Status == InProgress {
						fmt.Printf("Requeueing job %d from worker %d\n", job.ID, id)
						job.Status = Queued
						job.WorkerID = 0
						s.jobQueue <- job
					}
				}
				delete(s.workerStatus, id)
			}
		}
		s.mu.Unlock()
	}
}
