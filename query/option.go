package query

import "time"

type option func(*query)

type numeric interface {
	uint8 | uint16 | uint32 | uint64 | int8 | int16 | int32 | int64 | float32 | float64 | complex64 | complex128
}

func And(opts ...option) option {
	return logicOpt(and, opts...)
}

func Or(opts ...option) option {
	return logicOpt(or, opts...)
}

func logicOpt(oper logicalOperator, opts ...option) option {
	return func(q *query) {
		for i, opt := range opts {
			opt(q)
			if i+1 < len(opts) {
				q.AddValue(oper)
			}
		}
	}
}

type startsWithOptAllows interface {
	string
}

func conditionOpt[T any](field string, oper relationalOperator, value T) option {
	return func(q *query) {
		q.AddValue(newCondition(field, oper, &value))
	}
}

func StartsWith[T startsWithOptAllows](field string, value T) option {
	return conditionOpt(field, startsWith, value)
}

type endsWithOptAllows interface {
	string
}

func EndsWith[t endsWithOptAllows](field string, value t) option {
	return func(q *query) {
		q.AddValue(newCondition(field, endsWith, &value))
	}
}

type containsOptAllows interface {
	string
}

func Contains[t containsOptAllows](field string, value t) option {
	return func(q *query) {
		q.AddValue(newCondition(field, contains, &value))
	}
}

type doesNotContainOptAllows interface {
	string
}

func DoesNotContain[t doesNotContainOptAllows](field string, value t) option {
	return func(q *query) {
		q.AddValue(newCondition(field, doesNotContain, &value))
	}
}

type isOptAllows interface {
	string | int | bool | numeric
}

func Is[t isOptAllows](field string, value t) option {
	return func(q *query) {
		q.AddValue(newCondition(field, is, &value))
	}
}

type isNotOptAllows interface {
	string | int | bool | numeric
}

func IsNot[t isNotOptAllows](field string, value t) option {
	return func(q *query) {
		q.AddValue(newCondition(field, isNot, &value))
	}
}

type isEmptyOptAllows interface {
	string | int | time.Time | bool | numeric
}

func IsEmpty[t isEmptyOptAllows](field string) option {
	return func(q *query) {
		q.AddValue(newCondition[t](field, isEmpty, nil))
	}
}

type isNotEmptyOptAllows interface {
	string | int | time.Time | bool | numeric
}

func IsNotEmpty[t isNotEmptyOptAllows](field string) option {
	return func(q *query) {
		q.AddValue(newCondition[t](field, isNotEmpty, nil))
	}
}

type isAnythingOptAllows interface {
	string | int | time.Time | bool | numeric
}

func IsAnything[t isAnythingOptAllows](field string) option {
	return func(q *query) {
		q.AddValue(newCondition[t](field, isAnything, nil))
	}
}

type isEmptyStringOptAllows interface {
	string
}

func IsEmptyString[t isEmptyStringOptAllows](field string) option {
	return func(q *query) {
		q.AddValue(newCondition[t](field, isEmptyString, nil))
	}
}

type lessThanOrIsOptAllows interface {
	string | int | numeric
}

func LessThanOrIs[t lessThanOrIsOptAllows](field string, value t) option {
	return func(q *query) {
		q.AddValue(newCondition(field, lessThanOrIs, &value))
	}
}

type greaterThanOrIsOptAllows interface {
	string | int | numeric
}

func GreaterThanOrIs[t greaterThanOrIsOptAllows](field string, value t) option {
	return func(q *query) {
		q.AddValue(newCondition(field, greaterThanOrIs, &value))
	}
}

type betweenOptAllows interface {
	string | int | time.Time | numeric
}

func Between[t betweenOptAllows](field string, start, end t) option {
	list := newList("%", start, end)
	return func(q *query) {
		q.AddValue(newCondition(field, between, &list))
	}
}

type isSameOptAllows interface {
	string | int | time.Time | bool | numeric
}

func IsSame[t isSameOptAllows](field1, field2 string) option {
	return func(q *query) {
		q.AddValue(newCondition(field1, isSame, &field2))
	}
}

func IsDifferent(field1, field2 string) option {
	return func(q *query) {
		q.AddValue(newCondition(field1, isDifferent, &field2))
	}
}

type isOneOfOptAllows interface {
	int
}

func IsOneOf[t isOneOfOptAllows](field string, values []t) option {
	return func(q *query) {
		q.AddValue(newCondition(field, isOneOf, &values))
	}
}

type isNotOneOfOptAllows interface {
	int
}

func IsNotOneOf[t isNotOneOfOptAllows](field string, values []t) option {
	return func(q *query) {
		q.AddValue(newCondition(field, isNotOneOf, &values))
	}
}

type lessThanOptAllows interface {
	int | numeric
}

func LessThan[t lessThanOptAllows](field string, value t) option {
	return func(q *query) {
		q.AddValue(newCondition(field, lessThan, &value))
	}
}

type greaterThanOptAllows interface {
	int | numeric
}

func GreaterThan[t greaterThanOptAllows](field string, value t) option {
	return func(q *query) {
		q.AddValue(newCondition(field, greaterThan, &value))
	}
}

type onOptAllows interface {
	time.Time
}

func On[t onOptAllows](field string, value t) option {
	return func(q *query) {
		q.AddValue(newCondition(field, on, &value))
	}
}

type notOnOptAllows interface {
	time.Time
}

func NotOn[t notOnOptAllows](field string, value t) option {
	return func(q *query) {
		q.AddValue(newCondition(field, notOn, &value))
	}
}

type beforeOptAllows interface {
	time.Time
}

func Before[t beforeOptAllows](field string, value t) option {
	return func(q *query) {
		q.AddValue(newCondition(field, before, &value))
	}
}

type atOrBeforeOptAllows interface {
	time.Time | time.Duration
}

func AtOrBefore[t atOrBeforeOptAllows](field string, value t) option {
	return func(q *query) {
		q.AddValue(newCondition(field, atOrBefore, &value))
	}
}

type afterOptAllows interface {
	time.Time | time.Duration
}

func After[t afterOptAllows](field string, value t) option {
	return func(q *query) {
		q.AddValue(newCondition(field, after, &value))
	}
}

type relativeAfterOptAllows interface {
	time.Time | time.Duration
}

func RelativeAfter[T relativeAfterOptAllows](field string, value T) option {
	return func(q *query) {
		q.AddValue(newCondition(field, relativeAfter, &value))
	}
}

type relativeBeforeOptAllows interface {
	time.Time | time.Duration
}

func RelativeBefore[T relativeBeforeOptAllows](field string, value T) option {
	return func(q *query) {
		q.AddValue(newCondition(field, relativeBefore, &value))
	}
}

type trendOnOrAfterOptAllows interface {
	time.Time | time.Duration
}

func TrendOnOrAfter[T trendOnOrAfterOptAllows](field string, value T) option {
	return func(q *query) {
		q.AddValue(newCondition(field, trendOnOrAfter, &value))
	}
}

type trendOnOrBeforeOptAllows interface {
	time.Time | time.Duration
}

func TrendOnOrBefore[T trendOnOrBeforeOptAllows](field string, value T) option {
	return func(q *query) {
		q.AddValue(newCondition(field, trendOnOrBefore, &value))
	}
}

type trendAfterOptAllows interface {
	time.Time | time.Duration
}

func TrendAfter[T trendAfterOptAllows](field string, value T) option {
	return func(q *query) {
		q.AddValue(newCondition(field, trendAfter, &value))
	}
}

type trendBeforeOptAllows interface {
	time.Time | time.Duration
}

func TrendBefore[T trendBeforeOptAllows](field string, value T) option {
	return func(q *query) {
		q.AddValue(newCondition(field, trendBefore, &value))
	}
}

type trendOnOptAllows interface {
	time.Time | time.Duration
}

func TrendOn[T trendOnOptAllows](field string, value T) option {
	return func(q *query) {
		q.AddValue(newCondition(field, trendOn, &value))
	}
}

type atOrAfterOptAllows interface {
	time.Time
}

func AtOrAfter[t atOrAfterOptAllows](field string, value t) option {
	return func(q *query) {
		q.AddValue(newCondition(field, atOrAfter, &value))
	}
}

type greaterThanOrIsFieldOptAllows interface {
	numeric
}

func GreaterThanOrIsField[t greaterThanOrIsFieldOptAllows](field1 string, field2 string) option {
	return func(q *query) {
		q.AddValue(newCondition(field1, greaterThanOrIsField, &field2))
	}
}

type greaterThanFieldOptAllows interface {
	numeric
}

func GreaterThanField[t greaterThanFieldOptAllows](field1 string, field2 string) option {
	return func(q *query) {
		q.AddValue(newCondition(field1, greaterThanField, &field2))
	}
}

type lessThanFieldOptAllows interface {
	numeric
}

func LessThanField[t lessThanFieldOptAllows](field1 string, field2 string) option {
	return func(q *query) {
		q.AddValue(newCondition(field1, lessThanField, &field2))
	}
}

type lessThanOrIsFieldOptAllows interface {
	numeric
}

func LessThanOrIsField[t lessThanOrIsFieldOptAllows](field1 string, field2 string) option {
	return func(q *query) {
		q.AddValue(newCondition(field1, lessThanOrIsField, &field2))
	}
}
