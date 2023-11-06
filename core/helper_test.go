package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yosida95/uritemplate/v3"
)

func TestAddParametersWithOriginalNames(t *testing.T) {

	params := map[string]string{"var1": "val1", "var2": "val2"}
	normalizedNames := map[string]string{"var1": "VaR1", "var2": "vAr2"}

	values := addParametersWithOriginalNames(params, normalizedNames, nil)
	expected := uritemplate.Values{"VaR1": uritemplate.String("val1"), "vAr2": uritemplate.String("val2")}

	assert.Equal(t, expected, values)
}
