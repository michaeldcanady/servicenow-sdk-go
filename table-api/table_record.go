package tableapi

import (
	"encoding/json"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
)

// TableRecord represents a record with attributes.
type TableRecord interface {
	Get(string) RecordElement
	Set(string, interface{})
	HasAttribute(string) bool
}

// tableRecord is an implementation of TableRecord.
type tableRecord struct {
	record      map[string]interface{}
	changedKeys []string
}

// Get retrieves a RecordElement for the specified field.
func (tR *tableRecord) Get(field string) RecordElement {
	if internal.IsNil(tR) || len(tR.record) == 0 || !tR.HasAttribute(field) {
		return nil
	}

	value := tR.record[field]

	elem := recordElement{}

	switch v := value.(type) {
	case map[string]interface{}:
		elem.displayValue = v["displayValue"]
		elem.value = v["value"]
		elem.link = v["link"].(string)
	case interface{}:
		elem.value = v
	}

	return &elem
}

// Set updates the value for the specified field.
func (tR *tableRecord) Set(field string, value interface{}) {
	if internal.IsNil(tR) || len(tR.record) == 0 || !tR.HasAttribute(field) {
		return
	}

	tR.record[field] = value
	tR.changedKeys = append(tR.changedKeys, field)
}

// HasAttribute checks if the field exists in the record.
func (tR *tableRecord) HasAttribute(field string) bool {
	if internal.IsNil(tR) || len(tR.record) == 0 {
		return false
	}
	_, ok := tR.record[field]
	return ok
}

func (tR *tableRecord) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &tR.record)
}
