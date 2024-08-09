package types

type String interface {
	Type
}

type stringValue struct {
	str string
}

func NewString(str string) String {
	return &stringValue{str}
}

func (s *stringValue) String() string {
	return s.str
}
