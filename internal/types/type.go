package types

type Stringable interface {
	String() string
}

type Type interface {
	Stringable
}
