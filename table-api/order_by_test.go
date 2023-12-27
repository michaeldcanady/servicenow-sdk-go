package tableapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/stretchr/testify/assert"
)

func TestNewOrderBy(t *testing.T) {
	orderBy := NewOrderBy()

	assert.IsType(t, &core.OrderBy{}, orderBy)
	assert.IsType(t, &OrderBy{}, orderBy)
}
