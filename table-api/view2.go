package tableapi

// View2 UI view for which to render the data. Determines the fields returned in the response.
type View2 int64

const (
	// View2Unknown represents unknown UI view
	View2Unknown View2 = iota - 1
	// View2Desktop represents desktop UI view
	View2Desktop
	// View2Mobile represents mobile UI view
	View2Mobile
	// View2Both represents both UI view
	View2Both
)

// String returns string representation
func (v View2) String() string {
	view, ok := map[View2]string{
		View2Unknown: "unknown",
		View2Desktop: "desktop",
		View2Mobile:  "mobile",
		View2Both:    "both",
	}[v]
	if !ok {
		return View2Unknown.String()
	}
	return view
}
