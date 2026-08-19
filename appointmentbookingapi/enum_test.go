package appointmentbookingapi

import (
	"strings"
	"testing"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// enumCase describes one enum's parse/render pair so every enum in the package is
// held to the same contract: canonical rendering, case-insensitive parsing, and an
// unknown sentinel plus error for anything unrecognized.
type enumCase struct {
	name      string
	parse     func(string) (interface{}, error)
	unknown   interface{}
	canonical map[string]interface{} // canonical wire string -> expected value
	render    func(interface{}) string
}

func enumCases() []enumCase {
	return []enumCase{
		{
			name:    "View",
			parse:   ParseView,
			unknown: ViewUnknown,
			canonical: map[string]interface{}{
				"platform": ViewPlatform,
				"portal":   ViewPortal,
			},
			render: func(v interface{}) string { return v.(View).String() },
		},
		{
			name:    "ViewScale",
			parse:   ParseViewScale,
			unknown: ViewScaleUnknown,
			canonical: map[string]interface{}{
				"day":  ViewScaleDay,
				"week": ViewScaleWeek,
			},
			render: func(v interface{}) string { return v.(ViewScale).String() },
		},
		{
			name:    "TimeFormat",
			parse:   ParseTimeFormat,
			unknown: TimeFormatUnknown,
			canonical: map[string]interface{}{
				"12hr": TimeFormat12Hr,
				"24hr": TimeFormat24Hr,
			},
			render: func(v interface{}) string { return v.(TimeFormat).String() },
		},
		{
			name:    "DefaultTimeZone",
			parse:   ParseDefaultTimeZone,
			unknown: DefaultTimeZoneUnknown,
			canonical: map[string]interface{}{
				"user":     DefaultTimeZoneUser,
				"location": DefaultTimeZoneLocation,
			},
			render: func(v interface{}) string { return v.(DefaultTimeZone).String() },
		},
		{
			name:    "ShortMonth",
			parse:   ParseShortMonth,
			unknown: ShortMonthUnknown,
			canonical: map[string]interface{}{
				"Jan": ShortMonthJan,
				"Feb": ShortMonthFeb,
				"Mar": ShortMonthMar,
				"Apr": ShortMonthApr,
				"May": ShortMonthMay,
				"Jun": ShortMonthJun,
				"Jul": ShortMonthJul,
				"Aug": ShortMonthAug,
				"Sep": ShortMonthSep,
				"Oct": ShortMonthOct,
				"Nov": ShortMonthNov,
				"Dec": ShortMonthDec,
			},
			render: func(v interface{}) string { return v.(ShortMonth).String() },
		},
		{
			name:    "ShortWeekday",
			parse:   ParseShortWeekday,
			unknown: ShortWeekdayUnknown,
			canonical: map[string]interface{}{
				"Mon": ShortWeekdayMon,
				"Tue": ShortWeekdayTue,
				"Wed": ShortWeekdayWed,
				"Thu": ShortWeekdayThu,
				"Fri": ShortWeekdayFri,
				"Sat": ShortWeekdaySat,
				"Sun": ShortWeekdaySun,
			},
			render: func(v interface{}) string { return v.(ShortWeekday).String() },
		},
	}
}

// TestEnums_ParseCanonical is the regression guard for parsers that lower-cased their
// input while comparing against title-case constants, so no value could ever match.
func TestEnums_ParseCanonical(t *testing.T) {
	for _, enum := range enumCases() {
		t.Run(enum.name, func(t *testing.T) {
			for str, want := range enum.canonical {
				got, err := enum.parse(str)

				require.NoError(t, err, "parsing %q", str)
				assert.Equal(t, want, got, "parsing %q", str)
			}
		})
	}
}

func TestEnums_ParseIsCaseInsensitive(t *testing.T) {
	for _, enum := range enumCases() {
		t.Run(enum.name, func(t *testing.T) {
			for str, want := range enum.canonical {
				for _, variant := range []string{strings.ToUpper(str), strings.ToLower(str)} {
					got, err := enum.parse(variant)

					require.NoError(t, err, "parsing %q", variant)
					assert.Equal(t, want, got, "parsing %q", variant)
				}
			}
		})
	}
}

func TestEnums_ParseUnknown(t *testing.T) {
	for _, enum := range enumCases() {
		t.Run(enum.name, func(t *testing.T) {
			got, err := enum.parse("not-a-real-value")

			require.ErrorIs(t, err, snerrors.ErrUnknownEnumValue)
			assert.Equal(t, enum.unknown, got)
			assert.Contains(t, err.Error(), "not-a-real-value", "the offending value should be echoed")
		})
	}
}

func TestUnknownEnumValueError(t *testing.T) {
	err := unknownEnumValueError("short month", "Smarch")

	require.ErrorIs(t, err, snerrors.ErrUnknownEnumValue)
	assert.EqualError(t, err, `unknown enum value: "Smarch" is not a valid short month`)
}

// TestEnums_StringRoundTrips pins each value's rendered form to the canonical wire
// string it parses from, so String() stays usable as the serialized representation.
func TestEnums_StringRoundTrips(t *testing.T) {
	for _, enum := range enumCases() {
		t.Run(enum.name, func(t *testing.T) {
			for str, value := range enum.canonical {
				assert.Equal(t, str, enum.render(value))

				reparsed, err := enum.parse(enum.render(value))
				require.NoError(t, err)
				assert.Equal(t, value, reparsed)
			}
		})
	}
}

func TestEnums_UnknownRendersUnknown(t *testing.T) {
	assert.Equal(t, "unknown", ViewUnknown.String())
	assert.Equal(t, "unknown", ViewScaleUnknown.String())
	assert.Equal(t, "unknown", TimeFormatUnknown.String())
	assert.Equal(t, "unknown", DefaultTimeZoneUnknown.String())
	assert.Equal(t, "unknown", ShortMonthUnknown.String())
	assert.Equal(t, "unknown", ShortWeekdayUnknown.String())
}

func TestInvertEnumStrings(t *testing.T) {
	strs := map[ShortMonth]string{
		ShortMonthUnknown: shortMonthUnknown,
		ShortMonthJan:     shortMonthJan,
		ShortMonthMay:     shortMonthMay,
	}

	values := invertEnumStrings(strs, ShortMonthUnknown)

	assert.Equal(t, map[string]ShortMonth{"jan": ShortMonthJan, "may": ShortMonthMay}, values)
	assert.NotContains(t, values, shortMonthUnknown, "the unknown sentinel must not be parseable")
}
