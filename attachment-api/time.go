package attachmentapi

import (
	"fmt"
	"strings"
	"time"
)

type Time time.Time

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, time.Time(t).Format("2006-01-02 15:04:05"))), nil
}

func (t *Time) UnmarshalJSON(data []byte) error {

	parsedTime, err := time.Parse("2006-01-02 15:04:05", strings.Replace(string(data), "\"", "", -1))
	if err != nil {
		return err
	}
	*t = Time(parsedTime)
	return nil
}
