package internal

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
)

func TestServiceNowCollectionResponseFromDiscriminatorValue(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Valid collection response",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			strct := mocking.NewMockParsableFactory()
			parsableFactory := strct.Factory
			parseNode := mocking.NewMockParseNode()

			factory := ServiceNowCollectionResponseFromDiscriminatorValue[*mocking.MockParsable](parsableFactory)
			parsable, err := factory(parseNode)

			assert.Nil(t, err)
			assert.IsType(t, &BaseServiceNowCollectionResponse[*mocking.MockParsable]{}, parsable)
		})
	}
}

func TestServiceNowItemResponseFromDiscriminatorValue(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Valid item response",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			strct := mocking.NewMockParsableFactory()
			parsableFactory := strct.Factory
			parseNode := mocking.NewMockParseNode()

			factory := ServiceNowItemResponseFromDiscriminatorValue[*mocking.MockParsable](parsableFactory)
			parsable, err := factory(parseNode)

			assert.Nil(t, err)
			assert.IsType(t, &BaseServiceNowItemResponse[*mocking.MockParsable]{}, parsable)
		})
	}
}
