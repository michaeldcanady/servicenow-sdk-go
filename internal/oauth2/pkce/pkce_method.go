package pkce

type Method int

const (
	MethodUnknown Method = iota - 1
	MethodUnset
	MethodPlain
	MethodS256
)

func (m Method) String() string {
	switch m {
	case MethodPlain:
		return "plain"
	case MethodS256:
		return "S256"
	case MethodUnset:
		return "unset"
	default:
		return "unknown"
	}
}
