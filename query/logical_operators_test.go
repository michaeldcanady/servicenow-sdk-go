package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogicalOperatorString(t *testing.T) {
	assert.Equal(t, "", unset.String())
	assert.Equal(t, "^", and.String())
	assert.Equal(t, "^OR", or.String())
	assert.Equal(t, "^NQ", newQuery.String())
}

func TestAndFunction(t *testing.T) {
	q := new()
	v := 42
	opt := And(func(q *query) { q.AddValue(newCondition("age", is, &v)) })

	opt(q)

	assert.Equal(t, newCondition("age", is, &v), q.GetHead())
	assert.Equal(t, and, q.GetTail())
}

func TestOrFunction(t *testing.T) {
	q := new()
	v := 42
	opt := Or(func(q *query) { q.AddValue(newCondition("age", is, &v)) })

	opt(q)

	assert.Equal(t, newCondition("age", is, &v), q.GetHead())
	assert.Equal(t, or, q.GetTail())
}
