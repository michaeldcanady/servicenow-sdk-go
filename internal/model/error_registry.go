package model

import (
	"sync"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

var (
	errorRegistryInstance utils.Dictionary[string, abstractions.ErrorMappings]
	once                  sync.Once
)

// GetErrorRegistryInstance returns the singleton instance of the error registry.
func GetErrorRegistryInstance() utils.Dictionary[string, abstractions.ErrorMappings] {
	once.Do(func() {
		errorRegistryInstance = utils.NewConcurrentDictionary[string, abstractions.ErrorMappings]()
	})
	return errorRegistryInstance
}
