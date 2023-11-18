package core

type RelationalOperator string

const (
	Null           RelationalOperator = ""
	Is             RelationalOperator = "="
	IsNot          RelationalOperator = "!="
	GreaterThan    RelationalOperator = ">"
	GreaterOrEqual RelationalOperator = ">="
	LessThan       RelationalOperator = "<"
	LessOrEqual    RelationalOperator = "<="
	Contains       RelationalOperator = "CONTAINS"
	NotContains    RelationalOperator = "!CONTAINS"
	StartsWith     RelationalOperator = "STARTSWITH"
	EndsWith       RelationalOperator = "ENDSWITH"
	Between        RelationalOperator = "BETWEEN"
	IsSame         RelationalOperator = "SAMEAS"
	IsDifferent    RelationalOperator = "NSAMEAS"
	IsEmpty        RelationalOperator = "ISEMPTY"
)
