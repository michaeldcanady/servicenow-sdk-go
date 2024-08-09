package types

import "time"

type DateTime interface {
	Type
}

type datetime struct {
	time.Time
}

func NewDateTime(dT time.Time) DateTime {
	return &datetime{dT}
}

//TODO: Stringify
func (dT *datetime) String() string {
	return ""
}
