package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"server/middleware"
)

type Health struct {
	Message string        `json:"message"`
	Date    string        `json:"date"`
	Uptime  time.Duration `json:"uptime"`
}

func main() {
	const PORT string = ":5000"
	startTime := time.Now()

	router := http.NewServeMux()

	// Setup logger
	logger := slog.New(slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	))

	slog.SetDefault(logger)

	// Setup routes
	router.HandleFunc("GET /", hello)
	router.HandleFunc("GET /health", healthCheckHandler)

	// Create middleware stack
	middlewareStack := middleware.CreateMiddlewareStack(
		middleware.LoggingMiddleware,
	)

	// Configure server
	server := &http.Server{
		Addr:        PORT,
		Handler:     middlewareStack(router),
		IdleTimeout: 3 * time.Minute,
		BaseContext: func(l net.Listener) context.Context {
			ctx := context.Background()
			ctx = context.WithValue(ctx, "startTime", startTime)
			return ctx
		},
	}

	// Start server in a goroutine
	go func() {
		slog.Info("Server started running on", slog.String("port", PORT))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Server failed to start", "error", err)
		}
	}()

	// Signal context for graceful shutdown
	signalCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	// Create a context with timeout for shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	select {
	case <-signalCtx.Done():
		slog.Info("Received shutdown signal")
	case <-shutdownCtx.Done():
		slog.Info("Shutdown timeout reached")
	}

	// Attempt graceful shutdown
	if err := server.Shutdown(shutdownCtx); err != nil {
		slog.Error("Server forced to shutdown", "error", err)
	}

	slog.Info("Server exiting")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("health route")

	startTime := r.Context().Value("startTime").(time.Time)

	health := Health{
		Message: "Server is running fine",
		Date:    time.Now().Format("01-02-2006 Monday 15:04:05"),
		Uptime:  time.Since(startTime),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(health); err != nil {
		slog.Error("Failed to encode health response", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
