package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	query := new()

	assert.NotNil(t, query.list)
}

func TestQuery(t *testing.T) {
	val := "valu1"
	va1 := "valu2"

	assert.Equal(t, "field1=valu1^fiel2=valu2", Query(And(Is("field1", val), Is("fiel2", va1))))
}

func TestQueryAddValue(t *testing.T) {
	query := new()
	query.AddValue("values")
	assert.Equal(t, "values", query.list.GetHead().GetValue())
}

func TestQueryGetHead(t *testing.T) {
	query := new()
	query.AddValue("values")
	assert.Equal(t, "values", query.GetHead())
}

func TestQueryGetTail(t *testing.T) {
	query := new()
	query.AddValue("values")
	query.AddValue("values1")
	assert.Equal(t, "values1", query.GetTail())
}

func TestQueryString(t *testing.T) {
	query := new()
	val := "valu1"
	va1 := "valu2"
	query.AddValue(newCondition("field1", is, &val))
	query.AddValue(and)
	query.AddValue(newCondition("fiel2", is, &va1))

	assert.Equal(t, "field1=valu1^fiel2=valu2", query.String())
}
