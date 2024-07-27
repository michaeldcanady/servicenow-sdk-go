package query

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFragment(t *testing.T) {
	value := 42
	frag := newFragment("age", is, &value)

	typedFrag := frag.(*fragmentImpl[int])

	assert.NotNil(t, frag)
	assert.Equal(t, "age", typedFrag.Field)
	assert.Equal(t, is, typedFrag.RelationalOperator)
	assert.Equal(t, &value, typedFrag.Value)
	assert.Nil(t, frag.getNext())
	assert.Equal(t, unset, frag.getLogicalOperator())
}

func TestSetNext(t *testing.T) {
	value1 := 42
	value2 := 30
	frag1 := newFragment("age", is, &value1)
	frag2 := newFragment("salary", is, &value2)

	frag1.setNext(frag2, and)

	assert.Equal(t, frag2, frag1.getNext())
	assert.Equal(t, and, frag1.getLogicalOperator())
}

func TestSetLogicalOperator(t *testing.T) {
	value := 42
	frag := newFragment("age", is, &value)

	frag.setLogicalOperator(or)

	assert.Equal(t, or, frag.getLogicalOperator())
}

func TestString(t *testing.T) {
	value := 42
	frag := newFragment("age", is, &value)

	expected := fmt.Sprintf("age%v%v", is, value)
	assert.Equal(t, expected, frag.String())

	fragNilValue := newFragment[any]("age", is, nil)
	expectedNil := fmt.Sprintf("age%v", is)
	assert.Equal(t, expectedNil, fragNilValue.String())
}
