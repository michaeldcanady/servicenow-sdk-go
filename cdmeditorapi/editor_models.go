package cdmeditorapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// MessageResult represents a simple string result.
type MessageResult struct {
	core.BaseModel
	Message *string
}

// NewMessageResult instantiates a new MessageResult.
func NewMessageResult(message *string) *MessageResult {
	return &MessageResult{
		BaseModel: *core.NewBaseModel(),
		Message:   message,
	}
}

// Serialize writes the object's properties to the given writer.
func (m *MessageResult) Serialize(_ serialization.SerializationWriter) error { return nil }

// GetFieldDeserializers returns the deserializers for this object's fields.
func (m *MessageResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return nil
}

// CreateMessageResultFromDiscriminatorValue creates a new MessageResult from a ParseNode.
func CreateMessageResultFromDiscriminatorValue(node serialization.ParseNode) (serialization.Parsable, error) {
	val, err := node.GetStringValue()
	if err != nil {
		return nil, err
	}
	return NewMessageResult(val), nil
}
