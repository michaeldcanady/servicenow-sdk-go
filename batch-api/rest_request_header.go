package batchapi

import (
	"strings"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/kiota"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/model"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

const (
	nameKey  = "name"
	valueKey = "value"
)

// RestRequestHeader represents a Service-Now batch header
type RestRequestHeader interface {
	// GetName returns the key value for the header
	GetName() (*string, error)
	// setName sets the key value for the header
	SetName(*string) error
	// GetValue returns the value of the header
	GetValue() (*string, error)
	// SetValue sets the value of the header
	SetValue(*string) error
	serialization.Parsable
	model.Model
}

// RestRequestHeaderModel implementation of RestRequestHeader
type RestRequestHeaderModel struct {
	model.Model
}

// NewRestRequestHeader creates new instance of BatchHeader
func NewRestRequestHeader() *RestRequestHeaderModel {
	return &RestRequestHeaderModel{
		model.NewBaseModel(),
	}
}

// CreateRestRequestHeaderFromDiscriminatorValue is a parsable factory for creating a BatchRequest
func CreateRestRequestHeaderFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewRestRequestHeader(), nil
}

// Serialize writes the objects properties to the current writer.
func (bH *RestRequestHeaderModel) Serialize(writer serialization.SerializationWriter) error {
	if utils.IsNil(bH) {
		return nil
	}

	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(nameKey)(bH.GetName),
		internalSerialization.SerializeStringFunc(valueKey)(bH.GetValue),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (bH *RestRequestHeaderModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	if utils.IsNil(bH) {
		return nil
	}

	return map[string]func(serialization.ParseNode) error{
		nameKey:  internalSerialization.DeserializeStringFunc()(bH.SetName),
		valueKey: internalSerialization.DeserializeStringFunc()(bH.SetValue),
	}
}

// GetName returns the name of the header
func (bH *RestRequestHeaderModel) GetName() (*string, error) {
	if utils.IsNil(bH) {
		return nil, nil
	}

	backingStore := bH.GetBackingStore()
	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, nameKey)
}

// SetName sets name to provided value
func (bH *RestRequestHeaderModel) SetName(name *string) error {
	if utils.IsNil(bH) {
		return nil
	}

	return kiota.DefaultBackedModelMutatorFunc(bH.GetBackingStore(), nameKey, name)
}

// GetValue returns the value of the header
func (bH *RestRequestHeaderModel) GetValue() (*string, error) {
	if utils.IsNil(bH) {
		return nil, nil
	}

	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](bH.GetBackingStore(), valueKey)
}

// SetValue sets the value to the provided value
func (bH *RestRequestHeaderModel) SetValue(value *string) error {
	if utils.IsNil(bH) {
		return nil
	}

	return kiota.DefaultBackedModelMutatorFunc(bH.GetBackingStore(), valueKey, value)
}

// headers support headers types
type headers interface {
	*abstractions.RequestHeaders
}

// createRestRequestHeaderFromHeaders converts headers to RestRequestHeader
func createRestRequestHeaderFromHeaders[h headers](headers h) ([]RestRequestHeader, error) {
	batchHeaders := make([]RestRequestHeader, 0)

	if requestHeaders, ok := interface{}(headers).(*abstractions.RequestHeaders); ok {
		if utils.IsNil(requestHeaders) {
			return batchHeaders, nil
		}
		for _, key := range requestHeaders.ListKeys() {
			batchHeader := NewRestRequestHeader()
			values := requestHeaders.Get(key)
			if err := batchHeader.SetName(&key); err != nil {
				return nil, err
			}
			valuesString := strings.Join(values, ", ")
			if err := batchHeader.SetValue(&valuesString); err != nil {
				return nil, err
			}
			batchHeaders = append(batchHeaders, batchHeader)
		}
		return batchHeaders, nil
	}

	return nil, nil
}
