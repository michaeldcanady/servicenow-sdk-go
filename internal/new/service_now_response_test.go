package internal

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
)

func TestServiceNowResponseFromDiscriminatorValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Collection response",
			test: func(t *testing.T) {
				strct := mocking.NewMockParsableFactory()
				parsableFactory := strct.Factory

				parseNode := mocking.NewMockParseNode()
				parseNode.On("GetRawValue").Return(any([]*mocking.MockParsable{}), nil)

				factory := ServiceNowResponseFromDiscriminatorValue[*mocking.MockParsable](parsableFactory)
				parsable, err := factory(parseNode)
				assert.Nil(t, err)
				assert.IsType(t, &BaseServiceNowCollectionResponse[*mocking.MockParsable]{}, parsable)
			},
		},
		{
			name: "Item response",
			test: func(t *testing.T) {
				strct := mocking.NewMockParsableFactory()
				parsableFactory := strct.Factory

				parseNode := mocking.NewMockParseNode()
				parseNode.On("GetRawValue").Return(any(&mocking.MockParsable{}), nil)

				factory := ServiceNowResponseFromDiscriminatorValue[*mocking.MockParsable](parsableFactory)
				parsable, err := factory(parseNode)
				assert.Nil(t, err)
				assert.IsType(t, &BaseServiceNowItemResponse[*mocking.MockParsable]{}, parsable)
			},
		},
		{
			name: "GetRawValue error",
			test: func(t *testing.T) {
				strct := mocking.NewMockParsableFactory()
				parsableFactory := strct.Factory

				parseNode := mocking.NewMockParseNode()
				parseNode.On("GetRawValue").Return(nil, errors.New("retrieval error"))

				factory := ServiceNowResponseFromDiscriminatorValue[*mocking.MockParsable](parsableFactory)
				parsable, err := factory(parseNode)
				assert.Equal(t, errors.New("retrieval error"), err)
				assert.Nil(t, parsable)
			},
		},
		{
			name: "Unsupported kind",
			test: func(t *testing.T) {
				strct := mocking.NewMockParsableFactory()
				parsableFactory := strct.Factory

				parseNode := mocking.NewMockParseNode()
				parseNode.On("GetRawValue").Return(any(true), nil)

				factory := ServiceNowResponseFromDiscriminatorValue[*mocking.MockParsable](parsableFactory)
				parsable, err := factory(parseNode)
				assert.Equal(t, fmt.Errorf("unsupported type: %s", reflect.Bool), err)
				assert.Nil(t, parsable)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
