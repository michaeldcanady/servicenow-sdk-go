package internal

import "net/url"

type URLInformation interface {
	Query() map[string]string
	Path() map[string]string
	Template() string
	URL() (*url.URL, error)
}
