package attachmentapi

import (
	"strconv"
	"strings"
)

// Deprecated: deprecated since v1.8.0.
//
// Int ...
type Int int

func (i *Int) UnmarshalJSON(data []byte) error {
	cleanData := strings.ReplaceAll(string(data), "\"", "")

	if cleanData == "" {
		cleanData = "0"
	}

	cleanInt, err := strconv.Atoi(cleanData)

	*i = Int(cleanInt)

	return err
}
