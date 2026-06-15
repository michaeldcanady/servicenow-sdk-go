//go:build preview.query

package query2

import (
	"testing"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast2"
)

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
