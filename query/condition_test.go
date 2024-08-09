package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCondition(t *testing.T) {
	val := "value1"
	cond := newCondition("field1", is, &val)

	assert.Equal(t, "field1", cond.field)
	assert.Equal(t, is, cond.operator)
	assert.Equal(t, "value1", *cond.value)
}

func TestConditionString(t *testing.T) {
	val := "value1"
	cond := newCondition("field1", is, &val)

	assert.Equal(t, "field1=value1", cond.String())

	cond = newCondition[string]("field1", is, nil)

	assert.Equal(t, "field1=", cond.String())
}
