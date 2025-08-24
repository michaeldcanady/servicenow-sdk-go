//go:build preview.query

package query

import (
	"testing"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
	"github.com/stretchr/testify/assert"
)

func TestDateTimeConditionBuilder_On(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := time.Time{}
				mockQueryBuilder := &QueryBuilder{}

				mockBuilder := newMockConBuilder[time.Time]()
				mockBuilder.On("binaryCondition", ast.OperatorOn, ast.NewLiteralNode(input.String())).Return(mockQueryBuilder)
				builder := &DateTimeConditionBuilder{mockBuilder}

				queryBuilder := builder.On(input)

				assert.Equal(t, mockQueryBuilder, queryBuilder)
				mockBuilder.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestDateTimeConditionBuilder_NotOn(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := time.Time{}
				mockQueryBuilder := &QueryBuilder{}

				mockBuilder := newMockConBuilder[time.Time]()
				mockBuilder.On("binaryCondition", ast.OperatorNotOn, ast.NewLiteralNode(input.String())).Return(mockQueryBuilder)
				builder := &DateTimeConditionBuilder{mockBuilder}

				queryBuilder := builder.NotOn(input)

				assert.Equal(t, mockQueryBuilder, queryBuilder)
				mockBuilder.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestDateTimeConditionBuilder_Before(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := time.Time{}
				mockQueryBuilder := &QueryBuilder{}

				mockBuilder := newMockConBuilder[time.Time]()
				mockBuilder.On("binaryCondition", ast.OperatorBefore, ast.NewLiteralNode(input.String())).Return(mockQueryBuilder)
				builder := &DateTimeConditionBuilder{mockBuilder}

				queryBuilder := builder.Before(input)

				assert.Equal(t, mockQueryBuilder, queryBuilder)
				mockBuilder.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestDateTimeConditionBuilder_AtOrBefore(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := time.Time{}
				mockQueryBuilder := &QueryBuilder{}

				mockBuilder := newMockConBuilder[time.Time]()
				mockBuilder.On("binaryCondition", ast.OperatorAtOrBefore, ast.NewLiteralNode(input.String())).Return(mockQueryBuilder)
				builder := &DateTimeConditionBuilder{mockBuilder}

				queryBuilder := builder.AtOrBefore(input)

				assert.Equal(t, mockQueryBuilder, queryBuilder)
				mockBuilder.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestDateTimeConditionBuilder_AtOrAfter(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := time.Time{}
				mockQueryBuilder := &QueryBuilder{}

				mockBuilder := newMockConBuilder[time.Time]()
				mockBuilder.On("binaryCondition", ast.OperatorAtOrAfter, ast.NewLiteralNode(input.String())).Return(mockQueryBuilder)
				builder := &DateTimeConditionBuilder{mockBuilder}

				queryBuilder := builder.AtOrAfter(input)

				assert.Equal(t, mockQueryBuilder, queryBuilder)
				mockBuilder.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestDateTimeConditionBuilder_After(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := time.Time{}
				mockQueryBuilder := &QueryBuilder{}

				mockBuilder := newMockConBuilder[time.Time]()
				mockBuilder.On("binaryCondition", ast.OperatorAfter, ast.NewLiteralNode(input.String())).Return(mockQueryBuilder)
				builder := &DateTimeConditionBuilder{mockBuilder}

				queryBuilder := builder.After(input)

				assert.Equal(t, mockQueryBuilder, queryBuilder)
				mockBuilder.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestDateTimeConditionBuilder_TrendOnOrAfter(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := time.Time{}
				mockQueryBuilder := &QueryBuilder{}

				mockBuilder := newMockConBuilder[time.Time]()
				mockBuilder.On("binaryCondition", ast.OperatorTrendOnOrAfter, ast.NewLiteralNode(input.String())).Return(mockQueryBuilder)
				builder := &DateTimeConditionBuilder{mockBuilder}

				queryBuilder := builder.TrendOnOrAfter(input)

				assert.Equal(t, mockQueryBuilder, queryBuilder)
				mockBuilder.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestDateTimeConditionBuilder_TrendOnOrBefore(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := time.Time{}
				mockQueryBuilder := &QueryBuilder{}

				mockBuilder := newMockConBuilder[time.Time]()
				mockBuilder.On("binaryCondition", ast.OperatorTrendOnOrBefore, ast.NewLiteralNode(input.String())).Return(mockQueryBuilder)
				builder := &DateTimeConditionBuilder{mockBuilder}

				queryBuilder := builder.TrendOnOrBefore(input)

				assert.Equal(t, mockQueryBuilder, queryBuilder)
				mockBuilder.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestDateTimeConditionBuilder_TrendAfter(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := time.Time{}
				mockQueryBuilder := &QueryBuilder{}

				mockBuilder := newMockConBuilder[time.Time]()
				mockBuilder.On("binaryCondition", ast.OperatorTrendAfter, ast.NewLiteralNode(input.String())).Return(mockQueryBuilder)
				builder := &DateTimeConditionBuilder{mockBuilder}

				queryBuilder := builder.TrendAfter(input)

				assert.Equal(t, mockQueryBuilder, queryBuilder)
				mockBuilder.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestDateTimeConditionBuilder_TrendBefore(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := time.Time{}
				mockQueryBuilder := &QueryBuilder{}

				mockBuilder := newMockConBuilder[time.Time]()
				mockBuilder.On("binaryCondition", ast.OperatorTrendBefore, ast.NewLiteralNode(input.String())).Return(mockQueryBuilder)
				builder := &DateTimeConditionBuilder{mockBuilder}

				queryBuilder := builder.TrendBefore(input)

				assert.Equal(t, mockQueryBuilder, queryBuilder)
				mockBuilder.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestDateTimeConditionBuilder_TrendOn(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := time.Time{}
				mockQueryBuilder := &QueryBuilder{}

				mockBuilder := newMockConBuilder[time.Time]()
				mockBuilder.On("binaryCondition", ast.OperatorTrendOn, ast.NewLiteralNode(input.String())).Return(mockQueryBuilder)
				builder := &DateTimeConditionBuilder{mockBuilder}

				queryBuilder := builder.TrendOn(input)

				assert.Equal(t, mockQueryBuilder, queryBuilder)
				mockBuilder.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestDateTimeConditionBuilder_RelativeAfter(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := time.Time{}
				mockQueryBuilder := &QueryBuilder{}

				mockBuilder := newMockConBuilder[time.Time]()
				mockBuilder.On("binaryCondition", ast.OperatorRelativeAfter, ast.NewLiteralNode(input.String())).Return(mockQueryBuilder)
				builder := &DateTimeConditionBuilder{mockBuilder}

				queryBuilder := builder.RelativeAfter(input)

				assert.Equal(t, mockQueryBuilder, queryBuilder)
				mockBuilder.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestDateTimeConditionBuilder_RelativeBefore(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := time.Time{}
				mockQueryBuilder := &QueryBuilder{}

				mockBuilder := newMockConBuilder[time.Time]()
				mockBuilder.On("binaryCondition", ast.OperatorRelativeBefore, ast.NewLiteralNode(input.String())).Return(mockQueryBuilder)
				builder := &DateTimeConditionBuilder{mockBuilder}

				queryBuilder := builder.RelativeBefore(input)

				assert.Equal(t, mockQueryBuilder, queryBuilder)
				mockBuilder.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
