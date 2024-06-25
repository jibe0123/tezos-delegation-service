package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"technical-test/internal/api"
	"technical-test/internal/app"
	"technical-test/pkg/config"
	database "technical-test/pkg/storage"
	tezos "technical-test/pkg/tzkt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Build the DSN for MariaDB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	// Initialize the MariaDB database with retry
	var db database.Database
	var err error
	for i := 0; i < 10; i++ {
		db, err = database.NewMariaDB(dsn)
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database: %v", err)
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()
	fmt.Println("Database initialized successfully 💾")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tezosClient := tezos.NewClient(cfg.TezosAPIBaseURL)
	newApp := app.NewApp(db, tezosClient) // Use db.DB() to get *sql.DB

	go newApp.StartPolling(ctx)

	// Set up the Gin router using your existing NewRouter function
	router := api.NewRouter(newApp)

	// Create the HTTP server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-quit
		log.Println("Shutting down server... 👋")

		ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctxShutdown); err != nil {
			log.Fatalf("Server Shutdown Failed: %v", err)
		}
		log.Println("Server exited properly ✅")
	}()

	// Start the server using Gin
	fmt.Println("Server is starting on port 8080 🚀")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to run server: %v", err)
	}
}
