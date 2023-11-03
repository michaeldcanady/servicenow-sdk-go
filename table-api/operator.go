package tableapi

type Operator int

const (
	Null Operator = iota
	Is
	IsNot
	GreaterThan
	GreaterOrEqual
	LessThan
	LessOrEqual
	Contains
	NotContains
	StartsWith
	EndsWith
	Between
	IsSame
	IsDifferent
	IsEmpty
)

func (o Operator) String() string {

	ops := map[Operator]string{
		Is:             "=",
		IsNot:          "!=",
		GreaterThan:    ">",
		GreaterOrEqual: ">=",
		LessThan:       "<",
		LessOrEqual:    "<=",
		Contains:       "CONTAINS",
		NotContains:    "!CONTAINS",
		StartsWith:     "STARTSWITH",
		EndsWith:       "ENDSWITH",
		Between:        "BETWEEN",
		IsSame:         "SAMEAS",
		IsDifferent:    "NSAMEAS",
		IsEmpty:        "ISEMPTY",
	}

	val, isOk := ops[o]

	if !isOk {
		return ""
	}

	return val
}

// IsValidOperator checks if the given operator is a valid Operator value.
func IsValidOperator(op Operator) bool {
	return op >= Is && op <= IsEmpty
}
