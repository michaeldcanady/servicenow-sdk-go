package http

import (
	"fmt"
	"strings"
)

const (
	contentTypeSeparator         = ';'
	contentTypeApplicationJSON   = "application/json"
	contentTypeTextPlain         = "text/plain"
	contentTypeFormURLEncoded    = "application/x-www-form-urlencoded"
	contentTypeOctetStream       = "application/octet-stream"
	contentTypeMultipartFormData = "multipart/form-data"
	contentTypeAny               = "*/*"
	contentTypeUnknown           = "unknown"
)

type ContentType int64

const (
	ContentTypeUnknown ContentType = iota - 1
	ContentTypeApplicationJSON
	ContentTypeTextPlain
	ContentTypeFormURLEncoded
	ContentTypeOctetStream
	ContentTypeMultipartFormData
	ContentTypeAny
)

func (ct ContentType) String() string {
	switch ct {
	case ContentTypeApplicationJSON:
		return contentTypeApplicationJSON
	case ContentTypeTextPlain:
		return contentTypeTextPlain
	case ContentTypeFormURLEncoded:
		return contentTypeFormURLEncoded
	case ContentTypeOctetStream:
		return contentTypeOctetStream
	case ContentTypeMultipartFormData:
		return contentTypeMultipartFormData
	case ContentTypeAny:
		return contentTypeAny
	default:
		return contentTypeUnknown
	}
}

func ParseContentType(s string) (ContentType, error) {
	// Strip parameters: "application/json; charset=utf-8" → "application/json"
	if i := strings.IndexByte(s, contentTypeSeparator); i != -1 {
		s = strings.TrimSpace(s[:i])
	}
	switch strings.ToLower(s) {
	case contentTypeApplicationJSON:
		return ContentTypeApplicationJSON, nil
	case contentTypeTextPlain:
		return ContentTypeTextPlain, nil
	case contentTypeFormURLEncoded:
		return ContentTypeFormURLEncoded, nil
	case contentTypeOctetStream:
		return ContentTypeOctetStream, nil
	case contentTypeMultipartFormData:
		return ContentTypeMultipartFormData, nil
	case contentTypeAny:
		return ContentTypeAny, nil
	default:
		return ContentTypeUnknown, fmt.Errorf("http: unknown content type %q", s)
	}
}
