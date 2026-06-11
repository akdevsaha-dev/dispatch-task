package queue

import (
	"encoding/json"
	"time"
)

type JobStatus string

const (
	StatusPending    JobStatus = "pending"
	StatusRunning    JobStatus = "running"
	StatusCompleted  JobStatus = "completed"
	StatusFailed     JobStatus = "failed"
	StatusRetrying   JobStatus = "retrying"
	StatusDeadLetter JobStatus = "dead_letter"
)

type CreateJobRequest struct {
	Type         string          `json:"type"`
	Payload      json.RawMessage `json:"payload"`
	DelaySeconds int             `json:"delay_seconds,omitempty"` // Allows users to schedule jobs for later
}

type Job struct {
	Id         string          `json:"id"`
	Type       string          `json:"type"`
	Payload    json.RawMessage `json:"payload"`
	Status     JobStatus       `json:"job_status"`
	MaxRetries int             `json:"max_retries"`
	RetryCount int             `json:"retry_count"`
	CreatedAt  time.Time       `json:"created_at"`
}

type CreateJobResponse struct {
	JobID string `json:"job_id"`
}
