package internal

// Test is a generic type that represents a test case with a title, input, expected output, and error.
// It also has a prepare and a cleanup function that can be used to set up and tear down the test environment.
type Test[T any] struct {
	// Title is the title of the test.
	Title string
	// Prepare is a function that prepares the assets for the test.
	Prepare func()
	// Input is the input value for the test.
	Input interface{}
	// Cleanup is a function that cleans up the assets after the test.
	Cleanup func()
	// Expected is the expected output value for the test.
	Expected T
	// Error is the expected error value for the test.
	Error error
}
