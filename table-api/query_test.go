package tableapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/stretchr/testify/assert"
)

func TestNewQuery(t *testing.T) {
	query := NewQuery()

	assert.IsType(t, &core.Query{}, query)
	assert.IsType(t, &Query{}, query)
}
