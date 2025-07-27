package query

import (
	"strings"
)

// Operators for conditions
type Operator string

const (
	OpEq  Operator = "="
	OpNeq Operator = "!="
	OpGt  Operator = ">"
	OpLt  Operator = "<"
	// Add more as needed
)

type LogicalGroup struct {
	Operator string // "AND" or "OR"
	Children []QueryNode
}

func (lG *LogicalGroup) Serialize() string {
	builder := strings.Builder{}
	childrenCount := len(lG.Children)
	for index, child := range lG.Children {
		builder.WriteString(child.Serialize())
		if index != childrenCount-1 {
			builder.WriteString(lG.Operator)
		}
	}

	return builder.String()
}

// Logical wrapper for conditions
type QueryNode interface {
	Serialize() string
}

// A single AND-based query tree
type QueryTree struct {
	root *LogicalGroup
}

func NewQueryTree() *QueryTree {
	return &QueryTree{
		root: &LogicalGroup{
			Operator: "^",
			Children: make([]QueryNode, 0),
		},
	}
}

func (qt *QueryTree) AddCondition(c QueryNode) {
	qt.root.Children = append(qt.root.Children, c)
}

func (qt *QueryTree) Serialize() string {
	return qt.root.Serialize()
}

// Multiple QueryTrees ORed together
type QueryForest struct {
	Trees []*QueryTree
}

func NewQueryForest() *QueryForest {
	return &QueryForest{
		Trees: make([]*QueryTree, 0),
	}
}

func (qf *QueryForest) AddTree(tree *QueryTree) {
	qf.Trees = append(qf.Trees, tree)
}

func (qf *QueryForest) Serialize() string {
	var parts []string
	for _, tree := range qf.Trees {
		parts = append(parts, tree.Serialize())
	}
	return strings.Join(parts, "^NQ") // OR operator in ServiceNow
}
