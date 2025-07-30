package handler_test

import (
	"testing"

	"github.com/codeandlearn1991/newsapi/internal/handler"
)

func TestNewsPostReqBody_Validate(t *testing.T) {
	testCases := []struct {
		name        string
		req         handler.NewsPostReqBody
		expectedErr bool
	}{
		{
			name: "author empty",
			req: handler.NewsPostReqBody{},
			expectedErr: true,
		},
		{
			name: "title empty",
			req: handler.NewsPostReqBody{
				Author: "test-author",
			},
			expectedErr: true,
		},
		{
			name: "summary empty",
			req: handler.NewsPostReqBody{
				Author: "test-author",
				Title: "test-title",
			},
			expectedErr: true,
		},
		{
			name: "created_at invalid",
			req: handler.NewsPostReqBody{
				Author: "test-author",
				Title: "test-title",
				Summary: "test-summary",
				CreatedAt: "invalid",
			},
			expectedErr: true,
		},
		{
			name: "source invalid",
			req: handler.NewsPostReqBody{
				Author: "test-author",
				Title: "test-title",
				Summary: "test-summary",
				CreatedAt: "2025-07-30T15:30:45Z",
				Source: "invalid",
			},
			expectedErr: true,
		},
		{
			name: "tags empty",
			req: handler.NewsPostReqBody{
				Author: "test-author",
				Title: "test-title",
				Summary: "test-summary",
				CreatedAt: "2025-07-30T15:30:45Z",
				Source: "https://example.com",
			},
			expectedErr: true,
		},
		{
			name: "validate",
			req: handler.NewsPostReqBody{
				Author: "test-author",
				Title: "test-title",
				Summary: "test-summary",
				CreatedAt: "2025-07-30T15:30:45Z",
				Source: "https://example.com",
				Tags: []string{"tag1", "tag2"},
			},
			expectedErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.req.Validate()

			if tc.expectedErr && err == nil {
				t.Fatalf("expected error but got nil")
			} else if !tc.expectedErr && err != nil {
				t.Fatalf("expected nil but got error: %v", err)
			}
		})
	}
}
