package query

import "testing"

func TestBuildQuery(t *testing.T) {
	opt1 := func(q *query) {
		q.addFragment(newFragment[string]("field1", is, nil), and)
	}
	opt2 := func(q *query) {
		q.addFragment(newFragment[string]("field2", is, nil), or)
	}

	q := BuildQuery(opt1, opt2)

	if q.head.String() != "field1=" {
		t.Errorf("expected head fragment to be 'field1=', got %s", q.head.String())
	}

	if q.tail.String() != "field2=" {
		t.Errorf("expected tail fragment to be 'field2=', got %s", q.tail.String())
	}
}
func TestQuery(t *testing.T) {
	opt1 := func(q *query) {
		q.addFragment(newFragment[string]("field1", is, nil), and)
	}
	opt2 := func(q *query) {
		q.addFragment(newFragment[string]("field2", is, nil), or)
	}

	queryString := Query(opt1, opt2)

	expected := "field1=^ORfield2="
	if queryString != expected {
		t.Errorf("expected query string to be '%s', got '%s'", expected, queryString)
	}
}

func TestQueryIterate(t *testing.T) {

}

func TestAddFragment(t *testing.T) {
	q := &query{}
	frag1 := newFragment[string]("field1", is, nil)
	frag2 := newFragment[string]("field2", is, nil)

	q.addFragment(frag1, and)
	q.addFragment(frag2, or)

	if q.head.String() != "field1=" {
		t.Errorf("expected head fragment to be 'field1=', got %s", q.head.String())
	}

	if q.tail.String() != "field2=" {
		t.Errorf("expected tail fragment to be 'field2=', got %s", q.tail.String())
	}
}

func TestQueryExtend(t *testing.T) {
	q1 := &query{}
	q2 := &query{}

	frag1 := newFragment[string]("field1", is, nil)
	frag2 := newFragment[string]("field2", is, nil)
	frag3 := newFragment[string]("field3", is, nil)

	q1.addFragment(frag1, and)
	q2.addFragment(frag2, or)
	q2.addFragment(frag3, and)

	q1.extend(q2, and)

	expected := "field1=field2=field3="
	if q1.String() != expected {
		t.Errorf("expected extended query string to be '%s', got '%s'", expected, q1.String())
	}
}

func TestQueryString(t *testing.T) {

}
