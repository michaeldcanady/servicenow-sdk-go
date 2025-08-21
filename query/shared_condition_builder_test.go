//go:build preview.query

package query

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
	"github.com/stretchr/testify/assert"
)

func TestSharedConditionBuilder_Is(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				mockQueryBuilder := &QueryBuilder{}
				mockBuilder := newMockConditionBuilder()
				mockBuilder.On("binaryCondition", ast.OperatorIs, ast.NewLiteralNode("test")).Return(mockQueryBuilder)

				builder := &SharedConditionBuilder[string]{mockBuilder}

				queryBuilder := builder.Is("test")

				mockBuilder.AssertExpectations(t)
				assert.Equal(t, mockQueryBuilder, queryBuilder)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestSharedConditionBuilder_IsNot(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				mockQueryBuilder := &QueryBuilder{}
				mockBuilder := newMockConditionBuilder()
				mockBuilder.On("binaryCondition", ast.OperatorIsNot, ast.NewLiteralNode("test")).Return(mockQueryBuilder)

				builder := &SharedConditionBuilder[string]{mockBuilder}

				queryBuilder := builder.IsNot("test")

				mockBuilder.AssertExpectations(t)
				assert.Equal(t, mockQueryBuilder, queryBuilder)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestSharedConditionBuilder_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				mockQueryBuilder := &QueryBuilder{}
				mockBuilder := newMockConditionBuilder()
				mockBuilder.On("unaryCondition", ast.OperatorIsEmpty).Return(mockQueryBuilder)

				builder := &SharedConditionBuilder[int]{mockBuilder}

				queryBuilder := builder.IsEmpty()

				mockBuilder.AssertExpectations(t)
				assert.Equal(t, mockQueryBuilder, queryBuilder)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestSharedConditionBuilder_IsNotEmpty(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				mockQueryBuilder := &QueryBuilder{}
				mockBuilder := newMockConditionBuilder()
				mockBuilder.On("unaryCondition", ast.OperatorIsNotEmpty).Return(mockQueryBuilder)

				builder := &SharedConditionBuilder[int]{mockBuilder}

				queryBuilder := builder.IsNotEmpty()

				mockBuilder.AssertExpectations(t)
				assert.Equal(t, mockQueryBuilder, queryBuilder)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestSharedConditionBuilder_IsAnything(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				mockQueryBuilder := &QueryBuilder{}
				mockBuilder := newMockConditionBuilder()
				mockBuilder.On("unaryCondition", ast.OperatorIsAnything).Return(mockQueryBuilder)

				builder := &SharedConditionBuilder[int]{mockBuilder}

				queryBuilder := builder.IsAnything()

				mockBuilder.AssertExpectations(t)
				assert.Equal(t, mockQueryBuilder, queryBuilder)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestSharedConditionBuilder_IsDynamic(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				mockQueryBuilder := &QueryBuilder{}
				mockBuilder := newMockConditionBuilder()
				mockBuilder.On("binaryCondition", ast.OperatorIsDynamic, ast.NewLiteralNode("test")).Return(mockQueryBuilder)

				builder := &SharedConditionBuilder[int]{mockBuilder}

				queryBuilder := builder.IsDynamic("test")

				mockBuilder.AssertExpectations(t)
				assert.Equal(t, mockQueryBuilder, queryBuilder)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestSharedConditionBuilder_IsSame(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				mockQueryBuilder := &QueryBuilder{}
				mockBuilder := newMockConditionBuilder()
				mockBuilder.On("binaryCondition", ast.OperatorIsSame, ast.NewLiteralNode("test")).Return(mockQueryBuilder)

				builder := &SharedConditionBuilder[int]{mockBuilder}

				queryBuilder := builder.IsSame("test")

				mockBuilder.AssertExpectations(t)
				assert.Equal(t, mockQueryBuilder, queryBuilder)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
