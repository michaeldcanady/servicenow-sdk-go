package query

type relationalOperator int64

func (o relationalOperator) String() string {
	return map[relationalOperator]string{
		startsWith:           "STARTSWITH",
		endsWith:             "ENDSWITH",
		contains:             "LIKE",
		doesNotContain:       "NOT LIKE",
		is:                   "=",
		isNot:                "!=",
		isEmpty:              "ISEMPTY",
		isNotEmpty:           "ISNOTEMPTY",
		isAnything:           "ANYTHING",
		isEmptyString:        "EMPTYSTRING",
		lessThanOrIs:         "<=",
		greaterThanOrIs:      ">=",
		between:              "BETWEEN",
		isSame:               "SAMEAS",
		isDifferent:          "NSAMEAS",
		isOneOf:              "IN",
		isNotOneOf:           "NOT IN",
		lessThan:             "<",
		greaterThan:          ">",
		on:                   "ON",
		notOn:                "NOTON",
		before:               "<",
		atOrBefore:           "<=",
		after:                ">",
		atOrAfter:            ">=",
		greaterThanOrIsField: "GT_OR_EQUALS_FIELD",
		greaterThanField:     "GT_FIELD",
		lessThanField:        "LT_FIELD",
		lessThanOrIsField:    "LT_OR_EQUALS_FIELD",
		relativeAfter:        "RELATIVEGT",
		relativeBefore:       "RELATIVELT",
		trendOnOrAfter:       "DATEPART",
		trendOnOrBefore:      "DATEPART",
		trendAfter:           "DATEPART",
		trendBefore:          "DATEPART",
		trendOn:              "DATEPART",
	}[o]
}

const (
	startsWith relationalOperator = iota
	endsWith
	contains
	doesNotContain
	is
	isNot
	isEmpty
	isNotEmpty
	isAnything
	isEmptyString
	lessThanOrIs
	greaterThanOrIs
	between
	isSame
	isDifferent
	isOneOf
	isNotOneOf
	lessThan
	greaterThan
	on
	notOn
	before
	atOrBefore
	after
	relativeAfter
	relativeBefore
	atOrAfter
	greaterThanOrIsField
	greaterThanField
	lessThanField
	lessThanOrIsField
	trendOnOrAfter
	trendOnOrBefore
	trendAfter
	trendBefore
	trendOn
)
