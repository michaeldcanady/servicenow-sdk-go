package testing

type InternalTest struct {
	Title     string
	Input     interface{}
	Expected  interface{}
	ShouldErr bool
	Error     error
}
