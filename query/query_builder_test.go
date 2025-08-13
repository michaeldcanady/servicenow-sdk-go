package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryBuilder_AddFilter(t *testing.T) {
	builder := NewQueryBuilder()

	builder.
		AddFilter("example", GreaterThanCondition(7)).
		OrGroup(func(q *QueryBuilder) {
			q.AddFilter("example2", GreaterThanCondition(7))
			q.AddFilter("example3", IsCondition("random"))
		})

	assert.Equal(t, "example>7^example2>7^ORexample3=random", builder.String())
}
