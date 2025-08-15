//go:build preview

package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperator(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Unknown value",
			test: func(t *testing.T) {
				op := Operator(500)

				assert.Equal(t, "unknown", op.String())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
