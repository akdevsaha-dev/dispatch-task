package main

import (
	"log"
	"net/http"

	"github.com/akdevsaha-dev/dispatch-task/cmd/producer/routes"
)

func main() {
	mux := http.NewServeMux()

	routes.RegisterJobRoute(mux)

	log.Print("Server started on http://localhost:8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("Server failed", err)
	}
}
