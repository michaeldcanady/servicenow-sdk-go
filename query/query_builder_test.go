//go:build preview.query

package query

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
	"github.com/stretchr/testify/assert"
)

func TestNewQueryBuilder(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successfully",
			test: func(t *testing.T) {
				builder := NewQueryBuilder()

				assert.Nil(t, builder.query)
				assert.Nil(t, builder.Error)
				assert.Equal(t, ast.OperatorUnknown, builder.logicalOperator)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestNewQuery(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successfully",
			test: func(t *testing.T) {
				builder := NewQuery()

				assert.NotNil(t, builder.query)
				assert.Nil(t, builder.Error)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

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

func TestQueryBuilder_addCondition(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Nil query",
			test: func(t *testing.T) {
				builder := &QueryBuilder{
					logicalOperator: ast.OperatorUnknown,
				}

				node := ast.NewMockNode()
				_ = builder.addCondition(node)

				assert.Equal(t, node, builder.query)
			},
		},
		{
			name: "Unset logical operator",
			test: func(t *testing.T) {
				builder := &QueryBuilder{
					logicalOperator: ast.OperatorUnknown,
				}

				builder.query = ast.NewMockNode()

				node := ast.NewMockNode()
				_ = builder.addCondition(node)

				assert.Equal(t, errors.Join(errors.New("logicalOperator is unset")), builder.Error)
			},
		},
		{
			name: "Nil condition",
			test: func(t *testing.T) {
				builder := &QueryBuilder{
					logicalOperator: ast.OperatorUnknown,
				}

				_ = builder.addCondition(nil)

				assert.Equal(t, errors.Join(errors.New("condition is nil")), builder.Error)
			},
		},
		{
			name: "Extend query",
			test: func(t *testing.T) {
				builder := &QueryBuilder{
					logicalOperator: ast.OperatorAnd,
				}

				query := ast.NewMockNode()

				builder.query = query

				node := ast.NewMockNode()
				_ = builder.addCondition(node)

				assert.Equal(t, ast.NewBinaryNode(query, ast.OperatorAnd, node), builder.query)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestQueryBuilder_Build(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				query := ast.NewMockNode()
				builder := &QueryBuilder{
					logicalOperator: ast.OperatorUnknown,
					query:           query,
				}

				builtQuery, err := builder.Build()

				assert.Equal(t, query, builtQuery)
				assert.Nil(t, err)
			},
		},
		{
			name: "Build error",
			test: func(t *testing.T) {
				expectedErr := errors.New("logicalOperator is unset")
				builder := &QueryBuilder{
					logicalOperator: ast.OperatorUnknown,
					Error:           expectedErr,
				}

				builtQuery, err := builder.Build()

				assert.Nil(t, builtQuery)
				assert.Equal(t, expectedErr, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
