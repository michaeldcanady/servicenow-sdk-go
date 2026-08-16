package appserviceapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateAppServiceActionResponseFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateAppServiceActionResponseFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}
