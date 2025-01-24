package models

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

type Error struct {
	Message *string
	Detail  *string
}

func NewErrorFromDiscriminatorValue(parseNode serialization.ParseNode) (serialization.Parsable, error) {
	return &Error{}, nil
}

// Serialize writes the objects properties to the current writer.
func (eV *Error) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(eV) {
		return nil
	}

	return errors.New("Serialize not implemented")
}

// GetFieldDeserializers returns the deserialization information for this object.
func (eV *Error) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		"message": func(pn serialization.ParseNode) error {
			val, err := pn.GetStringValue()
			if err != nil {
				return err
			}
			eV.Message = val
			return nil
		},
		"detail": func(pn serialization.ParseNode) error {
			val, err := pn.GetStringValue()
			if err != nil {
				return err
			}
			eV.Detail = val
			return nil
		},
	}
}

type ServiceNowError struct {
	// Exception is the exception details in the error response.
	MainError *Error
	// Status is the status of the error response.
	Status *string
}

// Serialize writes the objects properties to the current writer.
func (eV *ServiceNowError) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(eV) {
		return nil
	}

	return errors.New("Serialize not implemented")
}

// GetFieldDeserializers returns the deserialization information for this object.
func (eV *ServiceNowError) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		"error": func(pn serialization.ParseNode) error {
			parsable, err := pn.GetObjectValue(NewErrorFromDiscriminatorValue)
			if err != nil {
				return err
			}

			val, ok := parsable.(*Error)
			if !ok {
				return errors.New("parsable is not Error")
			}

			eV.MainError = val
			return nil
		},
		"status": func(pn serialization.ParseNode) error {
			val, err := pn.GetStringValue()
			if err != nil {
				return err
			}
			eV.Status = val
			return nil
		},
	}
}

func (eV *ServiceNowError) Error() string {
	return *eV.MainError.Message
}

func NewServiceNowErrorFromDiscriminatorValue(parseNode serialization.ParseNode) (serialization.Parsable, error) {
	return &ServiceNowError{}, nil
}
