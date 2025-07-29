package router

import (
	"net/http"

	"github.com/codeandlearn1991/newsapi/internal/handler"
)

func New() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("POST /news", handler.PostNews())
	r.HandleFunc("GET /news", handler.GetAllNews())
	r.HandleFunc("GET /news/{id}", handler.GetNewsByID())
	r.HandleFunc("PUT /news/{id}", handler.UpdateNewsById())
	r.HandleFunc("DELETE /news/{id}", handler.DeleteNewsByID())

	return r
}