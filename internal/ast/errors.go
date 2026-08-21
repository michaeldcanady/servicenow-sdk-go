package ast

import "errors"

// Sentinels for structural faults surfaced through traversal. Visitors wrap
// these with node context (position, element index) so the resulting message
// identifies where traversal failed; callers can match with errors.Is.
var (
	// ErrNilNode is returned when Accept is invoked on a nil node.
	ErrNilNode = errors.New("node cannot be nil")

	// ErrNilChild is returned when a composite node references a nil child.
	ErrNilChild = errors.New("child cannot be nil")

	// ErrUnknownOperator is returned when a node carries an operator with no
	// defined encoded-query rendering.
	ErrUnknownOperator = errors.New("operator is unknown")
)
