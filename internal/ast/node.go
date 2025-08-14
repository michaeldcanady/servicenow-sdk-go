package ast

// Node Represents a tree node
type Node interface {
	// Left The leftmost (starting) position of the node in source text.
	Left() int
	// Right The rightmost (ending) position of the node in source text.
	Right() int
	// Pos The actual position of the node.
	Pos() int
}
