package tableapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/stretchr/testify/assert"
)

func TestTableRequestBuilderGetQueryParameters(t *testing.T) {

	params := &TableRequestBuilderGetQueryParameters{
		DisplayValue:             TRUE,
		ExcludeReferenceLink:     false,
		Fields:                   []string{"field1", "field2"},
		QueryNoDomain:            true,
		View:                     DESKTOP,
		Limit:                    10,
		NoCount:                  false,
		Offset:                   90,
		Query:                    "fdafdsfdsfasdfsda",
		QueryCategory:            "adfsdfasdfasdf",
		SuppressPaginationHeader: true,
	}

	queryMap, err := core.ToQueryMap(params)
	if err != nil {
		t.Error(err)
	}

	expectedValue := map[string]string{"sysparm_suppress_pagination_header": "true", "sysparm_display_value": "true", "sysparm_exclude_reference_link": "false", "sysparm_fields": "field1,field2", "sysparm_limit": "10", "sysparm_no_count": "false", "sysparm_offset": "90", "sysparm_query": "fdafdsfdsfasdfsda", "sysparm_query_category": "adfsdfasdfasdf", "sysparm_query_no_domain": "true", "sysparm_view": "desktop"}

	assert.Equal(t, expectedValue, queryMap)
}
