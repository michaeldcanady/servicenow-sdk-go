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
