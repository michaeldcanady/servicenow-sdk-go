package query

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/ast"
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

// invalidQueryPlaceholder is what a rejected condition renders as: an error
// condition carries this sentinel literal instead of a real subtree, and
// combinations render it in place of the rejected side (e.g.,
// "a=1^<invalid query>"). It deliberately renders as an unusable term rather
// than an empty string — an empty sysparm_query asks ServiceNow for every
// record, while this placeholder fails the request server-side instead.
// Callers that check Error() — as documented — never see it.
const invalidQueryPlaceholder = "<invalid query>" //nolint:gochecknoglobals

func (c baseCondition) String() string {
	if c.node == nil {
		return invalidQueryPlaceholder
	}
	visitor := ast.NewStringerVisitor()
	c.node.Accept(visitor)
	return visitor.String()
}

// NewCondition creates a condition from an AST node.
func NewCondition(node ast.Node) Condition {
	return baseCondition{node: node}
}

// NewErrorCondition creates a condition with an error. The condition carries a
// sentinel node rather than no node at all, so that combining it with other
// conditions still produces a safely traversable tree — rendering it yields
// invalidQueryPlaceholder in place of the rejected side.
func NewErrorCondition(err error) Condition {
	return baseCondition{
		node: ast.NewLiteralNode(invalidQueryPlaceholder),
		err:  err,
	}
}
