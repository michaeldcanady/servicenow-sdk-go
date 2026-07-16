package http

type HTTPHeader int64

const (
	HTTPHeaderUnknown HTTPHeader = iota - 1
	HTTPHeaderContentType
)

var httpHeaderStrings = map[HTTPHeader]string{
	HTTPHeaderUnknown:     "unknown",
	HTTPHeaderContentType: "content-type",
}

func (hH HTTPHeader) String() string {
	value, ok := httpHeaderStrings[hH]
	if !ok {
		return HTTPHeaderUnknown.String()
	}
	return value
}
