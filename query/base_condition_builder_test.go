//go:build preview.query

package query

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
)

func TestBaseConditionBuilder_unaryCondition(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				query := newMockConditionAdder[*QueryBuilder]()
				query.On("addCondition", ast.NewUnaryNode(ast.OperatorIsNotEmpty, ast.NewLiteralNode("example"))).Return(&QueryBuilder{})
				builder := BaseConditionBuilder[any]{
					field: "example",
					query: query,
				}

				_ = builder.unaryCondition(ast.OperatorIsNotEmpty)

				query.AssertExpectations(t)
			},
		},
		{
			name: "Unknown operator",
			test: func(t *testing.T) {
				query := newMockConditionAdder[*QueryBuilder]()
				query.On("addCondition", ast.NewUnaryNode(ast.OperatorUnknown, ast.NewLiteralNode("example"))).Return(&QueryBuilder{})
				query.On("addErrors", []error{errors.Join(UnknownOperatorErr)})
				builder := BaseConditionBuilder[any]{
					field: "example",
					query: query,
				}

				_ = builder.unaryCondition(ast.OperatorUnknown)

				query.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestBaseConditionBuilder_binaryCondition(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				value := ast.NewMockNode()
				query := newMockConditionAdder[*QueryBuilder]()
				query.On("addCondition", ast.NewBinaryNode(ast.NewLiteralNode("example"), ast.OperatorIs, value)).Return(&QueryBuilder{})
				builder := BaseConditionBuilder[any]{
					field: "example",
					query: query,
				}

				_ = builder.binaryCondition(ast.OperatorIs, value)

				query.AssertExpectations(t)
			},
		},
		{
			name: "Unknown operator",
			test: func(t *testing.T) {
				value := ast.NewMockNode()
				query := newMockConditionAdder[*QueryBuilder]()
				query.On("addCondition", ast.NewBinaryNode(ast.NewLiteralNode("example"), ast.OperatorUnknown, value)).Return(&QueryBuilder{})
				query.On("addErrors", []error{errors.Join(UnknownOperatorErr)})
				builder := BaseConditionBuilder[any]{
					field: "example",
					query: query,
				}

				_ = builder.binaryCondition(ast.OperatorUnknown, value)

				query.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
