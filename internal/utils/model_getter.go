package utils

type ModelGetter[T any] func() (T, error)
