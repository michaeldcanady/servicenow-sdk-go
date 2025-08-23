//go:build preview.query

package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldBuilder_Field(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				qb := &QueryBuilder{}
				builder := &FieldBuilder{query: qb}

				conBuilder := builder.Field("field1")

				assert.Equal(t, NewUnitedConditionBuilder("field1", qb), conBuilder)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestFieldBuilder_StringField(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				qb := &QueryBuilder{}
				builder := &FieldBuilder{query: qb}

				conBuilder := builder.StringField("field1")

				assert.Equal(t, NewStringConditionBuilder("field1", qb), conBuilder)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestFieldBuilder_DateTimeField(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				qb := &QueryBuilder{}
				builder := &FieldBuilder{query: qb}

				conBuilder := builder.DateTimeField("field1")

				assert.Equal(t, NewDateTimeConditionBuilder("field1", qb), conBuilder)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestFieldBuilder_NumericField(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				qb := &QueryBuilder{}
				builder := &FieldBuilder{query: qb}

				conBuilder := builder.NumericField("field1")

				assert.Equal(t, NewNumericConditionBuilder("field1", qb), conBuilder)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
