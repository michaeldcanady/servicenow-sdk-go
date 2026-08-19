package accountapi

import (
	"fmt"
	"math"
	"testing"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ParseRankTier is used as a Kiota EnumFactory by account.go's rank_tier field
// deserializer. rank_tier.go itself doesn't assert that (policyapi/state.go and
// policyapi/input_status.go both do), so pin it here.
var _ serialization.EnumFactory = ParseRankTier

// RankTier's String method must stay on a value receiver to satisfy fmt.Stringer
// for the value type — the backing store holds *RankTier, but fmt verbs and
// SerializeEnumFunc are reached through the value.
var _ fmt.Stringer = RankTierUnknown

// TestRankTierValues pins the numeric value of each constant. The declaration
// block interleaves the RankTier iota constants with the untyped string
// constants, so iota's position depends on that ordering being left alone;
// reordering or inserting a member silently shifts every value stored in a
// backing store.
func TestRankTierValues(t *testing.T) {
	tests := []struct {
		name string
		val  RankTier
		want int16
	}{
		{"Unknown", RankTierUnknown, -1},
		{"Blacklist", RankTierBlacklist, 0},
		{"Other", RankTierOther, 1},
		{"Strategic", RankTierStrategic, 2},
		{"Tactical", RankTierTactical, 3},
		{"Valued", RankTierValued, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, int16(tt.val))
		})
	}
}

func TestRankTierString(t *testing.T) {
	tests := []struct {
		name string
		val  RankTier
		want string
	}{
		{"Blacklist", RankTierBlacklist, "blacklist"},
		{"Other", RankTierOther, "other"},
		{"Strategic", RankTierStrategic, "strategic"},
		{"Tactical", RankTierTactical, "tactical"},
		{"Valued", RankTierValued, "valued"},
		{"Unknown", RankTierUnknown, "unknown"},
		// Everything outside the declared set collapses onto "unknown" rather
		// than rendering a numeric value.
		{"above the declared range", RankTier(999), "unknown"},
		{"below the declared range", RankTier(-2), "unknown"},
		{"int16 min", RankTier(math.MinInt16), "unknown"},
		{"int16 max", RankTier(math.MaxInt16), "unknown"},
		{"zero value is Blacklist, not Unknown", RankTier(0), "blacklist"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.val.String())
		})
	}
}

// TestRankTierStringViaFmt confirms String() is what fmt reaches for, so a
// RankTier interpolated into an error or log line reads as its wire value.
func TestRankTierStringViaFmt(t *testing.T) {
	assert.Equal(t, "strategic", fmt.Sprintf("%v", RankTierStrategic))
	assert.Equal(t, "an unknown tier: unknown", fmt.Sprintf("an unknown tier: %v", RankTier(999)))
}

func TestParseRankTier(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    RankTier
		wantErr bool
	}{
		{"blacklist", "blacklist", RankTierBlacklist, false},
		{"other", "other", RankTierOther, false},
		{"strategic", "strategic", RankTierStrategic, false},
		{"tactical", "tactical", RankTierTactical, false},
		{"valued", "valued", RankTierValued, false},

		// Every unrecognized input returns RankTierUnknown *and* a non-nil
		// error; callers that ignore the error still get a usable sentinel.
		{"empty string", "", RankTierUnknown, true},
		{"whitespace only", "   ", RankTierUnknown, true},
		{"unrecognized value", "something-else", RankTierUnknown, true},
		{"numeric string", "2", RankTierUnknown, true},

		// Matching is exact: no case folding and no trimming.
		{"capitalized", "Blacklist", RankTierUnknown, true},
		{"uppercased", "BLACKLIST", RankTierUnknown, true},
		{"leading space", " blacklist", RankTierUnknown, true},
		{"trailing space", "blacklist ", RankTierUnknown, true},
		{"trailing newline", "blacklist\n", RankTierUnknown, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseRankTier(tt.input)

			// The unknown sentinel comes back on both paths.
			assert.Equal(t, tt.want, got)
			// Kiota's EnumFactory contract is interface{}; account.go's
			// deserializer branches on whether that holds T or *T, so pin
			// which branch this factory exercises.
			assert.IsType(t, RankTier(0), got)

			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

// TestRankTierRoundTrip walks each declared value out through String() and back
// in through ParseRankTier — the exact path a value takes through Serialize and
// then GetFieldDeserializers.
func TestRankTierRoundTrip(t *testing.T) {
	for _, val := range []RankTier{
		RankTierBlacklist,
		RankTierOther,
		RankTierStrategic,
		RankTierTactical,
		RankTierValued,
	} {
		t.Run(val.String(), func(t *testing.T) {
			got, err := ParseRankTier(val.String())
			require.NoError(t, err)
			assert.Equal(t, val, got)
		})
	}
}

// TestRankTierUnknownDoesNotRoundTrip documents an asymmetry between String()
// and ParseRankTier: String() renders RankTierUnknown (and any out-of-range
// value) as "unknown", but ParseRankTier has no case for "unknown" and rejects
// it. So the SDK emits a rank_tier value it will then refuse to read back.
//
// Both sibling enums in this repo do accept their unknown string —
// policyapi.ParseState and policyapi.ParseInputStatus each have a
// `case stateUnknown:` / `case inputStatusUnknown:` arm — so this looks like an
// omission in ParseRankTier rather than a deliberate choice.
//
// BUG: ParseRankTier rejects "unknown", which String() can produce.
func TestRankTierUnknownDoesNotRoundTrip(t *testing.T) {
	// The value Serialize would write for an unknown or out-of-range tier.
	serialized := RankTierUnknown.String()
	require.Equal(t, "unknown", serialized)
	require.Equal(t, "unknown", RankTier(999).String())

	got, err := ParseRankTier(serialized)

	// Current behavior: rejected. Once ParseRankTier gains an "unknown" case,
	// this becomes require.NoError and the round trip closes.
	require.Error(t, err)
	assert.Equal(t, RankTierUnknown, got)
}

// TestParseRankTierStringRoundTrip is the reverse direction: a wire value
// ServiceNow sends must survive being parsed and re-rendered unchanged.
func TestParseRankTierStringRoundTrip(t *testing.T) {
	for _, s := range []string{"blacklist", "other", "strategic", "tactical", "valued"} {
		t.Run(s, func(t *testing.T) {
			parsed, err := ParseRankTier(s)
			require.NoError(t, err)

			tier, ok := parsed.(RankTier)
			require.True(t, ok, "ParseRankTier must return a RankTier")
			assert.Equal(t, s, tier.String())
		})
	}
}

// TestParseRankTierDistinctValues guards against two string cases accidentally
// mapping onto the same constant, which a copy-paste in either switch would
// otherwise hide.
func TestParseRankTierDistinctValues(t *testing.T) {
	inputs := []string{"blacklist", "other", "strategic", "tactical", "valued"}

	seen := make(map[RankTier]string, len(inputs))
	for _, s := range inputs {
		parsed, err := ParseRankTier(s)
		require.NoError(t, err)

		tier, ok := parsed.(RankTier)
		require.True(t, ok)

		if prev, dup := seen[tier]; dup {
			t.Fatalf("%q and %q both parse to %d", prev, s, int16(tier))
		}
		seen[tier] = s
	}

	assert.Len(t, seen, len(inputs))
}
