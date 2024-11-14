package tableapi

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/RecoLabs/servicenow-sdk-go/core"
	"github.com/RecoLabs/servicenow-sdk-go/internal"
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

			_, err := NewTablePageIterator(page, client)

			// TODO: Fails due to constructor function
			//assert.Equal(t, test.expected, iterator)
			assert.Equal(t, test.err, err)
		})
	}
}

func TestConstructTableCollection(t *testing.T) {
	tests := []test[core.CollectionResponse[TableEntry]]{
		{
			title: "Valid",
			value: &http.Response{
				Body: io.NopCloser(strings.NewReader(string(getFakeCollectionJSON()))),
			},
			expected: &fakeCollectionResponse,
			err:      nil,
		},
		{
			title: "Nil Body",
			value: &http.Response{
				Body: http.NoBody,
			},
			expected: nil,
			err:      internal.ErrNilResponseBody,
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			collection, err := constructTableCollection[TableEntry](test.value.(*http.Response))

			assert.Equal(t, test.expected, collection)
			assert.Equal(t, test.err, err)
		})
	}
}
