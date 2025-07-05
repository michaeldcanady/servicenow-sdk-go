package credentials

import (
	"context"
	"errors"
	"testing"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
)

func TestNewRequestAdapterAdapter(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				mockCredential := newMockCredential()

				requestAdapter := NewRequestAdapterAdapter(mockCredential)

				assert.IsType(t, &RequestAdapterAdapter{}, requestAdapter)
				assert.Equal(t, mockCredential, requestAdapter.credential)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestRequestAdapterAdapter_AuthenticateRequest(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				mockCredential := newMockCredential()
				mockCredential.On("GetAuthentication").Return("auth string", nil)

				expectedHeader := abstractions.NewRequestHeaders()
				expectedHeader.Add("Authorization", "auth string")

				requestInformation := abstractions.NewRequestInformation()

				client := &RequestAdapterAdapter{credential: mockCredential}

				err := client.AuthenticateRequest(context.Background(), requestInformation, map[string]interface{}{})

				assert.Equal(t, expectedHeader, requestInformation.Headers)
				assert.Nil(t, err)
				mockCredential.AssertExpectations(t)
			},
		},
		{
			name: "Auth error",
			test: func(t *testing.T) {
				mockCredential := newMockCredential()
				mockCredential.On("GetAuthentication").Return("", errors.New("auth error"))

				expectedHeader := abstractions.NewRequestHeaders()
				expectedHeader.Add("Authorization", "auth string")

				requestInformation := abstractions.NewRequestInformation()

				client := &RequestAdapterAdapter{credential: mockCredential}

				err := client.AuthenticateRequest(context.Background(), requestInformation, map[string]interface{}{})

				assert.Equal(t, errors.New("auth error"), err)
				mockCredential.AssertExpectations(t)
			},
		},
		{
			name: "Nil adapter",
			test: func(t *testing.T) {
				mockCredential := (Credential)(nil)

				requestInformation := abstractions.NewRequestInformation()

				client := &RequestAdapterAdapter{credential: mockCredential}

				err := client.AuthenticateRequest(context.Background(), requestInformation, map[string]interface{}{})

				assert.Equal(t, errors.New("credential is nil"), err)
			},
		},
		{
			name: "Nil adapter",
			test: func(t *testing.T) {
				requestInformation := abstractions.NewRequestInformation()

				client := (*RequestAdapterAdapter)(nil)

				err := client.AuthenticateRequest(context.Background(), requestInformation, map[string]interface{}{})

				assert.Equal(t, errors.New("credential is nil"), err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
