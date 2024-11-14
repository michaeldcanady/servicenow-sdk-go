package batchapi

import (
	"net/url"
	"testing"

	intBatch "github.com/RecoLabs/servicenow-sdk-go/batch-api/internal"
	"github.com/stretchr/testify/assert"
)

type mockHasBaseURL struct {
	baseURL string
}

func (m *mockHasBaseURL) GetBaseURL() string {
	if m == nil {
		return ""
	}
	return m.baseURL
}

func TestGetBaseURL(t *testing.T) {
	tests := []intBatch.Test[url.URL]{
		{
			Title:       "valid",
			Input:       &mockHasBaseURL{baseURL: "https://validurl.com"},
			ExpectedErr: nil,
			Expected: url.URL{
				Host:   "validurl.com",
				Scheme: "https",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			url, err := getBaseURL(test.Input.(hasBaseURL))
			assert.Equal(t, test.ExpectedErr, err)
			assert.Equal(t, test.Expected, *url)
		})
	}
}
