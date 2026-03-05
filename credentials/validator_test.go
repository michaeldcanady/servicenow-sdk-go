package credentials

import (
	"net/url"
	"testing"

	"github.com/microsoft/kiota-abstractions-go/authentication"
	"github.com/stretchr/testify/assert"
)

func TestValidator(t *testing.T) {
	//nolint: staticcheck // while being deprecated there doesn't seem to be a replacement
	v := authentication.NewAllowedHostsValidator([]string{"example.com"})
	u, _ := url.Parse("https://example.com/api")
	assert.True(t, v.IsUrlHostValid(u))

	u2, _ := url.Parse("https://example.com:443/api")
	assert.True(t, v.IsUrlHostValid(u2), "Should be valid with port if base matches")

	//nolint: staticcheck // while being deprecated there doesn't seem to be a replacement
	v2 := authentication.NewAllowedHostsValidator([]string{})
	assert.True(t, v2.IsUrlHostValid(u))
}
