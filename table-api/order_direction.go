package tableapi

// OrderDirection represents the order direction for sorting.
type OrderDirection string

const (
	Unset OrderDirection = ""
	Asc   OrderDirection = "^ORDERBY"
	Desc  OrderDirection = "^ORDERBYDESC"
)
