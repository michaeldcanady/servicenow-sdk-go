package internal

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	sharedCredential                = MockCredential{}
	sharedBaseAuthorizationProvider = BaseAuthorizationProvider{credential: &sharedCredential}
)

type MockCredential struct {
}

func (c *MockCredential) GetAuthentication() (string, error) {
	return "Bearer dfasdfdsfd", nil
}

func TestNewBaseAuthorizationProvider(t *testing.T) {
	tests := []Test[*BaseAuthorizationProvider]{
		{
			Title:       "Valid",
			Input:       &sharedCredential,
			Expected:    &sharedBaseAuthorizationProvider,
			expectedErr: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			provider, err := NewBaseAuthorizationProvider(test.Input.(Credential))
			assert.Equal(t, test.Expected, provider)
			assert.Equal(t, test.expectedErr, err)
		})
	}
}

func TestBaseAuthorizationProvider_AuthorizeRequest(t *testing.T) {
	tests := []Test[*MockRequestInformation]{
		{
			Title: "Valid",
			Input: &MockRequestInformation{
				Headers: http.Header{},
			},
			Expected: &MockRequestInformation{
				Headers: http.Header{
					authorizationHeader: []string{"Bearer dfasdfdsfd"},
				},
			},
			expectedErr: nil,
		},
		{
			Title:       "Nil Request",
			Input:       (*MockRequestInformation)(nil),
			Expected:    nil,
			expectedErr: ErrNilRequest,
		},
	}
	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			err := sharedBaseAuthorizationProvider.AuthorizeRequest(test.Input.(RequestInformation))
			assert.Equal(t, test.Expected, test.Input)
			assert.Equal(t, test.expectedErr, err)
		})
	}
}
