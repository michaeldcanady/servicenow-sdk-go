package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryBuilder_AddFilter(t *testing.T) {
	builder := NewQueryBuilder()

	builder.
		AddFilter("example", GreaterThan(7)).
		OrGroup(func(q *QueryBuilder) {
			q.AddFilter("example2", Equals("random"))
			q.AddFilter("example3", Equals("random1"))
		})

	assert.Equal(t, "example>7^example2=random^ORexample3=random1", builder.Build())
}
