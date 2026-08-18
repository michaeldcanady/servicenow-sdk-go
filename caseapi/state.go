package caseapi

import (
	"strings"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
)

const (
	stateUnknown                = "unknown"
	stateAvailable              = "available"
	stateAvailableConditionally = "available conditionally"
	stateNotAvailable           = "not available"
	statePending                = "pending"
)

// State specifies the UI state for which to render the data.
type State int8

const (
	// StateUnknown represents an unknown UI state.
	StateUnknown State = iota - 1
	// StateAvailable
	StateAvailable
	// StateAvailableConditionally
	StateAvailableConditionally
	// StateNotAvailable
	StateNotAvailable
	// StatePending
	StatePending
)

// ParseState resolves the wire representation of a state to a [State].
// Matching is case-insensitive.
func ParseState(s string) (interface{}, error) {
	if state, ok := stateValues[strings.ToLower(s)]; ok {
		return state, nil
	}
	return StateUnknown, unknownEnumValueError("state", s)
}

var stateStrings = map[State]string{
	StateUnknown:                stateUnknown,
	StateAvailable:              stateAvailable,
	StateAvailableConditionally: stateAvailableConditionally,
	StateNotAvailable:           stateNotAvailable,
	StatePending:                statePending,
}

// stateValues is the lower-cased inverse of [stateStrings], used by [ParseState].
var stateValues = invertEnumStrings(stateStrings, StateUnknown)

// String returns the string representation of the State.
func (e State) String() string {
	return conversion.EnumString(stateStrings, e, stateUnknown)
}
