package core

import (
	"net/http"
)

type Client interface {
	Send(requestInfo IRequestInformation, errorMapping ErrorMapping) (*http.Response, error)
}
