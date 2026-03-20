package utils

type SerializerFunc[T any] func(accessor ModelAccessor[T]) WriterFunc
