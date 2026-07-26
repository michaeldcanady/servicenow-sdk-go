package accountapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateAccountCollectionResponseFromDiscriminatorValue(t *testing.T) {
	instance, err := CreateAccountCollectionResponseFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, instance)
}

func TestCreateAccountItemResponseFromDiscriminatorValue(t *testing.T) {
	instance, err := CreateAccountItemResponseFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, instance)
}
