package oauth2

import (
	"testing"
)

func TestXXX(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestAuthMethod_String(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "ClientSecretPost returns correct string",
			test: func(t *testing.T) {
				got := AuthMethodClientSecretPost.String()
				want := "client_secret_post"
				if got != want {
					t.Errorf("got %q, want %q", got, want)
				}
			},
		},
		{
			name: "ClientSecretBasic returns correct string",
			test: func(t *testing.T) {
				got := AuthMethodClientSecretBasic.String()
				want := "client_secret_basic"
				if got != want {
					t.Errorf("got %q, want %q", got, want)
				}
			},
		},
		{
			name: "Unknown returns correct string",
			test: func(t *testing.T) {
				got := AuthMethodUnknown.String()
				want := "unknown"
				if got != want {
					t.Errorf("got %q, want %q", got, want)
				}
			},
		},
		{
			name: "Unsupported value falls back to PKCEMethodUnknown",
			test: func(t *testing.T) {
				got := AuthMethod(999).String()
				want := "unknown"
				if got != want {
					t.Errorf("got %q, want %q", got, want)
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
