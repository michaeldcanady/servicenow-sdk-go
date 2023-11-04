package tableapi

import "fmt"

// OrderBy represents an order-by clause.
type OrderBy struct {
	Direction OrderDirection
	Field     string
}

func NewOrderBy() *OrderBy {
	return &OrderBy{}
}

func (oB *OrderBy) String() string {

	if oB.Direction == Unset {
		return ""
	}

	return fmt.Sprintf("%s%s", oB.Direction, oB.Field)
}
