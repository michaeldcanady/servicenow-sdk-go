package attachmentapi

import (
	"strconv"
	"strings"
)

// Deprecated: deprecated since v1.8.0.
//
// Bool ...
type Bool bool

func (i *Bool) UnmarshalJSON(data []byte) error {
	cleanData := strings.ReplaceAll(string(data), "\"", "")

	cleanInt, err := strconv.ParseBool(cleanData)

	*i = Bool(cleanInt)

	return err
}
