package kiota

import (
	abstractions "github.com/microsoft/kiota-abstractions-go/serialization"
)

func DeserializeStringFunc(setter ModelSetter[*string]) abstractions.NodeParser {
	return func(node abstractions.ParseNode) error {
		return SetValueFromSource(node.GetStringValue, setter)
	}
}

func DeserializeInt64Func(setter ModelSetter[*int64]) abstractions.NodeParser {
	return func(node abstractions.ParseNode) error {
		return SetValueFromSource(node.GetInt64Value, setter)
	}
}
