package query

type ValueWrapper interface {
	ToCondition(field string) Condition
}
