package main

import "time"

type JobType string

const (
	ShortJob  JobType = "short"
	MediumJob JobType = "medium"
	LongJob   JobType = "long"
)

type JobStatus string

// define the core job struct and statuses
const (
	Queued     JobStatus = "queued"
	InProgress JobStatus = "in_progress"
	Completed  JobStatus = "completed"
	Failed     JobStatus = "failed"
)

type Job struct {
	ID         int       `json:"id"`
	JobType    JobType   `json:"job_type"`
	Payload    string    `json:"payload"`
	Status     JobStatus `json:"status"`
	WorkerID   int       `json:"worker_id"`
	AssignedAt time.Time `json:"assigned_at"`
}
