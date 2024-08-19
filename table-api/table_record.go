package tableapi

import (
	"encoding/json"
	"fmt"

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

// RecordElement represents an element within a record.
type RecordElement interface {
	GetDisplayValue() ElementValue
	GetValue() ElementValue
	GetLink() string
}

// recordElement is an implementation of RecordElement.
type recordElement struct {
	displayValue interface{}
	value        interface{}
	link         string
}

// GetDisplayValue returns the display value.
func (rE *recordElement) GetDisplayValue() ElementValue {
	if internal.IsNil(rE) || internal.IsNil(rE.displayValue) {
		return nil
	}

	return &elementValue{val: rE.displayValue}
}

// GetValue returns the raw value.
func (rE *recordElement) GetValue() ElementValue {
	if internal.IsNil(rE) || internal.IsNil(rE.value) {
		return nil
	}

	return &elementValue{val: rE.value}
}

// GetLink returns the link.
func (rE *recordElement) GetLink() string {
	if internal.IsNil(rE) {
		return ""
	}

	return rE.link
}

// ElementValue represents a generic value.
type ElementValue interface {
	GetInt64Value() (*int64, error)
	GetStringValue() (*string, error)
	GetBoolValue() (*bool, error)
	GetFloat64Value() (*float64, error)
}

// elementValue is an implementation of ElementValue.
type elementValue struct {
	val interface{}
}

func (eV *elementValue) GetInt64Value() (*int64, error) {
	if internal.IsNil(eV) || internal.IsNil(eV.val) {
		return nil, nil
	}

	var val int64

	if err := internal.As(eV.val, &val); err != nil {
		return nil, err
	}

	return &val, nil
}

func (eV *elementValue) GetStringValue() (*string, error) {
	if internal.IsNil(eV) || internal.IsNil(eV.val) {
		return nil, nil
	}

	val, ok := eV.val.(*string)
	if !ok {
		return nil, fmt.Errorf("type '%T' is not compatible with type string", eV.val)
	}
	return val, nil
}
func (eV *elementValue) GetBoolValue() (*bool, error) {
	if internal.IsNil(eV) || internal.IsNil(eV.val) {
		return nil, nil
	}

	val, ok := eV.val.(*bool)
	if !ok {
		return nil, fmt.Errorf("type '%T' is not compatible with type bool", eV.val)
	}
	return val, nil
}
func (eV *elementValue) GetFloat64Value() (*float64, error) {
	if internal.IsNil(eV) || internal.IsNil(eV.val) {
		return nil, nil
	}

	var val float64

	if err := internal.As(eV.val, &val); err != nil {
		return nil, err
	}

	return &val, nil
}
