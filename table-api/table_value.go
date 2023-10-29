package tableapi

import "reflect"

type TableValue struct {
	value interface{}
}

func (tV *TableValue) ToInt64() (int64, error) {
	return convertType[int64](tV.value)
}

func (tV *TableValue) ToFloat64() (float64, error) {
	return convertType[float64](tV.value)
}

func (tV *TableValue) ToString() (string, error) {
	return convertType[string](tV.value)
}

func (tV *TableValue) ToBool() (bool, error) {
	return convertType[bool](tV.value)
}

func (tV *TableValue) GetType() reflect.Type {
	return reflect.TypeOf(tV.value)
}
