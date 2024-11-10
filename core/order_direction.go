package core

// OrderDirection represents the order direction for sorting.
type OrderDirection string

const (
	// Unset ...
	Unset OrderDirection = ""
	// Asc ...
	Asc OrderDirection = "^ORDERBY"
	// Desc ...
	Desc OrderDirection = "^ORDERBYDESC"
)
