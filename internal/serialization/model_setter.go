package serialization

type ModelSetter[T any] func(val T) error

func SetValueFromSource[T any](source func() (T, error), setter ModelSetter[T]) error {
	return SetMutatedValueFromSource(source, setter, func(t T) (T, error) { return t, nil })
}
