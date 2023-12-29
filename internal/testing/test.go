package testing

type InternalTest struct {
	Title     string
	Input     interface{}
	Expected  interface{}
	ShouldErr bool
	Error     error
}

type Test struct {
	Title    string
	Expected string
	Actual   string
}
