package tableapi

type DisplayValue2 int

const (
	DisplayValue2Unknown DisplayValue2 = iota - 1
	DisplayValue2True
	DisplayValue2False
	DisplayValue2All
)

func (e DisplayValue2) String() string {
	str, ok := map[DisplayValue2]string{
		DisplayValue2Unknown: "unknown",
		DisplayValue2True:    "true",
		DisplayValue2False:   "false",
		DisplayValue2All:     "all",
	}[e]
	if !ok {
		return View2Unknown.String()
	}
	return str
}
