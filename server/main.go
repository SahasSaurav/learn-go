package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
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

	logger := slog.New(slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	))

	slog.SetDefault(logger)

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})

	router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("health route")

		str := r.Context().Value("hello")
		fmt.Println("string with", str)

		health := Health{
			Message: "Server is running fine",
			Date:    time.Now().Format("01-02-2006 Monday 15:04:05"),
			Uptime:  time.Since(startTime),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(health); err != nil {
			slog.Error("Server is not working fine", "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	middlewareStack := middleware.CreateMiddlewareStack(
		middleware.LoggingMiddleware,
	)

	server := &http.Server{
		Addr:        PORT,
		Handler:     middlewareStack(router),
		IdleTimeout: 3 * time.Minute,
	}

	slog.Info("Server stared running on", slog.String("port", PORT))
	if err := server.ListenAndServe(); err != nil {
		slog.Error("Server failed to start", "error", err)
	}

}
