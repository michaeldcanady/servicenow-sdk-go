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

// InputStatus represents the input status of a policy mapping.
type InputStatus int64

// InputStatusUnknown, InputStatusInvalid, and InputStatusValid are the possible values of InputStatus.
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

// String returns the string representation of the InputStatus.
func (i InputStatus) String() string {
	return conversion.EnumString(inputStatusStrings, i, inputStatusUnknown)
}

var _ serialization.EnumFactory = ParseInputStatus

// ParseInputStatus parses a string into an InputStatus.
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
