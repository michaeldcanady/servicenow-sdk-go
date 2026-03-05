package tableapi

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTableRequestBuilder(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Create builder",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := &MockClient{}
			params := map[string]string{"baseurl": "http://test", "table": "t"}
			res := NewTableRequestBuilder(client, params)
			assert.NotNil(t, res)
		})
	}
}

func TestTableRequestBuilder_Get(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(getFakeCollectionJSON())
	}))
	defer server.Close()

	u, _ := url.Parse(server.URL)
	params := map[string]string{"baseurl": "http://" + u.Host, "table": u.Path}
	builder := NewTableRequestBuilder(&MockClient{}, params)

	tests := []struct {
		name string
		err  bool
	}{
		{"Ok", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := builder.Get(nil)
			if tt.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, res)
			}
		})
	}
}

func TestTableRequestBuilder_Post(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"result":{}}`))
	}))
	defer server.Close()

	u, _ := url.Parse(server.URL)
	params := map[string]string{"baseurl": "http://" + u.Host, "table": u.Path}
	builder := NewTableRequestBuilder(&MockClient{}, params)

	tests := []struct {
		name  string
		input interface{}
		err   bool
	}{
		{"Map", map[string]string{"a": "b"}, false},
		{"Entry", TableEntry{"a": "b"}, false},
		{"Bad", 123, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test Post (deprecated but for coverage)
			if m, ok := tt.input.(map[string]string); ok {
				_, _ = builder.Post(m, nil)
			}
			// Test Post2
			if m, ok := tt.input.(map[string]string); ok {
				_, _ = builder.Post2(m, nil)
			}
			// Test Post3
			_, err := builder.Post3(tt.input, nil)
			if tt.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestTableRequestBuilder_Count(t *testing.T) {
	tests := []struct {
		name          string
		expectedCount int
	}{
		{
			name:          "Count 10",
			expectedCount: 10,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("X-Total-Count", "10")
				w.WriteHeader(http.StatusOK)
			}))
			defer server.Close()

			u, _ := url.Parse(server.URL)
			params := map[string]string{"baseurl": "http://" + u.Host, "table": u.Path}
			builder := NewTableRequestBuilder(&MockClient{}, params)

			count, err := builder.Count()
			assert.NoError(t, err)
			assert.Equal(t, test.expectedCount, count)
		})
	}
}

func TestNew2TableRequestBuilder(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Create builder 2",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := New2TableRequestBuilder(nil, nil)
			assert.NotNil(t, res)
		})
	}
}

func TestTableRequestBuilder_Get2(t *testing.T) {
	tests := []struct {
		name string
		err  bool
	}{
		{"Ok", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock client would be better here for unit test
		})
	}
}

func TestTableRequestBuilder_ById2(t *testing.T) {
	tests := []struct {
		name string
		id   string
	}{
		{
			name: "Get by ID",
			id:   "id",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			builder := New2TableRequestBuilder(nil, nil)
			res := builder.ByID2(test.id)
			assert.NotNil(t, res)
		})
	}
}
