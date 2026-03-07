package serialization

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetMutatedValueFromSource(t *testing.T) {
	tests := []struct {
		name    string
		source  func() (string, error)
		setter  ModelSetter[int]
		mutator Mutator[string, int]
		wantErr bool
		err     string
	}{
		{
			name:    "Success",
			source:  func() (string, error) { return "123", nil },
			setter:  func(v int) error { return nil },
			mutator: func(s string) (int, error) { return 123, nil },
			wantErr: false,
		},
		{
			name:    "Source is nil",
			source:  nil,
			setter:  func(v int) error { return nil },
			mutator: func(s string) (int, error) { return 123, nil },
			wantErr: true,
			err:     "source is nil",
		},
		{
			name:    "Mutator is nil",
			source:  func() (string, error) { return "123", nil },
			setter:  func(v int) error { return nil },
			mutator: nil,
			wantErr: true,
			err:     "mutator is nil",
		},
		{
			name:    "Setter is nil",
			source:  func() (string, error) { return "123", nil },
			setter:  nil,
			mutator: func(s string) (int, error) { return 123, nil },
			wantErr: true,
			err:     "setter is nil",
		},
		{
			name:    "Source error",
			source:  func() (string, error) { return "", errors.New("source error") },
			setter:  func(v int) error { return nil },
			mutator: func(s string) (int, error) { return 123, nil },
			wantErr: true,
			err:     "source error",
		},
		{
			name:    "Mutator error",
			source:  func() (string, error) { return "123", nil },
			setter:  func(v int) error { return nil },
			mutator: func(s string) (int, error) { return 0, errors.New("mutator error") },
			wantErr: true,
			err:     "mutator error",
		},
		{
			name:    "Setter error",
			source:  func() (string, error) { return "123", nil },
			setter:  func(v int) error { return errors.New("setter error") },
			mutator: func(s string) (int, error) { return 123, nil },
			wantErr: true,
			err:     "setter error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := SetMutatedValueFromSource(tt.source, tt.setter, tt.mutator)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.err, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestSetValueFromSource(t *testing.T) {
	tests := []struct {
		name    string
		source  func() (string, error)
		setter  ModelSetter[string]
		wantErr bool
	}{
		{
			name:    "Success",
			source:  func() (string, error) { return "test", nil },
			setter:  func(v string) error { return nil },
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := SetValueFromSource(tt.source, tt.setter)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestWriteMutatedValueToSource(t *testing.T) {
	tests := []struct {
		name     string
		writer   func(string) error
		accessor ModelAccessor[int]
		mutator  Mutator[int, string]
		wantErr  bool
		err      string
	}{
		{
			name:     "Success",
			writer:   func(v string) error { return nil },
			accessor: func() (int, error) { return 123, nil },
			mutator:  func(v int) (string, error) { return "123", nil },
			wantErr:  false,
		},
		{
			name:     "Writer is nil",
			writer:   nil,
			accessor: func() (int, error) { return 123, nil },
			mutator:  func(v int) (string, error) { return "123", nil },
			wantErr:  true,
			err:      "writer is nil",
		},
		{
			name:     "Mutator is nil",
			writer:   func(v string) error { return nil },
			accessor: func() (int, error) { return 123, nil },
			mutator:  nil,
			wantErr:  true,
			err:      "mutator is nil",
		},
		{
			name:     "Accessor is nil",
			writer:   func(v string) error { return nil },
			accessor: nil,
			mutator:  func(v int) (string, error) { return "123", nil },
			wantErr:  true,
			err:      "accessor is nil",
		},
		{
			name:     "Accessor error",
			writer:   func(v string) error { return nil },
			accessor: func() (int, error) { return 0, errors.New("accessor error") },
			mutator:  func(v int) (string, error) { return "123", nil },
			wantErr:  true,
			err:      "accessor error",
		},
		{
			name:     "Mutator error",
			writer:   func(v string) error { return nil },
			accessor: func() (int, error) { return 123, nil },
			mutator:  func(v int) (string, error) { return "", errors.New("mutator error") },
			wantErr:  true,
			err:      "mutator error",
		},
		{
			name:     "Writer error",
			writer:   func(v string) error { return errors.New("writer error") },
			accessor: func() (int, error) { return 123, nil },
			mutator:  func(v int) (string, error) { return "123", nil },
			wantErr:  true,
			err:      "writer error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WriteMutatedValueToSource(tt.writer, tt.accessor, tt.mutator)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.err, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestWriteValueToSource(t *testing.T) {
	tests := []struct {
		name     string
		writer   func(string) error
		accessor ModelAccessor[string]
		wantErr  bool
	}{
		{
			name:     "Success",
			writer:   func(v string) error { return nil },
			accessor: func() (string, error) { return "test", nil },
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WriteValueToSource(tt.writer, tt.accessor)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
