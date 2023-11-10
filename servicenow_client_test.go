package servicenowsdkgo

import (
	"context"
	"net/http"
	"net/url"
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

func (rI *MockRequestInformation) SetUri(url *url.URL) {
}

func (rI *MockRequestInformation) Url() (string, error) {
	return "https://www.example.com", nil
}

func (rI *MockRequestInformation) ToRequest() (*http.Request, error) {
	return http.NewRequest("GET", "https://www.example.com", nil)
}

func (rI *MockRequestInformation) ToRequestWithContext(ctx context.Context) (*http.Request, error) {
	return &http.Request{}, nil
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

func TestClient_Now(t *testing.T) {

	cred := credentials.NewUsernamePasswordCredential("username", "password")

	client := NewServiceNowClient(cred, "instance")

	nowBuilder := client.Now()

	assert.IsType(t, &NowRequestBuilder{}, nowBuilder)
	assert.Equal(t, client.BaseUrl+"/now", nowBuilder.PathParameters["baseurl"])
	assert.Equal(t, client, nowBuilder.Client)
}

func TestClient_toRequest(t *testing.T) {

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
}
