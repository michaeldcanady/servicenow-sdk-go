package servicenowsdkgo

type NowRequestBuilder struct {
	RequestBuilder
}

// NewNowRequestBuilder creates a new instance of the NowRequestBuilder associated with the given URL and Client.
// It accepts the URL and Client as parameters and returns a pointer to the created NowRequestBuilder.
func NewNowRequestBuilder(url string, client *Client) *NowRequestBuilder {
	pathParameters := map[string]string{"baseurl": url}
	requestBuilder := NewRequestBuilder(client, "{+baseurl}/Now", pathParameters)
	return &NowRequestBuilder{
		*requestBuilder,
	}
}

// Table returns a TableRequestBuilder associated with the NowRequestBuilder.
// It accepts a table name as a parameter and constructs the URL for table-related requests.
// The returned TableRequestBuilder can be used to build and execute table-related requests.
func (N *NowRequestBuilder) Table(tableName string) *TableRequestBuilder {
	N.PathParameters["table"] = tableName
	return NewTableRequestBuilder(N.Client, N.PathParameters)
}
