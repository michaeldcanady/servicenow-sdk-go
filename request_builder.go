package servicenowsdkgo

type RequestBuilder struct {
	PathParameters map[string]string
	Client         *Client
	UrlTemplate    string
}

// NewRequestBuilder creates a new instance of the RequestBuilder associated with the given URL and Client.
// It accepts the URL and Client as parameters and returns a pointer to the created RequestBuilder.
func NewRequestBuilder(client *Client, urlTemplate string, pathParameters map[string]string) *RequestBuilder {
	return &RequestBuilder{
		Client:         client,
		UrlTemplate:    urlTemplate,
		PathParameters: pathParameters,
	}
}
