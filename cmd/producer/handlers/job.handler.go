package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/akdevsaha-dev/dispatch-task/internal/queue"
)

type JobHandler struct {
	producer *queue.Producer
}

func NewJobHandler(p *queue.Producer) *JobHandler {
	return &JobHandler{producer: p}
}
func (h *JobHandler) CreateJob(w http.ResponseWriter, r *http.Request) {
	var req queue.CreateJobRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	resp, err := h.producer.Enqueue(r.Context(), req)
	if err != nil {
		http.Error(w, "failed to enqueue job", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(resp)
}
