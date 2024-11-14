package tableapi

import "github.com/RecoLabs/servicenow-sdk-go/internal"

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
