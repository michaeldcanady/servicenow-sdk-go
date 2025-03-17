package internal

import (
	"sync"

	abstractions "github.com/microsoft/kiota-abstractions-go"
)

var (
	errorRegistryInstance Dictionary[string, abstractions.ErrorMappings]
	once                  sync.Once
)

// GetErrorRegistryInstance returns the singleton instance of the error registry.
func GetErrorRegistryInstance() Dictionary[string, abstractions.ErrorMappings] {
	once.Do(func() {
		errorRegistryInstance = NewConcurrentDictionary[string, abstractions.ErrorMappings]()
	})
	return errorRegistryInstance
}
