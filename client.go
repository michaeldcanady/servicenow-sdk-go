package servicenowsdkgo

import (
	"errors"
	"net/http"
	"strings"
)

type Client struct {
	Credential *UsernamePasswordCredential
	BaseUrl    string
	Session    http.Client
}

// NewClient creates a new instance of the ServiceNow client.
// It accepts a UsernamePasswordCredential and an instance URL.
// If the instance URL does not end with ".service-now.com/api", it appends the suffix.
// It returns a pointer to the Client.
func NewClient(credential *UsernamePasswordCredential, instance string) *Client {
	if !strings.HasSuffix(instance, ".service-now.com/api") {
		instance += ".service-now.com/api"
	}

	return &Client{
		Credential: credential,
		BaseUrl:    instance,
		Session:    http.Client{},
	}
}

// Now returns a NowRequestBuilder associated with the Client.
// It prepares the NowRequestBuilder with the base URL for the ServiceNow instance.
func (C *Client) Now() *NowRequestBuilder {
	return NewNowRequestBuilder(C.BaseUrl+"/now", C)
}

func (C *Client) Send(requestInfo *RequestInformation) (*http.Response, error) {
	if requestInfo == nil {
		return nil, errors.New("requestInfo cannot be nil")
	}
	request, err := requestInfo.toRequest()
	if err != nil {
		return nil, err
	}
	return C.Session.Do(request)
}
