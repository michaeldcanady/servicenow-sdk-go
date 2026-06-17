package policyapi

import (
	"fmt"

	"github.com/microsoft/kiota-abstractions-go/serialization"
)

type InputStatus int64

const (
	InputStatusUnknown InputStatus = iota - 1
	InputStatusInvalid
	InputStatusValid
)

func (i InputStatus) String() string {
	str, ok := map[InputStatus]string{
		InputStatusUnknown: "unknown",
		InputStatusInvalid: "invalid",
		InputStatusValid:   "valid",
	}[i]
	if !ok {
		return InputStatusUnknown.String()
	}
	return str
}

var _ serialization.EnumFactory = ParseInputStatus

func ParseInputStatus(v string) (interface{}, error) {
	switch v {
	case "unknown":
		return InputStatusUnknown, nil
	case "invalid":
		return InputStatusInvalid, nil
	case "valid":
		return InputStatusValid, nil
	default:
		return InputStatusUnknown, fmt.Errorf("invalid input status: %s", v)
	}
}
