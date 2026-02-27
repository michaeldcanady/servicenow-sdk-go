package utils

import (
	"errors"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestStringPtrToInt64Ptr(t *testing.T) {
	s4 := "4"
	sBad := "not int"
	tests := []struct {
		name     string
		input    *string
		expected int64
		err      error
	}{
		{"Successful", &s4, 4, nil},
		{"NilInput", nil, 0, errors.New("input is nil")},
		{"InvalidSyntax", &sBad, 0, strconv.ErrSyntax},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := StringPtrToInt64Ptr(tt.input)
			if tt.err != nil {
				if err == nil {
					t.Fatal("Expected error")
				}
				if !errors.Is(err, tt.err) && err.Error() != tt.err.Error() {
					t.Errorf("Expected error %v, got %v", tt.err, err)
				}
			} else {
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}
				if *res != tt.expected {
					t.Errorf("Expected %v, got %v", tt.expected, *res)
				}
			}
		})
	}
}

func TestStringPtrToFloat64Ptr(t *testing.T) {
	s4 := "4.5"
	sBad := "not float"
	tests := []struct {
		name     string
		input    *string
		expected float64
		err      error
	}{
		{"Successful", &s4, 4.5, nil},
		{"NilInput", nil, 0, errors.New("input is nil")},
		{"InvalidSyntax", &sBad, 0, strconv.ErrSyntax},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := StringPtrToFloat64Ptr(tt.input)
			if tt.err != nil {
				if err == nil {
					t.Fatal("Expected error")
				}
			} else {
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}
				if *res != tt.expected {
					t.Errorf("Expected %v, got %v", tt.expected, *res)
				}
			}
		})
	}
}

func TestStringPtrToBoolPtr(t *testing.T) {
	sTrue := "true"
	sBad := "not bool"
	tests := []struct {
		name     string
		input    *string
		expected bool
		err      error
	}{
		{"Successful", &sTrue, true, nil},
		{"NilInput", nil, false, errors.New("input is nil")},
		{"InvalidSyntax", &sBad, false, strconv.ErrSyntax},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := StringPtrToBoolPtr(tt.input)
			if tt.err != nil {
				if err == nil {
					t.Fatal("Expected error")
				}
			} else {
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}
				if *res != tt.expected {
					t.Errorf("Expected %v, got %v", tt.expected, *res)
				}
			}
		})
	}
}

func TestStringPtrToTimePtr(t *testing.T) {
	format := "2006-01-02"
	sDate := "2024-01-01"
	sBad := "bad-date"
	mutator := StringPtrToTimePtr(format)

	tests := []struct {
		name     string
		input    *string
		expected time.Time
		err      bool
	}{
		{"Successful", &sDate, time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), false},
		{"NilInput", nil, time.Time{}, true},
		{"ParseError", &sBad, time.Time{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := mutator(tt.input)
			if tt.err {
				if err == nil {
					t.Fatal("Expected error")
				}
			} else {
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}
				if !res.Equal(tt.expected) {
					t.Errorf("Expected %v, got %v", tt.expected, *res)
				}
			}
		})
	}
}

func TestStringPtrToPrimitiveSlice(t *testing.T) {
	sList := "1,2,3"
	sBad := "1,invalid,3"
	mutator := StringPtrToPrimitiveSlice[int64](",", func(s string) (int64, error) {
		i, err := strconv.ParseInt(s, 10, 64)
		return i, err
	})

	tests := []struct {
		name     string
		input    *string
		expected []int64
		err      bool
	}{
		{"Successful", &sList, []int64{1, 2, 3}, false},
		{"MutatorError", &sBad, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := mutator(tt.input)
			if tt.err {
				if err == nil {
					t.Error("Expected error")
				}
			} else {
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}
				if !reflect.DeepEqual(res, tt.expected) {
					t.Errorf("Expected %v, got %v", tt.expected, res)
				}
			}
		})
	}
}
