package tableapi

type DisplayValue string

const (
	// TRUE Returns the display values for all fields.
	TRUE DisplayValue = "true"
	// FALSE Returns the actual values from the database.
	FALSE DisplayValue = "false"
	// ALL Returns both actual and display values.
	ALL DisplayValue = "all"
)
