package internal

import (
	"sync"
	"testing"

	abstractions "github.com/microsoft/kiota-abstractions-go"
)

func TestGetErrorRegistryInstance(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Single Instance",
			test: func(t *testing.T) {
				// Obtain two instances of the ErrorRegistry
				instance1 := GetErrorRegistryInstance()
				instance2 := GetErrorRegistryInstance()

				// Check if both instances point to the same memory address
				if instance1 != instance2 {
					t.Errorf("GetErrorRegistryInstance did not return the same instance")
				}
			},
		},
		{
			name: "Thread-Safety",
			test: func(t *testing.T) {
				// Use a WaitGroup to test concurrent access
				var wg sync.WaitGroup
				instances := make([]Dictionary[string, abstractions.ErrorMappings], 10)

				// Attempt to get the instance concurrently
				for i := 0; i < 10; i++ {
					wg.Add(1)
					go func(index int) {
						defer wg.Done()
						instances[index] = GetErrorRegistryInstance()
					}(i)
				}
				wg.Wait()

				// Verify all instances point to the same memory address
				for i := 1; i < 10; i++ {
					if instances[i] != instances[0] {
						t.Errorf("Singleton instance is not thread-safe: instances differ")
					}
				}
			},
		},
		{
			name: "Empty Initialization",
			test: func(t *testing.T) {
				// Obtain the instance and check for an empty dictionary
				registry := GetErrorRegistryInstance()
				if len(registry.(*ConcurrentDictionary[string, abstractions.ErrorMappings]).mapping) != 0 {
					t.Errorf("Expected empty dictionary on initialization, got %d elements", len(registry.(*ConcurrentDictionary[string, abstractions.ErrorMappings]).mapping))
				}
			},
		},
	}

	// Run each test case
	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
