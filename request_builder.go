package servicenowsdkgo

type RequestBuilder struct {
	Url    string
	Client *Client
}

func NewRequestBuilder(url string, client *Client) *RequestBuilder {
	return &RequestBuilder{
		Url:    url,
		Client: client,
	}
}
