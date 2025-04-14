package mocking

import (
	"time"

	"github.com/google/uuid"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/mock"
)

type MockParseNode struct {
	mock.Mock
}

func NewMockParseNode() *MockParseNode {
	return &MockParseNode{
		Mock: mock.Mock{},
	}
}

// GetChildNode returns a new parse node for the given identifier.
func (mPN *MockParseNode) GetChildNode(index string) (serialization.ParseNode, error) {
	args := mPN.Called(index)
	return args.Get(0).(serialization.ParseNode), args.Error(1)
}

// GetCollectionOfObjectValues returns the collection of Parsable values from the node.
func (mPN *MockParseNode) GetCollectionOfObjectValues(ctor serialization.ParsableFactory) ([]serialization.Parsable, error) {
	args := mPN.Called(ctor)
	return args.Get(0).([]serialization.Parsable), args.Error(1)
}

// GetCollectionOfPrimitiveValues returns the collection of primitive values from the node.
func (mPN *MockParseNode) GetCollectionOfPrimitiveValues(targetType string) ([]interface{}, error) {
	args := mPN.Called(targetType)
	return args.Get(0).([]interface{}), args.Error(1)
}

// GetCollectionOfEnumValues returns the collection of Enum values from the node.
func (mPN *MockParseNode) GetCollectionOfEnumValues(parser serialization.EnumFactory) ([]interface{}, error) {
	args := mPN.Called(parser)
	return args.Get(0).([]interface{}), args.Error(1)
}

// GetObjectValue returns the Parsable value from the node.
func (mPN *MockParseNode) GetObjectValue(ctor serialization.ParsableFactory) (serialization.Parsable, error) {
	args := mPN.Called(ctor)
	return args.Get(0).(serialization.Parsable), args.Error(1)
}

// GetStringValue returns a String value from the nodes.
func (mPN *MockParseNode) GetStringValue() (*string, error) {
	args := mPN.Called()
	return args.Get(0).(*string), args.Error(1)
}

// GetBoolValue returns a Bool value from the nodes.
func (mPN *MockParseNode) GetBoolValue() (*bool, error) {
	args := mPN.Called()
	return args.Get(0).(*bool), args.Error(1)
}

// GetInt8Value returns a int8 value from the nodes.
func (mPN *MockParseNode) GetInt8Value() (*int8, error) {
	args := mPN.Called()
	return args.Get(0).(*int8), args.Error(1)
}

// GetByteValue returns a Byte value from the nodes.
func (mPN *MockParseNode) GetByteValue() (*byte, error) {
	args := mPN.Called()
	return args.Get(0).(*byte), args.Error(1)
}

// GetFloat32Value returns a Float32 value from the nodes.
func (mPN *MockParseNode) GetFloat32Value() (*float32, error) {
	args := mPN.Called()
	return args.Get(0).(*float32), args.Error(1)
}

// GetFloat64Value returns a Float64 value from the nodes.
func (mPN *MockParseNode) GetFloat64Value() (*float64, error) {
	args := mPN.Called()
	return args.Get(0).(*float64), args.Error(1)
}

// GetInt32Value returns a Int32 value from the nodes.
func (mPN *MockParseNode) GetInt32Value() (*int32, error) {
	args := mPN.Called()
	return args.Get(0).(*int32), args.Error(1)
}

// GetInt64Value returns a Int64 value from the nodes.
func (mPN *MockParseNode) GetInt64Value() (*int64, error) {
	args := mPN.Called()
	return args.Get(0).(*int64), args.Error(1)
}

// GetTimeValue returns a Time value from the nodes.
func (mPN *MockParseNode) GetTimeValue() (*time.Time, error) {
	args := mPN.Called()
	return args.Get(0).(*time.Time), args.Error(1)
}

// GetISODurationValue returns a ISODuration value from the nodes.
func (mPN *MockParseNode) GetISODurationValue() (*serialization.ISODuration, error) {
	args := mPN.Called()
	return args.Get(0).(*serialization.ISODuration), args.Error(1)
}

// GetTimeOnlyValue returns a TimeOnly value from the nodes.
func (mPN *MockParseNode) GetTimeOnlyValue() (*serialization.TimeOnly, error) {
	args := mPN.Called()
	return args.Get(0).(*serialization.TimeOnly), args.Error(1)
}

// GetDateOnlyValue returns a DateOnly value from the nodes.
func (mPN *MockParseNode) GetDateOnlyValue() (*serialization.DateOnly, error) {
	args := mPN.Called()
	return args.Get(0).(*serialization.DateOnly), args.Error(1)
}

// GetUUIDValue returns a UUID value from the nodes.
func (mPN *MockParseNode) GetUUIDValue() (*uuid.UUID, error) {
	args := mPN.Called()
	return args.Get(0).(*uuid.UUID), args.Error(1)
}

// GetEnumValue returns a Enum value from the nodes.
func (mPN *MockParseNode) GetEnumValue(parser serialization.EnumFactory) (interface{}, error) {
	args := mPN.Called()
	return args.Get(0), args.Error(1)
}

// GetByteArrayValue returns a ByteArray value from the nodes.
func (mPN *MockParseNode) GetByteArrayValue() ([]byte, error) {
	args := mPN.Called()
	return args.Get(0).([]byte), args.Error(1)
}

// GetRawValue returns the values of the node as an interface of any type.
func (mPN *MockParseNode) GetRawValue() (interface{}, error) {
	args := mPN.Called()
	return args.Get(0), args.Error(1)
}

// GetOnBeforeAssignFieldValues returns a callback invoked before the node is deserialized.
func (mPN *MockParseNode) GetOnBeforeAssignFieldValues() serialization.ParsableAction {
	args := mPN.Called()
	return args.Get(0).(serialization.ParsableAction)
}

// SetOnBeforeAssignFieldValues sets a callback invoked before the node is deserialized.
func (mPN *MockParseNode) SetOnBeforeAssignFieldValues(action serialization.ParsableAction) error {
	args := mPN.Called(action)
	return args.Error(1)
}

// GetOnAfterAssignFieldValues returns a callback invoked after the node is deserialized.
func (mPN *MockParseNode) GetOnAfterAssignFieldValues() serialization.ParsableAction {
	args := mPN.Called()
	return args.Get(0).(serialization.ParsableAction)
}

// SetOnAfterAssignFieldValues sets a callback invoked after the node is deserialized.
func (mPN *MockParseNode) SetOnAfterAssignFieldValues(action serialization.ParsableAction) error {
	args := mPN.Called(action)
	return args.Error(1)
}
