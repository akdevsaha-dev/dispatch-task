package routes

import (
	"net/http"

	"github.com/akdevsaha-dev/dispatch-task/cmd/producer/handlers"
)

func RegisterJobRoute(mux *http.ServeMux, h *handlers.JobHandler) {
	mux.HandleFunc("POST /jobs", h.CreateJob)
}
