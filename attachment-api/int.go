package attachmentapi

import (
	"strconv"
	"strings"
)

// Deprecated: deprecated since v{unreleased}.
//
// Int ...
type Int int

func (i *Int) UnmarshalJSON(data []byte) error {
	cleanData := strings.Replace(string(data), "\"", "", -1)

	if cleanData == "" {
		cleanData = "0"
	}

	cleanInt, err := strconv.Atoi(cleanData)

	*i = Int(cleanInt)

	return err
}
