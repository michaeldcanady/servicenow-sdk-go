//go:build preview.query

package ast

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestKindOf(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "String",
			test: func(t *testing.T) {
				node := kindOf("value")

				assert.Equal(t, KindString, node)
			},
		},
		{
			name: "Reference",
			test: func(t *testing.T) {
				node := kindOf("fec4d97893d32210806af40bdd03d66f")

				assert.Equal(t, KindReference, node)
			},
		},
		{
			name: "bool",
			test: func(t *testing.T) {
				node := kindOf(true)

				assert.Equal(t, KindBoolean, node)
			},
		},
		{
			name: "int",
			test: func(t *testing.T) {
				node := kindOf(int(8))

				assert.Equal(t, KindNumeric, node)
			},
		},
		{
			name: "int32",
			test: func(t *testing.T) {
				node := kindOf(int32(8))

				assert.Equal(t, KindNumeric, node)
			},
		},
		{
			name: "int64",
			test: func(t *testing.T) {
				node := kindOf(int64(8))

				assert.Equal(t, KindNumeric, node)
			},
		},
		{
			name: "uint",
			test: func(t *testing.T) {
				node := kindOf(uint(8))

				assert.Equal(t, KindNumeric, node)
			},
		},
		{
			name: "uint8",
			test: func(t *testing.T) {
				node := kindOf(uint8(8))

				assert.Equal(t, KindNumeric, node)
			},
		},
		{
			name: "uint16",
			test: func(t *testing.T) {
				node := kindOf(uint16(8))

				assert.Equal(t, KindNumeric, node)
			},
		},
		{
			name: "uint32",
			test: func(t *testing.T) {
				node := kindOf(uint32(8))

				assert.Equal(t, KindNumeric, node)
			},
		},
		{
			name: "uint64",
			test: func(t *testing.T) {
				node := kindOf(uint64(8))

				assert.Equal(t, KindNumeric, node)
			},
		},
		{
			name: "float32",
			test: func(t *testing.T) {
				node := kindOf(float32(8))

				assert.Equal(t, KindNumeric, node)
			},
		},
		{
			name: "float64",
			test: func(t *testing.T) {
				node := kindOf(float64(8))

				assert.Equal(t, KindNumeric, node)
			},
		},
		{
			name: "time",
			test: func(t *testing.T) {
				node := kindOf(time.Time{})

				assert.Equal(t, KindDateTime, node)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
