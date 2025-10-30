package conversion

import (
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringPtrToInt64Ptr(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := "4"
				result, err := StringPtrToInt64Ptr(&input)

				assert.Nil(t, err)
				assert.Equal(t, int64(4), *result)
			},
		},
		{
			name: "Nil input",
			test: func(t *testing.T) {
				result, err := StringPtrToInt64Ptr(nil)

				assert.Nil(t, result)
				assert.Equal(t, errors.New("input is nil"), err)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				input := "not int"
				result, err := StringPtrToInt64Ptr(&input)

				assert.Nil(t, result)
				assert.Equal(t, &strconv.NumError{Func: "Atoi", Num: "not int", Err: errors.New("invalid syntax")}, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestStringPtrToFloat64Ptr(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := "4"
				result, err := StringPtrToFloat64Ptr(&input)

				assert.Nil(t, err)
				assert.Equal(t, float64(4), *result)
			},
		},
		{
			name: "Nil input",
			test: func(t *testing.T) {
				result, err := StringPtrToFloat64Ptr(nil)

				assert.Nil(t, result)
				assert.Equal(t, errors.New("input is nil"), err)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				input := "not int"
				result, err := StringPtrToFloat64Ptr(&input)

				assert.Nil(t, result)
				assert.Equal(t, &strconv.NumError{Func: "ParseFloat", Num: "not int", Err: errors.New("invalid syntax")}, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestStringPtrToBoolPtr(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := "true"
				result, err := StringPtrToBoolPtr(&input)

				assert.Nil(t, err)
				assert.Equal(t, true, *result)
			},
		},
		{
			name: "Nil input",
			test: func(t *testing.T) {
				result, err := StringPtrToBoolPtr(nil)

				assert.Nil(t, result)
				assert.Equal(t, errors.New("input is nil"), err)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				input := "not bool"
				result, err := StringPtrToBoolPtr(&input)

				assert.Nil(t, result)
				assert.Equal(t, &strconv.NumError{Func: "ParseBool", Num: "not bool", Err: errors.New("invalid syntax")}, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
