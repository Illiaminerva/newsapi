package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/codeandlearn1991/newsapi/internal/logger"
	"github.com/google/uuid"
)

type NewsStorer interface {
	// Create news from post request body.
	Create(NewsPostReqBody) (NewsPostReqBody, error)
	// FindByID news by its ID.
	FindByID(uuid.UUID) (NewsPostReqBody, error)
	// FindAll returns all the news in the store
	FindAll() ([]NewsPostReqBody, error)
	// DeleteByID deletes news by its ID.
	DeleteByID(uuid.UUID) error
	// Update updates news by its ID.
	Update(uuid.UUID, NewsPostReqBody) (NewsPostReqBody, error)
}

func PostNews(ns NewsStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.FromContext(r.Context())
		logger.Info("request received")
		var newsRequestBody NewsPostReqBody
		if err := json.NewDecoder(r.Body).Decode(&newsRequestBody); err != nil {
			logger.Error("failed to decode request body", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := newsRequestBody.Validate(); err != nil {
			logger.Error("failed to validate request body", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("validation failed: %s", err.Error())))
			return
		}

		if _, err := ns.Create(newsRequestBody); err != nil {
			logger.Error("failed to create news", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func GetAllNews(ns NewsStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.FromContext(r.Context())
		logger.Info("request received")

		news, err := ns.FindAll()
		if err != nil {
			logger.Error("failed to get all news", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(news); err != nil {
			logger.Error("failed to encode response", "error", err)
		}
	}
}

func GetNewsByID(ns NewsStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.FromContext(r.Context())
		logger.Info("request received")
		idStr := r.PathValue("id")
		if idStr == "" {
			logger.Error("missing id parameter")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("missing id parameter"))
			return
		}
		id, err := uuid.Parse(idStr)
		if err != nil {
			logger.Error("failed to parse id", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid id format"))
			return
		}
		news, err := ns.FindByID(id)
		if err != nil {
			logger.Error("failed to get news", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(news); err != nil {
			logger.Error("failed to encode response", "error", err)
		}
	}
}

func UpdateNewsById(ns NewsStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.FromContext(r.Context())
		logger.Info("request received")

		// Extract and validate ID from URL
		idStr := r.PathValue("id")
		if idStr == "" {
			logger.Error("missing id parameter")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("missing id parameter"))
			return
		}

		id, err := uuid.Parse(idStr)
		if err != nil {
			logger.Error("failed to parse id", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid id format"))
			return
		}

		// Parse and validate request body
		var newsRequestBody NewsPostReqBody
		if err := json.NewDecoder(r.Body).Decode(&newsRequestBody); err != nil {
			logger.Error("failed to decode request body", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := newsRequestBody.Validate(); err != nil {
			logger.Error("failed to validate request body", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("validation failed: %s", err.Error())))
			return
		}

		// Update news
		updatedNews, err := ns.Update(id, newsRequestBody)
		if err != nil {
			logger.Error("failed to update news", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(updatedNews); err != nil {
			logger.Error("failed to encode response", "error", err)
		}
	}
}

func DeleteNewsByID(ns NewsStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.FromContext(r.Context())
		logger.Info("request received")

		// Extract and validate ID from URL
		idStr := r.PathValue("id")
		if idStr == "" {
			logger.Error("missing id parameter")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("missing id parameter"))
			return
		}

		id, err := uuid.Parse(idStr)
		if err != nil {
			logger.Error("failed to parse id", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid id format"))
			return
		}

		// Delete news
		if err := ns.DeleteByID(id); err != nil {
			logger.Error("failed to delete news", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}