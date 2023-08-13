package servicenowsdkgo

import "strings"

type RequestBuilder struct {
	Url    string
	Client *Client
}

// NewRequestBuilder creates a new instance of the RequestBuilder associated with the given URL and Client.
// It accepts the URL and Client as parameters and returns a pointer to the created RequestBuilder.
func NewRequestBuilder(url string, client *Client) *RequestBuilder {
	return &RequestBuilder{
		Url:    url,
		Client: client,
	}
}

// AppendSegment appends a URL segment to the existing URL of the RequestBuilder.
// It accepts a URL segment as a parameter and constructs a new URL by concatenating the existing URL and the segment.
// If the provided URL segment does not start with a "/", the function adds it before concatenation.
// It returns the resulting URL after appending the segment.
func (R *RequestBuilder) AppendSegment(urlSegment string) string {
	if !strings.HasPrefix(urlSegment, "/") {
		urlSegment = "/" + urlSegment
	}
	return R.Url + urlSegment
}
