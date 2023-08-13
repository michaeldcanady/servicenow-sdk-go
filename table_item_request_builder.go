package servicenowsdkgo

type TableItemRequestBuilder struct {
	RequestBuilder
}

func NewTableItemRequestBuilder(url string, client *Client) *TableItemRequestBuilder {
	requestBuilder := NewRequestBuilder(url, client)
	return &TableItemRequestBuilder{
		*requestBuilder,
	}
}
