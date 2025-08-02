package router

import (
	"net/http"

	"github.com/codeandlearn1991/newsapi/internal/handler"
)

func New(ns handler.NewsStorer) *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("POST /news", handler.PostNews(ns))
	r.HandleFunc("GET /news", handler.GetAllNews(ns))
	r.HandleFunc("GET /news/{id}", handler.GetNewsByID(ns))
	r.HandleFunc("PUT /news/{id}", handler.UpdateNewsById())
	r.HandleFunc("DELETE /news/{id}", handler.DeleteNewsByID())

	return r
}