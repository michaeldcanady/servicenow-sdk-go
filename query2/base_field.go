//go:build preview.query

package query2

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast2"
)

//https://www.servicenow.com/docs/r/platform-user-interface/r_OpAvailableFiltersQueries.html

// BaseField represents the common properties and methods for all ServiceNow fields.
type BaseField struct {
	name string
}

func (f BaseField) unary(op ast2.Operator) Condition {
	return NewCondition(ast2.NewUnaryNode(op, ast2.NewLiteralNode(f.name)))
}

func (f BaseField) buildBinary(op ast2.Operator, right ast2.Node) Condition {
	return NewCondition(ast2.NewBinaryNode(ast2.NewLiteralNode(f.name), op, right))
}

func (f BaseField) binary(op ast2.Operator, val interface{}) Condition {
	return f.buildBinary(op, ast2.NewLiteralNode(val))
}

func (f BaseField) pair(op ast2.Operator, left, right interface{}) Condition {
	return f.buildBinary(op, ast2.NewPairNode(ast2.NewLiteralNode(left), ast2.NewLiteralNode(right)))
}

func (f BaseField) multi(op ast2.Operator, nodes *ast2.ArrayNode) Condition {
	return f.buildBinary(op, nodes)
}

// IsAnything query that field is anything.
func (f BaseField) IsAnything() Condition { return f.unary(ast2.OperatorIsAnything) }

// IsEmpty query that field is empty.
func (f BaseField) IsEmpty() Condition { return f.unary(ast2.OperatorIsEmpty) }

// IsNotEmpty query that field is not empty.
func (f BaseField) IsNotEmpty() Condition { return f.unary(ast2.OperatorIsNotEmpty) }

// IsDynamic query that field is dynamically the sysID provided.
func (f BaseField) IsDynamic(sysID string) Condition {
	return f.binary(ast2.OperatorIsDynamic, sysID)
}

// IsSame query that field is the same as the provided value.
func (f BaseField) IsSame(sysID string) Condition {
	return f.binary(ast2.OperatorIsSame, sysID)
}

// IsDifferent query that field is different from the provided value.
func (f BaseField) IsDifferent(sysID string) Condition {
	return f.binary(ast2.OperatorIsDifferent, sysID)
}

// IsInHierarchy query that reference field is in hierarchy.
func (f BaseField) IsInHierarchy() Condition {
	return f.unary(ast2.OperatorIsInHierarchy)
}
