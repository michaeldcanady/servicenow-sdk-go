package servicenowsdkgo

type TableItemRequestBuilder struct {
	RequestBuilder
}

// NewTableItemRequestBuilder creates a new instance of the TableItemRequestBuilder associated with the given URL and Client.
// It accepts the URL and Client as parameters and returns a pointer to the created TableItemRequestBuilder.
func NewTableItemRequestBuilder(url string, client *Client) *TableItemRequestBuilder {
	requestBuilder := NewRequestBuilder(url, client)
	return &TableItemRequestBuilder{
		*requestBuilder,
	}
}
