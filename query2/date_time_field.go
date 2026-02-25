//go:build preview.query

package query2

import (
	"fmt"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast2"
)

// DateTimeField represents a date-time field in ServiceNow.
type DateTimeField struct {
	BaseField
}

func (f DateTimeField) dateTimeBinary(op ast2.Operator, val any) Condition {
	var literal string
	switch v := val.(type) {
	case DateTimeValue:
		literal = v.String()
	case time.Time:
		literal = v.Format("2006-01-02 15:04:05")
	case string:
		literal = v
	default:
		literal = fmt.Sprintf("%v", v)
	}
	return f.binary(op, literal)
}

// On query the date-time field is on a specific date-time.
func (f DateTimeField) On(val DateTimeValue) Condition {
	return f.dateTimeBinary(ast2.OperatorOn, val)
}

// NotOn query the date-time field is not on a specific date-time.
func (f DateTimeField) NotOn(val DateTimeValue) Condition {
	return f.dateTimeBinary(ast2.OperatorNotOn, val)
}

// Before query the date-time field is before a specific date-time.
func (f DateTimeField) Before(val DateTimeValue) Condition {
	return f.dateTimeBinary(ast2.OperatorBefore, val)
}

// AtOrBefore query the date-time field is at or before a specific date-time.
func (f DateTimeField) AtOrBefore(val DateTimeValue) Condition {
	return f.dateTimeBinary(ast2.OperatorAtOrBefore, val)
}

// After query the date-time field is after a specific date-time.
func (f DateTimeField) After(val DateTimeValue) Condition {
	return f.dateTimeBinary(ast2.OperatorAfter, val)
}

// AtOrAfter query the date-time field is at or after a specific date-time.
func (f DateTimeField) AtOrAfter(val DateTimeValue) Condition {
	return f.dateTimeBinary(ast2.OperatorAtOrAfter, val)
}

// Between query the date-time field is between the provided start and end values.
func (f DateTimeField) Between(start, end time.Time) Condition {
	if start.After(end) {
		return NewErrorCondition(fmt.Errorf("start time %v is after end time %v", start, end))
	}
	return f.pair(ast2.OperatorBetween, NewDateTimeValue(start), NewDateTimeValue(end))
}

// Javascript allows using a custom JavaScript expression (e.g., gs.daysAgoStart(0)).
func (f DateTimeField) Javascript(expr string) Condition {
	return f.On(JS(expr))
}

// Today query that field is today.
func (f DateTimeField) Today() Condition {
	return f.OnSpecialty("Today", "gs.beginningOfToday()", "gs.endOfToday()")
}

// Yesterday query that field is yesterday.
func (f DateTimeField) Yesterday() Condition {
	return f.OnSpecialty("Yesterday", "gs.beginningOfYesterday()", "gs.endOfYesterday()")
}

// Tomorrow query that field is tomorrow.
func (f DateTimeField) Tomorrow() Condition {
	return f.OnSpecialty("Tomorrow", "gs.beginningOfTomorrow()", "gs.endOfTomorrow()")
}

// ThisWeek query that field is this week.
func (f DateTimeField) ThisWeek() Condition {
	return f.OnSpecialty("This week", "gs.beginningOfThisWeek()", "gs.endOfThisWeek()")
}

// LastWeek query that field is last week.
func (f DateTimeField) LastWeek() Condition {
	return f.OnSpecialty("Last week", "gs.beginningOfLastWeek()", "gs.endOfLastWeek()")
}

// ThisMonth query that field is this month.
func (f DateTimeField) ThisMonth() Condition {
	return f.OnSpecialty("This month", "gs.beginningOfThisMonth()", "gs.endOfThisMonth()")
}

// LastMonth query that field is last month.
func (f DateTimeField) LastMonth() Condition {
	return f.OnSpecialty("Last month", "gs.beginningOfLastMonth()", "gs.endOfLastMonth()")
}

// ThisYear query that field is this year.
func (f DateTimeField) ThisYear() Condition {
	return f.OnSpecialty("This year", "gs.beginningOfThisYear()", "gs.endOfThisYear()")
}

// LastYear query that field is last year.
func (f DateTimeField) LastYear() Condition {
	return f.OnSpecialty("Last year", "gs.beginningOfLastYear()", "gs.endOfLastYear()")
}

// OnSpecialty builds a specialty date-time condition (e.g., ONToday@javascript:...).
func (f DateTimeField) OnSpecialty(label, startExpr, endExpr string) Condition {
	val := fmt.Sprintf("%s@javascript:%s@javascript:%s", label, startExpr, endExpr)
	return f.On(DateTimeValue{literal: val})
}

// IsMoreThan query that field is more than the provided value.
func (f DateTimeField) IsMoreThan(val string) Condition {
	return f.binary(ast2.OperatorIsMoreThan, val)
}

// IsLessThan query that field is less than the provided value.
func (f DateTimeField) IsLessThan(val string) Condition {
	return f.binary(ast2.OperatorIsLessThan, val)
}

// TrendOnOrAfter query the date-time field that trends on or after a specific value.
func (f DateTimeField) TrendOnOrAfter(val string) Condition {
	return f.binary(ast2.OperatorTrendOnOrAfter, val)
}

// TrendOnOrBefore query the date-time field that trends on or before a specific value.
func (f DateTimeField) TrendOnOrBefore(val string) Condition {
	return f.binary(ast2.OperatorTrendOnOrBefore, val)
}

// TrendAfter query the date-time field that trends after a specific value.
func (f DateTimeField) TrendAfter(val string) Condition {
	return f.binary(ast2.OperatorTrendAfter, val)
}

// TrendBefore query the date-time field that trends before a specific value.
func (f DateTimeField) TrendBefore(val string) Condition {
	return f.binary(ast2.OperatorTrendBefore, val)
}

// TrendOn query the date-time field that trends on a specific value.
func (f DateTimeField) TrendOn(val string) Condition {
	return f.binary(ast2.OperatorTrendOn, val)
}

// RelativeAfter query the date-time field that is relatively after a specific value.
func (f DateTimeField) RelativeAfter(val string) Condition {
	return f.binary(ast2.OperatorRelativeAfter, val)
}

// RelativeBefore query the date-time field that is relatively before a specific value.
func (f DateTimeField) RelativeBefore(val string) Condition {
	return f.binary(ast2.OperatorRelativeBefore, val)
}
