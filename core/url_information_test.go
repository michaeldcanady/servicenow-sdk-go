package core

import (
	"testing"
)

func TestNewURLInformation(t *testing.T) {
	ui := NewURLInformation()
	if ui == nil {
		t.Fatal("returned nil")
	}
}

func TestNewUrlInformation(t *testing.T) {
	// Deprecated alias
	ui := NewUrlInformation()
	if ui == nil {
		t.Fatal("returned nil")
	}
}

func TestUrlInformation_validateURLTemplate(t *testing.T) {
	tests := []struct {
		name     string
		template string
		err      error
	}{
		{"Ok", "http://{+baseurl}/t", nil},
		{"Empty", "", ErrEmptyURI},
		{"MissingBase", "http://t", ErrMissingBasePathTemplate},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ui := &UrlInformation{UrlTemplate: tt.template}
			if err := ui.validateURLTemplate(); err != tt.err {
				t.Errorf("got %v, expected %v", err, tt.err)
			}
		})
	}
}

func TestUrlInformation_validatePathParams(t *testing.T) {
	tests := []struct {
		name   string
		params map[string]string
		err    error
	}{
		{"Ok", map[string]string{"baseurl": "v"}, nil},
		{"Nil", nil, ErrNilPathParameters},
		{"MissingBase", map[string]string{"a": "b"}, ErrMissingBasePathParam},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ui := &UrlInformation{PathParameters: tt.params}
			if err := ui.validatePathParams(); err != tt.err {
				t.Errorf("got %v, expected %v", err, tt.err)
			}
		})
	}
}

func TestUrlInformation_validateQueryParams(t *testing.T) {
	tests := []struct {
		name   string
		params map[string]string
		err    error
	}{
		{"Ok", map[string]string{}, nil},
		{"Nil", nil, ErrNilQueryParamters},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ui := &UrlInformation{QueryParameters: tt.params}
			if err := ui.validateQueryParams(); err != tt.err {
				t.Errorf("got %v, expected %v", err, tt.err)
			}
		})
	}
}

func TestUrlInformation_validateParams(t *testing.T) {
	tests := []struct {
		name string
		ui   *UrlInformation
		err  bool
	}{
		{"Valid", &UrlInformation{UrlTemplate: "{+baseurl}", PathParameters: map[string]string{"baseurl": "v"}, QueryParameters: map[string]string{}}, false},
		{"InvalidTemplate", &UrlInformation{UrlTemplate: ""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ui.validateParams(); (err != nil) != tt.err {
				t.Errorf("got %v, expected err %v", err, tt.err)
			}
		})
	}
}

func TestUrlInformation_parseRawURL(t *testing.T) {
	tests := []struct {
		name  string
		input string
		err   error
	}{
		{"Ok", "http://test", nil},
		{"Empty", "", ErrEmptyRawUrl},
		{"NoSchema", "test", ErrMissingSchema},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ui := &UrlInformation{}
			_, err := ui.parseRawURL(tt.input)
			if err != tt.err {
				t.Errorf("got %v, expected %v", err, tt.err)
			}
		})
	}
}

func TestUrlInformation_ToUrl(t *testing.T) {
	ui := NewURLInformation()
	ui.UrlTemplate = "http://{+baseurl}/t"
	ui.PathParameters["baseurl"] = "test.com"
	
	uiRaw := NewURLInformation()
	uiRaw.PathParameters[rawURLKey] = "http://raw"

	tests := []struct {
		name     string
		ui       *UrlInformation
		expected string
		err      bool
	}{
		{"FromTemplate", ui, "http://test.com/t", false},
		{"FromRaw", uiRaw, "http://raw", false},
		{"Invalid", &UrlInformation{}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := tt.ui.ToUrl()
			if (err != nil) != tt.err {
				t.Errorf("got err %v, expected err %v", err, tt.err)
			}
			if !tt.err && res.String() != tt.expected {
				t.Errorf("got %s, expected %s", res.String(), tt.expected)
			}
		})
	}
}

func TestUrlInformation_AddQueryParameters(t *testing.T) {
	type Params struct {
		A string `url:"a"`
	}
	ui := NewURLInformation()
	var nilUI *UrlInformation

	tests := []struct {
		name   string
		ui     *UrlInformation
		source any
		err    bool
	}{
		{"Ok", ui, Params{A: "v"}, false},
		{"NilSource", ui, nil, true},
		{"NilUI", nilUI, Params{A: "v"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.ui.AddQueryParameters(tt.source)
			if (err != nil) != tt.err {
				t.Errorf("got err %v, expected err %v", err, tt.err)
			}
		})
	}
}

func TestUrlInformation_getURIFromRaw(t *testing.T) {
	ui := NewURLInformation()
	ui.PathParameters[rawURLKey] = "http://test"
	res, _ := ui.getURIFromRaw()
	if res.String() != "http://test" {
		t.Error("failed")
	}
	
	uiEmpty := NewURLInformation()
	res2, _ := uiEmpty.getURIFromRaw()
	if res2 != nil {
		t.Error("expected nil")
	}
}

func TestUrlInformation_buildURIFromTemplate(t *testing.T) {
	ui := NewURLInformation()
	ui.UrlTemplate = "http://{+baseurl}/t"
	ui.PathParameters["baseurl"] = "test.com"
	res, _ := ui.buildURIFromTemplate()
	if res != "http://test.com/t" {
		t.Errorf("got %s", res)
	}
	
	uiBad := NewURLInformation()
	uiBad.UrlTemplate = "{" // Invalid template
	_, err := uiBad.buildURIFromTemplate()
	if err == nil {
		t.Error("expected error")
	}
}
