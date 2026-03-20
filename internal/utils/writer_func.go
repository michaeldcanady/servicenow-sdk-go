package utils

import "github.com/microsoft/kiota-abstractions-go/serialization"

type WriterFunc func(serialization.SerializationWriter) error
