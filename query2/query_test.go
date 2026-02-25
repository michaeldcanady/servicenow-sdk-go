//go:build preview.query

package query2

import (
	"strings"
	"testing"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast2"
)

// --- Top-level Functions ---

func TestWhere(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		expected string
	}{
		{"Basic", "f", "f=v"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Where(tt.field).Is("v").String() != tt.expected {
				t.Errorf("got %s, expected %s", Where(tt.field).Is("v").String(), tt.expected)
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		expected string
	}{
		{"Basic", "f", "f=v"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).Is("v").String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).Is("v").String(), tt.expected)
			}
		})
	}
}

func TestNumber(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		expected string
	}{
		{"Basic", "f", "f=1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Number(tt.field).Is(1).String() != tt.expected {
				t.Errorf("got %s, expected %s", Number(tt.field).Is(1).String(), tt.expected)
			}
		})
	}
}

func TestBoolean(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		expected string
	}{
		{"Basic", "f", "f=true"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Boolean(tt.field).Is(true).String() != tt.expected {
				t.Errorf("got %s, expected %s", Boolean(tt.field).Is(true).String(), tt.expected)
			}
		})
	}
}

func TestDate(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		expected string
	}{
		{"Basic", "f", "fONjavascript:expr"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Date(tt.field).On(JS("expr")).String() != tt.expected {
				t.Errorf("got %s, expected %s", Date(tt.field).On(JS("expr")).String(), tt.expected)
			}
		})
	}
}

func TestDateTime(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		expected string
	}{
		{"Basic", "f", "fONjavascript:expr"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if DateTime(tt.field).On(JS("expr")).String() != tt.expected {
				t.Errorf("got %s, expected %s", DateTime(tt.field).On(JS("expr")).String(), tt.expected)
			}
		})
	}
}

func TestAnd(t *testing.T) {
	tests := []struct {
		name     string
		conds    []Condition
		expected string
	}{
		{"Empty", nil, ""},
		{"Multiple", []Condition{String("a").Is("1"), String("b").Is("2")}, "a=1^b=2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := And(tt.conds...)
			if tt.name == "Empty" {
				if res != nil {
					t.Error("Empty And should return nil")
				}
			} else {
				if res.String() != tt.expected {
					t.Errorf("got %s, expected %s", res.String(), tt.expected)
				}
			}
		})
	}
}

func TestOr(t *testing.T) {
	tests := []struct {
		name     string
		conds    []Condition
		expected string
	}{
		{"Empty", nil, ""},
		{"Multiple", []Condition{String("a").Is("1"), String("b").Is("2")}, "a=1^ORb=2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Or(tt.conds...)
			if tt.name == "Empty" {
				if res != nil {
					t.Error("Empty Or should return nil")
				}
			} else {
				if res.String() != tt.expected {
					t.Errorf("got %s, expected %s", res.String(), tt.expected)
				}
			}
		})
	}
}

func TestTime(t *testing.T) {
	now := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)
	tests := []struct {
		name     string
		input    time.Time
		expected string
	}{
		{"Standard", now, "2024-01-01 12:00:00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Time(tt.input).String() != tt.expected {
				t.Errorf("got %s, expected %s", Time(tt.input).String(), tt.expected)
			}
		})
	}
}

func TestJS(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Standard", "expr", "javascript:expr"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if JS(tt.input).String() != tt.expected {
				t.Errorf("got %s, expected %s", JS(tt.input).String(), tt.expected)
			}
		})
	}
}

func TestNewCondition(t *testing.T) {
	tests := []struct {
		name     string
		node     ast2.Node
		expected string
	}{
		{"Literal", ast2.NewLiteralNode("v"), "v"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if NewCondition(tt.node).String() != tt.expected {
				t.Errorf("got %s, expected %s", NewCondition(tt.node).String(), tt.expected)
			}
		})
	}
}

func TestNewErrorCondition(t *testing.T) {
	_, myErr := time.Parse("2006", "invalid")
	tests := []struct {
		name string
		err  error
	}{
		{"Stored", myErr},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if NewErrorCondition(tt.err).Error() != tt.err {
				t.Error("NewErrorCondition failed to store error")
			}
		})
	}
}

// --- Condition Methods ---

func TestCondition_And(t *testing.T) {
	tests := []struct {
		name     string
		c1, c2   Condition
		expected string
	}{
		{"Standard", String("a").Is("1"), String("b").Is("2"), "a=1^b=2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.c1.And(tt.c2).String() != tt.expected {
				t.Errorf("got %s, expected %s", tt.c1.And(tt.c2).String(), tt.expected)
			}
		})
	}
}

func TestCondition_Or(t *testing.T) {
	tests := []struct {
		name     string
		c1, c2   Condition
		expected string
	}{
		{"Standard", String("a").Is("1"), String("b").Is("2"), "a=1^ORb=2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.c1.Or(tt.c2).String() != tt.expected {
				t.Errorf("got %s, expected %s", tt.c1.Or(tt.c2).String(), tt.expected)
			}
		})
	}
}

func TestCondition_ToNode(t *testing.T) {
	tests := []struct {
		name string
		c    Condition
	}{
		{"Basic", String("f").Is("v")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.c.ToNode() == nil {
				t.Error("ToNode returned nil")
			}
		})
	}
}

func TestCondition_String(t *testing.T) {
	tests := []struct {
		name     string
		c        Condition
		expected string
	}{
		{"Standard", String("f").Is("v"), "f=v"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.c.String() != tt.expected {
				t.Errorf("got %s, expected %s", tt.c.String(), tt.expected)
			}
		})
	}
}

func TestCondition_Error(t *testing.T) {
	tests := []struct {
		name string
		c    Condition
	}{
		{"InvalidRange", Number("f").Between(10, 5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.c.Error() == nil {
				t.Error("Expected error")
			}
		})
	}
}

// --- BaseField Methods ---

func TestBaseField_IsAnything(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		expected string
	}{
		{"Basic", "f", "fANYTHING"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).IsAnything().String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).IsAnything().String(), tt.expected)
			}
		})
	}
}

func TestBaseField_IsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		expected string
	}{
		{"Basic", "f", "fISEMPTY"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).IsEmpty().String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).IsEmpty().String(), tt.expected)
			}
		})
	}
}

func TestBaseField_IsNotEmpty(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		expected string
	}{
		{"Basic", "f", "fISNOTEMPTY"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).IsNotEmpty().String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).IsNotEmpty().String(), tt.expected)
			}
		})
	}
}

func TestBaseField_IsDynamic(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		sid      string
		expected string
	}{
		{"Standard", "f", "sid", "fDYNAMICsid"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).IsDynamic(tt.sid).String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).IsDynamic(tt.sid).String(), tt.expected)
			}
		})
	}
}

func TestBaseField_IsSame(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		other    string
		expected string
	}{
		{"Standard", "f", "o", "fSAMEASo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).IsSame(tt.other).String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).IsSame(tt.other).String(), tt.expected)
			}
		})
	}
}

func TestBaseField_IsDifferent(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		other    string
		expected string
	}{
		{"Standard", "f", "o", "fNSAMEASo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).IsDifferent(tt.other).String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).IsDifferent(tt.other).String(), tt.expected)
			}
		})
	}
}

func TestBaseField_IsInHierarchy(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		expected string
	}{
		{"Basic", "f", "fIN HIERARCHY"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).IsInHierarchy().String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).IsInHierarchy().String(), tt.expected)
			}
		})
	}
}

// --- StringField Methods ---

func TestStringField_Is(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      string
		expected string
	}{
		{"Standard", "f", "v", "f=v"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).Is(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).Is(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestStringField_IsNot(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      string
		expected string
	}{
		{"Standard", "f", "v", "f!=v"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).IsNot(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).IsNot(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestStringField_StartsWith(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      string
		expected string
	}{
		{"Standard", "f", "v", "fSTARTSWITHv"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).StartsWith(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).StartsWith(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestStringField_EndsWith(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      string
		expected string
	}{
		{"Standard", "f", "v", "fENDSWITHv"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).EndsWith(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).EndsWith(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestStringField_Contains(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      string
		expected string
	}{
		{"Standard", "f", "v", "fLIKEv"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).Contains(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).Contains(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestStringField_DoesNotContain(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      string
		expected string
	}{
		{"Standard", "f", "v", "fNOT LIKEv"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).DoesNotContain(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).DoesNotContain(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestStringField_IsOneOf(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		vals     []string
		expected string
	}{
		{"Multiple", "f", []string{"a", "b"}, "fINa,b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).IsOneOf(tt.vals...).String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).IsOneOf(tt.vals...).String(), tt.expected)
			}
		})
	}
}

func TestStringField_IsNotOneOf(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		vals     []string
		expected string
	}{
		{"Multiple", "f", []string{"a", "b"}, "fNOT INa,b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).IsNotOneOf(tt.vals...).String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).IsNotOneOf(tt.vals...).String(), tt.expected)
			}
		})
	}
}

func TestStringField_IsEmptyString(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		expected string
	}{
		{"Basic", "f", "fEMPTYSTRING"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).IsEmptyString().String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).IsEmptyString().String(), tt.expected)
			}
		})
	}
}

func TestStringField_MatchesPattern(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      string
		expected string
	}{
		{"Standard", "f", "p", "fMATCHES PATTERNp"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).MatchesPattern(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).MatchesPattern(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestStringField_Between(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		l, u     string
		expected string
	}{
		{"Standard", "f", "a", "b", "fBETWEENa@b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).Between(tt.l, tt.u).String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).Between(tt.l, tt.u).String(), tt.expected)
			}
		})
	}
}

// --- NumberField Methods ---

func TestNumberField_Is(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      float64
		expected string
	}{
		{"Standard", "f", 1, "f=1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Number(tt.field).Is(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", Number(tt.field).Is(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestNumberField_IsNot(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      float64
		expected string
	}{
		{"Standard", "f", 1, "f!=1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Number(tt.field).IsNot(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", Number(tt.field).IsNot(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestNumberField_LessThan(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      float64
		expected string
	}{
		{"Standard", "f", 1, "f<1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Number(tt.field).LessThan(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", Number(tt.field).LessThan(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestNumberField_GreaterThan(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      float64
		expected string
	}{
		{"Standard", "f", 1, "f>1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Number(tt.field).GreaterThan(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", Number(tt.field).GreaterThan(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestNumberField_LessThanOrIs(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      float64
		expected string
	}{
		{"Standard", "f", 1, "f<=1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Number(tt.field).LessThanOrIs(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", Number(tt.field).LessThanOrIs(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestNumberField_GreaterThanOrIs(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      float64
		expected string
	}{
		{"Standard", "f", 1, "f>=1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Number(tt.field).GreaterThanOrIs(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", Number(tt.field).GreaterThanOrIs(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestNumberField_Between(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		l, u     float64
		expected string
		isErr    bool
	}{
		{"Valid", "f", 1, 2, "fBETWEEN1@2", false},
		{"Invalid", "f", 2, 1, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Number(tt.field).Between(tt.l, tt.u)
			if tt.isErr {
				if res.Error() == nil {
					t.Error("Expected error")
				}
			} else {
				if res.String() != tt.expected {
					t.Errorf("got %s, expected %s", res.String(), tt.expected)
				}
			}
		})
	}
}

func TestNumberField_IsOneOf(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		vals     []float64
		expected string
	}{
		{"Multiple", "f", []float64{1, 2}, "fIN1,2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Number(tt.field).IsOneOf(tt.vals...).String() != tt.expected {
				t.Errorf("got %s, expected %s", Number(tt.field).IsOneOf(tt.vals...).String(), tt.expected)
			}
		})
	}
}

func TestNumberField_IsNotOneOf(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		vals     []float64
		expected string
	}{
		{"Multiple", "f", []float64{1, 2}, "fNOT IN1,2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Number(tt.field).IsNotOneOf(tt.vals...).String() != tt.expected {
				t.Errorf("got %s, expected %s", Number(tt.field).IsNotOneOf(tt.vals...).String(), tt.expected)
			}
		})
	}
}

func TestNumberField_GreaterThanField(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		other    string
		expected string
	}{
		{"Standard", "f", "o", "fGT_FIELDo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Number(tt.field).GreaterThanField(tt.other).String() != tt.expected {
				t.Errorf("got %s, expected %s", Number(tt.field).GreaterThanField(tt.other).String(), tt.expected)
			}
		})
	}
}

func TestNumberField_LessThanField(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		other    string
		expected string
	}{
		{"Standard", "f", "o", "fLT_FIELDo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Number(tt.field).LessThanField(tt.other).String() != tt.expected {
				t.Errorf("got %s, expected %s", Number(tt.field).LessThanField(tt.other).String(), tt.expected)
			}
		})
	}
}

func TestNumberField_GreaterThanOrIsField(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		other    string
		expected string
	}{
		{"Standard", "f", "o", "fGT_OR_EQUALS_FIELDo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Number(tt.field).GreaterThanOrIsField(tt.other).String() != tt.expected {
				t.Errorf("got %s, expected %s", Number(tt.field).GreaterThanOrIsField(tt.other).String(), tt.expected)
			}
		})
	}
}

func TestNumberField_LessThanOrIsField(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		other    string
		expected string
	}{
		{"Standard", "f", "o", "fLT_OR_EQUALS_FIELDo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Number(tt.field).LessThanOrIsField(tt.other).String() != tt.expected {
				t.Errorf("got %s, expected %s", Number(tt.field).LessThanOrIsField(tt.other).String(), tt.expected)
			}
		})
	}
}

func TestNumberField_IsMoreThan(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      float64
		expected string
	}{
		{"Standard", "f", 1, "fMORETHAN1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Number(tt.field).IsMoreThan(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", Number(tt.field).IsMoreThan(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestNumberField_IsLessThan(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      float64
		expected string
	}{
		{"Standard", "f", 1, "fLESSTHAN1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Number(tt.field).IsLessThan(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", Number(tt.field).IsLessThan(tt.val).String(), tt.expected)
			}
		})
	}
}

// --- BooleanField Methods ---

func TestBooleanField_Is(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      bool
		expected string
	}{
		{"True", "f", true, "f=true"},
		{"False", "f", false, "f=false"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Boolean(tt.field).Is(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", Boolean(tt.field).Is(tt.val).String(), tt.expected)
			}
		})
	}
}

// --- DateTimeField Methods ---

func TestDateTimeField_On(t *testing.T) {
	now := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name     string
		field    string
		val      DateTimeValue
		expected string
	}{
		{"Standard", "f", Time(now), "fON2024-01-01 00:00:00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if DateTime(tt.field).On(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", DateTime(tt.field).On(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestDateTimeField_NotOn(t *testing.T) {
	now := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name     string
		field    string
		val      DateTimeValue
		expected string
	}{
		{"Standard", "f", Time(now), "fNOTON2024-01-01 00:00:00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if DateTime(tt.field).NotOn(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", DateTime(tt.field).NotOn(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestDateTimeField_Before(t *testing.T) {
	now := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name     string
		field    string
		val      DateTimeValue
		expected string
	}{
		{"Standard", "f", Time(now), "f<2024-01-01 00:00:00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if DateTime(tt.field).Before(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", DateTime(tt.field).Before(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestDateTimeField_AtOrBefore(t *testing.T) {
	now := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name     string
		field    string
		val      DateTimeValue
		expected string
	}{
		{"Standard", "f", Time(now), "f<=2024-01-01 00:00:00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if DateTime(tt.field).AtOrBefore(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", DateTime(tt.field).AtOrBefore(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestDateTimeField_After(t *testing.T) {
	now := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name     string
		field    string
		val      DateTimeValue
		expected string
	}{
		{"Standard", "f", Time(now), "f>2024-01-01 00:00:00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if DateTime(tt.field).After(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", DateTime(tt.field).After(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestDateTimeField_AtOrAfter(t *testing.T) {
	now := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name     string
		field    string
		val      DateTimeValue
		expected string
	}{
		{"Standard", "f", Time(now), "f>=2024-01-01 00:00:00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if DateTime(tt.field).AtOrAfter(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", DateTime(tt.field).AtOrAfter(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestDateTimeField_Between(t *testing.T) {
	t1 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name     string
		field    string
		st, en   time.Time
		expected string
		isErr    bool
	}{
		{"Valid", "f", t1, t2, "fBETWEEN2024-01-01 00:00:00@2024-01-02 00:00:00", false},
		{"Invalid", "f", t2, t1, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := DateTime(tt.field).Between(tt.st, tt.en)
			if tt.isErr {
				if res.Error() == nil {
					t.Error("Expected error")
				}
			} else {
				if res.String() != tt.expected {
					t.Errorf("got %s, expected %s", res.String(), tt.expected)
				}
			}
		})
	}
}

func TestDateTimeField_Javascript(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		expr     string
		expected string
	}{
		{"Standard", "f", "expr", "fONjavascript:expr"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if DateTime(tt.field).Javascript(tt.expr).String() != tt.expected {
				t.Errorf("got %s, expected %s", DateTime(tt.field).Javascript(tt.expr).String(), tt.expected)
			}
		})
	}
}

func TestDateTimeField_Today(t *testing.T) {
	expected := "fONToday@javascript:gs.beginningOfToday()@javascript:gs.endOfToday()"
	tests := []struct {
		name     string
		field    string
		expected string
	}{
		{"Basic", "f", expected},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if DateTime(tt.field).Today().String() != tt.expected {
				t.Errorf("got %s, expected %s", DateTime(tt.field).Today().String(), tt.expected)
			}
		})
	}
}

func TestDateTimeField_Yesterday(t *testing.T) {
	expected := "fONYesterday@javascript:gs.beginningOfYesterday()@javascript:gs.endOfYesterday()"
	tests := []struct {
		name     string
		field    string
		expected string
	}{
		{"Basic", "f", expected},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if DateTime(tt.field).Yesterday().String() != tt.expected {
				t.Errorf("got %s, expected %s", DateTime(tt.field).Yesterday().String(), tt.expected)
			}
		})
	}
}

func TestDateTimeField_Tomorrow(t *testing.T) {
	expected := "fONTomorrow@javascript:gs.beginningOfTomorrow()@javascript:gs.endOfTomorrow()"
	tests := []struct {
		name     string
		field    string
		expected string
	}{
		{"Basic", "f", expected},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if DateTime(tt.field).Tomorrow().String() != tt.expected {
				t.Errorf("got %s, expected %s", DateTime(tt.field).Tomorrow().String(), tt.expected)
			}
		})
	}
}

func TestDateTimeField_ThisWeek(t *testing.T) {
	expected := "fONThis week@javascript:gs.beginningOfThisWeek()@javascript:gs.endOfThisWeek()"
	tests := []struct {
		name     string
		field    string
		expected string
	}{
		{"Basic", "f", expected},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if DateTime(tt.field).ThisWeek().String() != tt.expected {
				t.Errorf("got %s, expected %s", DateTime(tt.field).ThisWeek().String(), tt.expected)
			}
		})
	}
}

func TestDateTimeField_LastWeek(t *testing.T) {
	expected := "fONLast week@javascript:gs.beginningOfLastWeek()@javascript:gs.endOfLastWeek()"
	tests := []struct {
		name     string
		field    string
		expected string
	}{
		{"Basic", "f", expected},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if DateTime(tt.field).LastWeek().String() != tt.expected {
				t.Errorf("got %s, expected %s", DateTime(tt.field).LastWeek().String(), tt.expected)
			}
		})
	}
}

func TestDateTimeField_ThisMonth(t *testing.T) {
	expected := "fONThis month@javascript:gs.beginningOfThisMonth()@javascript:gs.endOfThisMonth()"
	tests := []struct {
		name     string
		field    string
		expected string
	}{
		{"Basic", "f", expected},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if DateTime(tt.field).ThisMonth().String() != tt.expected {
				t.Errorf("got %s, expected %s", DateTime(tt.field).ThisMonth().String(), tt.expected)
			}
		})
	}
}

func TestDateTimeField_LastMonth(t *testing.T) {
	expected := "fONLast month@javascript:gs.beginningOfLastMonth()@javascript:gs.endOfLastMonth()"
	tests := []struct {
		name     string
		field    string
		expected string
	}{
		{"Basic", "f", expected},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if DateTime(tt.field).LastMonth().String() != tt.expected {
				t.Errorf("got %s, expected %s", DateTime(tt.field).LastMonth().String(), tt.expected)
			}
		})
	}
}

func TestDateTimeField_ThisYear(t *testing.T) {
	expected := "fONThis year@javascript:gs.beginningOfThisYear()@javascript:gs.endOfThisYear()"
	tests := []struct {
		name     string
		field    string
		expected string
	}{
		{"Basic", "f", expected},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if DateTime(tt.field).ThisYear().String() != tt.expected {
				t.Errorf("got %s, expected %s", DateTime(tt.field).ThisYear().String(), tt.expected)
			}
		})
	}
}

func TestDateTimeField_LastYear(t *testing.T) {
	expected := "fONLast year@javascript:gs.beginningOfLastYear()@javascript:gs.endOfLastYear()"
	tests := []struct {
		name     string
		field    string
		expected string
	}{
		{"Basic", "f", expected},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if DateTime(tt.field).LastYear().String() != tt.expected {
				t.Errorf("got %s, expected %s", DateTime(tt.field).LastYear().String(), tt.expected)
			}
		})
	}
}

func TestDateTimeField_OnSpecialty(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		l, s, e  string
		expected string
	}{
		{"Standard", "f", "L", "S", "E", "fONL@javascript:S@javascript:E"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if DateTime(tt.field).OnSpecialty(tt.l, tt.s, tt.e).String() != tt.expected {
				t.Errorf("got %s, expected %s", DateTime(tt.field).OnSpecialty(tt.l, tt.s, tt.e).String(), tt.expected)
			}
		})
	}
}

func TestDateTimeField_IsMoreThan(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      string
		expected string
	}{
		{"Standard", "f", "1", "fMORETHAN1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if DateTime(tt.field).IsMoreThan(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", DateTime(tt.field).IsMoreThan(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestDateTimeField_IsLessThan(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      string
		expected string
	}{
		{"Standard", "f", "1", "fLESSTHAN1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if DateTime(tt.field).IsLessThan(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", DateTime(tt.field).IsLessThan(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestDateTimeField_TrendOnOrAfter(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      string
		expected string
	}{
		{"Standard", "f", "v", "fDATEPARTv"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if DateTime(tt.field).TrendOnOrAfter(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", DateTime(tt.field).TrendOnOrAfter(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestDateTimeField_TrendOnOrBefore(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      string
		expected string
	}{
		{"Standard", "f", "v", "fDATEPARTv"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if DateTime(tt.field).TrendOnOrBefore(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", DateTime(tt.field).TrendOnOrBefore(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestDateTimeField_TrendAfter(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      string
		expected string
	}{
		{"Standard", "f", "v", "fDATEPARTv"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if DateTime(tt.field).TrendAfter(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", DateTime(tt.field).TrendAfter(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestDateTimeField_TrendBefore(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      string
		expected string
	}{
		{"Standard", "f", "v", "fDATEPARTv"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if DateTime(tt.field).TrendBefore(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", DateTime(tt.field).TrendBefore(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestDateTimeField_TrendOn(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      string
		expected string
	}{
		{"Standard", "f", "v", "fDATEPARTv"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if DateTime(tt.field).TrendOn(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", DateTime(tt.field).TrendOn(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestDateTimeField_RelativeAfter(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      string
		expected string
	}{
		{"Standard", "f", "v", "fDATEPARTv"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if DateTime(tt.field).RelativeAfter(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", DateTime(tt.field).RelativeAfter(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestDateTimeField_RelativeBefore(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      string
		expected string
	}{
		{"Standard", "f", "v", "fDATEPARTv"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if DateTime(tt.field).RelativeBefore(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", DateTime(tt.field).RelativeBefore(tt.val).String(), tt.expected)
			}
		})
	}
}

// --- Internal / Helper Tests ---

func TestDateTimeValue_String(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Standard", "v", "v"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := DateTimeValue{literal: tt.input}
			if v.String() != tt.expected {
				t.Errorf("got %s, expected %s", v.String(), tt.expected)
			}
		})
	}
}

func TestNewDateTimeValue(t *testing.T) {
	now := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name     string
		input    time.Time
		expected string
	}{
		{"Standard", now, "2024-01-01 00:00:00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewDateTimeValue(tt.input)
			if v.String() != tt.expected {
				t.Errorf("got %s, expected %s", v.String(), tt.expected)
			}
		})
	}
}

func TestComplexQueries(t *testing.T) {
	tests := []struct {
		name     string
		query    Condition
		expected string
	}{
		{
			"DeeplyNested",
			Or(
				And(
					Boolean("active").Is(true),
					Or(
						Number("priority").Is(1),
						Number("priority").Is(2),
					),
				),
				And(
					String("category").Is("software"),
					String("short_description").Contains("important"),
				),
			),
			"active=true^priority=1^ORpriority=2^ORcategory=software^short_descriptionLIKEimportant",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.query.String() != tt.expected {
				t.Errorf("got %s, expected %s", tt.query.String(), tt.expected)
			}
		})
	}
}

func TestDeepErrorPropagation(t *testing.T) {
	invalid1 := Number("a").Between(10, 5)
	invalid2 := DateTime("b").Between(time.Now().Add(time.Hour), time.Now())

	tests := []struct {
		name           string
		query          Condition
		expectedErrors []string
	}{
		{
			"MultiError",
			And(
				String("ok").Is("yes"),
				Or(invalid1, invalid2),
			),
			[]string{"is greater or equal to", "is after end time"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.query.Error()
			if err == nil {
				t.Fatal("Expected error")
			}
			errStr := err.Error()
			for _, exp := range tt.expectedErrors {
				if !strings.Contains(errStr, exp) {
					t.Errorf("Expected error to contain %q, but got: %s", exp, errStr)
				}
			}
		})
	}
}

func TestCoverageHacks(t *testing.T) {
	df := DateTime("f")
	tests := []struct {
		name     string
		op       ast2.Operator
		val      any
		expected string
	}{
		{"DefaultBranch", ast2.OperatorIs, 123, "f=123"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if df.dateTimeBinary(tt.op, tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", df.dateTimeBinary(tt.op, tt.val).String(), tt.expected)
			}
		})
	}
}
