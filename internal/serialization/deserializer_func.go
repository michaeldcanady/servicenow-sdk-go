package serialization

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeserializeMutatedStringFunc returns a deserializer function for a string value that is transformed into another type.
func DeserializeMutatedStringFunc[T any](mutator conversion.Mutator[*string, T], setter ModelSetter[T]) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		return SetMutatedValueFromSource(node.GetStringValue, setter, mutator)
	}
}

// deserializePrimitiveFunc returns a deserializer function that reads a value from a
// ParseNode via get and passes it to the setter.
func deserializePrimitiveFunc[T any](setter ModelSetter[T], get func(serialization.ParseNode) (T, error)) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		return SetValueFromSource(func() (T, error) { return get(node) }, setter)
	}
}

// DeserializeStringFunc returns a deserializer function for a string value.
func DeserializeStringFunc(setter ModelSetter[*string]) serialization.NodeParser {
	return deserializePrimitiveFunc(setter, func(n serialization.ParseNode) (*string, error) { return n.GetStringValue() })
}

// DeserializeInt64Func returns a deserializer function for an int64 value.
func DeserializeInt64Func(setter ModelSetter[*int64]) serialization.NodeParser {
	return deserializePrimitiveFunc(setter, func(n serialization.ParseNode) (*int64, error) { return n.GetInt64Value() })
}

// DeserializeInt32Func returns a deserializer function for an int32 value.
func DeserializeInt32Func(setter ModelSetter[*int32]) serialization.NodeParser {
	return deserializePrimitiveFunc(setter, func(n serialization.ParseNode) (*int32, error) { return n.GetInt32Value() })
}

// DeserializeBoolFunc returns a deserializer function for a bool value.
func DeserializeBoolFunc(setter ModelSetter[*bool]) serialization.NodeParser {
	return deserializePrimitiveFunc(setter, func(n serialization.ParseNode) (*bool, error) { return n.GetBoolValue() })
}

// DeserializeFloat64Func returns a deserializer function for a float64 value.
func DeserializeFloat64Func(setter ModelSetter[*float64]) serialization.NodeParser {
	return deserializePrimitiveFunc(setter, func(n serialization.ParseNode) (*float64, error) { return n.GetFloat64Value() })
}

// DeserializeFloat32Func returns a deserializer function for a float32 value.
func DeserializeFloat32Func(setter ModelSetter[*float32]) serialization.NodeParser {
	return deserializePrimitiveFunc(setter, func(n serialization.ParseNode) (*float32, error) { return n.GetFloat32Value() })
}

// DeserializeTimeFunc returns a deserializer function for a time.Time value.
func DeserializeTimeFunc(setter ModelSetter[*time.Time]) serialization.NodeParser {
	return deserializePrimitiveFunc(setter, func(n serialization.ParseNode) (*time.Time, error) { return n.GetTimeValue() })
}

// DeserializeTimeOnlyFunc returns a deserializer function for a TimeOnly value.
func DeserializeTimeOnlyFunc(setter ModelSetter[*serialization.TimeOnly]) serialization.NodeParser {
	return deserializePrimitiveFunc(setter, func(n serialization.ParseNode) (*serialization.TimeOnly, error) { return n.GetTimeOnlyValue() })
}

// DeserializeDateOnlyFunc returns a deserializer function for a DateOnly value.
func DeserializeDateOnlyFunc(setter ModelSetter[*serialization.DateOnly]) serialization.NodeParser {
	return deserializePrimitiveFunc(setter, func(n serialization.ParseNode) (*serialization.DateOnly, error) { return n.GetDateOnlyValue() })
}

// DeserializeUUIDFunc returns a deserializer function for a UUID value.
func DeserializeUUIDFunc(setter ModelSetter[*uuid.UUID]) serialization.NodeParser {
	return deserializePrimitiveFunc(setter, func(n serialization.ParseNode) (*uuid.UUID, error) { return n.GetUUIDValue() })
}

// DeserializeByteArrayFunc returns a deserializer function for a byte array value.
func DeserializeByteArrayFunc(setter ModelSetter[[]byte]) serialization.NodeParser {
	return deserializePrimitiveFunc(setter, func(n serialization.ParseNode) ([]byte, error) { return n.GetByteArrayValue() })
}

// DeserializeMutatedByteArrayFunc returns a deserializer function for a byte array value that is transformed into another type.
func DeserializeMutatedByteArrayFunc[T any](mutator conversion.Mutator[[]byte, T], setter ModelSetter[T]) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		return SetMutatedValueFromSource(node.GetByteArrayValue, setter, mutator)
	}
}

// DeserializeCollectionOfObjectValuesFunc returns a deserializer function for a collection of Parsable objects.
func DeserializeCollectionOfObjectValuesFunc[T serialization.Parsable](factory serialization.ParsableFactory, setter ModelSetter[[]T]) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		val, err := node.GetCollectionOfObjectValues(factory)
		if err != nil {
			return err
		}
		res := make([]T, len(val))
		for i, v := range val {
			res[i] = v.(T)
		}
		return setter(res)
	}
}

// DeserializeObjectValueFunc returns a deserializer function for a Parsable object.
func DeserializeObjectValueFunc[T serialization.Parsable](factory serialization.ParsableFactory, setter ModelSetter[T]) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		val, err := node.GetObjectValue(factory)
		if err != nil {
			return err
		}
		if val == nil {
			var nilT T
			return setter(nilT)
		}
		return setter(val.(T))
	}
}

// DeserializeEnumFunc returns a deserializer function for an enum value.
func DeserializeEnumFunc[T any](factory serialization.EnumFactory, setter ModelSetter[*T]) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		val, err := node.GetEnumValue(factory)
		if err != nil {
			return err
		}
		if val == nil {
			return setter(nil)
		}

		if v, ok := val.(*T); ok {
			return setter(v)
		}

		if v, ok := val.(T); ok {
			return setter(&v)
		}

		return errors.New("unexpected type from enum factory")
	}
}

// DeserializeAnyFunc returns a deserializer function for an any value.
func DeserializeAnyFunc(setter ModelSetter[any]) serialization.NodeParser {
	return deserializePrimitiveFunc(setter, func(n serialization.ParseNode) (any, error) { return n.GetRawValue() })
}

// DeserializeMutatedAnyFunc returns a deserializer function for an any value that is transformed into another type.
func DeserializeMutatedAnyFunc[T any](mutator conversion.Mutator[any, T], setter ModelSetter[T]) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		return SetMutatedValueFromSource(node.GetRawValue, setter, mutator)
	}
}

// DeserializeISODurationFunc returns a deserializer function for an ISODuration value.
func DeserializeISODurationFunc(setter ModelSetter[*serialization.ISODuration]) serialization.NodeParser {
	return deserializePrimitiveFunc(setter, func(n serialization.ParseNode) (*serialization.ISODuration, error) { return n.GetISODurationValue() })
}
