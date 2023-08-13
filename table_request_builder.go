package servicenowsdkgo

type TableRequestBuilder struct {
	RequestBuilder
}

type TableCollectionResponse struct {
	Result []*TableEntry
}

type TableEntry map[string]interface{}

func NewTableRequestBuilder(url string, client *Client) *TableRequestBuilder {
	requestBuilder := NewRequestBuilder(url, client)
	return &TableRequestBuilder{
		*requestBuilder,
	}
}

func (T *TableRequestBuilder) ById(sysId string) *TableItemRequestBuilder {
	return NewTableItemRequestBuilder(T.Url+"/"+sysId, T.Client)
}

func (T *TableRequestBuilder) Get() (*TableCollectionResponse, error) {

	resp := &TableCollectionResponse{}

	err := T.Client.Get(T.Url, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
