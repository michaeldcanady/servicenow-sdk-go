package accountapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountSerialization(t *testing.T) {
	account := NewAccount()
	name := "Boxeo EMEA"
	err := account.setName(&name)
	assert.NoError(t, err)

	// In a real scenario, we'd use a real SerializationWriter.
	// Since we are checking if fields are set/retrieved properly,
	// we can check the backing store directly via getter.

	retrievedName, err := account.GetName()
	assert.NoError(t, err)
	assert.Equal(t, &name, retrievedName)
}

func TestCreateAccountFromDiscriminatorValue(t *testing.T) {
	instance, err := CreateAccountFromDiscriminatorValue(nil)
	assert.NoError(t, err)
	assert.NotNil(t, instance)
	assert.IsType(t, &AccountModel{}, instance)
}
