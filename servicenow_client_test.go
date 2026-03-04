package servicenowsdkgo

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
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

// TODO: should be mocked
type MockRequestInformation struct {
	Headers     http.Header
	ReturnError error
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
	if rI.ReturnError != nil {
		return nil, rI.ReturnError
	}
	return http.NewRequest("GET", "https://www.example.com", nil)
}

func (rI *MockRequestInformation) ToRequestWithContext(ctx context.Context) (*http.Request, error) {
	if rI.ReturnError != nil {
		return nil, rI.ReturnError
	}
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

func TestNewServiceNowClient(t *testing.T) {
	cred := credentials.NewUsernamePasswordCredential("username", "password")

	client := NewServiceNowClient(cred, "instance")

	assert.NotNil(t, client)
}

func TestNewServiceNowClient2(t *testing.T) {
	authProvider, _ := internal.NewBaseAuthorizationProvider(sharedUsernameAndPasswordCred)

	tests := []test[*ServiceNowClient]{
		{
			Title: "Valid",
			Input: []interface{}{"instance", sharedUsernameAndPasswordCred},
			Expected: &ServiceNowClient{
				Credential:   sharedUsernameAndPasswordCred,
				authProvider: authProvider,
				BaseUrl:      "https://instance.service-now.com/api",
				Session:      &http.Client{},
			},
			expectedErr: nil,
		},
		{
			Title:       "Nil Credential",
			Input:       []interface{}{"instance", (*credentials.BasicCredential)(nil)},
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
			if test.Expected != nil {
				assert.Equal(t, test.Expected.authProvider, client.authProvider)
				assert.Equal(t, test.Expected.Credential, client.Credential)
				assert.Equal(t, test.Expected.BaseUrl, client.BaseUrl)
				assert.Equal(t, test.Expected.Session, client.Session)
			} else {
				assert.Equal(t, test.Expected, client)
			}
			assert.Equal(t, test.expectedErr, err)
		})
	}
}

func TestServiceNowClient_URL(t *testing.T) {
	cred := credentials.NewUsernamePasswordCredential("username", "password")

	client := NewServiceNowClient(cred, "instance")

	assert.Equal(t, client.BaseUrl, "https://instance.service-now.com/api")
}

type MockWebClient struct {
	Response *http.Response
	Err      error
}

func (m *MockWebClient) Do(req *http.Request) (*http.Response, error) {
	return m.Response, m.Err
}

func TestServiceNowClient_Now(t *testing.T) {
	cred := credentials.NewUsernamePasswordCredential("username", "password")

	client := NewServiceNowClient(cred, "instance")

	nowBuilder := client.Now()

	assert.IsType(t, &NowRequestBuilder{}, nowBuilder)
	assert.Equal(t, client.BaseUrl+"/now", nowBuilder.PathParameters["baseurl"])
	assert.Equal(t, client, nowBuilder.Client)
}

func TestServiceNowClient_Now2(t *testing.T) {
	cred := credentials.NewUsernamePasswordCredential("username", "password")

	client := NewServiceNowClient(cred, "instance")

	nowBuilder := client.Now2()

	assert.IsType(t, &NowRequestBuilder{}, nowBuilder)
	assert.Equal(t, client.BaseUrl+"/now", nowBuilder.PathParameters["baseurl"])
	assert.Equal(t, client, nowBuilder.Client)
}

func TestServiceNowClient_ToRequest(t *testing.T) {
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

type ErrorReader struct{}

func (e *ErrorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("read error")
}

func (e *ErrorReader) Close() error {
	return nil
}

func TestServiceNowClient_UnmarshallError(t *testing.T) {
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

	// Read Error Test
	errorRespReadError := &http.Response{
		StatusCode: 404,
		Body:       &ErrorReader{},
	}

	err = client.unmarshallError(errorRespReadError)
	assert.Error(t, err)
	assert.Equal(t, "read error", err.Error())
}

func TestServiceNowClient_ToRequestWithContext_Error(t *testing.T) {
	// Setup mock credential to fail
	cred := mocking.NewMockCredential()
	cred.On("GetAuthentication").Return("", errors.New("auth error"))

	client := NewServiceNowClient(cred, "instance")

	requestInfo := &MockRequestInformation{
		Headers: http.Header{},
	}

	ctx := context.TODO()

	_, err := client.toRequestWithContext(ctx, requestInfo)
	assert.Error(t, err)
	assert.Equal(t, "auth error", err.Error())
}

func TestServiceNowClient_ToRequestWithContext(t *testing.T) {
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

	_, err = client.toRequestWithContext(context.Background(), requestInfo)
	assert.Error(t, ErrNilContext, err)
}

// TODO: add tests
func TestServiceNowClient_ThrowIfFailedResponse(t *testing.T) {
	tests := []struct {
		name         string
		response     *http.Response
		errorMapping core.ErrorMapping
		expectedErr  bool
	}{
		{
			name: "Success Status",
			response: &http.Response{
				StatusCode: 200,
			},
			expectedErr: false,
		},
		{
			name: "Error Status with Mapping",
			response: &http.Response{
				StatusCode: 404,
				Body:       io.NopCloser(strings.NewReader(`{"error":{"message":"Not Found"}}`)),
			},
			errorMapping: core.ErrorMapping{
				"404": "Not Found",
			},
			expectedErr: true, // It returns the unmarshalled error
		},
		{
			name: "Error Status without Mapping",
			response: &http.Response{
				StatusCode: 500,
				Body:       io.NopCloser(strings.NewReader(`{"error":{"message":"Internal Server Error"}}`)),
			},
			expectedErr: true, // Returns ApiError
		},
		{
			name: "Error Status with Mapping but no match",
			response: &http.Response{
				StatusCode: 404,
				Body:       io.NopCloser(strings.NewReader(`{"error":{"message":"Not Found"}}`)),
			},
			errorMapping: core.ErrorMapping{
				"500": "Internal Server Error",
			},
			expectedErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := &ServiceNowClient{}
			err := client.throwIfFailedResponse(test.response, test.errorMapping)
			if test.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestServiceNowClient_SendWithContext(t *testing.T) {
	tests := []struct {
		name        string
		client      *ServiceNowClient
		requestInfo core.IRequestInformation
		ctx         context.Context
		expectedErr bool
	}{
		{
			name: "Successful Send",
			client: func() *ServiceNowClient {
				c := NewServiceNowClient(credentials.NewUsernamePasswordCredential("user", "pass"), "instance")
				c.Session = &MockWebClient{
					Response: &http.Response{
						StatusCode: 200,
						Body:       io.NopCloser(strings.NewReader("{}")),
					},
				}
				return c
			}(),
			requestInfo: &MockRequestInformation{Headers: http.Header{}},
			ctx:         context.Background(),
			expectedErr: false,
		},
		{
			name: "Send Error",
			client: func() *ServiceNowClient {
				c := NewServiceNowClient(credentials.NewUsernamePasswordCredential("user", "pass"), "instance")
				c.Session = &MockWebClient{
					Err: errors.New("network error"),
				}
				return c
			}(),
			requestInfo: &MockRequestInformation{Headers: http.Header{}},
			ctx:         context.Background(),
			expectedErr: true,
		},
		{
			name: "Failed Status Error",
			client: func() *ServiceNowClient {
				c := NewServiceNowClient(credentials.NewUsernamePasswordCredential("user", "pass"), "instance")
				c.Session = &MockWebClient{
					Response: &http.Response{
						StatusCode: 404,
						Body:       io.NopCloser(strings.NewReader(`{"error":{"message":"Not Found"}}`)),
					},
				}
				return c
			}(),
			requestInfo: &MockRequestInformation{Headers: http.Header{}},
			ctx:         context.Background(),
			expectedErr: true,
		},
		{
			name: "ToRequest Error",
			client: func() *ServiceNowClient {
				c := NewServiceNowClient(credentials.NewUsernamePasswordCredential("user", "pass"), "instance")
				return c
			}(),
			requestInfo: &MockRequestInformation{Headers: http.Header{}, ReturnError: errors.New("to request error")},
			ctx:         context.Background(),
			expectedErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := test.client.SendWithContext(test.ctx, test.requestInfo, nil)
			if test.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestServiceNowClient_Send(t *testing.T) {
	client := NewServiceNowClient(credentials.NewUsernamePasswordCredential("user", "pass"), "instance")
	client.Session = &MockWebClient{
		Response: &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader("{}")),
		},
	}
	requestInfo := &MockRequestInformation{Headers: http.Header{}}

	_, err := client.Send(requestInfo, nil)
	assert.NoError(t, err)
}
