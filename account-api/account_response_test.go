package accountapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccountCollectionResponseFromDiscriminatorValue(t *testing.T) {
	instance, err := CreateAccountCollectionResponseFromDiscriminatorValue(nil)
	assert.NoError(t, err)
	assert.NotNil(t, instance)
}

func TestCreateAccountItemResponseFromDiscriminatorValue(t *testing.T) {
	instance, err := CreateAccountItemResponseFromDiscriminatorValue(nil)
	assert.NoError(t, err)
	assert.NotNil(t, instance)
}
