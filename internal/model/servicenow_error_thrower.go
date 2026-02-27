package model

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/kiota"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

type ServiceNowErrorThrower struct {
	errorRegistry utils.Dictionary[string, abstractions.ErrorMappings]
	deserializer  kiota.Deserializer
}

func NewServiceNowErrorThrower(errorRegistry utils.Dictionary[string, abstractions.ErrorMappings], deserializer kiota.Deserializer) *ServiceNowErrorThrower {
	return &ServiceNowErrorThrower{
		errorRegistry: errorRegistry,
		deserializer:  deserializer,
	}
}

func (et *ServiceNowErrorThrower) resolveErrorFactory(typeName string, statusCode int64) (serialization.ParsableFactory, error) {
	if utils.IsNil(et) {
		return nil, NewNilPointerError("et is nil")
	}

	var errorCtor serialization.ParsableFactory
	statusAsString := strconv.Itoa(int(statusCode))

	errorMappings, err := et.errorRegistry.Get(typeName)
	if err != nil {
		return nil, err
	}

	if factory, ok := errorMappings[statusAsString]; ok {
		errorCtor = factory
	} else if factory, ok := errorMappings["XXX"]; ok {
		errorCtor = factory
	} else if factory, ok := errorMappings["4XX"]; statusCode >= 400 && statusCode < 500 && ok {
		errorCtor = factory
	} else if factory, ok := errorMappings["5XX"]; statusCode >= 500 && statusCode < 600 && ok {
		errorCtor = factory
	}

	if errorCtor == nil {
		return nil, errors.New("no error factory registered")
	}

	return errorCtor, nil
}

func (et *ServiceNowErrorThrower) Throw(typeName string, statusCode int64, contentType string, content []byte) error {
	if utils.IsNil(et) {
		return NewNilPointerError("et is nil")
	}

	factory, err := et.resolveErrorFactory(typeName, statusCode)
	if err != nil {
		return err
	}

	parsable, err := et.deserializer.Deserialize(contentType, content, factory)
	if err != nil {
		return err
	}

	errValue, ok := parsable.(error)
	if !ok {
		return fmt.Errorf("%T is not error", parsable)
	}

	return errValue
}
