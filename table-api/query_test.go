package tableapi

import "testing"

type Test struct {
	Title    string
	Expected string
	Actual   string
}

func TestAddQuery(t *testing.T) {

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
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			if result := test.Actual; result != test.Expected {
				t.Errorf("Expected: %s, Got: %s", test.Expected, test.Actual)
			}
		})
	}
}

func TestAddOrQuery(t *testing.T) {
	tests := []Test{
		{
			Title:    "Test AddOrQuery",
			Expected: "field1=value1^ORfield2!=value2",
			Actual:   NewQuery().AddOrQuery("field1", Is, "value1").AddOrQuery("field2", IsNot, "value2").String(),
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

func TestIsEmpty(t *testing.T) {
	q := NewQuery()
	q.IsEmpty("field3")
	expected := "field3ISEMPTY"
	if result := q.String(); result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}
}

func TestOrderBy(t *testing.T) {
	q := NewQuery()
	q.AddOrderBy("field4")
	expected := ""
	if result := q.String(); result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}
}
