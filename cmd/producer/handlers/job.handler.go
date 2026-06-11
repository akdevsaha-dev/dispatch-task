package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
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
	Type    string          `json:"type"`
	Paylaod json.RawMessage `json:"payload"`
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

func JobHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateJobRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	job := Job{
		Id:         uuid.NewString(),
		Type:       req.Type,
		Payload:    req.Paylaod,
		Status:     StatusPending,
		MaxRetries: 3,
		RetryCount: 0,
		CreatedAt:  time.Now(),
	}

	data, err := json.Marshal(job)
	if err != nil {
		http.Error(w, "Filed to create job", http.StatusInternalServerError)
	}
	response := CreateJobResponse{
		JobID: job.Id,
	}
	json.NewEncoder(w).Encode(response)
}
