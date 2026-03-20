package tableapi

// View specifies the UI view for which to render the data.
type View int

const (
	// View2Unknown represents an unknown UI view.
	View2Unknown View = iota - 1
	// View2Desktop renders the data for the desktop UI view.
	View2Desktop
	// View2Mobile renders the data for the mobile UI view.
	View2Mobile
	// View2Both renders the data for both desktop and mobile UI views.
	View2Both
)

// String returns the string representation of the View2.
func (e View) String() string {
	str, ok := map[View]string{
		View2Unknown: "unknown",
		View2Desktop: "desktop",
		View2Mobile:  "mobile",
		View2Both:    "both",
	}[e]
	if !ok {
		return View2Unknown.String()
	}
	return str
}
