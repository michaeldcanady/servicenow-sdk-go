package abstraction

import "net/http"

type Client interface {
	Send(requestInfo *RequestInformation, errorMapping ErrorMapping) (*http.Response, error)
}
