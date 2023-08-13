package servicenowsdkgo

type TableRequestBuilder struct {
	RequestBuilder
}

type TableCollectionResponse struct {
	Result []*TableEntry
}

type TableEntry map[string]interface{}

// NewTableRequestBuilder creates a new instance of the TableRequestBuilder associated with the given URL and Client.
// It accepts the URL and Client as parameters and returns a pointer to the created TableRequestBuilder.
func NewTableRequestBuilder(url string, client *Client) *TableRequestBuilder {
	requestBuilder := NewRequestBuilder(url, client)
	return &TableRequestBuilder{
		*requestBuilder,
	}
}

// ById returns a TableItemRequestBuilder for a specific record in the table.
// It accepts the sysId of the record as a parameter and constructs the URL for the record.
// The returned TableItemRequestBuilder can be used to build and execute requests for the specific record.
func (T *TableRequestBuilder) ById(sysId string) *TableItemRequestBuilder {
	return NewTableItemRequestBuilder(T.AppendSegment(sysId), T.Client)
}

// Get performs an HTTP GET request to the table URL using the Client's session.
// It retrieves a collection of records from the table and decodes the response into a TableCollectionResponse.
// It returns the TableCollectionResponse and any errors encountered during the request or decoding.
func (T *TableRequestBuilder) Get() (*TableCollectionResponse, error) {
	resp := &TableCollectionResponse{}
	err := T.Client.Get(T.Url, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
