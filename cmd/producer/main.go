package main

import (
	"log"
	"net/http"

	"github.com/akdevsaha-dev/dispatch-task/cmd/producer/handlers"
	"github.com/akdevsaha-dev/dispatch-task/cmd/producer/routes"
	"github.com/akdevsaha-dev/dispatch-task/internal/queue"
	"github.com/akdevsaha-dev/dispatch-task/internal/storage"
)

func main() {

	rdb := storage.RedisClient()
	producer := queue.NewProducer(rdb)
	jobHandler := handlers.NewJobHandler(producer)

	mux := http.NewServeMux()

	routes.RegisterJobRoute(mux, jobHandler)

	log.Print("Server started on http://localhost:8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("Server failed", err)
	}
}
