package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	Title    string
	Expected string
	Actual   string
}

func TestQueryAddQuery(t *testing.T) {
	tests := []Test{
		{
			Title:    "Test AddQuery",
			Expected: "field1=value1",
			Actual:   NewQuery().AddQuery("field1", Is, "value1").String(),
		},
		{
			Title:    "Test AddEqual",
			Expected: "field2=value2",
			Actual:   NewQuery().AddEqual("field2", "value2").String(),
		},
		{
			Title:    "Test AddNotEqual",
			Expected: "field!=true",
			Actual:   NewQuery().AddNotEqual("field", true).String(),
		},
		{
			Title:    "Test AddGreaterThan",
			Expected: "field>true",
			Actual:   NewQuery().AddGreaterThan("field", true).String(),
		},
		{
			Title:    "Test AddLessThan",
			Expected: "field<true",
			Actual:   NewQuery().AddLessThan("field", true).String(),
		},
		{
			Title:    "Test AddContains",
			Expected: "fieldCONTAINStrue",
			Actual:   NewQuery().AddContains("field", true).String(),
		},
		{
			Title:    "Test AddNotContains",
			Expected: "field!CONTAINStrue",
			Actual:   NewQuery().AddNotContains("field", true).String(),
		},
		{
			Title:    "Test AddStartsWith",
			Expected: "fieldSTARTSWITHtrue",
			Actual:   NewQuery().AddStartsWith("field", true).String(),
		},
		{
			Title:    "Test AddEndsWith",
			Expected: "fieldENDSWITHtrue",
			Actual:   NewQuery().AddEndsWith("field", true).String(),
		},
		{
			Title:    "Test AddBetween",
			Expected: "fieldBETWEEN1@10",
			Actual:   NewQuery().AddBetween("field", 1, 10).String(),
		},
		{
			Title:    "Test AddIsSame",
			Expected: "fieldSAMEASfield1",
			Actual:   NewQuery().AddIsSame("field", "field1").String(),
		},
		{
			Title:    "Test AddIsDifferent",
			Expected: "fieldNSAMEASfield1",
			Actual:   NewQuery().AddIsDifferent("field", "field1").String(),
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			if result := test.Actual; result != test.Expected {
				t.Errorf("Expected: %s, Got: %s", test.Expected, test.Actual)
			}
		})
	}
}

func TestQueryAddOrQuery(t *testing.T) {
	tests := []Test{
		{
			Title:    "Test AddOrQuery",
			Expected: "field1=value1^ORfield2!=value2",
			Actual:   NewQuery().AddOrQuery("field1", Is, "value1").AddOrQuery("field2", IsNot, "value2").String(),
		},
		{
			Title:    "Test AddOrEqual",
			Expected: "field1=value1^ORfield2=value2",
			Actual:   NewQuery().AddOrEqual("field1", "value1").AddOrEqual("field2", "value2").String(),
		},
		{
			Title:    "Test AddOrNotEqual",
			Expected: "field1!=value1^ORfield2!=value2",
			Actual:   NewQuery().AddOrNotEqual("field1", "value1").AddOrNotEqual("field2", "value2").String(),
		},
		{
			Title:    "Test AddOrGreaterThan",
			Expected: "field1>value1^ORfield2>value2",
			Actual:   NewQuery().AddOrGreaterThan("field1", "value1").AddOrGreaterThan("field2", "value2").String(),
		},
		{
			Title:    "Test AddOrLessThan",
			Expected: "field1<value1^ORfield2<value2",
			Actual:   NewQuery().AddOrLessThan("field1", "value1").AddOrLessThan("field2", "value2").String(),
		},
		{
			Title:    "Test AddOrContains",
			Expected: "field1CONTAINSvalue1^ORfield2CONTAINSvalue2",
			Actual:   NewQuery().AddOrContains("field1", "value1").AddOrContains("field2", "value2").String(),
		},
		{
			Title:    "Test AddOrNotContains",
			Expected: "field1!CONTAINSvalue1^ORfield2!CONTAINSvalue2",
			Actual:   NewQuery().AddOrNotContains("field1", "value1").AddOrNotContains("field2", "value2").String(),
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			if result := test.Actual; result != test.Expected {
				t.Errorf("Expected: %s, Got: %s", test.Expected, test.Actual)
			}
		})
	}
}

func TestQueryIsEmpty(t *testing.T) {
	q := NewQuery()
	q.IsEmpty("field3")
	expected := "field3ISEMPTY"
	if result := q.String(); result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}
}

func TestQueryOrderBy(t *testing.T) {
	q := NewQuery()
	q.AddEqual("field1", true)
	q.AddOrderBy("field4")
	expected := "field1=true^ORDERBYfield4"
	if result := q.String(); result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}

	q = NewQuery()
	q.AddEqual("field1", true)
	q.AddOrderByDesc("field4")
	expected = "field1=true^ORDERBYDESCfield4"
	if result := q.String(); result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}
}

func TestQueryString(t *testing.T) {
	query := NewQuery()
	actual := query.String()
	expected := ""
	assert.Equal(t, expected, actual)
}

func TestQueryEncode(t *testing.T) {
	query := NewQuery()

	actual := query.Encoded()
	expected := ""
	assert.Equal(t, expected, actual)
}
