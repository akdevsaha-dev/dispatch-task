package routes

import (
	"net/http"

	"github.com/akdevsaha-dev/dispatch-task/cmd/producer/handlers"
)

func RegisterJobRoute(mux *http.ServeMux) {

	jobHandler := http.HandlerFunc(handlers.JobHandler)
	mux.Handle("/add-job", jobHandler)
}
