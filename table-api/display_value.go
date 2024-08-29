package tableapi

// DisplayValue Determines the type of data returned, either the actual values from the database or the display values of the fields. Display values are manipulated based on the actual value in the database and user or system settings and preferences.
// If returning display values, the value that is returned is dependent on the field type.
type DisplayValue string

const (
	// Deprecated: deprecated since v{unreleased}. Use `DisplayValueTrue` instead.
	//TRUE Returns the display values for all fields.
	TRUE DisplayValue = "true"
	// Deprecated: deprecated since v{unreleased}. Use `DisplayValueFalse` instead.
	//FALSE Returns the actual values from the database.
	FALSE DisplayValue = "false"
	// Deprecated: deprecated since v{unreleased}. Use `DisplayValueAll` instead.
	//ALL Returns both actual and display values.
	ALL DisplayValue = "all"

	//DisplayValueTrue Returns the display values for all fields.
	DisplayValueTrue = TRUE
	//DisplayValueFalse Returns the actual values from the database.
	DisplayValueFalse = FALSE
	//DisplayValueAll Returns both actual and display values.
	DisplayValueAll = ALL
)
