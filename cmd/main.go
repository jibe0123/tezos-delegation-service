package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"technical-test/internal/api"
	"technical-test/internal/app"
	"technical-test/pkg/config"
	database "technical-test/pkg/sqlite"
	tezos "technical-test/pkg/tzkt"
	"time"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	db, err := database.InitDB(cfg.DatabasePath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("Database initialized successfully 💾")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tezosClient := tezos.NewClient(cfg.TezosAPIBaseURL)
	newApp := app.NewApp(db, tezosClient, ctx)
	go newApp.StartPolling()

	router := api.NewRouter(newApp)

	// explicit host to avoid warning message
	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: router,
	}

	go func() {
		log.Println("Server is starting on port 8080 🚀")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server... 👋")

	ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctxShutdown); err != nil {
		log.Fatalf("Server Shutdown Failed:%+s", err)
	}
	log.Println("Server exited properly ✅")
}
