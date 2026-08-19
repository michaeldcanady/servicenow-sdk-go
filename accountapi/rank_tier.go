package accountapi

import (
	"fmt"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
)

// RankTier represents the strategic ranking tier of an account.
type RankTier int16

const (
	RankTierUnknown RankTier = iota - 1
	RankTierBlacklist
	RankTierOther
	RankTierStrategic
	RankTierTactical
	RankTierValued

	rankTierUnknown   = "unknown"
	rankTierBlacklist = "blacklist"
	rankTierOther     = "other"
	rankTierStrategic = "strategic"
	rankTierTactical  = "tactical"
	rankTierValued    = "valued"
)

var rankTierStrings = map[RankTier]string{
	RankTierBlacklist: rankTierBlacklist,
	RankTierOther:     rankTierOther,
	RankTierStrategic: rankTierStrategic,
	RankTierTactical:  rankTierTactical,
	RankTierValued:    rankTierValued,
}

// ParseRankTier converts a string to a [RankTier] value, returning an error for unrecognized values.
func ParseRankTier(s string) (interface{}, error) {
	switch s {
	case rankTierBlacklist:
		return RankTierBlacklist, nil
	case rankTierOther:
		return RankTierOther, nil
	case rankTierStrategic:
		return RankTierStrategic, nil
	case rankTierTactical:
		return RankTierTactical, nil
	case rankTierValued:
		return RankTierValued, nil
	default:
		return RankTierUnknown, fmt.Errorf("invalid rank tier: %s", s)
	}
}

func (r RankTier) String() string {
	return conversion.EnumString(rankTierStrings, r, rankTierUnknown)
}
