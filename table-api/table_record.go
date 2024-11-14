package tableapi

import (
	"encoding/json"

	"github.com/RecoLabs/servicenow-sdk-go/internal"
)

var _ TableRecord = (*TableRecordImpl)(nil)

// TableRecord represents a record with attributes.
type TableRecord interface {
	Get(string) RecordElement
	Set(string, interface{})
	HasAttribute(string) bool
}

// TableRecordImpl is an implementation of TableRecord.
type TableRecordImpl struct {
	record      map[string]interface{}
	changedKeys []string
}

// Get retrieves a RecordElement for the specified field.
func (tR *TableRecordImpl) Get(field string) RecordElement {
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
func (tR *TableRecordImpl) Set(field string, value interface{}) {
	if internal.IsNil(tR) || len(tR.record) == 0 || !tR.HasAttribute(field) {
		return
	}

	tR.record[field] = value
	tR.changedKeys = append(tR.changedKeys, field)
}

// HasAttribute checks if the field exists in the record.
func (tR *TableRecordImpl) HasAttribute(field string) bool {
	if internal.IsNil(tR) || len(tR.record) == 0 {
		return false
	}
	_, ok := tR.record[field]
	return ok
}

func (tR *TableRecordImpl) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &tR.record)
}
