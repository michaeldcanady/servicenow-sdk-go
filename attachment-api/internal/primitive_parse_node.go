package internal

import (
	"encoding/base64"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/google/uuid"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	absser "github.com/microsoft/kiota-abstractions-go/serialization"
)

var _ absser.ParseNode = (*PrimitiveParseNode)(nil)

type PrimitiveParseNode struct {
	value                     interface{}
	onBeforeAssignFieldValues absser.ParsableAction
	onAfterAssignFieldValues  absser.ParsableAction
}

func NewPrimitiveParseNode(content interface{}) (*PrimitiveParseNode, error) {
	if content == nil {
		return nil, errors.New("content is nil")
	}

	return loadPrimitiveTree(content)
}

func loadPrimitiveTree(value interface{}) (*PrimitiveParseNode, error) {
	switch valueType := reflect.TypeOf(value); valueType.Kind() {
	case reflect.Map:
		parsedMap := make(map[interface{}]*PrimitiveParseNode)
		mapValue := reflect.ValueOf(value)
		for _, key := range mapValue.MapKeys() {
			elemValue := mapValue.MapIndex(key).Interface()
			parsedElem, err := loadPrimitiveTree(elemValue)
			if err != nil {
				return nil, err
			}
			parsedMap[key.Interface()] = parsedElem
		}
		return &PrimitiveParseNode{value: parsedMap}, nil

	case reflect.Slice:
		parsedSlice := make([]*PrimitiveParseNode, reflect.ValueOf(value).Len())
		sliceValue := reflect.ValueOf(value)
		for i := 0; i < sliceValue.Len(); i++ {
			elemValue := sliceValue.Index(i).Interface()
			parsedElem, err := loadPrimitiveTree(elemValue)
			if err != nil {
				return nil, err
			}
			parsedSlice[i] = parsedElem
		}
		return &PrimitiveParseNode{value: parsedSlice}, nil
	case reflect.Struct:
		return nil, errors.New("struct not implemented")
	case reflect.Bool:
		val := value.(bool)
		return &PrimitiveParseNode{value: val}, nil
	case reflect.Int:
		val := value.(int)
		return &PrimitiveParseNode{value: val}, nil
	case reflect.Int8:
		val := value.(int8)
		return &PrimitiveParseNode{value: val}, nil
	case reflect.Int16:
		val := value.(int16)
		return &PrimitiveParseNode{value: val}, nil
	case reflect.Int32:
		val := value.(int32)
		return &PrimitiveParseNode{value: val}, nil
	case reflect.Int64:
		val := value.(int64)
		return &PrimitiveParseNode{value: val}, nil
	case reflect.Uint:
		val := value.(uint)
		return &PrimitiveParseNode{value: val}, nil
	case reflect.Uint8:
		val := value.(uint8)
		return &PrimitiveParseNode{value: val}, nil
	case reflect.Uint16:
		val := value.(uint16)
		return &PrimitiveParseNode{value: val}, nil
	case reflect.Uint32:
		val := value.(uint32)
		return &PrimitiveParseNode{value: val}, nil
	case reflect.Uint64:
		val := value.(uint64)
		return &PrimitiveParseNode{value: val}, nil
	case reflect.Float32:
		val := value.(float32)
		return &PrimitiveParseNode{value: val}, nil
	case reflect.Float64:
		val := value.(float64)
		return &PrimitiveParseNode{value: val}, nil
	case reflect.String:
		val := value.(string)
		return &PrimitiveParseNode{value: val}, nil
	default:
		return nil, errors.New("unsupported type")
	}
}

// setValue sets the value represented by the node
func (n *PrimitiveParseNode) setValue(value interface{}) { //nolint:unused
	n.value = value
}

// GetChildNode returns a new parse node for the given identifier.
func (n *PrimitiveParseNode) GetChildNode(index string) (absser.ParseNode, error) {
	if index == "" {
		return nil, errors.New("index is empty")
	}

	childNodes, ok := n.value.(map[string]*PrimitiveParseNode)
	if !ok || len(childNodes) == 0 {
		return nil, nil
	}

	childNode := childNodes[index]
	if childNode != nil {
		err := childNode.SetOnBeforeAssignFieldValues(n.GetOnBeforeAssignFieldValues())
		if err != nil {
			return nil, err
		}
		err = childNode.SetOnAfterAssignFieldValues(n.GetOnAfterAssignFieldValues())
		if err != nil {
			return nil, err
		}
	}

	return childNode, nil
}

// GetCollectionOfObjectValues returns the collection of Parsable values from the node.
func (n *PrimitiveParseNode) GetCollectionOfObjectValues(ctor absser.ParsableFactory) ([]absser.Parsable, error) {
	if ctor == nil {
		return nil, errors.New("ctor is nil")
	}
	nodes, ok := n.value.([]*PrimitiveParseNode)
	if !ok {
		return nil, errors.New("value is not a collection")
	}
	result := make([]absser.Parsable, len(nodes))
	for i, v := range nodes {
		if v != nil {
			val, err := (*v).GetObjectValue(ctor)
			if err != nil {
				return nil, err
			}
			result[i] = val
		} else {
			result[i] = nil
		}
	}
	return result, nil
}

// GetCollectionOfPrimitiveValues returns the collection of primitive values from the node.
func (n *PrimitiveParseNode) GetCollectionOfPrimitiveValues(targetType string) ([]interface{}, error) {
	if targetType == "" {
		return nil, errors.New("targetType is empty")
	}
	nodes, ok := n.value.([]*PrimitiveParseNode)
	if !ok {
		return nil, errors.New("value is not a collection")
	}
	result := make([]interface{}, len(nodes))
	for i, v := range nodes {
		if v != nil {
			val, err := v.getPrimitiveValue(targetType)
			if err != nil {
				return nil, err
			}
			result[i] = val
		} else {
			result[i] = nil
		}
	}
	return result, nil
}

func (n *PrimitiveParseNode) getPrimitiveValue(targetType string) (interface{}, error) {
	switch targetType {
	case "string":
		return n.GetStringValue()
	case "bool":
		return n.GetBoolValue()
	case "uint8":
		return n.GetInt8Value()
	case "byte":
		return n.GetByteValue()
	case "float32":
		return n.GetFloat32Value()
	case "float64":
		return n.GetFloat64Value()
	case "int32":
		return n.GetInt32Value()
	case "int64":
		return n.GetInt64Value()
	case "time":
		return n.GetTimeValue()
	case "timeonly":
		return n.GetTimeOnlyValue()
	case "dateonly":
		return n.GetDateOnlyValue()
	case "isoduration":
		return n.GetISODurationValue()
	case "uuid":
		return n.GetUUIDValue()
	case "base64":
		return n.GetByteArrayValue()
	default:
		return nil, fmt.Errorf("targetType %s is not supported", targetType)
	}
}

// GetCollectionOfEnumValues returns the collection of Enum values from the node.
func (n *PrimitiveParseNode) GetCollectionOfEnumValues(parser absser.EnumFactory) ([]interface{}, error) {
	if parser == nil {
		return nil, errors.New("parser is nil")
	}
	nodes, ok := n.value.([]*PrimitiveParseNode)
	if !ok {
		return nil, errors.New("value is not a collection")
	}
	result := make([]interface{}, len(nodes))
	for i, v := range nodes {
		if v != nil {
			val, err := v.GetEnumValue(parser)
			if err != nil {
				return nil, err
			}
			result[i] = val
		} else {
			result[i] = nil
		}
	}
	return result, nil
}

// GetObjectValue returns the Parsable value from the node.
func (n *PrimitiveParseNode) GetObjectValue(ctor absser.ParsableFactory) (absser.Parsable, error) { //nolint:gocognit
	if ctor == nil {
		return nil, errors.New("constructor is nil")
	}
	result, err := ctor(n)
	if err != nil {
		return nil, err
	}
	_, isUntypedNode := result.(absser.UntypedNodeable)
	if isUntypedNode {
		switch value := n.value.(type) {
		case *bool:
			return absser.NewUntypedBoolean(*value), nil
		case *string:
			return absser.NewUntypedString(*value), nil
		case *float32:
			return absser.NewUntypedFloat(*value), nil
		case *float64:
			return absser.NewUntypedDouble(*value), nil
		case *int32:
			return absser.NewUntypedInteger(*value), nil
		case *int64:
			return absser.NewUntypedLong(*value), nil
		case nil:
			return absser.NewUntypedNull(), nil
		case map[string]*PrimitiveParseNode:
			properties := make(map[string]absser.UntypedNodeable)
			for key, value := range value {
				parsable, err := value.GetObjectValue(absser.CreateUntypedNodeFromDiscriminatorValue)
				if err != nil {
					return nil, errors.New("cannot parse object value")
				}
				if parsable == nil {
					parsable = absser.NewUntypedNull()
				}
				property, ok := parsable.(absser.UntypedNodeable)
				if ok {
					properties[key] = property
				}
			}
			return absser.NewUntypedObject(properties), nil
		case []*PrimitiveParseNode:
			collection := make([]absser.UntypedNodeable, len(value))
			for index, node := range value {
				parsable, err := node.GetObjectValue(absser.CreateUntypedNodeFromDiscriminatorValue)
				if err != nil {
					return nil, errors.New("cannot parse object value")
				}
				if parsable == nil {
					parsable = absser.NewUntypedNull()
				}
				property, ok := parsable.(absser.UntypedNodeable)
				if ok {
					collection[index] = property
				}
			}
			return absser.NewUntypedArray(collection), nil
		default:
			return absser.NewUntypedNode(value), nil
		}
	}

	abstractions.InvokeParsableAction(n.GetOnBeforeAssignFieldValues(), result)
	properties, ok := n.value.(map[string]*PrimitiveParseNode)
	fields := result.GetFieldDeserializers()
	if ok && len(properties) != 0 {
		itemAsHolder, isHolder := result.(absser.AdditionalDataHolder)
		var itemAdditionalData map[string]interface{}
		if isHolder {
			itemAdditionalData = itemAsHolder.GetAdditionalData()
			if itemAdditionalData == nil {
				itemAdditionalData = make(map[string]interface{})
				itemAsHolder.SetAdditionalData(itemAdditionalData)
			}
		}

		for key, value := range properties {
			field := fields[key]
			if value != nil {
				err := value.SetOnBeforeAssignFieldValues(n.GetOnBeforeAssignFieldValues())
				if err != nil {
					return nil, err
				}
				err = value.SetOnAfterAssignFieldValues(n.GetOnAfterAssignFieldValues())
				if err != nil {
					return nil, err
				}
			}
			if field == nil {
				if value != nil && isHolder {
					rawValue, err := value.GetRawValue()
					if err != nil {
						return nil, err
					}
					itemAdditionalData[key] = rawValue
				}
			} else {
				err := field(value)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	abstractions.InvokeParsableAction(n.GetOnAfterAssignFieldValues(), result)
	return result, nil
}

// GetStringValue returns a String value from the node.
func (n *PrimitiveParseNode) GetStringValue() (*string, error) {
	if str, ok := n.value.(string); ok {
		return &str, nil
	}
	return nil, errors.New("value is not a string")
}

// GetBoolValue returns a Bool value from the node.
func (n *PrimitiveParseNode) GetBoolValue() (*bool, error) {
	if b, ok := n.value.(bool); ok {
		return &b, nil
	}
	return nil, errors.New("value is not a bool")
}

// GetInt8Value returns an int8 value from the node.
func (n *PrimitiveParseNode) GetInt8Value() (*int8, error) {
	if i, ok := n.value.(int8); ok {
		return &i, nil
	}
	return nil, errors.New("value is not an int8")
}

// GetByteValue returns a Byte value from the node.
func (n *PrimitiveParseNode) GetByteValue() (*byte, error) {
	if b, ok := n.value.(byte); ok {
		return &b, nil
	}
	return nil, errors.New("value is not a byte")
}

// GetFloat32Value returns a Float32 value from the node.
func (n *PrimitiveParseNode) GetFloat32Value() (*float32, error) {
	if f, ok := n.value.(float32); ok {
		return &f, nil
	}
	return nil, errors.New("value is not a float32")
}

// GetFloat64Value returns a Float64 value from the node.
func (n *PrimitiveParseNode) GetFloat64Value() (*float64, error) {
	if f, ok := n.value.(float64); ok {
		return &f, nil
	}
	return nil, errors.New("value is not a float64")
}

// GetInt32Value returns an Int32 value from the node.
func (n *PrimitiveParseNode) GetInt32Value() (*int32, error) {
	if i, ok := n.value.(int32); ok {
		return &i, nil
	}
	return nil, errors.New("value is not an int32")
}

// GetInt64Value returns an Int64 value from the node.
func (n *PrimitiveParseNode) GetInt64Value() (*int64, error) {
	if i, ok := n.value.(int64); ok {
		return &i, nil
	}
	return nil, errors.New("value is not an int64")
}

// GetTimeValue returns a Time value from the node.
func (n *PrimitiveParseNode) GetTimeValue() (*time.Time, error) {
	if t, ok := n.value.(time.Time); ok {
		return &t, nil
	}
	return nil, errors.New("value is not a time.Time")
}

// GetISODurationValue returns a ISODuration value from the nodes.
func (n *PrimitiveParseNode) GetISODurationValue() (*absser.ISODuration, error) {
	v, err := n.GetStringValue()
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return absser.ParseISODuration(*v)
}

// GetTimeOnlyValue returns a TimeOnly value from the nodes.
func (n *PrimitiveParseNode) GetTimeOnlyValue() (*absser.TimeOnly, error) {
	v, err := n.GetStringValue()
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return absser.ParseTimeOnly(*v)
}

// GetDateOnlyValue returns a DateOnly value from the nodes.
func (n *PrimitiveParseNode) GetDateOnlyValue() (*absser.DateOnly, error) {
	v, err := n.GetStringValue()
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return absser.ParseDateOnly(*v)
}

// GetUUIDValue returns a UUID value from the nodes.
func (n *PrimitiveParseNode) GetUUIDValue() (*uuid.UUID, error) {
	v, err := n.GetStringValue()
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	parsed, err := uuid.Parse(*v)
	return &parsed, err
}

// GetEnumValue returns a Enum value from the nodes.
func (n *PrimitiveParseNode) GetEnumValue(parser absser.EnumFactory) (interface{}, error) {
	if parser == nil {
		return nil, errors.New("parser is nil")
	}
	s, err := n.GetStringValue()
	if err != nil {
		return nil, err
	}
	if s == nil {
		return nil, nil
	}
	return parser(*s)
}

// GetByteArrayValue returns a ByteArray value from the nodes.
func (n *PrimitiveParseNode) GetByteArrayValue() ([]byte, error) {
	s, err := n.GetStringValue()
	if err != nil {
		return nil, err
	}
	if s == nil {
		return nil, nil
	}
	return base64.StdEncoding.DecodeString(*s)
}

// GetRawValue returns the values of the node as an interface of any type.
func (n *PrimitiveParseNode) GetRawValue() (interface{}, error) {
	return n.value, nil
}

// GetOnBeforeAssignFieldValues returns a callback invoked before the node is deserialized.
func (n *PrimitiveParseNode) GetOnBeforeAssignFieldValues() absser.ParsableAction {
	return n.onBeforeAssignFieldValues
}

// SetOnBeforeAssignFieldValues sets a callback invoked before the node is deserialized.
func (n *PrimitiveParseNode) SetOnBeforeAssignFieldValues(action absser.ParsableAction) error {
	n.onBeforeAssignFieldValues = action
	return nil
}

// GetOnAfterAssignFieldValues returns a callback invoked after the node is deserialized.
func (n *PrimitiveParseNode) GetOnAfterAssignFieldValues() absser.ParsableAction {
	return n.onAfterAssignFieldValues
}

// SetOnAfterAssignFieldValues sets a callback invoked after the node is deserialized.
func (n *PrimitiveParseNode) SetOnAfterAssignFieldValues(action absser.ParsableAction) error {
	n.onAfterAssignFieldValues = action
	return nil
}
