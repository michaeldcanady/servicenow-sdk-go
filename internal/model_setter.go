package internal

type ModelSetter[T any] func(val T) error
