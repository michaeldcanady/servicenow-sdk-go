package internal

import (
	"net/http"
	"testing"

	"github.com/RecoLabs/servicenow-sdk-go/internal/mocking"
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
	tests := []mocking.Test[*BaseAuthorizationProvider]{
		{
			Title:       "Valid",
			Input:       &sharedCredential,
			Expected:    &sharedBaseAuthorizationProvider,
			ExpectedErr: nil,
		},
		{
			Title:       "Nil Credential",
			Input:       (*MockCredential)(nil),
			Expected:    nil,
			ExpectedErr: ErrNilCredential,
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			provider, err := NewBaseAuthorizationProvider(test.Input.(Credential))
			assert.Equal(t, test.Expected, provider)
			assert.Equal(t, test.ExpectedErr, err)
		})
	}
}

func TestBaseAuthorizationProvider_AuthorizeRequest(t *testing.T) {
	reqInfo := &mocking.MockRequestInformation{}

	tests := []mocking.Test[*mocking.MockRequestInformation]{
		{
			Title: "Valid",
			Setup: func() {
				reqInfo.On("AddHeaders", http.Header{authorizationHeader: {"Bearer dfasdfdsfd"}}).Return(nil)
			},
			Input:       reqInfo,
			Expected:    nil,
			ExpectedErr: nil,
		},
		{
			Title:       "Nil Request",
			Input:       (*mocking.MockRequestInformation)(nil),
			Expected:    nil,
			ExpectedErr: ErrNilRequest,
		},
	}
	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			if test.Setup != nil {
				test.Setup()
			}

			_ = sharedBaseAuthorizationProvider.AuthorizeRequest(test.Input.(RequestInformation))
			//assert.Equal(t, test.Expected, test.Input)
			//assert.Equal(t, test.ExpectedErr, err)

			if test.Cleanup != nil {
				test.Cleanup()
			}
		})
	}
}
