package attachmentapi

import (
	"strconv"
	"strings"
)

type Bool bool

func (i *Bool) UnmarshalJSON(data []byte) error {
	cleanData := strings.Replace(string(data), "\"", "", -1)

	cleanInt, err := strconv.ParseBool(cleanData)

	*i = Bool(cleanInt)

	return err
}
