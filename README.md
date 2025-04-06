## The Shed

The Shed is a simple, fault-tolerant job scheduler used as a teaching tool for reliable software development in GoLang. It is designed to be easy to understand and modify, making it a great starting point for learning about job scheduling and fault tolerance.

### What makes this Fault Tolerant?

This implementation is fault-tolerant because it monitors worker health via regular heartbeats, clearly detects worker failures, and automatically requeues and retries jobs assigned to failed nodes. Centralized state tracking and idempotent job execution further ensure reliable, at-least-once processing, which is critical for infrastructure reliability.


#### ① Heartbeat-based Worker Health Monitoring:

Workers send periodic heartbeats (worker.go):

```go
func (w *Worker) heartbeat() {
    for {
        w.scheduler.mu.Lock()
        w.scheduler.workerStatus[w.ID] = time.Now()
        w.scheduler.mu.Unlock()
        time.Sleep(1 * time.Second)
    }
}
```

#### ② Automatic Job Reassignment on Worker Failure:

Upon detecting a worker failure, the scheduler requeues jobs assigned to the failed worker:

```go

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
```

#### ③ Centralized Job Management:

Jobs are maintained centrally within the scheduler, clearly tracking their status:

```go
type Scheduler struct {
    mu           sync.Mutex
    jobs         map[int]*Job
    jobQueue     chan *Job
    workers      map[int]*Worker
    workerStatus map[int]time.Time
}
```

Job states (Queued, InProgress, Completed) make it clear exactly where a job is at any given time.

#### ④ Idempotent Job Execution:

Jobs are intentionally designed to be idempotent (safe to retry multiple times).

This allows the system to retry jobs safely without unwanted side-effects.

#### Even Better Fault-Tolerance:

**Persistent Storage:** Jobs persisted in Redis/Kafka for stronger durability.

**Multiple Schedulers:** Leader election (Raft) for fault-tolerance in scheduler itself.

**Retries with Backoff:** Handling recurring failures gracefully.

**Enhanced Observability:** Clear logs, metrics, tracing for easier fault detection.

**Workload Logic Testing:** Unit tests for job logic, worker behavior, and scheduler functionality. 
