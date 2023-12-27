package tableapi

type DisplayValue string

const (
	TRUE  DisplayValue = "true"
	FALSE DisplayValue = "false"
	ALL   DisplayValue = "all"
)

type DisplayOption int64

func (dO DisplayOption) String() string {
	return []string{
		"",
		"true",
		"false",
		"all",
	}[dO]
}

const (
	DisplayNil DisplayOption = iota
	// DisplayTrue Returns the display values for all fields.
	DisplayTrue
	// DisplayFalse Returns the actual values from the database.
	DisplayFalse
	// DisplayAll Returns both actual and display values.
	DisplayAll
)
