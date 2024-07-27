package query

import (
	"time"
)

type numeric interface {
	uint8
	uint16
	uint32
	uint64
	int8
	int16
	int32
	int64
	float32
	float64
	complex64
	complex128
}

// how to handle:
// - references
// - choice
// https://docs.servicenow.com/bundle/washingtondc-platform-user-interface/page/use/common-ui-elements/reference/r_OpAvailableFiltersQueries.html

type startsWithOptAllows interface {
	string
}

func StartsWith[t startsWithOptAllows](field string, value t) queryOption {
	return func(q *query) {
		q.addFragment(newFragment(field, startsWith, &value), unset)
	}
}

type endsWithOptAllows interface {
	string
}

func EndsWith[t endsWithOptAllows](field string, value t) queryOption {
	return func(q *query) {
		q.addFragment(newFragment(field, endsWith, &value), unset)
	}
}

type containsOptAllows interface {
	string
}

func Contains[t containsOptAllows](field string, value t) queryOption {
	return func(q *query) {
		q.addFragment(newFragment(field, contains, &value), unset)
	}
}

type doesNotContainOptAllows interface {
	string
}

func DoesNotContain[t doesNotContainOptAllows](field string, value t) queryOption {
	return func(q *query) {
		q.addFragment(newFragment(field, doesNotContain, &value), unset)
	}
}

type isOptAllows interface {
	string | int | bool | numeric
}

func Is[t isOptAllows](field string, value t) queryOption {
	return func(q *query) {
		q.addFragment(newFragment(field, is, &value), unset)
	}
}

type isNotOptAllows interface {
	string
	int
	bool
	numeric
}

func IsNot[t isNotOptAllows](field string, value t) queryOption {
	return func(q *query) {
		q.addFragment(newFragment(field, isNot, &value), unset)
	}
}

type isEmptyOptAllows interface {
	string
	int
	time.Time
	bool
	numeric
}

func IsEmpty[t isEmptyOptAllows](field string) queryOption {
	return func(q *query) {
		q.addFragment(newFragment[t](field, isEmpty, nil), unset)
	}
}

type isNotEmptyOptAllows interface {
	string
	int
	time.Time
	bool
	numeric
}

func IsNotEmpty[t isNotEmptyOptAllows](field string) queryOption {
	return func(q *query) {
		q.addFragment(newFragment[t](field, isNotEmpty, nil), unset)
	}
}

type isAnythingOptAllows interface {
	string
	int
	time.Time
	bool
	numeric
}

func IsAnything[t isAnythingOptAllows](field string) queryOption {
	return func(q *query) {
		q.addFragment(newFragment[t](field, isAnything, nil), unset)
	}
}

type isEmptyStringOptAllows interface {
	string
}

func IsEmptyString[t isEmptyStringOptAllows](field string) queryOption {
	return func(q *query) {
		q.addFragment(newFragment[t](field, isEmptyString, nil), unset)
	}
}

type lessThanOrIsOptAllows interface {
	string
	int
	numeric
}

func LessThanOrIs[t lessThanOrIsOptAllows](field string, value t) queryOption {
	return func(q *query) {
		q.addFragment(newFragment(field, lessThanOrIs, &value), unset)
	}
}

type greaterThanOrIsOptAllows interface {
	string
	int
	numeric
}

func GreaterThanOrIs[t greaterThanOrIsOptAllows](field string, value t) queryOption {
	return func(q *query) {
		q.addFragment(newFragment(field, greaterThanOrIs, &value), unset)
	}
}

//type betweenOptAllows interface {
//	string
//	int
//	time.Time
//	numeric
//}

//func Between[t betweenOptAllows](field string, start, end t) queryOption {
//	return func(q *query) {
//how to handle?
//		q.AddFragment(NewFragment(field, BETWEEN), unset)
//	}
//}

type isSameOptAllows interface {
	string
	int
	time.Time
	bool
	numeric
}

func IsSame[t isSameOptAllows](field1, field2 string) queryOption {
	return func(q *query) {
		q.addFragment(newFragment(field1, isSame, &field2), unset)
	}
}

type isDifferentOptAllows interface {
	string
	int
	time.Time
	bool
	numeric
}

func IsDifferent[t isDifferentOptAllows](field1, field2 string) queryOption {
	return func(q *query) {
		q.addFragment(newFragment(field1, isDifferent, &field2), unset)
	}
}

type isOneOfOptAllows interface {
	int
}

func IsOneOf[t isOneOfOptAllows](field string, values []t) queryOption {
	return func(q *query) {
		q.addFragment(newFragment(field, isOneOf, &values), unset)
	}
}

type isNotOneOfOptAllows interface {
	int
}

func IsNotOneOf[t isNotOneOfOptAllows](field string, values []t) queryOption {
	return func(q *query) {
		q.addFragment(newFragment(field, isNotOneOf, &values), unset)
	}
}

type lessThanOptAllows interface {
	int
	numeric
}

func LessThan[t lessThanOptAllows](field string, value t) queryOption {
	return func(q *query) {
		q.addFragment(newFragment(field, lessThan, &value), unset)
	}
}

type greaterThanOptAllows interface {
	int
	numeric
}

func GreaterThan[t greaterThanOptAllows](field string, value t) queryOption {
	return func(q *query) {
		q.addFragment(newFragment(field, greaterThan, &value), unset)
	}
}

type onOptAllows interface {
	time.Time
}

func On[t onOptAllows](field string, value t) queryOption {
	return func(q *query) {
		q.addFragment(newFragment(field, on, &value), unset)
	}
}

type notOnOptAllows interface {
	time.Time
}

func NotOn[t notOnOptAllows](field string, value t) queryOption {
	return func(q *query) {
		q.addFragment(newFragment(field, notOn, &value), unset)
	}
}

type beforeOptAllows interface {
	time.Time
}

func Before[t beforeOptAllows](field string, value t) queryOption {
	return func(q *query) {
		q.addFragment(newFragment(field, before, &value), unset)
	}
}

type atOrBeforeOptAllows interface {
	time.Time | time.Duration
}

func AtOrBefore[t atOrBeforeOptAllows](field string, value t) queryOption {
	return func(q *query) {
		q.addFragment(newFragment(field, atOrBefore, &value), unset)
	}
}

type afterOptAllows interface {
	time.Time | time.Duration
}

func After[t afterOptAllows](field string, value t) queryOption {
	return func(q *query) {
		q.addFragment(newFragment(field, after, &value), unset)
	}
}

type relativeAfterOptAllows interface {
	time.Time | time.Duration
}

func RelativeAfter[T relativeAfterOptAllows](field string, value T) queryOption {
	return func(q *query) {
		q.addFragment(newFragment(field, relativeAfter, &value), unset)
	}
}

type relativeBeforeOptAllows interface {
	time.Time | time.Duration
}

func RelativeBefore[T relativeBeforeOptAllows](field string, value T) queryOption {
	return func(q *query) {
		q.addFragment(newFragment(field, relativeBefore, &value), unset)
	}
}

type trendOnOrAfterOptAllows interface {
	time.Time | time.Duration
}

func TrendOnOrAfter[T trendOnOrAfterOptAllows](field string, value T) queryOption {
	return func(q *query) {
		q.addFragment(newFragment(field, trendOnOrAfter, &value), unset)
	}
}

type trendOnOrBeforeOptAllows interface {
	time.Time | time.Duration
}

func TrendOnOrBefore[T trendOnOrBeforeOptAllows](field string, value T) queryOption {
	return func(q *query) {
		q.addFragment(newFragment(field, trendOnOrBefore, &value), unset)
	}
}

type trendAfterOptAllows interface {
	time.Time | time.Duration
}

func TrendAfter[T trendAfterOptAllows](field string, value T) queryOption {
	return func(q *query) {
		q.addFragment(newFragment(field, trendAfter, &value), unset)
	}
}

type trendBeforeOptAllows interface {
	time.Time | time.Duration
}

func TrendBefore[T trendBeforeOptAllows](field string, value T) queryOption {
	return func(q *query) {
		q.addFragment(newFragment(field, trendBefore, &value), unset)
	}
}

type trendOnOptAllows interface {
	time.Time | time.Duration
}

func TrendOn[T trendOnOptAllows](field string, value T) queryOption {
	return func(q *query) {
		q.addFragment(newFragment(field, trendOn, &value), unset)
	}
}

type atOrAfterOptAllows interface {
	time.Time
}

func AtOrAfter[t atOrAfterOptAllows](field string, value t) queryOption {
	return func(q *query) {
		q.addFragment(newFragment(field, atOrAfter, &value), unset)
	}
}

type greaterThanOrIsFieldOptAllows interface {
	numeric
}

func GreaterThanOrIsField[t greaterThanOrIsFieldOptAllows](field1 string, field2 string) queryOption {
	return func(q *query) {
		q.addFragment(newFragment(field1, greaterThanOrIsField, &field2), unset)
	}
}

type greaterThanFieldOptAllows interface {
	numeric
}

func GreaterThanField[t greaterThanFieldOptAllows](field1 string, field2 string) queryOption {
	return func(q *query) {
		q.addFragment(newFragment(field1, greaterThanField, &field2), unset)
	}
}

type lessThanFieldOptAllows interface {
	numeric
}

func LessThanField[t lessThanFieldOptAllows](field1 string, field2 string) queryOption {
	return func(q *query) {
		q.addFragment(newFragment(field1, lessThanField, &field2), unset)
	}
}

type lessThanOrIsFieldOptAllows interface {
	numeric
}

func LessThanOrIsField[t lessThanOrIsFieldOptAllows](field1 string, field2 string) queryOption {
	return func(q *query) {
		q.addFragment(newFragment(field1, lessThanOrIsField, &field2), unset)
	}
}

// email notifications?
