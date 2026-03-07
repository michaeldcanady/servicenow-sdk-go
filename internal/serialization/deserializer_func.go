package serialization

import (
	"errors"

	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeserializeMutatedStringFunc returns a deserializer function for a string value that is transformed into another type.
func DeserializeMutatedStringFunc[T any](mutator Mutator[*string, T]) DeserializerFunc[T] {
	return func(setter ModelSetter[T]) serialization.NodeParser {
		return func(node serialization.ParseNode) error {
			return SetMutatedValueFromSource(node.GetStringValue, setter, mutator)
		}
	}
}

// DeserializeStringFunc returns a deserializer function for a string value.
func DeserializeStringFunc() DeserializerFunc[*string] {
	return func(setter ModelSetter[*string]) serialization.NodeParser {
		return func(node serialization.ParseNode) error {
			return SetValueFromSource(node.GetStringValue, setter)
		}
	}
}

// DeserializeInt64Func returns a deserializer function for an int64 value.
func DeserializeInt64Func() DeserializerFunc[*int64] {
	return func(setter ModelSetter[*int64]) serialization.NodeParser {
		return func(node serialization.ParseNode) error {
			return SetValueFromSource(node.GetInt64Value, setter)
		}
	}
}

// DeserializeInt32Func returns a deserializer function for an int32 value.
func DeserializeInt32Func() DeserializerFunc[*int32] {
	return func(setter ModelSetter[*int32]) serialization.NodeParser {
		return func(node serialization.ParseNode) error {
			return SetValueFromSource(node.GetInt32Value, setter)
		}
	}
}

// DeserializeBoolFunc returns a deserializer function for a bool value.
func DeserializeBoolFunc() DeserializerFunc[*bool] {
	return func(setter ModelSetter[*bool]) serialization.NodeParser {
		return func(node serialization.ParseNode) error {
			return SetValueFromSource(node.GetBoolValue, setter)
		}
	}
}

// DeserializeFloat64Func returns a deserializer function for a float64 value.
func DeserializeFloat64Func() DeserializerFunc[*float64] {
	return func(setter ModelSetter[*float64]) serialization.NodeParser {
		return func(node serialization.ParseNode) error {
			return SetValueFromSource(node.GetFloat64Value, setter)
		}
	}
}

// DeserializeFloat32Func returns a deserializer function for a float32 value.
func DeserializeFloat32Func() DeserializerFunc[*float32] {
	return func(setter ModelSetter[*float32]) serialization.NodeParser {
		return func(node serialization.ParseNode) error {
			return SetValueFromSource(node.GetFloat32Value, setter)
		}
	}
}

// DeserializeByteArrayFunc returns a deserializer function for a byte array value.
func DeserializeByteArrayFunc() DeserializerFunc[[]byte] {
	return func(setter ModelSetter[[]byte]) serialization.NodeParser {
		return func(node serialization.ParseNode) error {
			return SetValueFromSource(node.GetByteArrayValue, setter)
		}
	}
}

// DeserializeMutatedByteArrayFunc returns a deserializer function for a byte array value that is transformed into another type.
func DeserializeMutatedByteArrayFunc[T any](mutator Mutator[[]byte, T]) DeserializerFunc[T] {
	return func(setter ModelSetter[T]) serialization.NodeParser {
		return func(node serialization.ParseNode) error {
			return SetMutatedValueFromSource(node.GetByteArrayValue, setter, mutator)
		}
	}
}

// DeserializeCollectionOfObjectValuesFunc returns a deserializer function for a collection of Parsable objects.
func DeserializeCollectionOfObjectValuesFunc[T serialization.Parsable](factory serialization.ParsableFactory) DeserializerFunc[[]T] {
	return func(setter ModelSetter[[]T]) serialization.NodeParser {
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
}

// DeserializeObjectValueFunc returns a deserializer function for a Parsable object.
func DeserializeObjectValueFunc[T serialization.Parsable](factory serialization.ParsableFactory) DeserializerFunc[T] {
	return func(setter ModelSetter[T]) serialization.NodeParser {
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
}

// DeserializeEnumFunc returns a deserializer function for an enum value.
func DeserializeEnumFunc[T any](factory serialization.EnumFactory) DeserializerFunc[*T] {
	return func(setter ModelSetter[*T]) serialization.NodeParser {
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
}

// DeserializeAnyFunc returns a deserializer function for an any value.
func DeserializeAnyFunc() DeserializerFunc[any] {
	return func(setter ModelSetter[any]) serialization.NodeParser {
		return func(node serialization.ParseNode) error {
			return SetValueFromSource(node.GetRawValue, setter)
		}
	}
}

// DeserializeMutatedAnyFunc returns a deserializer function for an any value that is transformed into another type.
func DeserializeMutatedAnyFunc[T any](mutator Mutator[any, T]) DeserializerFunc[T] {
	return func(setter ModelSetter[T]) serialization.NodeParser {
		return func(node serialization.ParseNode) error {
			return SetMutatedValueFromSource(node.GetRawValue, setter, mutator)
		}
	}
}

// DeserializeISODurationFunc returns a deserializer function for an ISODuration value.
func DeserializeISODurationFunc() DeserializerFunc[*serialization.ISODuration] {
	return func(setter ModelSetter[*serialization.ISODuration]) serialization.NodeParser {
		return func(node serialization.ParseNode) error {
			return SetValueFromSource(node.GetISODurationValue, setter)
		}
	}
}
