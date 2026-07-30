package internalhttp

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestContentType_String(t *testing.T) {
	tests := []struct {
		name string
		ct   ContentType
		want string
	}{
		{name: "ApplicationJSON", ct: ContentTypeApplicationJSON, want: "application/json"},
		{name: "TextPlain", ct: ContentTypeTextPlain, want: "text/plain"},
		{name: "FormURLEncoded", ct: ContentTypeFormURLEncoded, want: "application/x-www-form-urlencoded"},
		{name: "OctetStream", ct: ContentTypeOctetStream, want: "application/octet-stream"},
		{name: "MultipartFormData", ct: ContentTypeMultipartFormData, want: "multipart/form-data"},
		{name: "Any", ct: ContentTypeAny, want: "*/*"},
		{name: "Unknown", ct: ContentTypeUnknown, want: "unknown"},
		{name: "OutOfRange", ct: ContentType(9999), want: "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.ct.String())
		})
	}
}

func TestParseContentType(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    ContentType
		wantErr bool
	}{
		{name: "application/json", input: "application/json", want: ContentTypeApplicationJSON},
		{name: "text/plain", input: "text/plain", want: ContentTypeTextPlain},
		{name: "form-urlencoded", input: "application/x-www-form-urlencoded", want: ContentTypeFormURLEncoded},
		{name: "octet-stream", input: "application/octet-stream", want: ContentTypeOctetStream},
		{name: "multipart/form-data", input: "multipart/form-data", want: ContentTypeMultipartFormData},
		{name: "any", input: "*/*", want: ContentTypeAny},
		{name: "with parameters stripped", input: "application/json; charset=utf-8", want: ContentTypeApplicationJSON},
		{name: "case insensitive", input: "APPLICATION/JSON", want: ContentTypeApplicationJSON},
		{name: "unknown type", input: "application/xml", want: ContentTypeUnknown, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseContentType(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.input)
			} else {
				require.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
