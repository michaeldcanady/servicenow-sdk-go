package testutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadJSON(t *testing.T) {
	type User struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	tests := []struct {
		name     string
		path     string
		expected User
		wantErr  bool
	}{
		{
			name: "Valid JSON",
			path: "testdata/sample.json",
			expected: User{
				Name:  "Test User",
				Email: "test@example.com",
			},
			wantErr: false,
		},
		{
			name:    "Missing file",
			path:    "testdata/missing.json",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var user User
			err := LoadJSON(tt.path, &user)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, user)
			}
		})
	}
}
