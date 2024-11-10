package attachmentapi

import (
	"strings"
	"time"
)

const (
	DateFormat     = "2006-01-02"
	TimeFormat     = "15:04:05"
	DateTimeFormat = DateFormat + " " + TimeFormat
)

type Time time.Time

func (t *Time) UnmarshalJSON(data []byte) error {
	parsedTime, err := time.Parse(DateTimeFormat, strings.Replace(string(data), "\"", "", -1))
	if err != nil {
		return err
	}
	*t = Time(parsedTime)
	return nil
}
