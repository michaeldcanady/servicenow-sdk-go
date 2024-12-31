package tableapi

import (
	"testing"

	jsonserialization "github.com/microsoft/kiota-serialization-json-go"
	"github.com/stretchr/testify/assert"
)

func TestNewElementValue(t *testing.T) {
	tests := []struct {
		Title    string
		Input    interface{}
		Err      error
		Expected interface{}
	}{}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			elemValue := newElementValue(test.Input)
			assert.Equal(t, test.Expected, elemValue)
		})
	}
}

func TestCreateElementValueFromDiscriminatorValue(t *testing.T) {
	tests := []struct {
		Title    string
		Input    interface{}
		Err      error
		Expected interface{}
	}{}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			parsable, err := CreateElementValueFromDiscriminatorValue(nil)
			assert.Equal(t, test.Err, err)
			assert.Equal(t, test.Expected, parsable)
		})
	}
}

func TestElementValue_Serialize(t *testing.T) {
	tests := []struct {
		Title    string
		Input    interface{}
		Err      error
		Expected interface{}
	}{}

	writer := jsonserialization.NewJsonSerializationWriter()

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			input, ok := test.Input.(elementValue)
			if !ok {
				t.Error("test.Input is not elementValue")
			}
			assert.Equal(t, test.Err, input.Serialize(writer))
			content, err := writer.GetSerializedContent()
			assert.Equal(t, test.Err, err)
			assert.Equal(t, test.Expected, content)
		})
	}
}

func TestElementValue_IsNil(t *testing.T) {
	tests := []struct {
		Title    string
		Input    interface{}
		Err      error
		Expected interface{}
	}{}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {})
	}
}

func TestElementValue_setValue(t *testing.T) {
	tests := []struct {
		Title    string
		Input    interface{}
		Err      error
		Expected interface{}
	}{}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {})
	}
}

func TestElementValue_GetStringValue(t *testing.T) {
	tests := []struct {
		Title    string
		Input    interface{}
		Err      error
		Expected interface{}
	}{}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {})
	}
}

func TestElementValue_GetBoolValue(t *testing.T) {
	tests := []struct {
		Title    string
		Input    interface{}
		Err      error
		Expected interface{}
	}{}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {})
	}
}

func TestElementValue_GetInt8Value(t *testing.T) {
	tests := []struct {
		Title    string
		Input    interface{}
		Err      error
		Expected interface{}
	}{}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {})
	}
}

func TestElementValue_GetByteValue(t *testing.T) {
	tests := []struct {
		Title    string
		Input    interface{}
		Err      error
		Expected interface{}
	}{}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {})
	}
}

func TestElementValue_GetFloat32Value(t *testing.T) {
	tests := []struct {
		Title    string
		Input    interface{}
		Err      error
		Expected interface{}
	}{}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {})
	}
}

func TestElementValue_GetFloat64Value(t *testing.T) {
	tests := []struct {
		Title    string
		Input    interface{}
		Err      error
		Expected interface{}
	}{}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {})
	}
}

func TestElementValue_GetInt32Value(t *testing.T) {
	tests := []struct {
		Title    string
		Input    interface{}
		Err      error
		Expected interface{}
	}{}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {})
	}
}

func TestElementValue_GetInt64Value(t *testing.T) {
	tests := []struct {
		Title    string
		Input    interface{}
		Err      error
		Expected interface{}
	}{}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {})
	}
}

func TestElementValue_GetTimeValue(t *testing.T) {
	tests := []struct {
		Title    string
		Input    interface{}
		Err      error
		Expected interface{}
	}{}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {})
	}
}

func TestElementValue_GetTimeOnlyValue(t *testing.T) {
	tests := []struct {
		Title    string
		Input    interface{}
		Err      error
		Expected interface{}
	}{}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {})
	}
}

func TestElementValue_GetDateOnlyValue(t *testing.T) {
	tests := []struct {
		Title    string
		Input    interface{}
		Err      error
		Expected interface{}
	}{}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {})
	}
}

func TestElementValue_GetEnumValue(t *testing.T) {
	tests := []struct {
		Title    string
		Input    interface{}
		Err      error
		Expected interface{}
	}{}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {})
	}
}

func TestElementValue_GetCollectionOfPrimitiveValues(t *testing.T) {
	tests := []struct {
		Title    string
		Input    interface{}
		Err      error
		Expected interface{}
	}{}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {})
	}
}

func TestElementValue_GetPrimitiveValue(t *testing.T) {
	tests := []struct {
		Title    string
		Input    interface{}
		Err      error
		Expected interface{}
	}{}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {})
	}
}

func TestElementValue_GetRawValue(t *testing.T) {
	tests := []struct {
		Title    string
		Input    interface{}
		Err      error
		Expected interface{}
	}{}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {})
	}
}
