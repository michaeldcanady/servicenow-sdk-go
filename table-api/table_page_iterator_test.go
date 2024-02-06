package tableapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/stretchr/testify/assert"
)

var (
	sharedTableCollectionResponse2 = TableCollectionResponse2[TableEntry]{
		// Populate with test data.
	}
	sharedClient      = &MockClient{}
	sharedIterator, _ = NewTablePageIterator[TableEntry](&sharedTableCollectionResponse2, sharedClient)
)

func TestNewTablePageIterator(t *testing.T) {
	tests := []test[*TablePageIterator[TableEntry]]{
		{
			title: "Valid",
			value: []interface{}{
				&sharedTableCollectionResponse2,
				sharedClient,
			},
			expected:  sharedIterator,
			expectErr: false,
			err:       nil,
		},
		{
			title: "Missing Client",
			value: []interface{}{
				&sharedTableCollectionResponse2,
				(*MockClient)(nil),
			},
			expected:  nil,
			expectErr: false,
			err:       core.ErrNilClient,
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			value, _ := test.value.([]interface{})

			page := value[0].(*TableCollectionResponse2[TableEntry])
			client := value[1].(*MockClient)

			iterator, err := NewTablePageIterator(page, client)

			assert.Equal(t, test.expected, iterator)
			assert.Equal(t, test.err, err)
		})
	}
}
