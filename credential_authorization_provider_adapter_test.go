package servicenowsdkgo

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
)

func TestNewCredentialAuthenticationProviderAdapter(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "successful",
			test: func(t *testing.T) {
				cred := mocking.NewMockCredential()

				authenticationProvider, err := newCredentialAuthenticationProviderAdapter(cred)
				assert.Nil(t, err)
				assert.Equal(t, &credentialAuthenticationProviderAdapter{
					cred: cred,
				}, authenticationProvider)
			},
		},
		{
			name: "nil credential",
			test: func(t *testing.T) {
				authenticationProvider, err := newCredentialAuthenticationProviderAdapter(nil)
				assert.Equal(t, errors.New("credential is nil"), err)
				assert.Nil(t, authenticationProvider)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestCredentialAuthenticationProviderAdapter_AuthenticateRequest(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "successful",
			test: func(t *testing.T) {
				credential := mocking.NewMockCredential()
				credential.On("GetAuthentication").Return("authentication", nil)

				ctx := mocking.NewMockContext()
				requestInformation := abstractions.NewRequestInformation()
				additionalInfo := map[string]interface{}{}

				requestAdapter := &credentialAuthenticationProviderAdapter{
					cred: credential,
				}

				err := requestAdapter.AuthenticateRequest(ctx, requestInformation, additionalInfo)

				assert.Nil(t, err)
				assert.Equal(t, []string{"authentication"}, requestInformation.Headers.Get("Authorization"))
			},
		},
		{
			name: "nil provider",
			test: func(t *testing.T) {
				credential := mocking.NewMockCredential()
				credential.On("GetAuthentication").Return("authentication", nil)

				ctx := mocking.NewMockContext()
				requestInformation := abstractions.NewRequestInformation()
				additionalInfo := map[string]interface{}{}

				requestAdapter := (*credentialAuthenticationProviderAdapter)(nil)

				err := requestAdapter.AuthenticateRequest(ctx, requestInformation, additionalInfo)

				assert.Nil(t, err)
			},
		},
		{
			name: "nil request",
			test: func(t *testing.T) {
				credential := mocking.NewMockCredential()
				ctx := mocking.NewMockContext()
				additionalInfo := map[string]interface{}{}

				requestAdapter := &credentialAuthenticationProviderAdapter{
					cred: credential,
				}

				err := requestAdapter.AuthenticateRequest(ctx, nil, additionalInfo)

				assert.Equal(t, errors.New("request is nil"), err)
			},
		},
		{
			name: "nil headers",
			test: func(t *testing.T) {
				credential := mocking.NewMockCredential()
				credential.On("GetAuthentication").Return("authentication", nil)

				ctx := mocking.NewMockContext()
				requestInformation := &abstractions.RequestInformation{}
				additionalInfo := map[string]interface{}{}

				requestAdapter := &credentialAuthenticationProviderAdapter{
					cred: credential,
				}

				err := requestAdapter.AuthenticateRequest(ctx, requestInformation, additionalInfo)

				assert.Nil(t, err)
				assert.Equal(t, []string{"authentication"}, requestInformation.Headers.Get("Authorization"))
			},
		},
		{
			name: "authorization error",
			test: func(t *testing.T) {
				credential := mocking.NewMockCredential()
				credential.On("GetAuthentication").Return("", errors.New("bad auth"))

				ctx := mocking.NewMockContext()
				requestInformation := abstractions.NewRequestInformation()
				additionalInfo := map[string]interface{}{}

				requestAdapter := &credentialAuthenticationProviderAdapter{
					cred: credential,
				}

				err := requestAdapter.AuthenticateRequest(ctx, requestInformation, additionalInfo)

				assert.Equal(t, errors.New("bad auth"), err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
