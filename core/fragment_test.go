package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFragment(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		operator RelationalOperator
		value    interface{}
	}{
		{
			name:     "Standard fragment",
			field:    "exampleField",
			operator: Is,
			value:    "exampleValue",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fragment := NewFragment(test.field, test.operator, test.value)

			assert.Equal(t, test.field, fragment.Field)
			assert.Equal(t, test.operator, fragment.RelationalOperator)
			assert.Equal(t, test.value, fragment.Value)
		})
	}
}

func TestFragmentSetNext(t *testing.T) {
	tests := []struct {
		name            string
		fragment1       *Fragment
		fragment2       *Fragment
		logicalOperator LogicalOperator
	}{
		{
			name:            "Set next with And",
			fragment1:       NewFragment("field1", Is, "value1"),
			fragment2:       NewFragment("field2", IsNot, "value2"),
			logicalOperator: And,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.fragment1.SetNext(test.fragment2, test.logicalOperator)

			assert.Equal(t, test.fragment2, test.fragment1.next)
			assert.Equal(t, test.logicalOperator, test.fragment1.LogicalOperator)
		})
	}
}

func TestFragmentIterate(t *testing.T) {
	f1 := NewFragment("field1", Is, "value1")
	f2 := NewFragment("field2", IsNot, "value2")
	f3 := NewFragment("field3", GreaterThan, "value3")

	f1.SetNext(f2, And)
	f2.SetNext(f3, Or)

	tests := []struct {
		name           string
		startFragment  *Fragment
		iterator       func(f *Fragment) bool
		expectedLength int
	}{
		{
			name:          "Iterate all",
			startFragment: f1,
			iterator: func(f *Fragment) bool {
				return true
			},
			expectedLength: 3,
		},
		{
			name:          "Stop early",
			startFragment: f1,
			iterator: func(f *Fragment) bool {
				return false
			},
			expectedLength: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var fragments []*Fragment
			test.startFragment.Iterate(func(f *Fragment) bool {
				fragments = append(fragments, f)
				return test.iterator(f)
			})

			assert.Equal(t, test.expectedLength, len(fragments))
		})
	}
}

func TestFragmentString(t *testing.T) {
	tests := []struct {
		name     string
		fragment *Fragment
		expected string
	}{
		{
			name:     "IsNot string",
			fragment: NewFragment("exampleField", IsNot, "exampleValue"),
			expected: "exampleField!=exampleValue",
		},
		{
			name:     "IsEmpty fragment",
			fragment: NewFragment("exampleField", IsEmpty, nil),
			expected: "exampleFieldISEMPTY",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.fragment.String())
		})
	}
}
