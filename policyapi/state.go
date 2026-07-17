package policyapi

import (
	"fmt"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

const (
	stateUnknown  = "unknown"
	stateActive   = "active"
	stateInactive = "inactive"
)

type State int64

const (
	StateUnknown State = iota - 1
	StateActive
	StateInactive
)

var stateStrings = map[State]string{
	StateUnknown:  stateUnknown,
	StateActive:   stateActive,
	StateInactive: stateInactive,
}

func (i State) String() string {
	return conversion.EnumString(stateStrings, i, stateUnknown)
}

var _ serialization.EnumFactory = ParseState

func ParseState(v string) (interface{}, error) {
	switch v {
	case stateUnknown:
		return StateUnknown, nil
	case stateActive:
		return StateActive, nil
	case stateInactive:
		return StateInactive, nil
	default:
		return StateUnknown, fmt.Errorf("invalid state: %s", v)
	}
}
