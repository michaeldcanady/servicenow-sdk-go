//go:build preview.query

package query

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
	"github.com/stretchr/testify/assert"
)

func TestStringConditionBuilder_StartsWith(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := "test"
				mockQueryBuilder := &QueryBuilder{}

				mockBuilder := newMockConBuilder[string]()
				mockBuilder.On("binaryCondition", ast.OperatorStartsWith, ast.NewLiteralNode(input)).Return(mockQueryBuilder)
				builder := &StringConditionBuilder{mockBuilder}

				queryBuilder := builder.StartsWith(input)

				assert.Equal(t, mockQueryBuilder, queryBuilder)
				mockBuilder.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestStringConditionBuilder_EndsWith(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := "test"
				mockQueryBuilder := &QueryBuilder{}

				mockBuilder := newMockConBuilder[string]()
				mockBuilder.On("binaryCondition", ast.OperatorEndsWith, ast.NewLiteralNode(input)).Return(mockQueryBuilder)
				builder := &StringConditionBuilder{mockBuilder}

				queryBuilder := builder.EndsWith(input)

				assert.Equal(t, mockQueryBuilder, queryBuilder)
				mockBuilder.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestStringConditionBuilder_Contains(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := "test"
				mockQueryBuilder := &QueryBuilder{}

				mockBuilder := newMockConBuilder[string]()
				mockBuilder.On("binaryCondition", ast.OperatorContains, ast.NewLiteralNode(input)).Return(mockQueryBuilder)
				builder := &StringConditionBuilder{mockBuilder}

				queryBuilder := builder.Contains(input)

				assert.Equal(t, mockQueryBuilder, queryBuilder)
				mockBuilder.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestStringConditionBuilder_DoesNotContain(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := "test"
				mockQueryBuilder := &QueryBuilder{}

				mockBuilder := newMockConBuilder[string]()
				mockBuilder.On("binaryCondition", ast.OperatorDoesNotContain, ast.NewLiteralNode(input)).Return(mockQueryBuilder)
				builder := &StringConditionBuilder{mockBuilder}

				queryBuilder := builder.DoesNotContain(input)

				assert.Equal(t, mockQueryBuilder, queryBuilder)
				mockBuilder.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestStringConditionBuilder_IsEmptyString(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				mockQueryBuilder := &QueryBuilder{}

				mockBuilder := newMockConBuilder[string]()
				mockBuilder.On("unaryCondition", ast.OperatorIsEmptyString).Return(mockQueryBuilder)
				builder := &StringConditionBuilder{mockBuilder}

				queryBuilder := builder.IsEmptyString()

				assert.Equal(t, mockQueryBuilder, queryBuilder)
				mockBuilder.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
