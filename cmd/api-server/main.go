package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/codeandlearn1991/newsapi/internal/router"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))
	logger.Info("server starting on port 8080")
	r := router.New()
	if err := http.ListenAndServe(":8080", r); err != nil {
		logger.Error("failed to start server", "error", err)
	}
}
