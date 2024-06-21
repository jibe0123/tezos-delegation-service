package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"technical-test/internal/api"
	"technical-test/internal/repository"
	"technical-test/internal/service"
	"technical-test/internal/sync"
	"technical-test/pkg/tzkt"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	baseURL := os.Getenv("TZKT_API_BASE_URL")
	if baseURL == "" {
		log.Fatalf("TZKT_API_BASE_URL environment variable is not set")
	}

	repo := repository.NewMemoryRepository()
	svc := service.NewDelegationService(repo)
	handler := api.NewHandler(svc)

	client := tzkt.NewClient("https://api.ghostnet.tzkt.io")
	poller := sync.NewPoller(client, svc)

	go poller.StartPolling()

	http.HandleFunc("/xtz/delegations", handler.GetDelegations)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
