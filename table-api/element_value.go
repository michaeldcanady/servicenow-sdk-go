package tableapi

import (
	"fmt"

	"github.com/RecoLabs/servicenow-sdk-go/internal"
)

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
