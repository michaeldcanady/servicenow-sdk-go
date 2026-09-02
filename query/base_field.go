// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package query

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/ast"
)

//https://www.servicenow.com/docs/r/platform-user-interface/r_OpAvailableFiltersQueries.html

// BaseField represents the common properties and methods for all ServiceNow fields.
type BaseField struct {
	name string
}

func (f BaseField) unary(op ast.Operator) Condition {
	return NewCondition(ast.NewUnaryNode(op, ast.NewLiteralNode(f.name)))
}

func (f BaseField) buildBinary(op ast.Operator, right ast.Node) Condition {
	return NewCondition(ast.NewBinaryNode(ast.NewLiteralNode(f.name), op, right))
}

func (f BaseField) binary(op ast.Operator, val interface{}) Condition {
	if err := validateQueryValue(f.name, op, val); err != nil {
		return NewErrorCondition(err)
	}
	return f.buildBinary(op, ast.NewLiteralNode(val))
}

func (f BaseField) pair(op ast.Operator, left, right interface{}) Condition {
	if err := validateQueryValue(f.name, op, left); err != nil {
		return NewErrorCondition(err)
	}
	if err := validateQueryValue(f.name, op, right); err != nil {
		return NewErrorCondition(err)
	}
	return f.buildBinary(op, ast.NewPairNode(ast.NewLiteralNode(left), ast.NewLiteralNode(right)))
}

// multi builds a condition over multiple primitive values (for example,
// IN / NOT IN),
// validating each value before building the array node. It is a package-level
// function because Go methods cannot declare type parameters.
func multi[T ast.Primitive](f BaseField, op ast.Operator, values ...T) Condition {
	for _, value := range values {
		if err := validateQueryValue(f.name, op, value); err != nil {
			return NewErrorCondition(err)
		}
	}
	return f.buildBinary(op, convertSliceToArrayNode(values...))
}

// IsAnything query that field is anything.
func (f BaseField) IsAnything() Condition { return f.unary(ast.OperatorIsAnything) }

// IsEmpty query that field is empty.
func (f BaseField) IsEmpty() Condition { return f.unary(ast.OperatorIsEmpty) }

// IsNotEmpty query that field is not empty.
func (f BaseField) IsNotEmpty() Condition { return f.unary(ast.OperatorIsNotEmpty) }

// IsDynamic query that field is dynamically the sysID provided.
func (f BaseField) IsDynamic(sysID string) Condition {
	return f.binary(ast.OperatorIsDynamic, sysID)
}

// IsSame query that field is the same as the provided value.
func (f BaseField) IsSame(sysID string) Condition {
	return f.binary(ast.OperatorIsSame, sysID)
}

// IsDifferent query that field is different from the provided value.
func (f BaseField) IsDifferent(sysID string) Condition {
	return f.binary(ast.OperatorIsDifferent, sysID)
}

// IsInHierarchy query that reference field is in hierarchy.
func (f BaseField) IsInHierarchy() Condition {
	return f.unary(ast.OperatorIsInHierarchy)
}
