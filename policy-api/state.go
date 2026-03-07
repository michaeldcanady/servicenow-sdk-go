package policyapi

import (
	"fmt"

	"github.com/microsoft/kiota-abstractions-go/serialization"
)

type State int64

const (
	StateUnknown State = iota - 1
	StateActive
	StateInactive
)

func (i State) String() string {
	str, ok := map[State]string{
		StateUnknown:  "unknown",
		StateActive:   "active",
		StateInactive: "inactive",
	}[i]
	if !ok {
		return StateUnknown.String()
	}
	return str
}

var _ serialization.EnumFactory = ParseState

func ParseState(v string) (interface{}, error) {
	switch v {
	case "unknown":
		return StateUnknown, nil
	case "active":
		return StateActive, nil
	case "inactive":
		return StateInactive, nil
	default:
		return StateUnknown, fmt.Errorf("invalid state: %s", v)
	}
}
