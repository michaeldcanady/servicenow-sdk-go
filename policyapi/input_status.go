package policyapi

import (
	"fmt"

	"github.com/microsoft/kiota-abstractions-go/serialization"
)

const (
	inputStatusUnknown = "unknown"
	inputStatusInvalid = "invalid"
	inputStatusValid   = "valid"
)

type InputStatus int64

const (
	InputStatusUnknown InputStatus = iota - 1
	InputStatusInvalid
	InputStatusValid
)

func (i InputStatus) String() string {
	str, ok := map[InputStatus]string{
		InputStatusUnknown: inputStatusUnknown,
		InputStatusInvalid: inputStatusInvalid,
		InputStatusValid:   inputStatusValid,
	}[i]
	if !ok {
		return InputStatusUnknown.String()
	}
	return str
}

var _ serialization.EnumFactory = ParseInputStatus

func ParseInputStatus(v string) (interface{}, error) {
	switch v {
	case inputStatusUnknown:
		return InputStatusUnknown, nil
	case inputStatusInvalid:
		return InputStatusInvalid, nil
	case inputStatusValid:
		return InputStatusValid, nil
	default:
		return InputStatusUnknown, fmt.Errorf("invalid input status: %s", v)
	}
}
