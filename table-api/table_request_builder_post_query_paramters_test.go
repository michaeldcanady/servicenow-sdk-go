package tableapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTableRequestBuilderPostQueryParametersDefaults(t *testing.T) {
	params := &TableRequestBuilderPostQueryParameters{}

	assert.Equal(t, "", params.DisplayValue)
	assert.Equal(t, false, params.ExcludeReferenceLink)
	assert.Equal(t, 0, len(params.Fields))
	assert.Equal(t, "", params.View)
}

func TestTableRequestBuilderPostQueryParamtersDefaults(t *testing.T) {
	params := &TableRequestBuilderPostQueryParamters{}

	assert.Equal(t, "", params.DisplayValue)
	assert.Equal(t, false, params.ExcludeReferenceLink)
	assert.Equal(t, 0, len(params.Fields))
	assert.Equal(t, "", params.View)
}
