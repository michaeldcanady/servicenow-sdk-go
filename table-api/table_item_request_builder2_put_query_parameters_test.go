package tableapi

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/google/go-querystring/query"
)

func TestTableItemRequestBuilder2PutQueryParameters(T *testing.T) {
	tests := []struct {
		Title string
		Input TableItemRequestBuilder2PutQueryParameters
		Err   error
		Ret   interface{}
	}{
		{
			Title: "",
			Input: TableItemRequestBuilder2PutQueryParameters{
				SysparmDisplayValue:         of(DisplayValue2All),
				SysparmExcludeReferenceLink: of(true),
			},
			Err: nil,
			Ret: url.Values{
				"sysparm_display_value":          []string{"all"},
				"sysparm_exclude_reference_link": []string{"true"},
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
