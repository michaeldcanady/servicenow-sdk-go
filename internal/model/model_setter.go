package model

type ModelSetter[T any] func(val T) error
