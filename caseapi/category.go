package caseapi

import (
	"strings"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
)

// TODO: is an int enum not string
const (
	categoryUnknown  = "unknown"
	categoryQuestion = "question"
	categoryIssue    = "issue"
	categoryFeature  = "feature"
)

// Category
type Category int8

const (
	// CategoryUnknown represents
	CategoryUnknown Category = iota - 1
	CategoryQuestion
	CategoryIssue
	CategoryFeature
)

// ParseCategory resolves the wire representation of a category to a [Category].
// Matching is case-insensitive.
func ParseCategory(s string) (interface{}, error) {
	if category, ok := categoryValues[strings.ToLower(s)]; ok {
		return category, nil
	}
	return CategoryUnknown, unknownEnumValueError("category", s)
}

var categoryStrings = map[Category]string{
	CategoryUnknown:  categoryUnknown,
	CategoryQuestion: categoryQuestion,
	CategoryIssue:    categoryIssue,
	CategoryFeature:  categoryFeature,
}

// categoryValues is the lower-cased inverse of [categoryStrings], used by [ParseCategory].
var categoryValues = invertEnumStrings(categoryStrings, CategoryUnknown)

// String returns the string representation of the Category.
func (e Category) String() string {
	return conversion.EnumString(categoryStrings, e, categoryUnknown)
}
