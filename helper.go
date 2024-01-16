package servicenowsdkgo

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"slices"
	"strings"
)

const (
	defaultSchema      = "https"
	expectedPath       = "api"
	baseServiceNowHost = "service-now"
)

var (
	ErrInvalidURIPath = errors.New("uri path is invalid")
	ErrMissingHost    = errors.New("uri host missing")
)

// validatePath validates the path parameter is "/api".
func validatePath(path string) (string, error) {
	path = strings.Trim(path, "/")

	pathSlice := strings.Split(path, "/")
	if pathSlice[0] == "" {
		pathSlice = remove(pathSlice, 0)
	}
	switch len(pathSlice) {
	case 0:
		pathSlice = append(pathSlice, expectedPath)
	default:
		if pathSlice[0] != expectedPath {
			return "", ErrInvalidURIPath
		}
	}
	return strings.Join(pathSlice, "/"), nil
}

// remove slice function to remove a specified index.
func remove[T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}

// validateHost validates the host string is {instance}.service-now.com or can be corrected to it.
func validateHost(host string) (string, error) {

	if host == "" {
		return "", ErrMissingHost
	}

	if !strings.Contains(host, ":") {
		host = host + ":"
	}

	host, port, err := net.SplitHostPort(host)
	if err != nil {
		return "", err
	}

	fmt.Println(host)

	hostSlice := strings.Split(host, ".")

	if hostSlice[0] == "www" {
		hostSlice = remove(hostSlice, 0)
	}

	switch len(hostSlice) {
	case 0:
		return "", errors.New("")
	case 1:
		hostSlice = append(hostSlice, baseServiceNowHost)
	default:
		if !slices.Contains(hostSlice, baseServiceNowHost) {
			return "", errors.New("")
		}
	}

	if hostSlice[len(hostSlice)-1] != "com" {
		hostSlice = append(hostSlice, "com")
	}

	host = strings.Join(hostSlice, ".")

	if port != "" {
		host = fmt.Sprintf("%s:%s", host, port)
	}

	return host, nil
}

// validateURI validates that the provided URI is as it should be.
func validateURI(uri string) (string, error) {

	parsedURI, err := url.Parse(uri)
	if err != nil {
		return "", err
	}

	if parsedURI.Scheme == "" {
		parsedURI.Scheme = defaultSchema
	}

	if parsedURI.Host == "" && parsedURI.Path != "" {
		parsedURI.Host = parsedURI.Path
		parsedURI.Path = ""
	}

	parsedURI.Host, err = validateHost(parsedURI.Host)
	if err != nil {
		return "", err
	}

	parsedURI.Path, err = validatePath(parsedURI.Path)
	if err != nil {
		return "", err
	}

	return parsedURI.String(), nil
}

func isUrl(str string) bool {
	_, err := url.ParseRequestURI(str)
	return err == nil
}
