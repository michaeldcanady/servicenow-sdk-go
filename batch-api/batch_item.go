package batchapi

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
)

type BatchItem interface {
	GetID() *string
	SetID(value *string)
	GetMethod() *string
	SetMethod(value *string)
	GetURL() *string
	SetURL(value *string)
	GetHeaders() internal.RequestHeader
	SetHeaders(value internal.RequestHeader)
	GetBody() internal.RequestBody
	SetBody(value internal.RequestBody)
	GetExcludeResponseHeaders() bool
	SetExcludeResponseHeaders(bool)
}

type batchItem struct {
	ID                     *string       `json:"id"`
	Method                 *string       `json:"method"`
	ExcludeResponseHeaders bool          `json:"exclude_response_headers"`
	URL                    *string       `json:"url"`
	Headers                []batchHeader `json:"headers"`
	Body                   string        `json:"body,omitempty"`
}

// NewBatchItem creates an instance of BatchItem
func NewBatchItem(excludeResponseHeaders bool) BatchItem {
	return &batchItem{
		ExcludeResponseHeaders: excludeResponseHeaders,
		Headers:                make([]batchHeader, 0),
	}
}

// GetId returns batch item `id` property
func (bi *batchItem) GetID() *string {
	return bi.ID
}

// SetID sets string value as batch item `id` property
func (bi *batchItem) SetID(value *string) {
	bi.ID = value
}

// GetMethod returns batch item `Method` property
func (bi *batchItem) GetMethod() *string {
	return bi.Method
}

// SetMethod sets string value as batch item `Method` property
func (bi *batchItem) SetMethod(value *string) {
	bi.Method = value
}

// GetURL returns batch item `Url` property
func (bi *batchItem) GetURL() *string {
	return bi.URL
}

// SetURL sets string value as batch item `Url` property
func (bi *batchItem) SetURL(value *string) {
	bi.URL = value
}

// GetHeaders returns batch item `Header` as a map[string]string
func (bi *batchItem) GetHeaders() internal.RequestHeader {
	headers := internal.NewRequestHeader()

	for _, header := range bi.Headers {
		headers.Set(header.Name, header.Value)
	}

	return headers
}

// SetHeaders sets map[string]string value as batch item `Header` property
func (bi *batchItem) SetHeaders(value internal.RequestHeader) {
	if value.Get(internal.ContentTypeHeader) == "" {
		value.Set(internal.ContentTypeHeader, internal.JSONContentType)
	}

	if value.Get(internal.AcceptHeader) == "" {
		value.Set(internal.AcceptHeader, internal.JSONContentType)
	}

	headers := make([]batchHeader, 0)

	value.Iterate(func(name string, value []string) bool {
		headers = append(headers, batchHeader{
			Name:  name,
			Value: strings.Join(value, ","),
		})
		return true
	})

	bi.Headers = headers
}

// GetBody returns batch item `RequestBody` property
func (bi *batchItem) GetBody() internal.RequestBody {
	var ret internal.RequestBody

	data, _ := base64.StdEncoding.DecodeString(bi.Body)

	_ = json.Unmarshal(data, &ret)

	return ret
}

// SetBody sets map[string]string value as batch item `RequestBody` property
func (bi *batchItem) SetBody(value internal.RequestBody) {
	data, _ := json.Marshal(value)

	bi.Body = base64.StdEncoding.EncodeToString(data)
}

// GetExcludeResponseHeaders returns if batch item wasn't response headers excluded
func (bi *batchItem) GetExcludeResponseHeaders() bool {
	return bi.ExcludeResponseHeaders
}

// GetExcludeResponseHeaders sets if to exclude batch item response headers
func (bi *batchItem) SetExcludeResponseHeaders(excludeResponseHeaders bool) {
	bi.ExcludeResponseHeaders = excludeResponseHeaders
}
