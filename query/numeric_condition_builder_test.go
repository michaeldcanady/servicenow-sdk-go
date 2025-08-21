//go:build preview.query

package query

import (
	"fmt"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
	"github.com/stretchr/testify/assert"
)

func TestNumericConditionBuilder_LessThan(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := float64(5)
				mockQueryBuilder := &QueryBuilder{}

				mockBuilder := newMockConBuilder[float64]()
				mockBuilder.On("binaryCondition", ast.OperatorLessThan, ast.NewLiteralNode(input)).Return(mockQueryBuilder)
				builder := &NumericConditionBuilder{mockBuilder}

				queryBuilder := builder.LessThan(input)

				assert.Equal(t, mockQueryBuilder, queryBuilder)
				mockBuilder.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestNumericConditionBuilder_LessThanOrIs(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := float64(5)
				mockQueryBuilder := &QueryBuilder{}

				mockBuilder := newMockConBuilder[float64]()
				mockBuilder.On("binaryCondition", ast.OperatorLessThanOrIs, ast.NewLiteralNode(input)).Return(mockQueryBuilder)
				builder := &NumericConditionBuilder{mockBuilder}

				queryBuilder := builder.LessThanOrIs(input)

				assert.Equal(t, mockQueryBuilder, queryBuilder)
				mockBuilder.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestNumericConditionBuilder_GreaterThanOrIs(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := float64(5)
				mockQueryBuilder := &QueryBuilder{}

				mockBuilder := newMockConBuilder[float64]()
				mockBuilder.On("binaryCondition", ast.OperatorGreaterThanOrIs, ast.NewLiteralNode(input)).Return(mockQueryBuilder)
				builder := &NumericConditionBuilder{mockBuilder}

				queryBuilder := builder.GreaterThanOrIs(input)

				assert.Equal(t, mockQueryBuilder, queryBuilder)
				mockBuilder.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestNumericConditionBuilder_Between(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				lower := float64(5)
				upper := float64(6)
				mockQueryBuilder := &QueryBuilder{}

				mockBuilder := newMockConBuilder[float64]()
				mockBuilder.On("binaryCondition", ast.OperatorBetween, ast.NewPairNode(
					ast.NewLiteralNode(lower),
					ast.NewLiteralNode(upper),
				)).Return(mockQueryBuilder)
				builder := &NumericConditionBuilder{mockBuilder}

				queryBuilder := builder.Between(lower, upper)

				assert.Equal(t, mockQueryBuilder, queryBuilder)
				mockBuilder.AssertExpectations(t)
			},
		},
		{
			name: "lower greater than upper",
			test: func(t *testing.T) {
				lower := float64(5)
				upper := float64(4)
				mockQueryBuilder := &QueryBuilder{}

				mockBuilder := newMockConBuilder[float64]()
				mockBuilder.On("binaryCondition", ast.OperatorBetween, ast.NewPairNode(
					ast.NewLiteralNode(lower),
					ast.NewLiteralNode(upper),
				)).Return(mockQueryBuilder)
				mockBuilder.On("addErrors", []error{fmt.Errorf("%v is greater or equal to %v", lower, upper)})
				builder := &NumericConditionBuilder{mockBuilder}

				queryBuilder := builder.Between(lower, upper)

				assert.Equal(t, mockQueryBuilder, queryBuilder)
				mockBuilder.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestNumericConditionBuilder_IsDifferent(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := float64(5)
				mockQueryBuilder := &QueryBuilder{}

				mockBuilder := newMockConBuilder[float64]()
				mockBuilder.On("binaryCondition", ast.OperatorIsDifferent, ast.NewLiteralNode(input)).Return(mockQueryBuilder)
				builder := &NumericConditionBuilder{mockBuilder}

				queryBuilder := builder.IsDifferent(input)

				assert.Equal(t, mockQueryBuilder, queryBuilder)
				mockBuilder.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestNumericConditionBuilder_GreaterThanField(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := "sysID"
				mockQueryBuilder := &QueryBuilder{}

				mockBuilder := newMockConBuilder[float64]()
				mockBuilder.On("binaryCondition", ast.OperatorGreaterThanField, ast.NewLiteralNode(input)).Return(mockQueryBuilder)
				builder := &NumericConditionBuilder{mockBuilder}

				queryBuilder := builder.GreaterThanField(input)

				assert.Equal(t, mockQueryBuilder, queryBuilder)
				mockBuilder.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestNumericConditionBuilder_LessThanField(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := "sysID"
				mockQueryBuilder := &QueryBuilder{}

				mockBuilder := newMockConBuilder[float64]()
				mockBuilder.On("binaryCondition", ast.OperatorLessThanField, ast.NewLiteralNode(input)).Return(mockQueryBuilder)
				builder := &NumericConditionBuilder{mockBuilder}

				queryBuilder := builder.LessThanField(input)

				assert.Equal(t, mockQueryBuilder, queryBuilder)
				mockBuilder.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestNumericConditionBuilder_GreaterThanOrIsField(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := "sysID"
				mockQueryBuilder := &QueryBuilder{}

				mockBuilder := newMockConBuilder[float64]()
				mockBuilder.On("binaryCondition", ast.OperatorGreaterThanOrIsField, ast.NewLiteralNode(input)).Return(mockQueryBuilder)
				builder := &NumericConditionBuilder{mockBuilder}

				queryBuilder := builder.GreaterThanOrIsField(input)

				assert.Equal(t, mockQueryBuilder, queryBuilder)
				mockBuilder.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestNumericConditionBuilder_LessThanOrIsField(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := "sysID"
				mockQueryBuilder := &QueryBuilder{}

				mockBuilder := newMockConBuilder[float64]()
				mockBuilder.On("binaryCondition", ast.OperatorLessThanOrIsField, ast.NewLiteralNode(input)).Return(mockQueryBuilder)
				builder := &NumericConditionBuilder{mockBuilder}

				queryBuilder := builder.LessThanOrIsField(input)

				assert.Equal(t, mockQueryBuilder, queryBuilder)
				mockBuilder.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestNumericConditionBuilder_IsMoreThan(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := float64(10)
				mockQueryBuilder := &QueryBuilder{}

				mockBuilder := newMockConBuilder[float64]()
				mockBuilder.On("binaryCondition", ast.OperatorIsMoreThan, ast.NewLiteralNode(input)).Return(mockQueryBuilder)
				builder := &NumericConditionBuilder{mockBuilder}

				queryBuilder := builder.IsMoreThan(input)

				assert.Equal(t, mockQueryBuilder, queryBuilder)
				mockBuilder.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestNumericConditionBuilder_IsLessThan(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := float64(10)
				mockQueryBuilder := &QueryBuilder{}

				mockBuilder := newMockConBuilder[float64]()
				mockBuilder.On("binaryCondition", ast.OperatorIsLessThan, ast.NewLiteralNode(input)).Return(mockQueryBuilder)
				builder := &NumericConditionBuilder{mockBuilder}

				queryBuilder := builder.IsLessThan(input)

				assert.Equal(t, mockQueryBuilder, queryBuilder)
				mockBuilder.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestNumericConditionBuilder_IsOneOf(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := float64(10)
				mockQueryBuilder := &QueryBuilder{}

				mockBuilder := newMockConBuilder[float64]()
				mockBuilder.On("binaryCondition", ast.OperatorIsOneOf, ast.NewArrayNode(ast.NewLiteralNode(input))).Return(mockQueryBuilder)
				builder := &NumericConditionBuilder{mockBuilder}

				queryBuilder := builder.IsOneOf(input)

				assert.Equal(t, mockQueryBuilder, queryBuilder)
				mockBuilder.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestNumericConditionBuilder_IsNotOneOf(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := float64(10)
				mockQueryBuilder := &QueryBuilder{}

				mockBuilder := newMockConBuilder[float64]()
				mockBuilder.On("binaryCondition", ast.OperatorIsNotOneOf, ast.NewArrayNode(ast.NewLiteralNode(input))).Return(mockQueryBuilder)
				builder := &NumericConditionBuilder{mockBuilder}

				queryBuilder := builder.IsNotOneOf(input)

				assert.Equal(t, mockQueryBuilder, queryBuilder)
				mockBuilder.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
