//go:build preview.query

package query

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
	"github.com/stretchr/testify/assert"
)

func TestQueryBuilder_And(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				builder := &QueryBuilder{
					logicalOperator: ast.OperatorUnknown,
				}

				_ = builder.And()

				assert.Equal(t, ast.OperatorAnd, builder.logicalOperator)
			},
		},
		{
			name: "Operator already set",
			test: func(t *testing.T) {
				builder := &QueryBuilder{
					logicalOperator: ast.OperatorOr,
				}

				_ = builder.And()

				assert.Equal(t, errors.Join(errors.New("logicalOperator already is set")), builder.Error)
				assert.Equal(t, ast.OperatorAnd, builder.logicalOperator)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestQueryBuilder_Or(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				builder := &QueryBuilder{
					logicalOperator: ast.OperatorUnknown,
				}

				_ = builder.Or()

				assert.Equal(t, ast.OperatorOr, builder.logicalOperator)
			},
		},
		{
			name: "Operator already set",
			test: func(t *testing.T) {
				builder := &QueryBuilder{
					logicalOperator: ast.OperatorAnd,
				}

				_ = builder.Or()

				assert.Equal(t, errors.Join(errors.New("logicalOperator already is set")), builder.Error)
				assert.Equal(t, ast.OperatorOr, builder.logicalOperator)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
