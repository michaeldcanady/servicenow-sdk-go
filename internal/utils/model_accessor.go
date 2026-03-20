package utils

type ModelAccessor[T any] func() (T, error)
