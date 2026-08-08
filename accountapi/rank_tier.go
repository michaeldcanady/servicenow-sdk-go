package accountapi

import "fmt"

type RankTier int16

const (
	RankTierUnknown RankTier = iota - 1
	RankTierBlacklist
	RankTierOther
	RankTierStrategic
	RankTierTactical
	RankTierValued

	rankTierUnknownString   = "unknown"
	rankTierBlacklistString = "blacklist"
	rankTierOtherString     = "other"
	rankTierStrategicString = "strategic"
	rankTierTacticalString  = "tactical"
	rankTierValuedString    = "valued"
)

func ParseRankTier(s string) (interface{}, error) {
	switch s {
	case rankTierBlacklistString:
		return RankTierBlacklist, nil
	case rankTierOtherString:
		return RankTierOther, nil
	case rankTierStrategicString:
		return RankTierStrategic, nil
	case rankTierTacticalString:
		return RankTierTactical, nil
	case rankTierValuedString:
		return RankTierValued, nil
	default:
		return RankTierUnknown, fmt.Errorf("invalid rank tier: %s", s)
	}
}

func (r RankTier) String() string {
	switch r {
	case RankTierBlacklist:
		return rankTierBlacklistString
	case RankTierOther:
		return rankTierOtherString
	case RankTierStrategic:
		return rankTierStrategicString
	case RankTierTactical:
		return rankTierTacticalString
	case RankTierValued:
		return rankTierValuedString
	default:
		return rankTierUnknownString
	}
}
