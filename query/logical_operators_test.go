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
	q := &query{}
	v := 42
	opt := And(func(q *query) { q.addFragment(newFragment[int]("age", is, &v), unset) })

	opt(q)

	assert.Equal(t, and, q.head.getLogicalOperator())
}

func TestOrFunction(t *testing.T) {
	q := &query{}
	v := 42
	opt := Or(func(q *query) { q.addFragment(newFragment[int]("age", is, &v), unset) })

	opt(q)

	assert.Equal(t, or, q.head.getLogicalOperator())
}
