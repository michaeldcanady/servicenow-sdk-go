package attachmentapi

import (
	"strings"
	"time"
)

type Time time.Time

func (t *Time) UnmarshalJSON(data []byte) error {

	parsedTime, err := time.Parse("2006-01-02 15:04:05", strings.Replace(string(data), "\"", "", -1))
	if err != nil {
		return err
	}
	*t = Time(parsedTime)
	return nil
}
