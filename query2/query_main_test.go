//go:build preview.query

package query2

import (
	"strings"
	"testing"
	"time"
)

func TestTopLevel(t *testing.T) {
	t.Run("Where", func(t *testing.T) {
		if Where("f").Is("v").String() != "f=v" {
			t.Error("Where failed")
		}
	})
	t.Run("String", func(t *testing.T) {
		if String("f").Is("v").String() != "f=v" {
			t.Error("String failed")
		}
	})
	t.Run("Number", func(t *testing.T) {
		if Number("f").Is(1).String() != "f=1" {
			t.Error("Number failed")
		}
	})
	t.Run("Boolean", func(t *testing.T) {
		if Boolean("f").Is(true).String() != "f=true" {
			t.Error("Boolean failed")
		}
	})
	t.Run("Date", func(t *testing.T) {
		if Date("f").On(JS("expr")).String() != "fONjavascript:expr" {
			t.Error("Date failed")
		}
	})
	t.Run("DateTime", func(t *testing.T) {
		if DateTime("f").On(JS("expr")).String() != "fONjavascript:expr" {
			t.Error("DateTime failed")
		}
	})
	t.Run("JS", func(t *testing.T) {
		if JS("expr").String() != "javascript:expr" {
			t.Error("JS failed")
		}
	})
}

func TestAndTop(t *testing.T) {
	tests := []struct {
		name     string
		conds    []Condition
		expected string
	}{
		{"Empty", nil, ""},
		{"Multiple", []Condition{String("a").Is("1"), String("b").Is("2")}, "a=1^b=2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := And(tt.conds...)
			if tt.name == "Empty" {
				if res != nil {
					t.Error("Empty And should return nil")
				}
			} else {
				if res.String() != tt.expected {
					t.Errorf("got %s, expected %s", res.String(), tt.expected)
				}
			}
		})
	}
}

func TestOrTop(t *testing.T) {
	tests := []struct {
		name     string
		conds    []Condition
		expected string
	}{
		{"Empty", nil, ""},
		{"Multiple", []Condition{String("a").Is("1"), String("b").Is("2")}, "a=1^ORb=2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Or(tt.conds...)
			if tt.name == "Empty" {
				if res != nil {
					t.Error("Empty Or should return nil")
				}
			} else {
				if res.String() != tt.expected {
					t.Errorf("got %s, expected %s", res.String(), tt.expected)
				}
			}
		})
	}
}

func TestComplexQueries(t *testing.T) {
	tests := []struct {
		name     string
		query    Condition
		expected string
	}{
		{
			"DeeplyNested",
			Or(
				And(
					Boolean("active").Is(true),
					Or(
						Number("priority").Is(1),
						Number("priority").Is(2),
					),
				),
				And(
					String("category").Is("software"),
					String("short_description").Contains("important"),
				),
			),
			"active=true^priority=1^ORpriority=2^ORcategory=software^short_descriptionLIKEimportant",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.query.String() != tt.expected {
				t.Errorf("got %s, expected %s", tt.query.String(), tt.expected)
			}
		})
	}
}

func TestDeepErrorPropagation(t *testing.T) {
	invalid1 := Number("a").Between(10, 5)
	invalid2 := DateTime("b").Between(time.Now().Add(time.Hour), time.Now())

	tests := []struct {
		name           string
		query          Condition
		expectedErrors []string
	}{
		{
			"MultiError",
			And(
				String("ok").Is("yes"),
				Or(invalid1, invalid2),
			),
			[]string{"is greater or equal to", "is after end time"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.query.Error()
			if err == nil {
				t.Fatal("Expected error")
			}
			errStr := err.Error()
			for _, exp := range tt.expectedErrors {
				if !strings.Contains(errStr, exp) {
					t.Errorf("Expected error to contain %q, but got: %s", exp, errStr)
				}
			}
		})
	}
}
