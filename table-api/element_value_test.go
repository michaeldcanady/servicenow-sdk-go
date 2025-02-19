package tableapi

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	jsonserialization "github.com/microsoft/kiota-serialization-json-go"
	"github.com/stretchr/testify/assert"
)

func TestNewElementValue(t *testing.T) {
	tests := []struct {
		Title    string
		Input    interface{}
		Err      error
		Expected interface{}
	}{
		{
			Title:    "Integer value",
			Input:    10,
			Err:      nil,
			Expected: &ElementValueImpl{val: 10},
		},
		{
			Title:    "Nil value",
			Input:    nil,
			Err:      nil,
			Expected: &ElementValueImpl{val: nil},
		},
		{
			Title:    "Typed nil value",
			Input:    (*bool)(nil),
			Err:      nil,
			Expected: &ElementValueImpl{val: (*bool)(nil)},
		},
		{
			Title:    "Integer slice",
			Input:    []int{10, 20},
			Err:      nil,
			Expected: &ElementValueImpl{val: []*ElementValueImpl{{val: 10}, {val: 20}}},
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			elemValue := newElementValue(test.Input)
			assert.IsType(t, &ElementValueImpl{}, test.Expected)
			assert.Equal(t, *test.Expected.(*ElementValueImpl), *elemValue)
		})
	}
}

func TestCreateElementValueFromDiscriminatorValue(t *testing.T) {
	tests := []struct {
		Title    string
		Input    interface{}
		Err      error
		Expected interface{}
	}{
		{
			Title:    "Successful",
			Input:    nil,
			Err:      nil,
			Expected: &ElementValueImpl{val: nil},
		},
	}

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
	}{
		{
			Title:    "Successful",
			Input:    &ElementValueImpl{},
			Err:      errors.New("Serialize not implemented"),
			Expected: []uint8{},
		},
		{
			Title:    "nil value",
			Input:    (*ElementValueImpl)(nil),
			Expected: []uint8{},
		},
	}

	writer := jsonserialization.NewJsonSerializationWriter()

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			input, ok := test.Input.(*ElementValueImpl)
			if !ok {
				t.Error("test.Input is not elementValue")
			}
			assert.Equal(t, test.Err, input.Serialize(writer))
			content, err := writer.GetSerializedContent()
			assert.Nil(t, err)
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
	}{
		{
			Title:    "Nil value",
			Input:    &ElementValueImpl{val: nil},
			Err:      nil,
			Expected: true,
		},
		{
			Title:    "Nil ElementValue",
			Input:    (*ElementValueImpl)(nil),
			Err:      nil,
			Expected: true,
		},
		{
			Title:    "Not nil value",
			Input:    &ElementValueImpl{val: "string"},
			Err:      nil,
			Expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			input, ok := test.Input.(*ElementValueImpl)
			if !ok {
				t.Error("test.Input is not elementValue")
			}
			assert.Equal(t, input.IsNil(), test.Expected)
		})
	}
}

func TestElementValue_setValue(t *testing.T) {
	tests := []struct {
		Title string
		Test  func(*testing.T)
	}{
		{
			Title: "Nil ElementValue",
			Test: func(t *testing.T) {
				value := (*ElementValueImpl)(nil)
				value.setValue("str")
				assert.Equal(t, (*ElementValueImpl)(nil), value)
			},
		},
		{
			Title: "String value",
			Test: func(t *testing.T) {
				value := &ElementValueImpl{val: nil}
				value.setValue("str")
				assert.Equal(t, "str", value.val)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Title, test.Test)
	}
}

func TestElementValue_GetStringValue(t *testing.T) {
	tests := []struct {
		Title string
		Test  func(*testing.T)
	}{
		{
			Title: "Successful",
			Test: func(t *testing.T) {
				input := "string"
				value := &ElementValueImpl{val: &input}
				strVal, err := value.GetStringValue()
				assert.Equal(t, "string", *strVal)
				assert.Nil(t, err)
			},
		},
		{
			Title: "Not String Pointer",
			Test: func(t *testing.T) {
				input := true
				value := &ElementValueImpl{val: input}
				strVal, err := value.GetStringValue()
				assert.Equal(t, (*string)(nil), strVal)
				assert.Equal(t, fmt.Errorf("value '%v' is not compatible with type string", input), err)
			},
		},
		{
			Title: "Nil ElementValue",
			Test: func(t *testing.T) {
				value := (*ElementValueImpl)(nil)
				strVal, err := value.GetStringValue()
				assert.Equal(t, (*string)(nil), strVal)
				assert.Nil(t, err)
			},
		},
		{
			Title: "Nil ElementValue.val",
			Test: func(t *testing.T) {
				value := &ElementValueImpl{val: nil}
				strVal, err := value.GetStringValue()
				assert.Equal(t, (*string)(nil), strVal)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Title, test.Test)
	}
}

func TestElementValue_GetBoolValue(t *testing.T) {
	tests := []struct {
		Title string
		Test  func(*testing.T)
	}{
		{
			Title: "Successful",
			Test: func(t *testing.T) {
				input := true
				value := &ElementValueImpl{val: &input}
				strVal, err := value.GetBoolValue()
				assert.Equal(t, true, *strVal)
				assert.Nil(t, err)
			},
		},
		{
			Title: "Not Bool Pointer",
			Test: func(t *testing.T) {
				input := "dfasfd"
				value := &ElementValueImpl{val: input}
				strVal, err := value.GetBoolValue()
				assert.Equal(t, (*bool)(nil), strVal)
				assert.Equal(t, fmt.Errorf("value '%v' is not compatible with type bool", input), err)
			},
		},
		{
			Title: "Nil ElementValue",
			Test: func(t *testing.T) {
				value := (*ElementValueImpl)(nil)
				strVal, err := value.GetBoolValue()
				assert.Equal(t, (*bool)(nil), strVal)
				assert.Nil(t, err)
			},
		},
		{
			Title: "Nil ElementValue.val",
			Test: func(t *testing.T) {
				value := &ElementValueImpl{val: nil}
				strVal, err := value.GetBoolValue()
				assert.Equal(t, (*bool)(nil), strVal)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Title, test.Test)
	}
}

func TestElementValue_GetInt8Value(t *testing.T) {
	tests := []struct {
		Title string
		Test  func(*testing.T)
	}{
		{
			Title: "Successful",
			Test: func(t *testing.T) {
				input := int8(4)
				value := &ElementValueImpl{val: &input}
				strVal, err := value.GetInt8Value()
				assert.Equal(t, int8(4), *strVal)
				assert.Nil(t, err)
			},
		},
		{
			Title: "Not Bool Pointer",
			Test: func(t *testing.T) {
				input := "str"
				value := &ElementValueImpl{val: input}
				strVal, err := value.GetInt8Value()
				assert.Equal(t, (*int8)(nil), strVal)
				assert.Equal(t, fmt.Errorf("value '%v' is not compatible with type int8", input), err)
			},
		},
		{
			Title: "Nil ElementValue",
			Test: func(t *testing.T) {
				value := (*ElementValueImpl)(nil)
				strVal, err := value.GetInt8Value()
				assert.Equal(t, (*int8)(nil), strVal)
				assert.Nil(t, err)
			},
		},
		{
			Title: "Nil ElementValue.val",
			Test: func(t *testing.T) {
				value := &ElementValueImpl{val: nil}
				strVal, err := value.GetInt8Value()
				assert.Equal(t, (*int8)(nil), strVal)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Title, test.Test)
	}
}

func TestElementValue_GetByteValue(t *testing.T) {
	tests := []struct {
		Title string
		Test  func(*testing.T)
	}{
		{
			Title: "Successful",
			Test: func(t *testing.T) {
				input := uint8(4)
				value := &ElementValueImpl{val: &input}
				strVal, err := value.GetByteValue()
				assert.Equal(t, uint8(4), *strVal)
				assert.Nil(t, err)
			},
		},
		{
			Title: "Not Bool Pointer",
			Test: func(t *testing.T) {
				input := "str"
				value := &ElementValueImpl{val: input}
				strVal, err := value.GetByteValue()
				assert.Equal(t, (*uint8)(nil), strVal)
				assert.Equal(t, fmt.Errorf("value '%v' is not compatible with type uint8", input), err)
			},
		},
		{
			Title: "Nil ElementValue",
			Test: func(t *testing.T) {
				value := (*ElementValueImpl)(nil)
				strVal, err := value.GetByteValue()
				assert.Equal(t, (*uint8)(nil), strVal)
				assert.Nil(t, err)
			},
		},
		{
			Title: "Nil ElementValue.val",
			Test: func(t *testing.T) {
				value := &ElementValueImpl{val: nil}
				strVal, err := value.GetByteValue()
				assert.Equal(t, (*uint8)(nil), strVal)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Title, test.Test)
	}
}

func TestElementValue_GetFloat32Value(t *testing.T) {
	tests := []struct {
		Title string
		Test  func(*testing.T)
	}{
		{
			Title: "Successful",
			Test: func(t *testing.T) {
				input := float32(4)
				value := &ElementValueImpl{val: &input}
				strVal, err := value.GetFloat32Value()
				assert.Equal(t, float32(4), *strVal)
				assert.Nil(t, err)
			},
		},
		{
			Title: "Not Bool Pointer",
			Test: func(t *testing.T) {
				input := "str"
				value := &ElementValueImpl{val: input}
				strVal, err := value.GetFloat32Value()
				assert.Equal(t, (*float32)(nil), strVal)
				assert.Equal(t, fmt.Errorf("value '%v' is not compatible with type float32", input), err)
			},
		},
		{
			Title: "Nil ElementValue",
			Test: func(t *testing.T) {
				value := (*ElementValueImpl)(nil)
				strVal, err := value.GetFloat32Value()
				assert.Equal(t, (*float32)(nil), strVal)
				assert.Nil(t, err)
			},
		},
		{
			Title: "Nil ElementValue.val",
			Test: func(t *testing.T) {
				value := &ElementValueImpl{val: nil}
				strVal, err := value.GetFloat32Value()
				assert.Equal(t, (*float32)(nil), strVal)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Title, test.Test)
	}
}

func TestElementValue_GetFloat64Value(t *testing.T) {
	tests := []struct {
		Title string
		Test  func(*testing.T)
	}{
		{
			Title: "Successful",
			Test: func(t *testing.T) {
				input := float64(4)
				value := &ElementValueImpl{val: &input}
				strVal, err := value.GetFloat64Value()
				assert.Equal(t, float64(4), *strVal)
				assert.Nil(t, err)
			},
		},
		{
			Title: "Not Bool Pointer",
			Test: func(t *testing.T) {
				input := "str"
				value := &ElementValueImpl{val: input}
				strVal, err := value.GetFloat64Value()
				assert.Equal(t, (*float64)(nil), strVal)
				assert.Equal(t, fmt.Errorf("value '%v' is not compatible with type float64", input), err)
			},
		},
		{
			Title: "Nil ElementValue",
			Test: func(t *testing.T) {
				value := (*ElementValueImpl)(nil)
				strVal, err := value.GetFloat64Value()
				assert.Equal(t, (*float64)(nil), strVal)
				assert.Nil(t, err)
			},
		},
		{
			Title: "Nil ElementValue.val",
			Test: func(t *testing.T) {
				value := &ElementValueImpl{val: nil}
				strVal, err := value.GetFloat64Value()
				assert.Equal(t, (*float64)(nil), strVal)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Title, test.Test)
	}
}

func TestElementValue_GetInt32Value(t *testing.T) {
	tests := []struct {
		Title string
		Test  func(*testing.T)
	}{
		{
			Title: "Successful",
			Test: func(t *testing.T) {
				input := int32(4)
				value := &ElementValueImpl{val: &input}
				strVal, err := value.GetInt32Value()
				assert.Equal(t, int32(4), *strVal)
				assert.Nil(t, err)
			},
		},
		{
			Title: "Not Bool Pointer",
			Test: func(t *testing.T) {
				input := "str"
				value := &ElementValueImpl{val: input}
				strVal, err := value.GetInt32Value()
				assert.Equal(t, (*int32)(nil), strVal)
				assert.Equal(t, fmt.Errorf("value '%v' is not compatible with type int32", input), err)
			},
		},
		{
			Title: "Nil ElementValue",
			Test: func(t *testing.T) {
				value := (*ElementValueImpl)(nil)
				strVal, err := value.GetInt32Value()
				assert.Equal(t, (*int32)(nil), strVal)
				assert.Nil(t, err)
			},
		},
		{
			Title: "Nil ElementValue.val",
			Test: func(t *testing.T) {
				value := &ElementValueImpl{val: nil}
				strVal, err := value.GetInt32Value()
				assert.Equal(t, (*int32)(nil), strVal)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Title, test.Test)
	}
}

func TestElementValue_GetInt64Value(t *testing.T) {
	tests := []struct {
		Title string
		Test  func(*testing.T)
	}{
		{
			Title: "Successful",
			Test: func(t *testing.T) {
				input := int64(4)
				value := &ElementValueImpl{val: &input}
				strVal, err := value.GetInt64Value()
				assert.Equal(t, int64(4), *strVal)
				assert.Nil(t, err)
			},
		},
		{
			Title: "Not Bool Pointer",
			Test: func(t *testing.T) {
				input := "str"
				value := &ElementValueImpl{val: input}
				strVal, err := value.GetInt64Value()
				assert.Equal(t, (*int64)(nil), strVal)
				assert.Equal(t, fmt.Errorf("value '%v' is not compatible with type int64", input), err)
			},
		},
		{
			Title: "Nil ElementValue",
			Test: func(t *testing.T) {
				value := (*ElementValueImpl)(nil)
				strVal, err := value.GetInt64Value()
				assert.Equal(t, (*int64)(nil), strVal)
				assert.Nil(t, err)
			},
		},
		{
			Title: "Nil ElementValue.val",
			Test: func(t *testing.T) {
				value := &ElementValueImpl{val: nil}
				strVal, err := value.GetInt64Value()
				assert.Equal(t, (*int64)(nil), strVal)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Title, test.Test)
	}
}

func TestElementValue_GetTimeValue(t *testing.T) {
	tests := []struct {
		Title string
		Test  func(*testing.T)
	}{
		{
			Title: "Successful",
			Test: func(t *testing.T) {
				input := time.Now()
				formattedInput := input.Format(time.RFC3339)
				value := &ElementValueImpl{val: &formattedInput}
				strVal, err := value.GetTimeValue()
				assert.Equal(t, input.Truncate(time.Second).UTC(), *strVal)
				assert.Nil(t, err)
			},
		},
		{
			Title: "Not Bool Pointer",
			Test: func(t *testing.T) {
				input := "str"
				value := &ElementValueImpl{val: input}
				strVal, err := value.GetTimeValue()
				assert.Equal(t, (*time.Time)(nil), strVal)
				assert.Equal(t, fmt.Errorf("value '%v' is not compatible with type time.Time", input), err)
			},
		},
		{
			Title: "Nil ElementValue",
			Test: func(t *testing.T) {
				value := (*ElementValueImpl)(nil)
				strVal, err := value.GetTimeValue()
				assert.Equal(t, (*time.Time)(nil), strVal)
				assert.Nil(t, err)
			},
		},
		{
			Title: "Not String Value",
			Test: func(t *testing.T) {
				input := true
				value := &ElementValueImpl{val: &input}
				strVal, err := value.GetTimeValue()
				assert.Equal(t, (*time.Time)(nil), strVal)
				assert.Equal(t, fmt.Errorf("value '%v' is not compatible with type string", input), err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Title, test.Test)
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
		Title string
		Test  func(*testing.T)
	}{
		{
			Title: "Successful",
			Test: func(t *testing.T) {
				input := ""
				output := (interface{})(nil)

				enumFactoryStruct := mocking.NewMockEnumFactory()
				enumFactoryStruct.On("EnumFactory", input).Return(output, nil)
				enumFactory := enumFactoryStruct.EnumFactory

				value := &ElementValueImpl{val: &input}
				val, err := value.GetEnumValue(enumFactory)
				assert.Equal(t, output, val)
				assert.Equal(t, nil, err)
			},
		},
		{
			Title: "Nil parser",
			Test: func(t *testing.T) {
				input := ""

				value := &ElementValueImpl{val: &input}
				val, err := value.GetEnumValue(nil)
				assert.Equal(t, nil, val)
				assert.Equal(t, errors.New("parser is nil"), err)
			},
		},
		{
			Title: "Nil ElementValue",
			Test: func(t *testing.T) {
				value := (*ElementValueImpl)(nil)
				strVal, err := value.GetEnumValue(nil)
				assert.Equal(t, nil, strVal)
				assert.Nil(t, err)
			},
		},
		{
			Title: "Not String Value",
			Test: func(t *testing.T) {
				input := true

				enumFactoryStruct := mocking.NewMockEnumFactory()
				enumFactory := enumFactoryStruct.EnumFactory

				value := &ElementValueImpl{val: &input}
				strVal, err := value.GetEnumValue(enumFactory)
				assert.Equal(t, nil, strVal)
				assert.Equal(t, fmt.Errorf("value '%v' is not compatible with type string", input), err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Title, test.Test)
	}
}

func toPointer[T any](val T) *T {
	return &val
}

func TestElementValue_GetCollectionOfPrimitiveValues(t *testing.T) {
	tests := []struct {
		Title string
		Test  func(*testing.T)
	}{
		{
			Title: "Successful bool",
			Test: func(t *testing.T) {
				input := []*ElementValueImpl{{val: true}, {val: false}, {val: true}}
				value := &ElementValueImpl{val: input}
				val, err := value.GetCollectionOfPrimitiveValues(PrimitiveBool)
				assert.ElementsMatch(t, []interface{}{toPointer(true), toPointer(false), toPointer(true)}, val)
				assert.Nil(t, err)
			},
		},
		{
			Title: "not primitive",
			Test: func(t *testing.T) {
				input := []*ElementValueImpl{{val: true}, {val: false}, {val: true}}
				value := &ElementValueImpl{val: input}
				val, err := value.GetCollectionOfPrimitiveValues(Primitive(-999))
				assert.Equal(t, ([]interface{})(nil), val)
				assert.Equal(t, errors.New("unknown primitive unknown"), err)
			},
		},
		{
			Title: "Unknown primitive",
			Test: func(t *testing.T) {
				input := []*ElementValueImpl{}
				value := &ElementValueImpl{val: input}
				val, err := value.GetCollectionOfPrimitiveValues(PrimitiveUnknown)
				assert.Equal(t, ([]interface{})(nil), val)
				assert.Equal(t, errors.New("target type can't be unknown"), err)
			},
		},
		{
			Title: "Unexpected collection",
			Test: func(t *testing.T) {
				input := []bool{true, false}
				value := &ElementValueImpl{val: input}
				val, err := value.GetCollectionOfPrimitiveValues(PrimitiveBool)
				assert.Equal(t, ([]interface{})(nil), val)
				assert.Equal(t, errors.New("val is not a collection"), err)
			},
		},
		{
			Title: "Nil ElementValue",
			Test: func(t *testing.T) {
				value := (*ElementValueImpl)(nil)
				strVal, err := value.GetCollectionOfPrimitiveValues(PrimitiveBool)
				assert.Equal(t, ([]interface{})(nil), strVal)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Title, test.Test)
	}
}

func TestElementValue_GetPrimitiveValue(t *testing.T) {
	tests := []struct {
		Title string
		Test  func(*testing.T)
	}{
		{
			Title: "Successful bool",
			Test: func(t *testing.T) {
				input := true
				value := &ElementValueImpl{val: &input}
				val, err := value.getPrimitiveValue(PrimitiveBool)
				assert.Equal(t, input, *(val.(*bool)))
				assert.Nil(t, err)
			},
		},
		{
			Title: "Successful string",
			Test: func(t *testing.T) {
				input := "string"
				value := &ElementValueImpl{val: &input}
				val, err := value.getPrimitiveValue(PrimitiveString)
				assert.Equal(t, input, *(val.(*string)))
				assert.Nil(t, err)
			},
		},
		{
			Title: "Successful byte",
			Test: func(t *testing.T) {
				input := byte(1)
				value := &ElementValueImpl{val: &input}
				val, err := value.getPrimitiveValue(PrimitiveByte)
				assert.Equal(t, input, *(val.(*byte)))
				assert.Nil(t, err)
			},
		},
		{
			Title: "Successful float32",
			Test: func(t *testing.T) {
				input := float32(1)
				value := &ElementValueImpl{val: &input}
				val, err := value.getPrimitiveValue(PrimitiveFloat32)
				assert.Equal(t, input, *(val.(*float32)))
				assert.Nil(t, err)
			},
		},
		{
			Title: "Successful float32",
			Test: func(t *testing.T) {
				input := float64(1)
				value := &ElementValueImpl{val: &input}
				val, err := value.getPrimitiveValue(PrimitiveFloat64)
				assert.Equal(t, input, *(val.(*float64)))
				assert.Nil(t, err)
			},
		},
		{
			Title: "Successful int32",
			Test: func(t *testing.T) {
				input := int32(1)
				value := &ElementValueImpl{val: &input}
				val, err := value.getPrimitiveValue(PrimitiveInt32)
				assert.Equal(t, input, *(val.(*int32)))
				assert.Nil(t, err)
			},
		},
		{
			Title: "Successful int64",
			Test: func(t *testing.T) {
				input := int64(1)
				value := &ElementValueImpl{val: &input}
				val, err := value.getPrimitiveValue(PrimitiveInt64)
				assert.Equal(t, input, *(val.(*int64)))
				assert.Nil(t, err)
			},
		},
		{
			Title: "Successful int8",
			Test: func(t *testing.T) {
				input := int8(1)
				value := &ElementValueImpl{val: &input}
				val, err := value.getPrimitiveValue(PrimitiveInt8)
				assert.Equal(t, input, *(val.(*int8)))
				assert.Nil(t, err)
			},
		},
		{
			Title: "Nil ElementValue",
			Test: func(t *testing.T) {
				value := &ElementValueImpl{val: nil}

				val, err := value.getPrimitiveValue(PrimitiveInt8)
				assert.Equal(t, nil, val)
				assert.Nil(t, err)
			},
		},
		{
			Title: "Unknown primitive",
			Test: func(t *testing.T) {
				input := int8(1)
				value := &ElementValueImpl{val: &input}
				val, err := value.getPrimitiveValue(PrimitiveUnknown)
				assert.Equal(t, nil, val)
				assert.Equal(t, fmt.Errorf("unknown primitive %s", PrimitiveUnknown), err)
			},
		},
		{
			Title: "Successful time value",
			Test: func(t *testing.T) {
				input := time.Now()
				formattedInput := input.Format(time.RFC3339)
				value := &ElementValueImpl{val: &formattedInput}
				val, err := value.getPrimitiveValue(PrimitiveTime)
				assert.Equal(t, input.Truncate(time.Second).UTC(), *(val.(*time.Time)))
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Title, test.Test)
	}
}

func TestElementValue_GetRawValue(t *testing.T) {
	tests := []struct {
		Title string
		Test  func(*testing.T)
	}{
		{
			Title: "Successful",
			Test: func(t *testing.T) {
				input := "test"
				pointerInput := &input

				value := &ElementValueImpl{val: pointerInput}

				val, err := value.GetRawValue()
				assert.Equal(t, pointerInput, val)
				assert.Nil(t, err)
			},
		},
		{
			Title: "Nil ElementValue",
			Test: func(t *testing.T) {
				value := &ElementValueImpl{val: nil}

				val, err := value.GetRawValue()
				assert.Equal(t, nil, val)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Title, test.Test)
	}
}
