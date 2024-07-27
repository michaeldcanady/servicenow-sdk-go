package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRelationalOperatorString(t *testing.T) {
	assert.Equal(t, "STARTSWITH", startsWith.String())
	assert.Equal(t, "ENDSWITH", endsWith.String())
	assert.Equal(t, "LIKE", contains.String())
	assert.Equal(t, "NOT LIKE", doesNotContain.String())
	assert.Equal(t, "=", is.String())
	assert.Equal(t, "!=", isNot.String())
	assert.Equal(t, "ISEMPTY", isEmpty.String())
	assert.Equal(t, "ISNOTEMPTY", isNotEmpty.String())
	assert.Equal(t, "ANYTHING", isAnything.String())
	assert.Equal(t, "EMPTYSTRING", isEmptyString.String())
	assert.Equal(t, "<=", lessThanOrIs.String())
	assert.Equal(t, ">=", greaterThanOrIs.String())
	assert.Equal(t, "BETWEEN", between.String())
	assert.Equal(t, "SAMEAS", isSame.String())
	assert.Equal(t, "NSAMEAS", isDifferent.String())
	assert.Equal(t, "IN", isOneOf.String())
	assert.Equal(t, "NOT IN", isNotOneOf.String())
	assert.Equal(t, "<", lessThan.String())
	assert.Equal(t, ">", greaterThan.String())
	assert.Equal(t, "ON", on.String())
	assert.Equal(t, "NOTON", notOn.String())
	assert.Equal(t, "<", before.String())
	assert.Equal(t, "<=", atOrBefore.String())
	assert.Equal(t, ">", after.String())
	assert.Equal(t, ">=", atOrAfter.String())
	assert.Equal(t, "GT_OR_EQUALS_FIELD", greaterThanOrIsField.String())
	assert.Equal(t, "GT_FIELD", greaterThanField.String())
	assert.Equal(t, "LT_FIELD", lessThanField.String())
	assert.Equal(t, "LT_OR_EQUALS_FIELD", lessThanOrIsField.String())
	assert.Equal(t, "RELATIVEGT", relativeAfter.String())
	assert.Equal(t, "RELATIVELT", relativeBefore.String())
	assert.Equal(t, "DATEPART", trendOnOrAfter.String())
	assert.Equal(t, "DATEPART", trendOnOrBefore.String())
	assert.Equal(t, "DATEPART", trendAfter.String())
	assert.Equal(t, "DATEPART", trendBefore.String())
	assert.Equal(t, "DATEPART", trendOn.String())
}
