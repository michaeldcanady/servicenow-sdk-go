package policyapi

import (
	"fmt"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
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

var inputStatusStrings = map[InputStatus]string{
	InputStatusUnknown: inputStatusUnknown,
	InputStatusInvalid: inputStatusInvalid,
	InputStatusValid:   inputStatusValid,
}

func (i InputStatus) String() string {
	return conversion.EnumString(inputStatusStrings, i, inputStatusUnknown)
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
