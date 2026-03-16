package pkce

import (
	"testing"
)

func TestMethod_String(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Plain returns correct string",
			test: func(t *testing.T) {
				got := MethodPlain.String()
				want := "plain"
				if got != want {
					t.Errorf("MethodPlain.String() = %q, want %q", got, want)
				}
			},
		},
		{
			name: "S256 returns correct string",
			test: func(t *testing.T) {
				got := MethodS256.String()
				want := "S256"
				if got != want {
					t.Errorf("MethodS256.String() = %q, want %q", got, want)
				}
			},
		},
		{
			name: "Unknown returns correct string",
			test: func(t *testing.T) {
				got := MethodUnknown.String()
				want := "unknown"
				if got != want {
					t.Errorf("MethodUnknown.String() = %q, want %q", got, want)
				}
			},
		},
		{
			name: "Unsupported value falls back to Unknown",
			test: func(t *testing.T) {
				got := Method(999).String()
				want := "unknown"
				if got != want {
					t.Errorf("Method(999).String() = %q, want %q", got, want)
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
