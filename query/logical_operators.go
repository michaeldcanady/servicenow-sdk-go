package query

// logicalOperator ...
type logicalOperator int64

func (o logicalOperator) String() string {
	return map[logicalOperator]string{
		unset:    "",
		and:      "^",
		or:       "^OR",
		newQuery: "^NQ",
	}[o]
}

const (
	unset logicalOperator = -1
	// and ...
	and logicalOperator = iota
	or
	newQuery
)

// And returns a queryOption for an And condition
func And(options ...queryOption) queryOption {
	return func(q *query) {
		subQuery := &query{}

		for _, opt := range options {
			opt(subQuery)
		}

		q.extend(subQuery, and)
	}
}

// Or returns a queryOption for an Or condition
func Or(options ...queryOption) queryOption {
	return func(q *query) {
		subQuery := &query{}

		for _, opt := range options {
			opt(subQuery)
		}

		q.extend(subQuery, or)
	}
}
