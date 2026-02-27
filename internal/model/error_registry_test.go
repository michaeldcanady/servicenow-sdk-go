//go:build preview.query

package model

import (
	"sync"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
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
			switch tt.name {
			case "Singleton":
				inst1 := GetErrorRegistryInstance()
				inst2 := GetErrorRegistryInstance()
				if inst1 != inst2 {
					t.Error("different instances returned")
				}
			case "ThreadSafe":
				var wg sync.WaitGroup
				const count = 100
				instances := make([]utils.Dictionary[string, abstractions.ErrorMappings], count)
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
