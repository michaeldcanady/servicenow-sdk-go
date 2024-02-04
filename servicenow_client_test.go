package servicenowsdkgo

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/RecoLabs/servicenow-sdk-go/core"
	"github.com/RecoLabs/servicenow-sdk-go/credentials"
	"github.com/RecoLabs/servicenow-sdk-go/internal"
	intCore "github.com/RecoLabs/servicenow-sdk-go/internal/core"
	"github.com/mozillazg/go-httpheader"
	"github.com/stretchr/testify/assert"
)

type test[T any] struct {
	Title string
	// Setup to make needed modifications for a specific test
	Setup func()
	// Cleanup to undo changes do to reusable items
	Cleanup     func()
	Input       interface{}
	Expected    T
	expectedErr error
}

var (
	sharedUsernameAndPasswordCred = credentials.NewUsernamePasswordCredential("username", "password")
)

type MockRequestInformation struct {
	Headers http.Header
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
	request, err := http.NewRequestWithContext(ctx, "GET", "https://www.example.com", nil)
	if err != nil {
		return nil, err
	}
	request.Header = rI.Headers

	return request, nil
}

func (rI *MockRequestInformation) AddHeaders(rawHeaders interface{}) error {
	var headers http.Header
	var err error

	val := reflect.ValueOf(rawHeaders)

	if val.Kind() == reflect.Struct {
		// use the httpheader.Encode function from the httpheader package
		// to encode the pointer value into an http.Header map
		headers, err = httpheader.Encode(rawHeaders)
		if err != nil {
			return err
		}
	} else if val.Type() == reflect.TypeOf(http.Header{}) {
		// if the value is already an http.Header map, just assign it to headers
		headers = rawHeaders.(http.Header)
	} else {
		// otherwise, return an error
		return core.ErrInvalidHeaderType
	}

	// iterate over the headers map and add each key-value pair to rI.Headers
	for key, values := range headers {
		for _, value := range values {
			rI.Headers.Add(key, value)
		}
	}
	return nil
}

func (rI *MockRequestInformation) GetContent() []byte {
	return nil
}
func (rI *MockRequestInformation) GetMethod() string {
	return ""
}
func (rI *MockRequestInformation) GetHeaders() intCore.RequestHeader {
	return &internal.RequestHeader{}
}

func TestNewClient(t *testing.T) {
	cred := credentials.NewUsernamePasswordCredential("username", "password")

	client := NewServiceNowClient(cred, "instance")

	assert.NotNil(t, client)
}

func TestNewClient2(t *testing.T) {
	authProvider, _ := internal.NewBaseAuthorizationProvider(sharedUsernameAndPasswordCred)

	tests := []test[*ServiceNowClient]{
		{
			Title: "Valid",
			Input: []interface{}{"instance", sharedUsernameAndPasswordCred},
			Expected: &ServiceNowClient{
				Credential:   sharedUsernameAndPasswordCred,
				authProvider: authProvider,
				BaseUrl:      "https://instance.service-now.com/api",
				Session:      http.Client{},
			},
			expectedErr: nil,
		},
		{
			Title:       "Nil Credential",
			Input:       []interface{}{"instance", (*credentials.UsernamePasswordCredential)(nil)},
			Expected:    nil,
			expectedErr: internal.ErrNilCredential,
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			values := test.Input.([]interface{})

			instance := values[0].(string)
			cred := values[1].(core.Credential)

			client, err := NewServiceNowClient2(cred, instance)
			assert.Equal(t, test.Expected, client)
			assert.Equal(t, test.expectedErr, err)
		})
	}
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
	requestInfo := &MockRequestInformation{
		Headers: http.Header{},
	}

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
	requestInfo := &MockRequestInformation{
		Headers: http.Header{},
	}

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
