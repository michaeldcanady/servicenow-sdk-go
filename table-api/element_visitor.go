//go:build preview.tableApiV2

package tableapi

import (
	"fmt"
	"reflect"
)

// elementVisitor represents a way to convert a value to element(s)
type elementVisitor interface {
	// Visit converts the provided value to an ElementValue
	Visit(any) (*ElementValue, error)
	// VisitSlice converts the provided slice into a slice of ElementValues
	VisitSlice(reflect.Value) (*ElementValue, error)
	// VisitMap converts the provided map into a map of ElementValues
	VisitMap(reflect.Value) (*ElementValue, error)
	// VisitPointer converts the pointer into a dereferenced ElementValue
	VisitPointer(reflect.Value) (*ElementValue, error)
	// VisitPrimitive converts the provided val to ElementValue
	VisitPrimitive(reflect.Value) (*ElementValue, error)
}

func newElementVis() *elementVis {
	return &elementVis{}
}

type elementVis struct{}

// Visit converts the provided value to an ElementValue
func (eV *elementVis) Visit(val any) (*ElementValue, error) {
	rv := reflect.ValueOf(val)

	switch kind := rv.Kind(); kind {
	case reflect.Array, reflect.Slice:
		return eV.VisitSlice(rv)
	case reflect.Map:
		return eV.VisitMap(rv)
	case reflect.Pointer:
		return eV.VisitPointer(rv)
	default:
		return nil, fmt.Errorf("unsupported kind %s", kind)
	}
}

// VisitSlice converts the provided slice into a slice of ElementValues
func (eV *elementVis) VisitSlice(val reflect.Value) (*ElementValue, error) {
	var err error
	array := make([]*ElementValue, val.Len())
	for i := 0; i < val.Len(); i++ {
		if array[i], err = eV.Visit(val.Index(i).Interface()); err != nil {
			return nil, err
		}
	}
	return &ElementValue{val: array}, nil
}

// VisitMap converts the provided map into a map of ElementValues
func (eV *elementVis) VisitMap(val reflect.Value) (*ElementValue, error) {
	var err error
	mapping := make(map[string]*ElementValue, val.Len())
	for _, valKey := range val.MapKeys() {
		key := valKey.Interface().(string)
		if mapping[key], err = eV.Visit(val.MapIndex(valKey).Interface()); err != nil {
			return nil, err
		}
	}
	return &ElementValue{val: mapping}, nil
}

// VisitPointer converts the pointer into a dereferenced ElementValue
func (eV *elementVis) VisitPointer(val reflect.Value) (*ElementValue, error) {
	if val.IsNil() {
		return nil, nil
	}
	return eV.Visit(val.Elem().Interface())
}

// VisitPrimitive converts the provided val to ElementValue
func (eV *elementVis) VisitPrimitive(val reflect.Value) (*ElementValue, error) {
	return &ElementValue{val: val.Interface()}, nil
}
