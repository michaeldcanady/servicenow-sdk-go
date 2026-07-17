package serialization

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// Serialize executes a list of serializers in sequence.
func Serialize(writer serialization.SerializationWriter, serializers ...WriterFunc) error {
	for _, s := range serializers {
		if err := s(writer); err != nil {
			return err
		}
	}
	return nil
}

// SerializeMutatedStringFunc returns a serializer function for a string value that is transformed from another type.
func SerializeMutatedStringFunc[T any](key string, mutator conversion.Mutator[T, *string]) SerializerFunc[T] {
	return func(accessor ModelAccessor[T]) WriterFunc {
		return func(sw serialization.SerializationWriter) error {
			return WriteMutatedValueToSource(func(v *string) error {
				if v != nil {
					return sw.WriteStringValue(key, v)
				}
				return nil
			}, accessor, mutator)
		}
	}
}

// serializePrimitiveFunc returns a serializer function that writes a value via write,
// skipping the write entirely when the value is nil.
func serializePrimitiveFunc[T any](key string, write func(serialization.SerializationWriter, string, T) error) SerializerFunc[T] {
	return func(accessor ModelAccessor[T]) WriterFunc {
		return func(sw serialization.SerializationWriter) error {
			return WriteValueToSource(func(v T) error {
				if conversion.IsNil(v) {
					return nil
				}
				return write(sw, key, v)
			}, accessor)
		}
	}
}

// SerializeStringFunc returns a serializer function for a string value.
func SerializeStringFunc(key string) SerializerFunc[*string] {
	return serializePrimitiveFunc(key, func(sw serialization.SerializationWriter, k string, v *string) error {
		return sw.WriteStringValue(k, v)
	})
}

// SerializeBoolFunc returns a serializer function for a bool value.
func SerializeBoolFunc(key string) SerializerFunc[*bool] {
	return serializePrimitiveFunc(key, func(sw serialization.SerializationWriter, k string, v *bool) error {
		return sw.WriteBoolValue(k, v)
	})
}

// SerializeInt64Func returns a serializer function for an int64 value.
func SerializeInt64Func(key string) SerializerFunc[*int64] {
	return serializePrimitiveFunc(key, func(sw serialization.SerializationWriter, k string, v *int64) error {
		return sw.WriteInt64Value(k, v)
	})
}

// SerializeInt32Func returns a serializer function for an int32 value.
func SerializeInt32Func(key string) SerializerFunc[*int32] {
	return serializePrimitiveFunc(key, func(sw serialization.SerializationWriter, k string, v *int32) error {
		return sw.WriteInt32Value(k, v)
	})
}

// SerializeFloat64Func returns a serializer function for a float64 value.
func SerializeFloat64Func(key string) SerializerFunc[*float64] {
	return serializePrimitiveFunc(key, func(sw serialization.SerializationWriter, k string, v *float64) error {
		return sw.WriteFloat64Value(k, v)
	})
}

// SerializeFloat32Func returns a serializer function for a float32 value.
func SerializeFloat32Func(key string) SerializerFunc[*float32] {
	return serializePrimitiveFunc(key, func(sw serialization.SerializationWriter, k string, v *float32) error {
		return sw.WriteFloat32Value(k, v)
	})
}

// SerializeTimeFunc returns a serializer function for a time.Time value.
func SerializeTimeFunc(key string) SerializerFunc[*time.Time] {
	return serializePrimitiveFunc(key, func(sw serialization.SerializationWriter, k string, v *time.Time) error {
		return sw.WriteTimeValue(k, v)
	})
}

// SerializeObjectValueFunc returns a serializer function for a Parsable object.
func SerializeObjectValueFunc[T serialization.Parsable](key string) SerializerFunc[T] {
	return serializePrimitiveFunc(key, func(sw serialization.SerializationWriter, k string, v T) error {
		return sw.WriteObjectValue(k, v)
	})
}

// SerializeByteArrayFunc returns a serializer function for a byte array value.
func SerializeByteArrayFunc(key string) SerializerFunc[[]byte] {
	return serializePrimitiveFunc(key, func(sw serialization.SerializationWriter, k string, v []byte) error {
		return sw.WriteByteArrayValue(k, v)
	})
}

// SerializeCollectionOfObjectValuesFunc returns a serializer function for a collection of Parsable objects.
func SerializeCollectionOfObjectValuesFunc[T serialization.Parsable](key string) SerializerFunc[[]T] {
	return func(accessor ModelAccessor[[]T]) WriterFunc {
		return func(sw serialization.SerializationWriter) error {
			return WriteValueToSource(func(v []T) error {
				if v != nil {
					parsables := make([]serialization.Parsable, len(v))
					for i, val := range v {
						parsables[i] = val
					}
					return sw.WriteCollectionOfObjectValues(key, parsables)
				}
				return nil
			}, accessor)
		}
	}
}

// SerializeCollectionOfStringValuesFunc returns a serializer function for a collection of string values.
func SerializeCollectionOfStringValuesFunc(key string) SerializerFunc[[]string] {
	return func(accessor ModelAccessor[[]string]) WriterFunc {
		return func(sw serialization.SerializationWriter) error {
			return WriteValueToSource(func(v []string) error {
				if v != nil {
					return sw.WriteCollectionOfStringValues(key, v)
				}
				return nil
			}, accessor)
		}
	}
}

// SerializeEnumFunc returns a serializer function for an enum value that has a String() method.
func SerializeEnumFunc[T any, P interface {
	*T
	String() string
}](key string) SerializerFunc[*T] {
	return func(accessor ModelAccessor[*T]) WriterFunc {
		return func(sw serialization.SerializationWriter) error {
			return WriteMutatedValueToSource(func(v *string) error {
				if v != nil {
					return sw.WriteStringValue(key, v)
				}
				return nil
			}, accessor, func(val *T) (*string, error) {
				if val == nil {
					return nil, nil
				}
				cast := P(val).String()
				return &cast, nil
			})
		}
	}
}

// SerializeStringToBoolFunc returns a serializer function for a bool value that is serialized as a string.
func SerializeStringToBoolFunc(key string) SerializerFunc[*bool] {
	return SerializeMutatedStringFunc(key, func(val *bool) (*string, error) {
		if val == nil {
			return nil, nil
		}
		str := fmt.Sprintf("%v", *val)
		return &str, nil
	})
}

// SerializeStringToFloat64Func returns a serializer function for a float64 value that is serialized as a string.
func SerializeStringToFloat64Func(key string) SerializerFunc[*float64] {
	return SerializeMutatedStringFunc(key, func(val *float64) (*string, error) {
		if val == nil {
			return nil, nil
		}
		str := strconv.FormatFloat(*val, 'f', -1, 64)
		return &str, nil
	})
}

// SerializeStringToInt64Func returns a serializer function for an int64 value that is serialized as a string.
func SerializeStringToInt64Func(key string) SerializerFunc[*int64] {
	return SerializeMutatedStringFunc(key, func(val *int64) (*string, error) {
		if val == nil {
			return nil, nil
		}
		str := fmt.Sprintf("%v", *val)
		return &str, nil
	})
}

// SerializeStringToTimeFunc returns a serializer function for a time.Time value that is serialized as a string.
func SerializeStringToTimeFunc(key string, layout string) SerializerFunc[*time.Time] {
	return SerializeMutatedStringFunc(key, func(val *time.Time) (*string, error) {
		if val == nil {
			return nil, nil
		}
		str := val.Format(layout)
		return &str, nil
	})
}

// SerializeISODurationFunc returns a serializer function for an ISODuration value.
func SerializeISODurationFunc(key string) SerializerFunc[*serialization.ISODuration] {
	return func(accessor ModelAccessor[*serialization.ISODuration]) WriterFunc {
		return func(sw serialization.SerializationWriter) error {
			return WriteValueToSource(func(v *serialization.ISODuration) error {
				if v != nil {
					return sw.WriteISODurationValue(key, v)
				}
				return nil
			}, accessor)
		}
	}
}

// SerializeAnyFunc returns a serializer function for an any value.
func SerializeAnyFunc(key string) SerializerFunc[any] {
	return func(accessor ModelAccessor[any]) WriterFunc {
		return func(sw serialization.SerializationWriter) error {
			return WriteValueToSource(func(v any) error {
				return sw.WriteAnyValue(key, v)
			}, accessor)
		}
	}
}

// SerializeStringToSliceFunc returns a serializer function for a slice of strings that is serialized as a single string joined by a separator.
func SerializeStringToSliceFunc(key string, separator string) SerializerFunc[[]string] {
	return SerializeMutatedStringFunc(key, func(val []string) (*string, error) {
		if val == nil {
			return nil, nil
		}
		str := strings.Join(val, separator)
		return &str, nil
	})
}
