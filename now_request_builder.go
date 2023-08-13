package servicenowsdkgo

type NowRequestBuilder struct {
	RequestBuilder
}

func NewNowRequestBuilder(url string, client *Client) *NowRequestBuilder {
	requestBuilder := NewRequestBuilder(url, client)
	return &NowRequestBuilder{
		*requestBuilder,
	}
}

func (N *NowRequestBuilder) Table(tableName string) *TableRequestBuilder {
	url := N.Url + "/table/" + tableName
	return NewTableRequestBuilder(url, N.Client)
}
