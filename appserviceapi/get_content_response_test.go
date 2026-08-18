package appserviceapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateGetContentResponseFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateGetContentResponseFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}
