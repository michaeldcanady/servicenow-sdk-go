//go:build preview.query

package query2

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
)

// Condition represents a part of a ServiceNow query.
type Condition interface {
	And(other Condition) Condition
	Or(other Condition) Condition
	ToNode() ast.Node
	String() string
	Error() error
}

type baseCondition struct {
	node ast.Node
	err  error
}

func (c baseCondition) ToNode() ast.Node {
	return c.node
}

func (c baseCondition) Error() error {
	return c.err
}

func (c baseCondition) And(other Condition) Condition {
	return baseCondition{
		node: ast.NewBinaryNode(c.ToNode(), ast.OperatorAnd, other.ToNode()),
		err:  errors.Join(c.Error(), other.Error()),
	}
}

func (c baseCondition) Or(other Condition) Condition {
	return baseCondition{
		node: ast.NewBinaryNode(c.ToNode(), ast.OperatorOr, other.ToNode()),
		err:  errors.Join(c.Error(), other.Error()),
	}
}

func (c baseCondition) String() string {
	visitor := ast.NewStringerVisitor()
	c.node.Accept(visitor)
	return visitor.String()
}

// NewCondition creates a condition from an AST node.
func NewCondition(node ast.Node) Condition {
	return baseCondition{node: node}
}

// NewErrorCondition creates a condition with an error.
func NewErrorCondition(err error) Condition {
	return baseCondition{err: err}
}
