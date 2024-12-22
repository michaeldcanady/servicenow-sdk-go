package core

import (
	"net/http"
)

// Deprecated: deprecated since v{unreleased}.
type Client interface {
	Send(requestInfo IRequestInformation, errorMapping ErrorMapping) (*http.Response, error)
}
