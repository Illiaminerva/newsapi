package handler_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/codeandlearn1991/newsapi/internal/handler"
	"github.com/google/uuid"
)

func Test_PostNews(t *testing.T) {
	testCases := []struct {
		name           string
		body           io.Reader
		store          handler.NewsStorer
		expectedStatus int
	}{
		{
			name: "invalid request body json",
			body: strings.NewReader("invalid"),
			store: mockNewsStore{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "invalid request body",
			body: strings.NewReader(`{"id": "3332388237", "author": "test-author", "title": "test-title", "summary": "test-summary", "created_at": "2025-07-30T15:30:45Z", "source": "https://example.com"}`),
			store: mockNewsStore{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "db error",
			body: strings.NewReader(`{"author": "test-author", "title": "test-title", "summary": "test-summary", "created_at": "2025-07-30T15:30:45Z", "source": "https://example.com", "tags": ["tag1", "tag2"]}`),
			store: mockNewsStore{errState: true},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name: "success",
			body: strings.NewReader(`{"author": "test-author", "title": "test-title", "summary": "test-summary", "created_at": "2025-07-30T15:30:45Z", "source": "https://example.com", "tags": ["tag1", "tag2"]}`),
			store: mockNewsStore{},
			expectedStatus: http.StatusCreated,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", tc.body)

			// Act
			handler.PostNews(tc.store)(w, r)

			// Assert
			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("expected status %d, got %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}

func Test_GetAllNews(t *testing.T) {
	testCases := []struct {
		name           string
		store          handler.NewsStorer
		expectedStatus int
	}{
		{
			name: "db error",
			store: mockNewsStore{errState: true},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name: "success",
			store: mockNewsStore{},
			expectedStatus: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)

			// Act
			handler.GetAllNews(tc.store)(w, r)

			// Assert
			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("expected status %d, got %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}

func Test_GetNewsByID(t *testing.T) {
	testCases := []struct {
		name           string
		url            string
		store          handler.NewsStorer
		expectedStatus int
	}{
		{
			name:           "missing id parameter",
			url:            "/news/",
			store:          mockNewsStore{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "invalid id format",
			url:            "/news/invalid-uuid",
			store:          mockNewsStore{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "db error",
			url:            "/news/123e4567-e89b-12d3-a456-426614174000",
			store:          mockNewsStore{errState: true},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "success",
			url:            "/news/123e4567-e89b-12d3-a456-426614174000",
			store:          mockNewsStore{},
			expectedStatus: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, tc.url, nil)
			r.SetPathValue("id", extractIDFromURL(tc.url))

			// Act
			handler.GetNewsByID(tc.store)(w, r)

			// Assert
			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("expected status %d, got %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}

// Helper function to extract ID from URL for testing
func extractIDFromURL(url string) string {
	parts := strings.Split(url, "/")
	if len(parts) >= 3 {
		return parts[2]
	}
	return ""
}

func Test_UpdateNewsById(t *testing.T) {
	testCases := []struct {
		name           string
		url            string
		body           io.Reader
		store          handler.NewsStorer
		expectedStatus int
	}{
		{
			name:           "missing id parameter",
			url:            "/news/",
			body:           strings.NewReader(`{"author": "test-author", "title": "test-title", "summary": "test-summary", "created_at": "2025-07-30T15:30:45Z", "source": "https://example.com", "tags": ["tag1", "tag2"]}`),
			store:          mockNewsStore{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "invalid id format",
			url:            "/news/invalid-uuid",
			body:           strings.NewReader(`{"author": "test-author", "title": "test-title", "summary": "test-summary", "created_at": "2025-07-30T15:30:45Z", "source": "https://example.com", "tags": ["tag1", "tag2"]}`),
			store:          mockNewsStore{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "invalid request body json",
			url:            "/news/123e4567-e89b-12d3-a456-426614174000",
			body:           strings.NewReader("invalid"),
			store:          mockNewsStore{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "invalid request body validation",
			url:            "/news/123e4567-e89b-12d3-a456-426614174000",
			body:           strings.NewReader(`{"author": "test-author", "title": "test-title", "summary": "test-summary", "created_at": "2025-07-30T15:30:45Z", "source": "https://example.com"}`),
			store:          mockNewsStore{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "db error",
			url:            "/news/123e4567-e89b-12d3-a456-426614174000",
			body:           strings.NewReader(`{"author": "test-author", "title": "test-title", "summary": "test-summary", "created_at": "2025-07-30T15:30:45Z", "source": "https://example.com", "tags": ["tag1", "tag2"]}`),
			store:          mockNewsStore{errState: true},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "success",
			url:            "/news/123e4567-e89b-12d3-a456-426614174000",
			body:           strings.NewReader(`{"author": "test-author", "title": "test-title", "summary": "test-summary", "created_at": "2025-07-30T15:30:45Z", "source": "https://example.com", "tags": ["tag1", "tag2"]}`),
			store:          mockNewsStore{},
			expectedStatus: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPut, tc.url, tc.body)
			r.SetPathValue("id", extractIDFromURL(tc.url))

			// Act
			handler.UpdateNewsById(tc.store)(w, r)

			// Assert
			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("expected status %d, got %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}

func Test_DeleteNewsByID(t *testing.T) {
	testCases := []struct {
		name           string
		url            string
		store          handler.NewsStorer
		expectedStatus int
	}{
		{
			name:           "missing id parameter",
			url:            "/news/",
			store:          mockNewsStore{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "invalid id format",
			url:            "/news/invalid-uuid",
			store:          mockNewsStore{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "db error",
			url:            "/news/123e4567-e89b-12d3-a456-426614174000",
			store:          mockNewsStore{errState: true},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "success",
			url:            "/news/123e4567-e89b-12d3-a456-426614174000",
			store:          mockNewsStore{},
			expectedStatus: http.StatusNoContent,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodDelete, tc.url, nil)
			r.SetPathValue("id", extractIDFromURL(tc.url))

			// Act
			handler.DeleteNewsByID(tc.store)(w, r)

			// Assert
			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("expected status %d, got %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}

type mockNewsStore struct {
	errState bool
}

func (m mockNewsStore) Create(handler.NewsPostReqBody) (news handler.NewsPostReqBody, err error) {
	if m.errState {
		return news, errors.New("error")
	}
	return news, nil
}

func (m mockNewsStore) FindByID(_ uuid.UUID) (news handler.NewsPostReqBody, err error) {
	if m.errState {
		return news, errors.New("error")
	}
	return news, nil
}

func (m mockNewsStore) FindAll() (news []handler.NewsPostReqBody, err error) {
	if m.errState {
		return news, errors.New("error")
	}
	return news, nil
}

func (m mockNewsStore) DeleteByID(_ uuid.UUID) error {
	if m.errState {
		return errors.New("error")
	}
	return nil
}

func (m mockNewsStore) Update(_ uuid.UUID, req handler.NewsPostReqBody) (handler.NewsPostReqBody, error) {
	if m.errState {
		return handler.NewsPostReqBody{}, errors.New("error")
	}
	return req, nil
}