package query

// Operators for conditions
type Operator string

const (
	OpEq  Operator = "="
	OpNeq Operator = "!="
	OpGt  Operator = ">"
	OpLt  Operator = "<"
	// Add more as needed
)
