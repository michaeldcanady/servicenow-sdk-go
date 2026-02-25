//go:build preview.query

package internal

import (
	"sync"
	"testing"

	abstractions "github.com/microsoft/kiota-abstractions-go"
)

func TestGetErrorRegistryInstance(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Singleton"},
		{"ThreadSafe"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Singleton" {
				inst1 := GetErrorRegistryInstance()
				inst2 := GetErrorRegistryInstance()
				if inst1 != inst2 {
					t.Error("different instances returned")
				}
			} else if tt.name == "ThreadSafe" {
				var wg sync.WaitGroup
				const count = 100
				instances := make([]Dictionary[string, abstractions.ErrorMappings], count)
				for i := 0; i < count; i++ {
					wg.Add(1)
					go func(idx int) {
						defer wg.Done()
						instances[idx] = GetErrorRegistryInstance()
					}(i)
				}
				wg.Wait()
				for i := 1; i < count; i++ {
					if instances[i] != instances[0] {
						t.Errorf("instance %d differs", i)
					}
				}
			}
		})
	}
}
