package tableapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/stretchr/testify/assert"
)

func TestNewFragment(t *testing.T) {
	fragment := NewFragment("field1", LessOrEqual, 5)
	assert.IsType(t, &core.Fragment{}, fragment)
	assert.IsType(t, &Fragment{}, fragment)
}
