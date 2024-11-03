package tableapi

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/google/go-querystring/query"
)

func TestTableItemRequestBuilder2DeleteQueryParameters(T *testing.T) {
	tests := []struct {
		Title string
		Input TableItemRequestBuilder2DeleteQueryParameters
		Err   error
		Ret   interface{}
	}{
		{
			Title: "",
			Input: TableItemRequestBuilder2DeleteQueryParameters{
				SysparmQueryNoDomain: of(true),
			},
			Err: nil,
			Ret: url.Values{
				"sysparm_query_no_domain": []string{"true"},
			},
		},
	}

	for _, test := range tests {
		T.Run(test.Title, func(t *testing.T) {
			ret, err := query.Values(test.Input)
			assert.Equal(t, test.Err, err)
			assert.Equal(t, test.Ret, ret)
		})
	}
}
