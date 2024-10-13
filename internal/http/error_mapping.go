package http

import (
	"regexp"
	"strconv"
	"strings"
)

type ErrorMapping interface {
	Set(code string, parser Parsable)
	Len() int
	Get(code int) Parsable
}

type errorMapping struct {
	completeCodes map[string]Parsable
	patterns      map[*regexp.Regexp]Parsable
}

func NewErrorMapping() ErrorMapping {
	return &errorMapping{
		completeCodes: make(map[int]Parsable),
		patterns:      make(map[*regexp.Regexp]Parsable),
	}
}

func (eM *errorMapping) Set(code string, parser Parsable) {
	if !strings.Contains(code, "X") {
		eM.completeCodes[code] = parser
		return
	}
	exp := strings.Replace(code, "X", "[0-9]", -1)
	regExp := regexp.MustCompile(exp)
	eM.patterns[regExp] = parser
}

func (eM *errorMapping) Len() int {
	return len(eM.completeCodes) + len(eM.patterns)
}

func (eM *errorMapping) Get(code int) Parsable {
	stringCode := strconv.Itoa(code)

	if parser, exists := eM.completeCodes[stringCode]; exists {
		return parser
	}

	for exp, parsable := range eM.patterns {
		if exp.Match([]byte(stringCode)) {
			return parsable
		}
	}
	return nil
}
