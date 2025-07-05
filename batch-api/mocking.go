package batchapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
	"github.com/stretchr/testify/mock"
)

type mockRestRequest struct {
	mocking.MockModel
}

func newMockRestRequest() *mockRestRequest {
	return &mockRestRequest{
		*mocking.NewMockModel(),
	}
}

// Serialize writes the objects properties to the current writer.
func (mRR *mockRestRequest) Serialize(writer serialization.SerializationWriter) error {
	args := mRR.Called(writer)
	return args.Error(0)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (mRR *mockRestRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	args := mRR.Called()
	return args.Get(0).(map[string]func(serialization.ParseNode) error)
}
func (mRR *mockRestRequest) GetBody() ([]byte, error) {
	args := mRR.Called()
	return args.Get(0).([]byte), args.Error(1)
}
func (mRR *mockRestRequest) SetBodyFromParsable(contentType string, parsable serialization.Parsable) error {
	args := mRR.Called(contentType, parsable)
	return args.Error(0)
}
func (mRR *mockRestRequest) SetBody(body []byte) error {
	args := mRR.Called(body)
	return args.Error(0)
}
func (mRR *mockRestRequest) GetExcludeResponseHeaders() (*bool, error) {
	args := mRR.Called()
	return args.Get(0).(*bool), args.Error(1)
}
func (mRR *mockRestRequest) SetExcludeResponseHeaders(exclude *bool) error {
	args := mRR.Called(exclude)
	return args.Error(0)
}
func (mRR *mockRestRequest) GetHeaders() ([]RestRequestHeader, error) {
	args := mRR.Called()
	return args.Get(0).([]RestRequestHeader), args.Error(1)
}
func (mRR *mockRestRequest) SetHeaders(headers []RestRequestHeader) error {
	args := mRR.Called(headers)
	return args.Error(0)
}
func (mRR *mockRestRequest) GetID() (*string, error) {
	args := mRR.Called()
	return args.Get(0).(*string), args.Error(1)
}
func (mRR *mockRestRequest) SetID(id *string) error {
	args := mRR.Called(id)
	return args.Error(0)
}
func (mRR *mockRestRequest) GetMethod() (*abstractions.HttpMethod, error) {
	args := mRR.Called()
	return args.Get(0).(*abstractions.HttpMethod), args.Error(1)
}
func (mRR *mockRestRequest) SetMethod(method *abstractions.HttpMethod) error {
	args := mRR.Called(method)
	return args.Error(0)
}
func (mRR *mockRestRequest) GetURL() (*string, error) {
	args := mRR.Called()
	return args.Get(0).(*string), args.Error(1)
}
func (mRR *mockRestRequest) SetURL(url *string) error {
	args := mRR.Called(url)
	return args.Error(0)
}

type mockBatchRequest struct {
	mock.Mock
}

func newMockBatchRequest() *mockBatchRequest {
	return &mockBatchRequest{
		Mock: mock.Mock{},
	}
}

func (mock *mockBatchRequest) GetBatchRequestID() (*string, error) {
	args := mock.Called()
	return args.Get(0).(*string), args.Error(1)
}

func (mock *mockBatchRequest) SetBatchRequestID(id *string) error {
	args := mock.Called(id)
	return args.Error(0)
}

func (mock *mockBatchRequest) GetRestRequests() ([]RestRequest, error) {
	args := mock.Called()
	return args.Get(0).([]RestRequest), args.Error(1)
}

func (mock *mockBatchRequest) SetRestRequests(requests []RestRequest) error {
	args := mock.Called(requests)
	return args.Error(0)
}

func (mock *mockBatchRequest) AddRequest(request RestRequest) error {
	args := mock.Called(request)
	return args.Error(0)
}

// Serialize writes the objects properties to the current writer.
func (mock *mockBatchRequest) Serialize(writer serialization.SerializationWriter) error {
	args := mock.Called(writer)
	return args.Error(0)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (mock *mockBatchRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	args := mock.Called()
	return args.Get(0).(map[string]func(serialization.ParseNode) error)
}

func (mock *mockBatchRequest) GetBackingStore() store.BackingStore {
	args := mock.Called()
	return args.Get(0).(store.BackingStore)
}
