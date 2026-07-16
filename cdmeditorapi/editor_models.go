package cdmeditorapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// MessageResult represents a simple string result.
type MessageResult struct {
	core.BaseModel
	Message *string
}

func NewMessageResult(message *string) *MessageResult {
	return &MessageResult{
		BaseModel: *core.NewBaseModel(),
		Message:   message,
	}
}

func (m *MessageResult) Serialize(writer serialization.SerializationWriter) error { return nil }
func (m *MessageResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return nil
}

func CreateMessageResultFromDiscriminatorValue(node serialization.ParseNode) (serialization.Parsable, error) {
	val, err := node.GetStringValue()
	if err != nil {
		return nil, err
	}
	return NewMessageResult(val), nil
}
