package tableapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderByNew(t *testing.T) {

	orderBy := NewOrderBy()

	assert.IsType(t, &OrderBy{}, orderBy)
	assert.IsType(t, &OrderBy{}, orderBy)
}
