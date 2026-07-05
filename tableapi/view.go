package tableapi

const (
	viewUnknown = "unknown"
	viewDesktop = "desktop"
	viewMobile  = "mobile"
	viewBoth    = "both"
)

// View specifies the UI view for which to render the data.
type View int

const (
	// ViewUnknown represents an unknown UI view.
	ViewUnknown View = iota - 1
	// ViewDesktop renders the data for the desktop UI view.
	ViewDesktop
	// ViewMobile renders the data for the mobile UI view.
	ViewMobile
	// ViewBoth renders the data for both desktop and mobile UI views.
	ViewBoth
)

// String returns the string representation of the View.
func (e View) String() string {
	str, ok := map[View]string{
		ViewUnknown: viewUnknown,
		ViewDesktop: viewDesktop,
		ViewMobile:  viewMobile,
		ViewBoth:    viewBoth,
	}[e]
	if !ok {
		return ViewUnknown.String()
	}
	return str
}
