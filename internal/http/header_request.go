package internal

// RequestHeader A request header is an HTTP header that can be used in an HTTP request to provide information about
// the request context, so that the server can tailor the response.
type RequestHeader int64

const (
	// RequestHeaderUnknown the provided value is an unknown request header.
	RequestHeaderUnknown RequestHeader = iota - 1
	RequestHeaderAuthorization
	RequestHeaderAccept
)

// String returns string representation of the request header
func (rH RequestHeader) String() string {
	stringVal, ok := map[RequestHeader]string{
		RequestHeaderUnknown:       "unknown",
		RequestHeaderAuthorization: "authorization",
		RequestHeaderAccept:        "accept",
	}[rH]
	if !ok {
		return RequestHeaderUnknown.String()
	}
	return stringVal
}
