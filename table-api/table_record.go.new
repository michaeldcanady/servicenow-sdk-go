package tableapi

// TableRecord defines an interface for managing structured records in a table.
//
// This interface provides methods for retrieving, updating, and checking the existence
// of attributes stored within a table entry.
//
// Implementing types should ensure proper storage and retrieval of RecordElement instances.
//
// Example usage:
//
//	var record TableRecord
//	element := record.Get("status")
//	fmt.Println(element)
type TableRecord interface {
	// Get retrieves a RecordElement associated with the specified key.
	//
	// This method returns the stored element for the given field name.
	//
	// Example:
	//
	//      element := record.Get("status")
	//      fmt.Println(element)
	Get(string) RecordElement
	// SetElement assigns a RecordElement to the specified key.
	//
	// Example:
	//
	//      element := NewRecordElement()
	//      err := record.SetElement("status", element)
	SetElement(string, RecordElement) error
	// SetValue assigns a value to the specified key using a RecordElement wrapper.
	//
	// This method ensures that the stored data conforms to the RecordElement interface.
	//
	// Example:
	//
	//      err := record.SetValue("status", "active")
	SetValue(string, any) error
	// HasAttribute checks whether the specified key exists in the record.
	//
	// Example:
	//
	//      exists := record.HasAttribute("status")
	//      fmt.Println(exists) // Output: true or false
	HasAttribute(string) bool
}
