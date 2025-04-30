package tableapi

// RecordElement defines an interface for structured table entry elements.
//
// This interface provides methods to retrieve structured values associated with
// a table entry, including display values, stored values, and an optional link.
//
// Implementing types should define logic for extracting and managing
// these values, ensuring compatibility with TableEntry storage.
//
// Example usage:
//
//	var element RecordElement
//	displayValue, err := element.GetDisplayValue()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(displayValue)
type RecordElement interface {
	// GetDisplayValue retrieves the display value associated with the element.
	//
	// The display value represents a human-readable format of the stored value.
	//
	// Example:
	//
	//      displayValue, err := element.GetDisplayValue()
	//      if err != nil {
	//          log.Fatal(err)
	//      }
	//      fmt.Println(displayValue)
	GetDisplayValue() (ElementValue, error)

	// GetValue retrieves the stored value of the element.
	//
	// This method returns the actual underlying value stored in the table entry.
	//
	// Example:
	//
	//      value, err := element.GetValue()
	//      if err != nil {
	//          log.Fatal(err)
	//      }
	//      fmt.Println(value)
	GetValue() (ElementValue, error)

	// GetLink retrieves an optional link associated with the element.
	//
	// Some elements may have associated links pointing to related records or metadata.
	//
	// Example:
	//
	//      link, err := element.GetLink()
	//      if err != nil {
	//          log.Fatal(err)
	//      }
	//      fmt.Println(*link)
	GetLink() (*string, error)
}
