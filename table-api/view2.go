package tableapi

type View2 int

const (
	View2Unknown View2 = iota - 1
	View2Desktop
	View2Mobile
	View2Both
)

func (e View2) String() string {
	str, ok := map[View2]string{
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
