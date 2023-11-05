package attachmentapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttachmentRequestBuilderGetQueryParameters(t *testing.T) {

	expected := map[string]string{
		"sysparm_limit":  "1000",
		"sysparm_offset": "500",
		"sysparm_query":  "field1=value1",
	}

	params := AttachmentRequestBuilderGetQueryParameters{
		Limit:  1000,
		Offset: 500,
		Query:  "field1=value1",
	}

	actual, err := toQueryMap(params)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expected, actual)
}
