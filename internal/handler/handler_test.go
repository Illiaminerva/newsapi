package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/codeandlearn1991/newsapi/internal/handler"
)

func Test_PostNews(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus int
	}{
		{
			name: "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)

			// Act
			handler.PostNews()(w, r)

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
		expectedStatus int
	}{
		{
			name: "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)

			// Act
			handler.GetAllNews()(w, r)

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
		expectedStatus int
	}{
		{
			name: "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)

			// Act
			handler.GetNewsByID()(w, r)

			// Assert
			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("expected status %d, got %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}

func Test_UpdateNewsById(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus int
	}{
		{
			name: "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)

			// Act
			handler.UpdateNewsById()(w, r)

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
		expectedStatus int
	}{
		{
			name: "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)

			// Act
			handler.DeleteNewsByID()(w, r)

			// Assert
			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("expected status %d, got %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}