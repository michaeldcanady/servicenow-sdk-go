package servicenowsdkgo

import (
	"encoding/json"
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

// Get performs an HTTP GET request to the specified URL using the Client's session.
// It sets the Authorization header using the Credential's authentication method.
// The response body is decoded into the provided target interface.
// It returns any errors encountered during the request or decoding.
func (C *Client) Get(url string, target interface{}) error {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	request.Header.Add("Authorization", C.Credential.GetAuthentication())
	resp, err := C.Session.Do(request)
	if err != nil {
		return err
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

func (C *Client) Put(url string, input, target interface{}) error {
	request, err := http.NewRequest(http.MethodPut, url, input)
	if err != nil {
		return err
	}
	request.Header.Add("Authorization", C.Credential.GetAuthentication())
	resp, err := C.Session.Do(request)
	if err != nil {
		return err
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

func (C *Client) Delete(url string, target interface{}) error {
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	request.Header.Add("Authorization", C.Credential.GetAuthentication())
	resp, err := C.Session.Do(request)
	if err != nil {
		return err
	}

	return json.NewDecoder(resp.Body).Decode(target)
}
