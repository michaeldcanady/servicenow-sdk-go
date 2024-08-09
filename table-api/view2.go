package tableapi

type View2 int64

const (
	ViewDesktop2 View2 = iota
	ViewMobile2
	ViewBoth2
)

func (v View2) String() string {
	views := map[View2]string{
		ViewDesktop2: "desktop",
		ViewMobile2:  "mobile",
		ViewBoth2:    "both",
	}

	str, ok := views[v]
	if !ok {
		return "invalid"
	}
	return str
}
