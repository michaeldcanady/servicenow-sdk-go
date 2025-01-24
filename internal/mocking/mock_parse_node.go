package mocking

import (
	"time"

	"github.com/google/uuid"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/mock"
)

var _ serialization.ParseNode = (*MockParseNode)(nil)

type MockParseNode struct {
	mock.Mock
}

func NewMockParseNode() *MockParseNode {
	return &MockParseNode{Mock: mock.Mock{}}
}

// GetChildNode returns a new parse node for the given identifier.
func (pN *MockParseNode) GetChildNode(index string) (serialization.ParseNode, error) {
	args := pN.Called(index)
	return args.Get(0).(serialization.ParseNode), args.Error(1)
}

// GetCollectionOfObjectValues returns the collection of Parsable values from the node.
func (pN *MockParseNode) GetCollectionOfObjectValues(ctor serialization.ParsableFactory) ([]serialization.Parsable, error) {
	args := pN.Called(ctor)
	return args.Get(0).([]serialization.Parsable), args.Error(1)
}

// GetCollectionOfPrimitiveValues returns the collection of primitive values from the node.
func (pN *MockParseNode) GetCollectionOfPrimitiveValues(targetType string) ([]interface{}, error) {
	args := pN.Called(targetType)
	return args.Get(0).([]interface{}), args.Error(1)
}

// GetCollectionOfEnumValues returns the collection of Enum values from the node.
func (pN *MockParseNode) GetCollectionOfEnumValues(parser serialization.EnumFactory) ([]interface{}, error) {
	args := pN.Called(parser)
	return args.Get(0).([]interface{}), args.Error(1)
}

// GetObjectValue returns the Parsable value from the node.
func (pN *MockParseNode) GetObjectValue(ctor serialization.ParsableFactory) (serialization.Parsable, error) {
	args := pN.Called(ctor)
	return args.Get(0).(serialization.Parsable), args.Error(1)
}

// GetStringValue returns a String value from the nodes.
func (pN *MockParseNode) GetStringValue() (*string, error) {
	args := pN.Called()
	return args.Get(0).(*string), args.Error(1)
}

// GetBoolValue returns a Bool value from the nodes.
func (pN *MockParseNode) GetBoolValue() (*bool, error) {
	args := pN.Called()
	return args.Get(0).(*bool), args.Error(1)
}

// GetInt8Value returns a int8 value from the nodes.
func (pN *MockParseNode) GetInt8Value() (*int8, error) {
	args := pN.Called()
	return args.Get(0).(*int8), args.Error(1)
}

// GetByteValue returns a Byte value from the nodes.
func (pN *MockParseNode) GetByteValue() (*byte, error) {
	args := pN.Called()
	return args.Get(0).(*byte), args.Error(1)
}

// GetFloat32Value returns a Float32 value from the nodes.
func (pN *MockParseNode) GetFloat32Value() (*float32, error) {
	args := pN.Called()
	return args.Get(0).(*float32), args.Error(1)
}

// GetFloat64Value returns a Float64 value from the nodes.
func (pN *MockParseNode) GetFloat64Value() (*float64, error) {
	args := pN.Called()
	return args.Get(0).(*float64), args.Error(1)
}

// GetInt32Value returns a Int32 value from the nodes.
func (pN *MockParseNode) GetInt32Value() (*int32, error) {
	args := pN.Called()
	return args.Get(0).(*int32), args.Error(1)
}

// GetInt64Value returns a Int64 value from the nodes.
func (pN *MockParseNode) GetInt64Value() (*int64, error) {
	args := pN.Called()
	return args.Get(0).(*int64), args.Error(1)
}

// GetTimeValue returns a Time value from the nodes.
func (pN *MockParseNode) GetTimeValue() (*time.Time, error) {
	args := pN.Called()
	return args.Get(0).(*time.Time), args.Error(1)
}

// GetISODurationValue returns a ISODuration value from the nodes.
func (pN *MockParseNode) GetISODurationValue() (*serialization.ISODuration, error) {
	args := pN.Called()
	return args.Get(0).(*serialization.ISODuration), args.Error(1)
}

// GetTimeOnlyValue returns a TimeOnly value from the nodes.
func (pN *MockParseNode) GetTimeOnlyValue() (*serialization.TimeOnly, error) {
	args := pN.Called()
	return args.Get(0).(*serialization.TimeOnly), args.Error(1)
}

// GetDateOnlyValue returns a DateOnly value from the nodes.
func (pN *MockParseNode) GetDateOnlyValue() (*serialization.DateOnly, error) {
	args := pN.Called()
	return args.Get(0).(*serialization.DateOnly), args.Error(1)
}

// GetUUIDValue returns a UUID value from the nodes.
func (pN *MockParseNode) GetUUIDValue() (*uuid.UUID, error) {
	args := pN.Called()
	return args.Get(0).(*uuid.UUID), args.Error(1)
}

// GetEnumValue returns a Enum value from the nodes.
func (pN *MockParseNode) GetEnumValue(parser serialization.EnumFactory) (interface{}, error) {
	args := pN.Called()
	return args.Get(0), args.Error(1)
}

// GetByteArrayValue returns a ByteArray value from the nodes.
func (pN *MockParseNode) GetByteArrayValue() ([]byte, error) {
	args := pN.Called()
	return args.Get(0).([]byte), args.Error(1)
}

// GetRawValue returns the values of the node as an interface of any type.
func (pN *MockParseNode) GetRawValue() (interface{}, error) {
	args := pN.Called()
	return args.Get(0), args.Error(1)
}

// GetOnBeforeAssignFieldValues returns a callback invoked before the node is deserialized.
func (pN *MockParseNode) GetOnBeforeAssignFieldValues() serialization.ParsableAction {
	args := pN.Called()
	return args.Get(0).(serialization.ParsableAction)
}

// SetOnBeforeAssignFieldValues sets a callback invoked before the node is deserialized.
func (pN *MockParseNode) SetOnBeforeAssignFieldValues(action serialization.ParsableAction) error {
	args := pN.Called(action)
	return args.Error(0)
}

// GetOnAfterAssignFieldValues returns a callback invoked after the node is deserialized.
func (pN *MockParseNode) GetOnAfterAssignFieldValues() serialization.ParsableAction {
	args := pN.Called()
	return args.Get(0).(serialization.ParsableAction)
}

// SetOnAfterAssignFieldValues sets a callback invoked after the node is deserialized.
func (pN *MockParseNode) SetOnAfterAssignFieldValues(action serialization.ParsableAction) error {
	args := pN.Called(action)
	return args.Error(0)
}
