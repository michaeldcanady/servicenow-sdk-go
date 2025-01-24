package mocking

import (
	"time"

	"github.com/google/uuid"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/mock"
)

var _ serialization.SerializationWriter = (*MockSerializationWriter)(nil)

type MockSerializationWriter struct {
	mock.Mock
}

func NewMockSerializationWriter() *MockSerializationWriter {
	return &MockSerializationWriter{Mock: mock.Mock{}}
}

func (sW *MockSerializationWriter) Close() error {
	args := sW.Called()
	return args.Error(0)
}

// WriteStringValue writes a String value to underlying the byte array.
func (sW *MockSerializationWriter) WriteStringValue(key string, value *string) error {
	args := sW.Called(key, value)
	return args.Error(0)
}

// WriteBoolValue writes a Bool value to underlying the byte array.
func (sW *MockSerializationWriter) WriteBoolValue(key string, value *bool) error {
	args := sW.Called(key, value)
	return args.Error(0)
}

// WriteInt8Value writes a int8 value to underlying the byte array.
func (sW *MockSerializationWriter) WriteInt8Value(key string, value *int8) error {
	args := sW.Called(key, value)
	return args.Error(0)
}

// WriteByteValue writes a Byte value to underlying the byte array.
func (sW *MockSerializationWriter) WriteByteValue(key string, value *byte) error {
	args := sW.Called(key, value)
	return args.Error(0)
}

// WriteInt32Value writes a Int32 value to underlying the byte array.
func (sW *MockSerializationWriter) WriteInt32Value(key string, value *int32) error {
	args := sW.Called(key, value)
	return args.Error(0)
}

// WriteInt64Value writes a Int64 value to underlying the byte array.
func (sW *MockSerializationWriter) WriteInt64Value(key string, value *int64) error {
	args := sW.Called(key, value)
	return args.Error(0)
}

// WriteFloat32Value writes a Float32 value to underlying the byte array.
func (sW *MockSerializationWriter) WriteFloat32Value(key string, value *float32) error {
	args := sW.Called(key, value)
	return args.Error(0)
}

// WriteFloat64Value writes a Float64 value to underlying the byte array.
func (sW *MockSerializationWriter) WriteFloat64Value(key string, value *float64) error {
	args := sW.Called(key, value)
	return args.Error(0)
}

// WriteByteArrayValue writes a ByteArray value to underlying the byte array.
func (sW *MockSerializationWriter) WriteByteArrayValue(key string, value []byte) error {
	args := sW.Called(key, value)
	return args.Error(0)
}

// WriteTimeValue writes a Time value to underlying the byte array.
func (sW *MockSerializationWriter) WriteTimeValue(key string, value *time.Time) error {
	args := sW.Called(key, value)
	return args.Error(0)
}

// WriteTimeOnlyValue writes the time part of a Time value to underlying the byte array.
func (sW *MockSerializationWriter) WriteTimeOnlyValue(key string, value *serialization.TimeOnly) error {
	args := sW.Called(key, value)
	return args.Error(0)
}

// WriteDateOnlyValue writes the date part of a Time value to underlying the byte array.
func (sW *MockSerializationWriter) WriteDateOnlyValue(key string, value *serialization.DateOnly) error {
	args := sW.Called(key, value)
	return args.Error(0)
}

// WriteISODurationValue writes a ISODuration value to underlying the byte array.
func (sW *MockSerializationWriter) WriteISODurationValue(key string, value *serialization.ISODuration) error {
	args := sW.Called(key, value)
	return args.Error(0)
}

// WriteUUIDValue writes a UUID value to underlying the byte array.
func (sW *MockSerializationWriter) WriteUUIDValue(key string, value *uuid.UUID) error {
	args := sW.Called(key, value)
	return args.Error(0)
}

// WriteObjectValue writes a Parsable value to underlying the byte array.
func (sW *MockSerializationWriter) WriteObjectValue(key string, item serialization.Parsable, additionalValuesToMerge ...serialization.Parsable) error {
	args := sW.Called(key, item, additionalValuesToMerge)
	return args.Error(0)
}

// WriteCollectionOfObjectValues writes a collection of Parsable values to underlying the byte array.
func (sW *MockSerializationWriter) WriteCollectionOfObjectValues(key string, collection []serialization.Parsable) error {
	args := sW.Called(key, collection)
	return args.Error(0)
}

// WriteCollectionOfStringValues writes a collection of String values to underlying the byte array.
func (sW *MockSerializationWriter) WriteCollectionOfStringValues(key string, collection []string) error {
	args := sW.Called(key, collection)
	return args.Error(0)
}

// WriteCollectionOfBoolValues writes a collection of Bool values to underlying the byte array.
func (sW *MockSerializationWriter) WriteCollectionOfBoolValues(key string, collection []bool) error {
	args := sW.Called(key, collection)
	return args.Error(0)
}

// WriteCollectionOfInt8Values writes a collection of Int8 values to underlying the byte array.
func (sW *MockSerializationWriter) WriteCollectionOfInt8Values(key string, collection []int8) error {
	args := sW.Called(key, collection)
	return args.Error(0)
}

// WriteCollectionOfByteValues writes a collection of Byte values to underlying the byte array.
func (sW *MockSerializationWriter) WriteCollectionOfByteValues(key string, collection []byte) error {
	args := sW.Called(key, collection)
	return args.Error(0)
}

// WriteCollectionOfInt32Values writes a collection of Int32 values to underlying the byte array.
func (sW *MockSerializationWriter) WriteCollectionOfInt32Values(key string, collection []int32) error {
	args := sW.Called(key, collection)
	return args.Error(0)
}

// WriteCollectionOfInt64Values writes a collection of Int64 values to underlying the byte array.
func (sW *MockSerializationWriter) WriteCollectionOfInt64Values(key string, collection []int64) error {
	args := sW.Called(key, collection)
	return args.Error(0)
}

// WriteCollectionOfFloat32Values writes a collection of Float32 values to underlying the byte array.
func (sW *MockSerializationWriter) WriteCollectionOfFloat32Values(key string, collection []float32) error {
	args := sW.Called(key, collection)
	return args.Error(0)
}

// WriteCollectionOfFloat64Values writes a collection of Float64 values to underlying the byte array.
func (sW *MockSerializationWriter) WriteCollectionOfFloat64Values(key string, collection []float64) error {
	args := sW.Called(key, collection)
	return args.Error(0)
}

// WriteCollectionOfTimeValues writes a collection of Time values to underlying the byte array.
func (sW *MockSerializationWriter) WriteCollectionOfTimeValues(key string, collection []time.Time) error {
	args := sW.Called(key, collection)
	return args.Error(0)
}

// WriteCollectionOfISODurationValues writes a collection of ISODuration values to underlying the byte array.
func (sW *MockSerializationWriter) WriteCollectionOfISODurationValues(key string, collection []serialization.ISODuration) error {
	args := sW.Called(key, collection)
	return args.Error(0)
}

// WriteCollectionOfDateOnlyValues writes a collection of DateOnly values to underlying the byte array.
func (sW *MockSerializationWriter) WriteCollectionOfDateOnlyValues(key string, collection []serialization.DateOnly) error {
	args := sW.Called(key, collection)
	return args.Error(0)
}

// WriteCollectionOfTimeOnlyValues writes a collection of TimeOnly values to underlying the byte array.
func (sW *MockSerializationWriter) WriteCollectionOfTimeOnlyValues(key string, collection []serialization.TimeOnly) error {
	args := sW.Called(key, collection)
	return args.Error(0)
}

// WriteCollectionOfUUIDValues writes a collection of UUID values to underlying the byte array.
func (sW *MockSerializationWriter) WriteCollectionOfUUIDValues(key string, collection []uuid.UUID) error {
	args := sW.Called(key, collection)
	return args.Error(0)
}

// GetSerializedContent returns the resulting byte array from the serialization writer.
func (sW *MockSerializationWriter) GetSerializedContent() ([]byte, error) {
	args := sW.Called()
	return args.Get(0).([]byte), args.Error(1)
}

// WriteNullValue writes a null value for the specified key.
func (sW *MockSerializationWriter) WriteNullValue(key string) error {
	args := sW.Called(key)
	return args.Error(0)
}

// WriteAdditionalData writes additional data to underlying the byte array.
func (sW *MockSerializationWriter) WriteAdditionalData(value map[string]interface{}) error {
	args := sW.Called(value)
	return args.Error(0)
}

// WriteAnyValue an object of unknown type as a json value
func (sW *MockSerializationWriter) WriteAnyValue(key string, value interface{}) error {
	args := sW.Called(key, value)
	return args.Error(0)
}

// GetOnBeforeSerialization returns a callback invoked before the serialization process starts.
func (sW *MockSerializationWriter) GetOnBeforeSerialization() serialization.ParsableAction {
	args := sW.Called()
	return args.Get(0).(serialization.ParsableAction)
}

// SetOnBeforeSerialization sets a callback invoked before the serialization process starts.
func (sW *MockSerializationWriter) SetOnBeforeSerialization(action serialization.ParsableAction) error {
	args := sW.Called(action)
	return args.Error(0)
}

// GetOnAfterObjectSerialization returns a callback invoked after the serialization process completes.
func (sW *MockSerializationWriter) GetOnAfterObjectSerialization() serialization.ParsableAction {
	args := sW.Called()
	return args.Get(0).(serialization.ParsableAction)
}

// SetOnAfterObjectSerialization sets a callback invoked after the serialization process completes.
func (sW *MockSerializationWriter) SetOnAfterObjectSerialization(action serialization.ParsableAction) error {
	args := sW.Called(action)
	return args.Error(0)
}

// GetOnStartObjectSerialization returns a callback invoked right after the serialization process starts.
func (sW *MockSerializationWriter) GetOnStartObjectSerialization() serialization.ParsableWriter {
	args := sW.Called()
	return args.Get(0).(serialization.ParsableWriter)
}

// SetOnStartObjectSerialization sets a callback invoked right after the serialization process starts.
func (sW *MockSerializationWriter) SetOnStartObjectSerialization(writer serialization.ParsableWriter) error {
	args := sW.Called(writer)
	return args.Error(0)
}
