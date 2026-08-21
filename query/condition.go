package query

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/ast"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
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
	if conversion.IsNil(other) {
		return NewErrorCondition(ErrNilCondition)
	}
	return baseCondition{
		node: ast.NewBinaryNode(c.ToNode(), ast.OperatorAnd, other.ToNode()),
		err:  errors.Join(c.Error(), other.Error()),
	}
}

func (c baseCondition) Or(other Condition) Condition {
	if conversion.IsNil(other) {
		return NewErrorCondition(ErrNilCondition)
	}
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
// Callers that check [Condition.Error] — as documented — never see it.
const invalidQueryPlaceholder = "<invalid query>" //nolint:gochecknoglobals

// render walks the condition's tree and returns its encoded-query fragment,
// reporting structural faults (nil node, nil child, unknown operator) that
// make the whole rendering untrustworthy. Construction-time faults are not
// reported here: they degrade locally to invalidQueryPlaceholder terms while
// the rest of the tree still renders, mirroring [Condition.String].
func (c baseCondition) render() (string, error) {
	if c.node == nil {
		return "", ast.ErrNilNode
	}

	visitor := ast.NewStringerVisitor()
	if err := c.node.Accept(visitor); err != nil {
		return "", err
	}

	return visitor.String(), nil
}

// String renders the condition as a ServiceNow encoded-query fragment.
// Construction faults degrade locally — rejected sides render as
// invalidQueryPlaceholder terms while the rest of the tree still renders —
// and structural faults (nil child, unknown operator) degrade globally to the
// placeholder, since a partially rendered query is worse than none. The
// underlying construction fault is available via [Condition.Error].
func (c baseCondition) String() string {
	rendered, rerr := c.render()
	if rerr != nil {
		// A structural fault means the tree has no valid rendering; degrade
		// to the placeholder rather than emit a truncated or corrupted
		// sysparm_query.
		return invalidQueryPlaceholder
	}
	return rendered
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
