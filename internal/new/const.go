package internal

import "regexp"

const (
	firstLinkHeaderKey = "first"
	prevLinkHeaderKey  = "prev"
	nextLinkHeaderKey  = "next"
	lastLinkHeaderKey  = "last"
)

var (
	linkHeaderRegex = regexp.MustCompile(`<([^>]+)>;rel="([^"]+)"`)
)
