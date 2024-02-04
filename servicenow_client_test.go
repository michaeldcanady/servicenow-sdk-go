package servicenowsdkgo

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
	"github.com/stretchr/testify/assert"
)

type MockRequestInformation struct {
}

func (rI *MockRequestInformation) AddRequestOptions(options []core.RequestOption) {
}

func (rI *MockRequestInformation) GetRequestOptions() []core.RequestOption {
	return []core.RequestOption{}
}

func (rI *MockRequestInformation) SetStreamContent(content []byte) {

}

func (rI *MockRequestInformation) AddQueryParameters(source interface{}) error {
	return nil
}

func (rI *MockRequestInformation) SetUri(url *url.URL) { //nolint:stylecheck
}

func (rI *MockRequestInformation) Url() (string, error) { //nolint:stylecheck
	return "https://www.example.com", nil
}

func (rI *MockRequestInformation) ToRequest() (*http.Request, error) {
	return http.NewRequest("GET", "https://www.example.com", nil)
}

func (rI *MockRequestInformation) ToRequestWithContext(ctx context.Context) (*http.Request, error) {
	return http.NewRequestWithContext(ctx, "GET", "https://www.example.com", nil)
}

func (rI *MockRequestInformation) AddHeaders(rawHeaders interface{}) error {
	return nil
}

func TestNewClient(t *testing.T) {
	cred := credentials.NewUsernamePasswordCredential("username", "password")

	client := NewServiceNowClient(cred, "instance")

	assert.NotNil(t, client)
}

func TestClientURL(t *testing.T) {
	cred := credentials.NewUsernamePasswordCredential("username", "password")

	client := NewServiceNowClient(cred, "instance")

	assert.Equal(t, client.BaseUrl, "https://instance.service-now.com/api")
}

func TestClientNow(t *testing.T) {
	cred := credentials.NewUsernamePasswordCredential("username", "password")

	client := NewServiceNowClient(cred, "instance")

	nowBuilder := client.Now(context.Background())

	assert.IsType(t, &NowRequestBuilder{}, nowBuilder)
	assert.Equal(t, client.BaseUrl+"/now", nowBuilder.PathParameters["baseurl"])
	assert.Equal(t, client, nowBuilder.Client)
}

func TestClientToRequest(t *testing.T) {
	requestInfo := &MockRequestInformation{}

	cred := credentials.NewUsernamePasswordCredential("username", "password")

	client := NewServiceNowClient(cred, "instance")

	request, err := client.toRequest(requestInfo)
	if err != nil {
		t.Error(err)
	}

	expected := &url.URL{Scheme: "https", Opaque: "", User: (*url.Userinfo)(nil), Host: "www.example.com", Path: "", RawPath: "", OmitHost: false, ForceQuery: false, RawQuery: "", Fragment: "", RawFragment: ""}
	expectedContentTypeHeader := "application/json"
	expectedAcceptHeader := "application/json"
	expectedAuthorizationHeader := "Basic dXNlcm5hbWU6cGFzc3dvcmQ="

	assert.Equal(t, expected, request.URL)
	assert.Equal(t, expectedContentTypeHeader, request.Header.Get("Content-Type"))
	assert.Equal(t, expectedAcceptHeader, request.Header.Get("Accept"))
	assert.Equal(t, expectedAuthorizationHeader, request.Header.Get("Authorization"))

	_, err = client.toRequest(nil)
	assert.Error(t, ErrNilRequestInfo, err)
}

func TestClientUnmarshallError(t *testing.T) {
	cred := credentials.NewUsernamePasswordCredential("username", "password")

	client := NewServiceNowClient(cred, "instance")

	properErr := core.ServiceNowError{
		Exception: core.Exception{
			Detail:  "Resource not found",
			Message: "Resource not found",
		},
		Status: "404",
	}

	jsonBytes, err := json.Marshal(properErr)
	if err != nil {
		t.Error(err)
	}

	errorResp := &http.Response{
		StatusCode: 404,
		Body:       io.NopCloser(strings.NewReader(string(jsonBytes))),
	}

	err = client.unmarshallError(errorResp)
	assert.IsType(t, &core.ServiceNowError{}, err)
	assert.Equal(t, &properErr, err)

	errorResp = &http.Response{
		StatusCode: 404,
		Body:       io.NopCloser(strings.NewReader("bad response")),
	}
	err = client.unmarshallError(errorResp)
	assert.IsType(t, &json.SyntaxError{}, err)
}

func TestClientToRequestWithContext(t *testing.T) {
	requestInfo := &MockRequestInformation{}

	cred := credentials.NewUsernamePasswordCredential("username", "password")

	ctx := context.TODO()

	client := NewServiceNowClient(cred, "instance")
	request, err := client.toRequestWithContext(ctx, requestInfo) //nolint:all
	if err != nil {
		t.Error(err)
	}

	expected := &url.URL{Scheme: "https", Opaque: "", User: (*url.Userinfo)(nil), Host: "www.example.com", Path: "", RawPath: "", OmitHost: false, ForceQuery: false, RawQuery: "", Fragment: "", RawFragment: ""}
	expectedContentTypeHeader := "application/json"
	expectedAcceptHeader := "application/json"
	expectedAuthorizationHeader := "Basic dXNlcm5hbWU6cGFzc3dvcmQ="

	assert.Equal(t, expected, request.URL)
	assert.Equal(t, expectedContentTypeHeader, request.Header.Get("Content-Type"))
	assert.Equal(t, expectedAcceptHeader, request.Header.Get("Accept"))
	assert.Equal(t, expectedAuthorizationHeader, request.Header.Get("Authorization"))
	assert.Equal(t, ctx, request.Context())

	_, err = client.toRequestWithContext(context.TODO(), nil)
	assert.Error(t, ErrNilRequestInfo, err)

	_, err = client.toRequestWithContext(nil, requestInfo) //nolint:all
	assert.Error(t, ErrNilContext, err)
}
