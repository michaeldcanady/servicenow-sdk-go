package internal

type Test[T any] struct {
	Title string
	// Setup to make needed modifications for a specific test
	Setup func()
	// Cleanup to undo changes do to reusable items
	Cleanup func()
	//input       interface{}
	Expected T
	//shouldErr   bool
	ExpectedErr error
}
