package internal

import (
	"errors"
	"sync"

	abstractions "github.com/microsoft/kiota-abstractions-go"
)

var lock = &sync.Mutex{}

var errorRegistryInstance *ErrorRegistry

type ErrorRegisterable interface {
	Put(key string, value abstractions.ErrorMappings) error
	Get(key string) (abstractions.ErrorMappings, bool)
}

type ErrorRegistry struct {
	registry map[string]abstractions.ErrorMappings
}

// Create a global thread safe singleton for global values
func getInstance() *ErrorRegistry {
	if errorRegistryInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if errorRegistryInstance == nil {
			errorRegistryInstance = &ErrorRegistry{
				registry: make(map[string]abstractions.ErrorMappings),
			}
		}
	}

	return errorRegistryInstance
}

func (eR *ErrorRegistry) Put(key string, value abstractions.ErrorMappings) error {
	if instance := getInstance(); IsNil(eR) || eR != instance {
		eR = instance
	}

	_, ok := eR.Get(key)
	if ok {
		return errors.New("key already exists")
	}

	eR.registry[key] = value
	return nil
}

func (eR *ErrorRegistry) Get(key string) (abstractions.ErrorMappings, bool) {
	if instance := getInstance(); IsNil(eR) || eR != instance {
		eR = instance
	}

	// returns only the value by default
	item, ok := eR.registry[key]
	return item, ok
}
