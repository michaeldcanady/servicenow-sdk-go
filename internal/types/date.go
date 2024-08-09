package types

import "time"

type Date interface {
	Type
}

type dateValue struct {
	t time.Time
}

func NewDate(t time.Time) Date {
	return &dateValue{t}
}

//TODO: Stringify
func (d *dateValue) String() string {
	return ""
}
