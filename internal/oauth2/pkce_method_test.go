package oauth2

import (
	"testing"
)

func TestPKCEMethod_String(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Plain returns correct string",
			test: func(t *testing.T) {
				got := PKCEMethodPlain.String()
				want := "plain"
				if got != want {
					t.Errorf("PKCEMethodPlain.String() = %q, want %q", got, want)
				}
			},
		},
		{
			name: "S256 returns correct string",
			test: func(t *testing.T) {
				got := PKCEMethodS256.String()
				want := "S256"
				if got != want {
					t.Errorf("PKCEMethodS256.String() = %q, want %q", got, want)
				}
			},
		},
		{
			name: "Unknown returns correct string",
			test: func(t *testing.T) {
				got := PKCEMethodUnknown.String()
				want := "unknown"
				if got != want {
					t.Errorf("PKCEMethodUnknown.String() = %q, want %q", got, want)
				}
			},
		},
		{
			name: "Unsupported value falls back to Unknown",
			test: func(t *testing.T) {
				got := PKCEMethod(999).String()
				want := "unknown"
				if got != want {
					t.Errorf("PKCEMethod(999).String() = %q, want %q", got, want)
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
