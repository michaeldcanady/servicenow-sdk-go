package tableapi

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
)

// DataValue is a struct that can represent different types of data, such as int, bool, string, or time.
type DataValue struct {
	value interface{}
}

// UnmarshalJSON unmarshals JSON data into a DataValue, parsing it as an int, bool, string, or time.
func (tV *DataValue) UnmarshalJSON(data []byte) error {
	var val interface{}

	err := json.Unmarshal(data, &val)
	if err != nil {
		return err
	}

	value := strings.Replace(val.(string), "\"", "", -1)

	valueInt, err := tV.parseInt(value)
	if err == nil {
		tV.value = valueInt
		return nil
	}
	valueBool, err := tV.parseBool(value)
	if err == nil {
		tV.value = valueBool
		return nil
	}
	valueTime, err := tV.parseTime(value)
	if err == nil {
		tV.value = valueTime
		return nil
	}

	return nil
}

// MarshalJSON marshals a DataValue into JSON data, using the underlying value type.
func (tV *DataValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(tV.value)
}

// parseInt tries to parse a string as an int, returning an int64 or an error.
func (tV *DataValue) parseInt(data string) (int64, error) {
	if data == "" {
		data = "0"
	}

	cleanInt, err := strconv.Atoi(data)
	if err != nil {
		return -1, err
	}

	return int64(cleanInt), nil
}

// parseBool tries to parse a string as a bool, returning a bool or an error.
func (tV *DataValue) parseBool(data string) (bool, error) {

	cleanBool, err := strconv.ParseBool(data)
	if err != nil {
		return false, err
	}

	return cleanBool, nil
}

// parseTime tries to parse a string as a time, using different formats, returning a time.Time or an error.
func (tV *DataValue) parseTime(data string) (time.Time, error) {
	// Parse as date time
	parsedTime, err := time.Parse(internal.DateTimeFormat, data)
	if err == nil {
		return parsedTime, nil
	}
	// Parse as date
	parsedTime, err = time.Parse(internal.DateFormat, data)
	if err == nil {
		return parsedTime, nil
	}
	// Parse as time
	parsedTime, err = time.Parse(internal.TimeFormat, data)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}

// Int returns tV's underlying value as an int64, or an error if the value is not an int.
func (tV *DataValue) Int() (int64, error) {
	switch v := tV.value.(type) {
	case int:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return v, nil
	default:
		return 0, fmt.Errorf("unable to convert %T to int64", tV.value)
	}
}

// Float returns tV's underlying value as a float64, or an error if the value is not a float.
func (tV *DataValue) Float() (float64, error) {
	switch v := tV.value.(type) {
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	default:
		return 0, fmt.Errorf("unable to convert %T to float64", tV.value)
	}
}

// String returns tV's underlying value as a string, or an error if the value is not a string.
func (tV *DataValue) String() (string, error) {
	return convertType[string](tV.value)
}

// Bool returns tV's underlying value as a bool, or an error if the value is not a bool.
func (tV *DataValue) Bool() (bool, error) {
	return convertType[bool](tV.value)
}

// Time returns tV's underlying value as a time, or an error if the value is not time.
func (tV *DataValue) Time() (time.Time, error) {
	return convertType[time.Time](tV.value)
}

// Type returns tV's underlying value type, using the reflect package.
func (tV *DataValue) Type() reflect.Type {
	return reflect.TypeOf(tV.value)
}
