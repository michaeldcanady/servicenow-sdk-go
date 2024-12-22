package tableapi

// Deprecated: deprecated since v{unreleased}.
//
// View UI view for which to render the data. Determines the fields returned in the response.
type View string

const (
	// Deprecated: deprecated since v{unreleased}. Use `ViewDesktop` instead.
	//
	// DESKTOP Returns fields for the desktop view.
	DESKTOP View = "desktop"
	// Deprecated: deprecated since v{unreleased}. Use `ViewMobile` instead.
	//
	// MOBILE Returns fields for the mobile view.
	MOBILE View = "mobile"
	// Deprecated: deprecated since v{unreleased}. Use `ViewBoth` instead.
	//
	// BOTH Returns fields for both views.
	BOTH View = "both"

	// Deprecated: deprecated since v{unreleased}.
	//
	// ViewDesktop Returns fields for the desktop view.
	ViewDesktop = DESKTOP
	// Deprecated: deprecated since v{unreleased}.
	//
	// ViewMobile Returns fields for the mobile view.
	ViewMobile = MOBILE
	// Deprecated: deprecated since v{unreleased}.
	//
	// ViewBoth Returns fields for both views.
	ViewBoth = BOTH
)
