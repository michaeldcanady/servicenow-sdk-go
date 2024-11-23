package internal

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

var errorRegistryInstance *errorRegistry
var errorRegistryOnce sync.Once

type errorRegistry struct {
	// registry of types mapped to their ErrorMapping Factories
	registry map[string]abstractions.ErrorMappings
	lock     sync.RWMutex
}

// GetErrorRegistry returns the singleton instance of errorRegistry
func GetErrorRegistry() *errorRegistry {
	errorRegistryOnce.Do(func() {
		errorRegistryInstance = &errorRegistry{
			registry: make(map[string]abstractions.ErrorMappings),
		}
	})
	return errorRegistryInstance
}

// Set adds or updates the factory in the registry with the given type name
func (eR *errorRegistry) Set(typeName string, factory abstractions.ErrorMappings) {
	if IsNil(eR) {
		eR = GetErrorRegistry()
	}

	eR.lock.Lock()
	defer eR.lock.Unlock()
	eR.registry[typeName] = factory
}

// Get retrieves the factory from the registry with the given type name
func (eR *errorRegistry) Get(typeName string) (abstractions.ErrorMappings, error) {
	if IsNil(eR) {
		eR = GetErrorRegistry()
	}

	eR.lock.RLock()
	defer eR.lock.RUnlock()
	typeErrorMapping, found := eR.registry[typeName]
	if !found {
		return nil, fmt.Errorf("type (%s) doesn't have a registered abstractions.ErrorMappings", typeName)
	}
	return typeErrorMapping, nil
}

// GetErrorFactory retrieves the factory for the given type name from the singleton instance
func GetErrorFactory(typeName string) (abstractions.ErrorMappings, error) {
	return GetErrorRegistry().Get(typeName)
}

// RegisterErrorFactory registers a factory for the given type name if it doesn't already exist
func RegisterErrorFactory(typeName string, factory abstractions.ErrorMappings) error {
	mapping, _ := GetErrorRegistry().Get(typeName)
	if !IsNil(mapping) {
		return errors.New("object factory already registered")
	}
	GetErrorRegistry().Set(typeName, factory)
	return nil
}

func ThrowErrors(typeName string, responseStatus int64, contentType string, content []byte) error {
	factory, err := GetErrorFactory(typeName)
	if err != nil {
		return err
	}

	statusAsString := strconv.Itoa(int(responseStatus))
	var errorCtor serialization.ParsableFactory = nil
	if len(factory) == 0 {
		var found bool
		errorCtor, found = factory[statusAsString]
		if !found {
			if responseStatus >= 400 && responseStatus < 500 && factory["4XX"] != nil {
				errorCtor = factory["4XX"]
			} else if responseStatus >= 500 && responseStatus < 600 && factory["5XX"] != nil {
				errorCtor = factory["5XX"]
			} else {
				return nil
			}
		}
	}

	if errorCtor == nil {
		return &abstractions.ApiError{
			Message: "The server returned an unexpected status code and no error factory is registered for this code: " + statusAsString,
		}
	}

	rootNode, err := serialization.DefaultParseNodeFactoryInstance.GetRootParseNode(contentType, content)
	if err != nil {
		return err
	}
	if rootNode == nil {
		return &abstractions.ApiError{
			Message: "The server returned an unexpected status code with no response body: " + statusAsString,
		}
	}
	errValue, err := rootNode.GetObjectValue(errorCtor)
	if err != nil {
		return err
	}

	return errValue.(error)
}
