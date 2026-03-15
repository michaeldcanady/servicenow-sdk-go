package servicenowsdkgo

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
)

func TestCdmRequestBuilder(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				url := "https://example.service-now.com/api"
				requestAdapter := mocking.NewMockRequestAdapter()
				builder := NewCdmRequestBuilder(url, requestAdapter)

				expected := map[string]string{
					"request-raw-url": url,
				}

				assert.NotNil(t, builder)
				assert.Equal(t, expected, builder.GetPathParameters())
				assert.Equal(t, cdmURLTemplate, builder.GetURLTemplate())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestCdmRequestBuilder_Policies(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				url := "https://example.service-now.com/api"
				requestAdapter := mocking.NewMockRequestAdapter()
				builder := NewCdmRequestBuilderInternal(map[string]string{"baseurl": url}, requestAdapter)

				policiesBuilder := builder.Policies()

				expected := map[string]string{
					"baseurl": url,
				}

				assert.NotNil(t, policiesBuilder)
				assert.Equal(t, expected, policiesBuilder.GetPathParameters())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
