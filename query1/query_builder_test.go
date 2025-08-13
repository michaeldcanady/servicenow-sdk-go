package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryBuilder(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				query := NewQuery().NumericField("example").GreaterThan(7).Or().NumericField("example2").GreaterThan(7).Or().StringField("example3").Is("random").Build()

				assert.Equal(t, "example>7^ORexample2>7^ORexample3=random", query)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
