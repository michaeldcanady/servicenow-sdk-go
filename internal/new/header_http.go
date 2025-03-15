package internal

type HTTPHeader int64

const (
	HTTPHeaderUnknown HTTPHeader = iota - 1
	HTTPHeaderContentType
)

func (hH HTTPHeader) String() string {
	value, ok := map[HTTPHeader]string{
		HTTPHeaderUnknown:     "unknown",
		HTTPHeaderContentType: "content-type",
	}[hH]
	if !ok {
		return HTTPHeaderUnknown.String()
	}
	return value
}
