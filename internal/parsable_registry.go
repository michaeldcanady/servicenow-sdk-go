package internal

import (
	"errors"
	"regexp"
	"sync"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/yosida95/uritemplate/v3"
)

var parsableRegistryInstance *parsableRegistry
var parsableRegistryOnce sync.Once

type parsableRegistry struct {
	// registry of types mapped to their Parsable Factories
	registry map[*regexp.Regexp]serialization.ParsableFactory
	lock     sync.RWMutex
}

// GetParsableRegistry returns the singleton instance of parsableRegistry
func GetParsableRegistry() *parsableRegistry {
	parsableRegistryOnce.Do(func() {
		parsableRegistryInstance = &parsableRegistry{
			registry: make(map[*regexp.Regexp]serialization.ParsableFactory),
		}
	})
	return parsableRegistryInstance
}

// Set adds or updates the factory in the registry with the given type name
func (eR *parsableRegistry) Set(templateString string, factory serialization.ParsableFactory) {
	if IsNil(eR) {
		eR = GetParsableRegistry()
	}

	template := uritemplate.MustNew(templateString)
	re := template.Regexp()

	eR.lock.Lock()
	defer eR.lock.Unlock()
	eR.registry[re] = factory
}

// Get retrieves the factory from the registry with the given type name
func (eR *parsableRegistry) Get(uri string) (serialization.ParsableFactory, error) {
	if IsNil(eR) {
		eR = GetParsableRegistry()
	}

	eR.lock.RLock()
	defer eR.lock.RUnlock()
	for re, parsable := range eR.registry {
		if re.MatchString(uri) {
			return parsable, nil
		}
	}
	return nil, errors.New("no object factory registered")
}

// GetParsableFactory retrieves the factory for the given type name from the singleton instance
func GetParsableFactory(uri string) (serialization.ParsableFactory, error) {
	return GetParsableRegistry().Get(uri)
}

// RegisterParsableFactory registers a factory for the given type name if it doesn't already exist
func RegisterParsableFactory(typeName string, factory serialization.ParsableFactory) error {
	mapping, _ := GetParsableRegistry().Get(typeName)
	if !IsNil(mapping) {
		return errors.New("object factory already registered")
	}
	GetParsableRegistry().Set(typeName, factory)
	return nil
}

// MustRegisterParsableFactory registers a factory for the given type name if it doesn't already exist
func MustRegisterParsableFactory(typeName string, factory serialization.ParsableFactory) {
	if err := RegisterParsableFactory(typeName, factory); err != nil {
		panic(err)
	}
}
