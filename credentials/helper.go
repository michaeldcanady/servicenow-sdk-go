package credentials

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func decodeAccessToken(response *http.Response) (*AccessToken, error) {
	defer response.Body.Close()
	var accessToken AccessToken
	if err := json.NewDecoder(response.Body).Decode(&accessToken); err != nil {
		return nil, err
	}

	accessToken.ExpiresAt = time.Now().Add(time.Duration(accessToken.ExpiresIn) * time.Second)
	return &accessToken, nil
}

// parsePort parses string port into int
func parsePort(port string) int {
	if port == "" {
		return 0
	}
	parsedPort, err := strconv.Atoi(port)
	if err != nil {
		fmt.Printf("Error parsing port: %v\n", err)
		return 0
	}
	return parsedPort
}

// makeRange makes a range from min to max
func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

// parseLocalURL parses the provided URI into a url.URL
func parseLocalURL(uri string) (*url.URL, error) {
	parsedURL, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	if parsedURL.Hostname() != "localhost" && parsedURL.Hostname() != "127.0.0.1" {
		return nil, ErrNotLocalHost
	}

	if parsedURL.Port() == "" {
		return nil, ErrMissingPort
	}
	return parsedURL, nil
}

// startLocalServer starts a local server based on the given target URL and configuration
func startLocalServer(targetURL string, config *serverConfig) (*HTTPServer, error) {
	var ports = makeRange(config.PortRangeStart, config.PortRangeEnd)

	if targetURL != "" {
		parsedURL, err := parseLocalURL(targetURL)
		if err != nil {
			return nil, err
		}
		port := parsedURL.Port()
		config.Host = parsedURL.Host
		// fmt.Printf("Attempting to start server on specified port: %s\n", port)
		ports = []int{parsePort(port)}
	}

	for _, port := range ports {
		addr := fmt.Sprintf("%s:%d", config.Host, port)
		server := NewHTTPServer(addr)

		err := server.Start2()
		if err == nil {
			return server, nil
		}
	}
	return nil, errors.New("Unable to start server on any of the specified ports.")
}
