package tableapi

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
)

func TestNewTablePageIterator(t *testing.T) {
	client := &MockClient{}
	resp := &TableCollectionResponse2[TableEntry]{}
	tests := []struct {
		name   string
		resp   *TableCollectionResponse2[TableEntry]
		client core.Client
		err    error
	}{
		{"Ok", resp, client, nil},
		{"NilClient", resp, nil, core.ErrNilClient},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewTablePageIterator(tt.resp, tt.client)
			if err != tt.err {
				t.Errorf("got err %v, expected %v", err, tt.err)
			}
		})
	}
}

func TestConstructTableCollection(t *testing.T) {
	tests := []struct {
		name  string
		input *http.Response
		err   error
	}{
		{"Ok", &http.Response{Body: io.NopCloser(strings.NewReader(string(getFakeCollectionJSON())))}, nil},
		{"NilBody", &http.Response{Body: http.NoBody}, internal.ErrNilResponseBody},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := constructTableCollection[TableEntry](tt.input)
			if err != tt.err {
				t.Errorf("got err %v, expected %v", err, tt.err)
			}
		})
	}
}
