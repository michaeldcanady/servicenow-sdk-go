package mocking

import (
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/mock"
)

var ParsableFactory serialization.ParsableFactory = (&parsableFactory{}).ParsableFactory

type parsableFactory struct {
	mock.Mock
}

func (m *parsableFactory) ParsableFactory(parseNode serialization.ParseNode) (serialization.Parsable, error) {
	args := m.Called(parseNode)
	return args.Get(0).(serialization.Parsable), args.Error(1)
}
