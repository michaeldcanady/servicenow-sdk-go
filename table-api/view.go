package tableapi

// View UI view for which to render the data. Determines the fields returned in the response.
type View string

const (
	// Deprecated: deprecated since v{unreleased}. Use `ViewDesktop` instead.
	// DESKTOP Returns fields for the desktop view.
	DESKTOP View = "desktop"
	// Deprecated: deprecated since v{unreleased}. Use `ViewMobile` instead.
	// MOBILE Returns fields for the mobile view.
	MOBILE View = "mobile"
	// Deprecated: deprecated since v{unreleased}. Use `ViewBoth` instead.
	// BOTH Returns fields for both views.
	BOTH View = "both"

	// ViewDesktop Returns fields for the desktop view.
	ViewDesktop = DESKTOP
	// ViewMobile Returns fields for the mobile view.
	ViewMobile = MOBILE
	// ViewBoth Returns fields for both views.
	ViewBoth = BOTH
)
