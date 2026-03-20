package utils

import "github.com/microsoft/kiota-abstractions-go/serialization"

type DeserializerFunc[T any] func(setter ModelSetter[T]) serialization.NodeParser
