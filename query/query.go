package query

import "fmt"

// Query represents a ServiceNow query
type query struct {
	head fragment
	tail fragment
}

// queryOption represents a query option
type queryOption func(*query)

// BuildQuery creates a new Query with given options
func BuildQuery(options ...queryOption) *query {
	q := &query{}
	for _, option := range options {
		option(q)
	}
	return q
}

// Query creates a new query string with given options
func Query(options ...queryOption) string {
	return BuildQuery(options...).String()
}

func (q *query) iterate(predicate func(fragment)) {
	current := q.head
	for ok := true; ok; ok = (current != nil) {
		predicate(current)
		current = current.getNext()
	}
}

// addFragment adds a fragment to the query.
func (q *query) addFragment(fragment fragment, operator logicalOperator) {
	if q.head == nil {
		q.head = fragment
	}

	if q.tail != nil {
		q.tail.setNext(fragment, operator)
	}

	q.tail = fragment
}

func (q *query) extend(sub *query, operator logicalOperator) {
	sub.iterate(func(f fragment) {
		if f == sub.head {
			q.addFragment(f, operator)
		}
		q.addFragment(f, f.getLogicalOperator())
	})
}

func (q *query) String() string {
	queryString := ""

	q.iterate(func(f fragment) {
		queryString += fmt.Sprintf("%v%v", f.String(), f.getLogicalOperator())
	})
	return queryString
}
