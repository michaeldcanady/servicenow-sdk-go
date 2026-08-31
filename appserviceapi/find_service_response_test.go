// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package appserviceapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateFindServiceResponseFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateFindServiceResponseFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}
