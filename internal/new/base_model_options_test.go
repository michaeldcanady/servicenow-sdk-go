package internal

import "testing"

// TODO: add tests
func TestWithBackingStoreFactory(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "",
			test: func(t *testing.T) {},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
