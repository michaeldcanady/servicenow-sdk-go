// Copyright (c) 2026 Michael Canady
// SPDX-License-Identifier: MIT

package query

import (
	"strings"
	"testing"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/ast"
)

// injectionPayload is the encoded-query injection payload from issue #645: a
// value that breaks out of its term and appends clauses that would bypass an
// "active=true" guard once rendered into sysparm_query.
const injectionPayload = "x^active=false^ORsys_id!=0"

// evilStringer renders to a value containing a reserved character, proving the
// validator inspects the rendered literal rather than the raw input.
type evilStringer struct{}

func (e evilStringer) String() string { return "a^b" }

func TestValidateQueryValue(t *testing.T) {
	tests := []struct {
		name    string
		field   string
		op      ast.Operator
		val     any
		wantErr bool
		char    string
	}{
		{"CleanString", "f", ast.OperatorIs, "hello world", false, ""},
		{"CleanNumber", "f", ast.OperatorIs, 1.5, false, ""},
		{"CleanBool", "f", ast.OperatorIs, true, false, ""},
		{"Caret", "f", ast.OperatorIs, "a^b", true, "^"},
		{"LeadingCaret", "f", ast.OperatorIs, "^b", true, "^"},
		{"TrailingCaret", "f", ast.OperatorIs, "a^", true, "^"},
		{"OrInjection", "f", ast.OperatorIs, "a^ORb=c", true, "^"},
		{"MultipleReserved", "f", ast.OperatorIs, "a^b,c@d", true, "^"},
		{"StringerMetacharacters", "f", ast.OperatorIs, evilStringer{}, true, "^"},
		// "," and "@" are literals outside the operator contexts that render
		// them: scalar values keep commas ("Smith, John") and at-signs
		// (email addresses).
		{"CommaLiteralInScalarValue", "f", ast.OperatorIs, "a,b", false, ""},
		{"AtSignLiteralInScalarValue", "f", ast.OperatorIs, "a@b", false, ""},
		// ...but they are structural where the operator renders them.
		{"CommaInListValue", "f", ast.OperatorIsOneOf, "a,b", true, ","},
		{"CommaInNotInListValue", "f", ast.OperatorIsNotOneOf, "a,b", true, ","},
		{"AtSignInPairValue", "f", ast.OperatorBetween, "a@b", true, "@"},
		// Each separator stays a literal in the other's context.
		{"AtSignLiteralInListValue", "f", ast.OperatorIsOneOf, "a@b", false, ""},
		{"CommaLiteralInPairValue", "f", ast.OperatorBetween, "a,b", false, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateQueryValue(tt.field, tt.op, tt.val)
			if !tt.wantErr {
				if err != nil {
					t.Errorf("validateQueryValue(%v) = %v, expected nil", tt.val, err)
				}
				return
			}
			if err == nil {
				t.Fatalf("validateQueryValue(%v) = nil, expected error", tt.val)
			}
			if !strings.Contains(err.Error(), tt.char) {
				t.Errorf("error %q does not mention offending character %q", err.Error(), tt.char)
			}
			if !strings.Contains(err.Error(), tt.field) {
				t.Errorf("error %q does not mention field %q", err.Error(), tt.field)
			}
		})
	}
}

func TestQueryInjectionRejected(t *testing.T) {
	t.Run("DirectCondition", func(t *testing.T) {
		c := Where("name").Is(injectionPayload)
		if c.Error() == nil {
			t.Fatal("Expected error for injected value")
		}
	})
	t.Run("GuardClauseAnd", func(t *testing.T) {
		q := Boolean("active").Is(true).And(Where("name").Is(injectionPayload))
		err := q.Error()
		if err == nil {
			t.Fatal("Expected combined condition to carry the validation error")
		}
		if !strings.Contains(err.Error(), "active=false") {
			t.Errorf("error %q does not identify the rejected payload", err.Error())
		}
	})
	t.Run("GuardClauseOr", func(t *testing.T) {
		q := Boolean("active").Is(true).Or(Where("name").Is(injectionPayload))
		if q.Error() == nil {
			t.Fatal("Expected combined condition to carry the validation error")
		}
	})
	t.Run("ErrorsJoinFromBothSides", func(t *testing.T) {
		q := Where("a").Is("x^y").And(Where("b").Is("c^d"))
		err := q.Error()
		if err == nil {
			t.Fatal("Expected joined validation errors")
		}
		for _, want := range []string{"x^y", "c^d"} {
			if !strings.Contains(err.Error(), want) {
				t.Errorf("joined error %q missing %q", err.Error(), want)
			}
		}
	})
}

func TestQueryInjectionRejectedPerOperator(t *testing.T) {
	tests := []struct {
		name string
		cond Condition
	}{
		{"StringIs", String("f").Is(injectionPayload)},
		{"StringIsNot", String("f").IsNot(injectionPayload)},
		{"StringStartsWith", String("f").StartsWith(injectionPayload)},
		{"StringEndsWith", String("f").EndsWith(injectionPayload)},
		{"StringContains", String("f").Contains(injectionPayload)},
		{"StringDoesNotContain", String("f").DoesNotContain(injectionPayload)},
		{"StringMatchesPattern", String("f").MatchesPattern(injectionPayload)},
		{"StringBetweenLower", String("f").Between(injectionPayload, "z")},
		{"StringBetweenUpper", String("f").Between("a", injectionPayload)},
		{"StringIsOneOf", String("f").IsOneOf("ok", injectionPayload)},
		{"StringIsNotOneOf", String("f").IsNotOneOf("ok", injectionPayload)},
		{"BaseIsDynamic", String("f").IsDynamic(injectionPayload)},
		{"BaseIsSame", String("f").IsSame(injectionPayload)},
		{"BaseIsDifferent", String("f").IsDifferent(injectionPayload)},
		{"NumberGreaterThanField", Number("f").GreaterThanField(injectionPayload)},
		{"NumberLessThanField", Number("f").LessThanField(injectionPayload)},
		{"NumberGreaterThanOrIsField", Number("f").GreaterThanOrIsField(injectionPayload)},
		{"NumberLessThanOrIsField", Number("f").LessThanOrIsField(injectionPayload)},
		{"DateTimeIsMoreThan", DateTime("f").IsMoreThan(injectionPayload)},
		{"DateTimeIsLessThan", DateTime("f").IsLessThan(injectionPayload)},
		{"DateTimeTrendOnOrAfter", DateTime("f").TrendOnOrAfter(injectionPayload)},
		{"DateTimeTrendBefore", DateTime("f").TrendBefore(injectionPayload)},
		{"DateTimeRelativeAfter", DateTime("f").RelativeAfter(injectionPayload)},
		{"DateTimeRelativeBefore", DateTime("f").RelativeBefore(injectionPayload)},
		// Trusted composites validate their caller-supplied fragments at
		// construction; the error must surface through the condition.
		{"DateTimeJavascript", DateTime("f").Javascript(injectionPayload)},
		{"DateTimeOnJSComposite", DateTime("f").On(JS(injectionPayload))},
		{"DateTimeOnSpecialtyLabel", DateTime("f").OnSpecialty(injectionPayload, "gs.a()", "gs.b()")},
		{"DateTimeOnSpecialtyStartExpr", DateTime("f").OnSpecialty("L", injectionPayload, "gs.b()")},
		{"DateTimeOnSpecialtyEndExpr", DateTime("f").OnSpecialty("L", "gs.a()", injectionPayload)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.cond.Error() == nil {
				t.Error("Expected validation error, got nil")
			}
		})
	}
}

func TestQueryInjectionRejectedEachCharacter(t *testing.T) {
	// Each metacharacter is rejected exactly where it would be structural:
	// "^" everywhere, "," in IN lists, "@" in BETWEEN pairs.
	tests := []struct {
		name               string
		value              string
		rejectedByContains bool
		rejectedByIsOneOf  bool
		rejectedByBetween  bool
	}{
		{"Caret", "a^b", true, true, true},
		{"CaretOr", "a^ORb", true, true, true},
		{"Comma", "a,b", false, true, false},
		{"AtSign", "a@b", false, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertRejected := func(opName string, err error, want bool) {
				switch {
				case want && err == nil:
					t.Errorf("%s: expected rejection of %q", opName, tt.value)
				case !want && err != nil:
					t.Errorf("%s: unexpected rejection of %q: %v", opName, tt.value, err)
				}
			}
			assertRejected("Contains", String("f").Contains(tt.value).Error(), tt.rejectedByContains)
			assertRejected("IsOneOf", String("f").IsOneOf("ok", tt.value).Error(), tt.rejectedByIsOneOf)
			assertRejected("Between", String("f").Between(tt.value, "z").Error(), tt.rejectedByBetween)
		})
	}
}

func TestQueryInjectionSafeValuesRenderUnchanged(t *testing.T) {
	tests := []struct {
		name     string
		cond     Condition
		expected string
	}{
		{"Spaces", String("name").Is("hello world"), "name=hello world"},
		{"Unicode", String("name").Is("café ☕"), "name=café ☕"},
		{"EmptyString", String("name").Is(""), "name="},
		{"EqualsInValue", String("name").Is("a=b"), "name=a=b"},
		{"BangInValue", String("name").Is("a!=b"), "name=a!=b"},
		{"PercentAndPunctuation", String("desc").Contains("50% off (today)!"), "descLIKE50% off (today)!"},
		{"CommaInFreeText", String("name").Is("Smith, John"), "name=Smith, John"},
		{"AtSignInEmail", String("email").EndsWith("@example.com"), "emailENDSWITH@example.com"},
		{"Colon", String("t").StartsWith("javascript:"), "tSTARTSWITHjavascript:"},
		{"NegativeFloat", Number("n").Is(-1.5), "n=-1.5"},
		{"ExponentFloat", Number("n").Is(1e6), "n=1e+06"},
		{"BooleanFalse", Boolean("b").Is(false), "b=false"},
		{"StructuralCommasStillEmitted", String("f").IsOneOf("a", "b"), "fINa,b"},
		{"StructuralPairStillEmitted", String("f").Between("a", "b"), "fBETWEENa@b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.cond.Error(); err != nil {
				t.Fatalf("Unexpected error for safe value: %v", err)
			}
			if got := tt.cond.String(); got != tt.expected {
				t.Errorf("got %q, expected %q", got, tt.expected)
			}
		})
	}
}

func TestTrustedDateCompositesUnaffected(t *testing.T) {
	ts := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name     string
		cond     Condition
		expected string
	}{
		{
			"OnSpecialtyToday",
			DateTime("f").Today(),
			"fONToday@javascript:gs.beginningOfToday()@javascript:gs.endOfToday()",
		},
		{
			"JavascriptValue",
			DateTime("f").Javascript("gs.daysAgoStart(0)"),
			"fONjavascript:gs.daysAgoStart(0)",
		},
		{
			"DateTimeValue",
			DateTime("f").On(NewDateTimeValue(ts)),
			"fON2024-01-01 00:00:00",
		},
		{
			"TimeValue",
			DateTime("f").dateTimeBinary(ast.OperatorBefore, ts),
			"f<2024-01-01 00:00:00",
		},
		{
			"BetweenPair",
			DateTime("f").Between(ts, ts.Add(24*time.Hour)),
			"fBETWEEN2024-01-01 00:00:00@2024-01-02 00:00:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.cond.Error(); err != nil {
				t.Fatalf("Trusted composite unexpectedly rejected: %v", err)
			}
			if got := tt.cond.String(); got != tt.expected {
				t.Errorf("got %q, expected %q", got, tt.expected)
			}
		})
	}
}

func TestMultiValueValidation(t *testing.T) {
	t.Run("NoValuesRendersBareOperator", func(t *testing.T) {
		c := String("f").IsOneOf()
		if err := c.Error(); err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if got := c.String(); got != "fIN" {
			t.Errorf("got %q, expected %q", got, "fIN")
		}
	})
	t.Run("LaterValueStillRejected", func(t *testing.T) {
		if String("f").IsOneOf("ok", "bad^x").Error() == nil {
			t.Error("Expected error")
		}
	})
	t.Run("NumbersAccepted", func(t *testing.T) {
		c := Number("f").IsOneOf(1, 2)
		if err := c.Error(); err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if got := c.String(); got != "fIN1,2" {
			t.Errorf("got %q, expected %q", got, "fIN1,2")
		}
	})
}
