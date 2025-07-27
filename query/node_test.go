package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNode(t *testing.T) {
	// First tree (status=active AND category=IT)
	q1 := NewQueryTree()
	q1.AddCondition(Condition{"active", OpEq, "true"})
	q1.AddCondition(Condition{"Created", OpEq, "IT"})
	q1.AddCondition(&LogicalGroup{
		Operator: "^OR",
		Children: []QueryNode{
			Condition{"active", OpEq, "true"},
			Condition{"Created", OpEq, "IT"},
		},
	})

	// Second tree (status=active AND category=HR)
	q2 := NewQueryTree()
	q2.AddCondition(Condition{"status", OpEq, "active"})
	q2.AddCondition(Condition{"category", OpEq, "HR"})

	// Combine into forest
	forest := &QueryForest{}
	forest.AddTree(q1)
	forest.AddTree(q2)

	// Output: status=active^category=IT^NQstatus=active^category=HR
	assert.Equal(t, "active=true^Created=IT^active=true^ORCreated=IT^NQstatus=active^category=HR", forest.Serialize())
}
