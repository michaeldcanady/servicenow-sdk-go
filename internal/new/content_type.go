package internal

import "strings"

// ContentType enum type for REST API content types
type ContentType int64

const (
	// ContentTypeUnknown represents unknown content type
	ContentTypeUnknown ContentType = iota - 1
	// ContentTypeApplicationJSON represents content type for application json
	ContentTypeApplicationJSON
)

// ParseContentType converts provided str to its ContentType or returns ContentTypeUnknown
func ParseContentType(str string) ContentType {
	str = strings.ToLower(str)

	for _, displayValue := range []ContentType{ContentTypeUnknown, ContentTypeApplicationJSON} {
		contentTypeString := strings.ToLower(displayValue.String())
		if str == contentTypeString {
			return displayValue
		}
	}

	return ContentTypeUnknown
}

// String return string representation
func (cT ContentType) String() string {
	value, ok := map[ContentType]string{
		ContentTypeUnknown:         "unknown",
		ContentTypeApplicationJSON: "application/json",
	}[cT]

	if !ok {
		return ContentTypeUnknown.String()
	}

	return value
}
