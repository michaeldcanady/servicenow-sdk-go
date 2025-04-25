package batchapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
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
