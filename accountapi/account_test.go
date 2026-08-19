package accountapi

import (
	"errors"
	"testing"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	// errWrite stands in for an error a serialization.SerializationWriter returns from a Write* call.
	errWrite = errors.New("write error")
	// errParse stands in for an error a serialization.ParseNode returns from a Get* call.
	errParse = errors.New("parse error")
)

// stringFields pairs every string-typed property key with its accessor/mutator, so the
// getter/setter, Serialize and GetFieldDeserializers tables can all iterate the same list
// and a key wired to the wrong accessor shows up as a failure.
func stringFields() []struct {
	name   string
	key    string
	getter func(*Account) (*string, error)
	setter func(*Account, *string) error
} {
	return []struct {
		name   string
		key    string
		getter func(*Account) (*string, error)
		setter func(*Account, *string) error
	}{
		{"BannerImageLight", bannerImageLightKey, (*Account).GetBannerImageLight, (*Account).SetBannerImageLight},
		{"Country", countryKey, (*Account).GetCountry, (*Account).SetCountry},
		{"Parent", parentKey, (*Account).GetParent, (*Account).SetParent},
		{"Notes", notesKey, (*Account).GetNotes, (*Account).SetNotes},
		{"StockSymbol", stockSymbolKey, (*Account).GetStockSymbol, (*Account).SetStockSymbol},
		{"ActiveEscalation", activeEscalationKey, (*Account).GetActiveEscalation, (*Account).SetActiveEscalation},
		{"SysUpdatedOn", sysUpdatedOnKey, (*Account).GetSysUpdatedOn, (*Account).SetSysUpdatedOn},
		{"AppleIcon", appleIconKey, (*Account).GetAppleIcon, (*Account).SetAppleIcon},
		{"Number", numberKey, (*Account).GetNumber, (*Account).SetNumber},
		{"SysUpdatedBy", sysUpdatedByKey, (*Account).GetSysUpdatedBy, (*Account).SetSysUpdatedBy},
		{"FiscalYear", fiscalYearKey, (*Account).GetFiscalYear, (*Account).SetFiscalYear},
		{"SysCreatedOn", sysCreatedOnKey, (*Account).GetSysCreatedOn, (*Account).SetSysCreatedOn},
		{"Contact", contactKey, (*Account).GetContact, (*Account).SetContact},
		{"StockPrice", stockPriceKey, (*Account).GetStockPrice, (*Account).SetStockPrice},
		{"State", stateKey, (*Account).GetState, (*Account).SetState},
		{"BannerImage", bannerImageKey, (*Account).GetBannerImage, (*Account).SetBannerImage},
		{"SysCreatedBy", sysCreatedByKey, (*Account).GetSysCreatedBy, (*Account).SetSysCreatedBy},
		{"Zip", zipKey, (*Account).GetZip, (*Account).SetZip},
		{"Phone", phoneKey, (*Account).GetPhone, (*Account).SetPhone},
		{"FaxPhone", faxPhoneKey, (*Account).GetFaxPhone, (*Account).SetFaxPhone},
		{"Name", nameKey, (*Account).GetName, (*Account).SetName},
		{"BannerText", bannerTextKey, (*Account).GetBannerText, (*Account).SetBannerText},
		{"AccountCode", accountCodeKey, (*Account).GetAccountCode, (*Account).SetAccountCode},
		{"City", cityKey, (*Account).GetCity, (*Account).SetCity},
		{"SysClassName", sysClassNameKey, (*Account).GetSysClassName, (*Account).SetSysClassName},
		{"AccountParent", accountParentKey, (*Account).GetAccountParent, (*Account).SetAccountParent},
		{"SysID", sysIDKey, (*Account).GetSysID, (*Account).SetSysID},
		{"Street", streetKey, (*Account).GetStreet, (*Account).SetStreet},
		{"LatLongError", latLongErrorKey, (*Account).GetLatLongError, (*Account).SetLatLongError},
		{"Theme", themeKey, (*Account).GetTheme, (*Account).SetTheme},
		{"VendorType", vendorTypeKey, (*Account).GetVendorType, (*Account).SetVendorType},
		{"Website", websiteKey, (*Account).GetWebsite, (*Account).SetWebsite},
		{"SysModCount", sysModCountKey, (*Account).GetSysModCount, (*Account).SetSysModCount},
		{"SysTags", sysTagsKey, (*Account).GetSysTags, (*Account).SetSysTags},
		{"RegistrationCode", registrationCodeKey, (*Account).GetRegistrationCode, (*Account).SetRegistrationCode},
		{"VendorManager", vendorManagerKey, (*Account).GetVendorManager, (*Account).SetVendorManager},
		{"AccountPath", accountPathKey, (*Account).GetAccountPath, (*Account).SetAccountPath},
		{"PrimaryContact", primaryContactKey, (*Account).GetPrimaryContact, (*Account).SetPrimaryContact},
	}
}

func int64Fields() []struct {
	name   string
	key    string
	getter func(*Account) (*int64, error)
	setter func(*Account, *int64) error
} {
	return []struct {
		name   string
		key    string
		getter func(*Account) (*int64, error)
		setter func(*Account, *int64) error
	}{
		{"Discount", discountKey, (*Account).GetDiscount, (*Account).SetDiscount},
		{"Profits", profitsKey, (*Account).GetProfits, (*Account).SetProfits},
		{"NumEmployees", numEmployeesKey, (*Account).GetNumEmployees, (*Account).SetNumEmployees},
		{"RevenuePerYear", revenuePerYearKey, (*Account).GetRevenuePerYear, (*Account).SetRevenuePerYear},
	}
}

func float64Fields() []struct {
	name   string
	key    string
	getter func(*Account) (*float64, error)
	setter func(*Account, *float64) error
} {
	return []struct {
		name   string
		key    string
		getter func(*Account) (*float64, error)
		setter func(*Account, *float64) error
	}{
		{"Longitude", longitudeKey, (*Account).GetLongitude, (*Account).SetLongitude},
		{"Latitude", latitudeKey, (*Account).GetLatitude, (*Account).SetLatitude},
		{"MarketCap", marketCapKey, (*Account).GetMarketCap, (*Account).SetMarketCap},
	}
}

func boolFields() []struct {
	name   string
	key    string
	getter func(*Account) (*bool, error)
	setter func(*Account, *bool) error
} {
	return []struct {
		name   string
		key    string
		getter func(*Account) (*bool, error)
		setter func(*Account, *bool) error
	}{
		{"Primary", primaryKey, (*Account).GetPrimary, (*Account).SetPrimary},
		{"Manufacturer", manufacturerKey, (*Account).GetManufacturer, (*Account).SetManufacturer},
		{"Vendor", vendorKey, (*Account).GetVendor, (*Account).SetVendor},
		{"PubliclyTraded", publiclyTradedKey, (*Account).GetPubliclyTraded, (*Account).SetPubliclyTraded},
		{"Partner", partnerKey, (*Account).GetPartner, (*Account).SetPartner},
		{"Customer", customerKey, (*Account).GetCustomer, (*Account).SetCustomer},
	}
}

// totalFields is the number of properties Account models; Serialize and
// GetFieldDeserializers must both cover exactly this many.
const totalFields = 52

func TestAccountFieldTablesAreComplete(t *testing.T) {
	// Guards the tables above against drift: if a property is added to account.go
	// without being added here, the per-field tables silently stop covering it.
	got := len(stringFields()) + len(int64Fields()) + len(float64Fields()) + len(boolFields()) + 1 // +1 for RankTier
	assert.Equal(t, totalFields, got)
}

func TestNewAccount(t *testing.T) {
	account := NewAccount()

	require.NotNil(t, account)
	assert.NotNil(t, account.GetBackingStore())
}

func TestCreateAccountFromDiscriminatorValue(t *testing.T) {
	tests := []struct {
		name string
		node *mocking.MockParseNode
	}{
		{"nil parse node", nil},
		{"non-nil parse node", mocking.NewMockParseNode()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var parsable, err = func() (any, error) {
				if tt.node == nil {
					return CreateAccountFromDiscriminatorValue(nil)
				}
				return CreateAccountFromDiscriminatorValue(tt.node)
			}()

			require.NoError(t, err)
			require.NotNil(t, parsable)
			assert.IsType(t, &Account{}, parsable)
		})
	}
}

func TestAccountStringProperties(t *testing.T) {
	for _, tt := range stringFields() {
		t.Run(tt.name, func(t *testing.T) {
			account := NewAccount()

			// unset property reads back as nil, not an error
			got, err := tt.getter(account)
			require.NoError(t, err)
			assert.Nil(t, got)

			// round trip through the backing store
			want := internal.ToPointer(tt.name + "-value")
			require.NoError(t, tt.setter(account, want))
			got, err = tt.getter(account)
			require.NoError(t, err)
			assert.Equal(t, want, got)

			// the value really landed under this property's key
			raw, err := account.GetBackingStore().Get(tt.key)
			require.NoError(t, err)
			assert.Equal(t, want, raw)

			// empty string is a distinct value from absent
			require.NoError(t, tt.setter(account, internal.ToPointer("")))
			got, err = tt.getter(account)
			require.NoError(t, err)
			require.NotNil(t, got)
			assert.Empty(t, *got)

			// explicit nil clears the property
			require.NoError(t, tt.setter(account, nil))
			got, err = tt.getter(account)
			require.NoError(t, err)
			assert.Nil(t, got)
		})
	}
}

func TestAccountInt64Properties(t *testing.T) {
	for _, tt := range int64Fields() {
		t.Run(tt.name, func(t *testing.T) {
			account := NewAccount()

			got, err := tt.getter(account)
			require.NoError(t, err)
			assert.Nil(t, got)

			for _, want := range []*int64{
				internal.ToPointer(int64(42)),
				internal.ToPointer(int64(0)),
				internal.ToPointer(int64(-1)),
				internal.ToPointer(int64(9223372036854775807)),
			} {
				require.NoError(t, tt.setter(account, want))
				got, err = tt.getter(account)
				require.NoError(t, err)
				assert.Equal(t, want, got)
			}

			raw, err := account.GetBackingStore().Get(tt.key)
			require.NoError(t, err)
			assert.Equal(t, internal.ToPointer(int64(9223372036854775807)), raw)

			require.NoError(t, tt.setter(account, nil))
			got, err = tt.getter(account)
			require.NoError(t, err)
			assert.Nil(t, got)
		})
	}
}

func TestAccountFloat64Properties(t *testing.T) {
	for _, tt := range float64Fields() {
		t.Run(tt.name, func(t *testing.T) {
			account := NewAccount()

			got, err := tt.getter(account)
			require.NoError(t, err)
			assert.Nil(t, got)

			for _, want := range []*float64{
				internal.ToPointer(41.8781),
				internal.ToPointer(0.0),
				internal.ToPointer(-87.6298),
			} {
				require.NoError(t, tt.setter(account, want))
				got, err = tt.getter(account)
				require.NoError(t, err)
				assert.Equal(t, want, got)
			}

			raw, err := account.GetBackingStore().Get(tt.key)
			require.NoError(t, err)
			assert.Equal(t, internal.ToPointer(-87.6298), raw)

			require.NoError(t, tt.setter(account, nil))
			got, err = tt.getter(account)
			require.NoError(t, err)
			assert.Nil(t, got)
		})
	}
}

func TestAccountBoolProperties(t *testing.T) {
	for _, tt := range boolFields() {
		t.Run(tt.name, func(t *testing.T) {
			account := NewAccount()

			got, err := tt.getter(account)
			require.NoError(t, err)
			assert.Nil(t, got)

			// false must be distinguishable from absent
			for _, want := range []*bool{internal.ToPointer(true), internal.ToPointer(false)} {
				require.NoError(t, tt.setter(account, want))
				got, err = tt.getter(account)
				require.NoError(t, err)
				require.NotNil(t, got)
				assert.Equal(t, *want, *got)
			}

			raw, err := account.GetBackingStore().Get(tt.key)
			require.NoError(t, err)
			assert.Equal(t, internal.ToPointer(false), raw)

			require.NoError(t, tt.setter(account, nil))
			got, err = tt.getter(account)
			require.NoError(t, err)
			assert.Nil(t, got)
		})
	}
}

func TestAccountRankTierProperty(t *testing.T) {
	tests := []struct {
		name string
		want *RankTier
	}{
		{"nil", nil},
		{"unknown", internal.ToPointer(RankTierUnknown)},
		{"blacklist", internal.ToPointer(RankTierBlacklist)},
		{"other", internal.ToPointer(RankTierOther)},
		{"strategic", internal.ToPointer(RankTierStrategic)},
		{"tactical", internal.ToPointer(RankTierTactical)},
		{"valued", internal.ToPointer(RankTierValued)},
		{"out of range value round trips unchanged", internal.ToPointer(RankTier(999))},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			account := NewAccount()

			require.NoError(t, account.SetRankTier(tt.want))
			got, err := account.GetRankTier()
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAccountRankTierStoredUnderRankTierKey(t *testing.T) {
	account := NewAccount()
	require.NoError(t, account.SetRankTier(internal.ToPointer(RankTierStrategic)))

	raw, err := account.GetBackingStore().Get(rankTierKey)
	require.NoError(t, err)
	assert.Equal(t, internal.ToPointer(RankTierStrategic), raw)
}

func TestAccountAccessorsNilReceiver(t *testing.T) {
	var account *Account

	t.Run("string getter and setter", func(t *testing.T) {
		got, err := account.GetName()
		require.ErrorIs(t, err, snerrors.ErrNilModel)
		assert.Nil(t, got)
		assert.ErrorIs(t, account.SetName(internal.ToPointer("x")), snerrors.ErrNilModel)
	})

	t.Run("int64 getter and setter", func(t *testing.T) {
		got, err := account.GetDiscount()
		require.ErrorIs(t, err, snerrors.ErrNilModel)
		assert.Nil(t, got)
		assert.ErrorIs(t, account.SetDiscount(internal.ToPointer(int64(1))), snerrors.ErrNilModel)
	})

	t.Run("float64 getter and setter", func(t *testing.T) {
		got, err := account.GetLatitude()
		require.ErrorIs(t, err, snerrors.ErrNilModel)
		assert.Nil(t, got)
		assert.ErrorIs(t, account.SetLatitude(internal.ToPointer(1.0)), snerrors.ErrNilModel)
	})

	t.Run("bool getter and setter", func(t *testing.T) {
		got, err := account.GetCustomer()
		require.ErrorIs(t, err, snerrors.ErrNilModel)
		assert.Nil(t, got)
		assert.ErrorIs(t, account.SetCustomer(internal.ToPointer(true)), snerrors.ErrNilModel)
	})

	t.Run("enum getter and setter", func(t *testing.T) {
		got, err := account.GetRankTier()
		require.ErrorIs(t, err, snerrors.ErrNilModel)
		assert.Nil(t, got)
		assert.ErrorIs(t, account.SetRankTier(internal.ToPointer(RankTierValued)), snerrors.ErrNilModel)
	})
}

func TestAccountGetterStoreTypeMismatch(t *testing.T) {
	// The backing store is untyped, so a value written under a property's key by
	// something other than the typed mutator surfaces as a conversion error.
	tests := []struct {
		name   string
		key    string
		stored any
		read   func(*Account) error
	}{
		{"string property holding an int", nameKey, 42, func(a *Account) error { _, err := a.GetName(); return err }},
		{"int64 property holding a string", discountKey, "nope", func(a *Account) error { _, err := a.GetDiscount(); return err }},
		{"float64 property holding a string", latitudeKey, "nope", func(a *Account) error { _, err := a.GetLatitude(); return err }},
		{"bool property holding a string", customerKey, "nope", func(a *Account) error { _, err := a.GetCustomer(); return err }},
		{"enum property holding a string", rankTierKey, "strategic", func(a *Account) error { _, err := a.GetRankTier(); return err }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			account := NewAccount()
			require.NoError(t, account.GetBackingStore().Set(tt.key, tt.stored))

			assert.Error(t, tt.read(account))
		})
	}
}

// populatedAccount returns an Account with one value set for every property.
func populatedAccount(t *testing.T) *Account {
	t.Helper()

	account := NewAccount()
	for _, f := range stringFields() {
		require.NoError(t, f.setter(account, internal.ToPointer(f.name+"-value")))
	}
	for _, f := range int64Fields() {
		require.NoError(t, f.setter(account, internal.ToPointer(int64(7))))
	}
	for _, f := range float64Fields() {
		require.NoError(t, f.setter(account, internal.ToPointer(1.5)))
	}
	for _, f := range boolFields() {
		require.NoError(t, f.setter(account, internal.ToPointer(true)))
	}
	require.NoError(t, account.SetRankTier(internal.ToPointer(RankTierStrategic)))

	return account
}

func TestAccountSerialize(t *testing.T) {
	tests := []struct {
		name      string
		account   func(*testing.T) *Account
		setupMock func(*mocking.MockSerializationWriter)
		assertOn  func(*testing.T, *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:    "nil model serializes to nothing without error",
			account: func(*testing.T) *Account { return nil },
			assertOn: func(t *testing.T, w *mocking.MockSerializationWriter) {
				w.AssertNotCalled(t, "WriteStringValue", mock.Anything, mock.Anything)
			},
		},
		{
			name:    "empty model writes no properties",
			account: func(*testing.T) *Account { return NewAccount() },
			assertOn: func(t *testing.T, w *mocking.MockSerializationWriter) {
				w.AssertNotCalled(t, "WriteStringValue", mock.Anything, mock.Anything)
				w.AssertNotCalled(t, "WriteInt64Value", mock.Anything, mock.Anything)
				w.AssertNotCalled(t, "WriteFloat64Value", mock.Anything, mock.Anything)
				w.AssertNotCalled(t, "WriteBoolValue", mock.Anything, mock.Anything)
			},
		},
		{
			name:    "fully populated model writes every property",
			account: populatedAccount,
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
				w.On("WriteInt64Value", mock.Anything, mock.Anything).Return(nil)
				w.On("WriteFloat64Value", mock.Anything, mock.Anything).Return(nil)
				w.On("WriteBoolValue", mock.Anything, mock.Anything).Return(nil)
			},
			assertOn: func(t *testing.T, w *mocking.MockSerializationWriter) {
				for _, f := range stringFields() {
					w.AssertCalled(t, "WriteStringValue", f.key, internal.ToPointer(f.name+"-value"))
				}
				for _, f := range int64Fields() {
					w.AssertCalled(t, "WriteInt64Value", f.key, internal.ToPointer(int64(7)))
				}
				for _, f := range float64Fields() {
					w.AssertCalled(t, "WriteFloat64Value", f.key, internal.ToPointer(1.5))
				}
				for _, f := range boolFields() {
					w.AssertCalled(t, "WriteBoolValue", f.key, internal.ToPointer(true))
				}
				// the enum is written as its String() form
				w.AssertCalled(t, "WriteStringValue", rankTierKey, internal.ToPointer("strategic"))
			},
		},
		{
			name: "out of range enum is written as its unknown string form",
			account: func(t *testing.T) *Account {
				account := NewAccount()
				require.NoError(t, account.SetRankTier(internal.ToPointer(RankTier(999))))
				return account
			},
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			},
			assertOn: func(t *testing.T, w *mocking.MockSerializationWriter) {
				w.AssertCalled(t, "WriteStringValue", rankTierKey, internal.ToPointer("unknown"))
			},
		},
		{
			name: "string write error propagates",
			account: func(t *testing.T) *Account {
				account := NewAccount()
				require.NoError(t, account.SetName(internal.ToPointer("acme")))
				return account
			},
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", nameKey, mock.Anything).Return(errWrite)
			},
			wantErr: errWrite,
		},
		{
			name: "int64 write error propagates",
			account: func(t *testing.T) *Account {
				account := NewAccount()
				require.NoError(t, account.SetDiscount(internal.ToPointer(int64(5))))
				return account
			},
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteInt64Value", discountKey, mock.Anything).Return(errWrite)
			},
			wantErr: errWrite,
		},
		{
			name: "float64 write error propagates",
			account: func(t *testing.T) *Account {
				account := NewAccount()
				require.NoError(t, account.SetLatitude(internal.ToPointer(1.5)))
				return account
			},
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteFloat64Value", latitudeKey, mock.Anything).Return(errWrite)
			},
			wantErr: errWrite,
		},
		{
			name: "bool write error propagates",
			account: func(t *testing.T) *Account {
				account := NewAccount()
				require.NoError(t, account.SetCustomer(internal.ToPointer(true)))
				return account
			},
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteBoolValue", customerKey, mock.Anything).Return(errWrite)
			},
			wantErr: errWrite,
		},
		{
			name: "enum write error propagates",
			account: func(t *testing.T) *Account {
				account := NewAccount()
				require.NoError(t, account.SetRankTier(internal.ToPointer(RankTierValued)))
				return account
			},
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", rankTierKey, mock.Anything).Return(errWrite)
			},
			wantErr: errWrite,
		},
		{
			name: "accessor error propagates",
			account: func(t *testing.T) *Account {
				account := NewAccount()
				// bannerImageLight is the first property Serialize touches, so the
				// accessor conversion failure is what surfaces.
				require.NoError(t, account.GetBackingStore().Set(bannerImageLightKey, 42))
				return account
			},
			wantErr: nil, // asserted via assert.Error below
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			if tt.setupMock != nil {
				tt.setupMock(writer)
			}

			err := tt.account(t).Serialize(writer)

			switch {
			case tt.name == "accessor error propagates":
				require.Error(t, err)
			case tt.wantErr != nil:
				require.ErrorIs(t, err, tt.wantErr)
			default:
				require.NoError(t, err)
			}

			if tt.assertOn != nil {
				tt.assertOn(t, writer)
			}
		})
	}
}

// TestAccountSerializeNilWriter verifies that Serialize(nil) returns
// snerrors.ErrNilWriter consistently, regardless of whether the model is empty
// or populated.
func TestAccountSerializeNilWriter(t *testing.T) {
	t.Run("empty model", func(t *testing.T) {
		err := NewAccount().Serialize(nil)
		require.ErrorIs(t, err, snerrors.ErrNilWriter)
	})

	t.Run("populated model", func(t *testing.T) {
		account := NewAccount()
		require.NoError(t, account.SetName(internal.ToPointer("acme")))

		err := account.Serialize(nil)
		require.ErrorIs(t, err, snerrors.ErrNilWriter)
	})
}

func TestAccountGetFieldDeserializersCoversEveryProperty(t *testing.T) {
	deserializers := NewAccount().GetFieldDeserializers()

	require.Len(t, deserializers, totalFields)

	keys := []string{rankTierKey}
	for _, f := range stringFields() {
		keys = append(keys, f.key)
	}
	for _, f := range int64Fields() {
		keys = append(keys, f.key)
	}
	for _, f := range float64Fields() {
		keys = append(keys, f.key)
	}
	for _, f := range boolFields() {
		keys = append(keys, f.key)
	}

	for _, key := range keys {
		assert.NotNil(t, deserializers[key], "missing deserializer for %s", key)
	}
}

func TestAccountGetFieldDeserializersStringProperties(t *testing.T) {
	for _, tt := range stringFields() {
		t.Run(tt.name, func(t *testing.T) {
			t.Run("deserializes into the matching property", func(t *testing.T) {
				account := NewAccount()
				node := mocking.NewMockParseNode()
				node.On("GetStringValue").Return(internal.ToPointer("deserialized"), nil)

				require.NoError(t, account.GetFieldDeserializers()[tt.key](node))

				got, err := tt.getter(account)
				require.NoError(t, err)
				assert.Equal(t, internal.ToPointer("deserialized"), got)
			})

			t.Run("parse error propagates", func(t *testing.T) {
				account := NewAccount()
				node := mocking.NewMockParseNode()
				node.On("GetStringValue").Return((*string)(nil), errParse)

				assert.ErrorIs(t, account.GetFieldDeserializers()[tt.key](node), errParse)
			})
		})
	}
}

func TestAccountGetFieldDeserializersInt64Properties(t *testing.T) {
	for _, tt := range int64Fields() {
		t.Run(tt.name, func(t *testing.T) {
			account := NewAccount()
			node := mocking.NewMockParseNode()
			node.On("GetInt64Value").Return(internal.ToPointer(int64(99)), nil)

			require.NoError(t, account.GetFieldDeserializers()[tt.key](node))

			got, err := tt.getter(account)
			require.NoError(t, err)
			assert.Equal(t, internal.ToPointer(int64(99)), got)

			errNode := mocking.NewMockParseNode()
			errNode.On("GetInt64Value").Return((*int64)(nil), errParse)
			assert.ErrorIs(t, account.GetFieldDeserializers()[tt.key](errNode), errParse)
		})
	}
}

func TestAccountGetFieldDeserializersFloat64Properties(t *testing.T) {
	for _, tt := range float64Fields() {
		t.Run(tt.name, func(t *testing.T) {
			account := NewAccount()
			node := mocking.NewMockParseNode()
			node.On("GetFloat64Value").Return(internal.ToPointer(2.5), nil)

			require.NoError(t, account.GetFieldDeserializers()[tt.key](node))

			got, err := tt.getter(account)
			require.NoError(t, err)
			assert.Equal(t, internal.ToPointer(2.5), got)

			errNode := mocking.NewMockParseNode()
			errNode.On("GetFloat64Value").Return((*float64)(nil), errParse)
			assert.ErrorIs(t, account.GetFieldDeserializers()[tt.key](errNode), errParse)
		})
	}
}

func TestAccountGetFieldDeserializersBoolProperties(t *testing.T) {
	for _, tt := range boolFields() {
		t.Run(tt.name, func(t *testing.T) {
			account := NewAccount()
			node := mocking.NewMockParseNode()
			node.On("GetBoolValue").Return(internal.ToPointer(true), nil)

			require.NoError(t, account.GetFieldDeserializers()[tt.key](node))

			got, err := tt.getter(account)
			require.NoError(t, err)
			assert.Equal(t, internal.ToPointer(true), got)

			errNode := mocking.NewMockParseNode()
			errNode.On("GetBoolValue").Return((*bool)(nil), errParse)
			assert.ErrorIs(t, account.GetFieldDeserializers()[tt.key](errNode), errParse)
		})
	}
}

func TestAccountGetFieldDeserializersRankTier(t *testing.T) {
	tests := []struct {
		name      string
		enumValue any
		enumErr   error
		want      *RankTier
		wantErr   error
		expectErr bool
	}{
		{
			name:      "value enum from factory",
			enumValue: RankTierStrategic,
			want:      internal.ToPointer(RankTierStrategic),
		},
		{
			name:      "pointer enum from factory",
			enumValue: internal.ToPointer(RankTierTactical),
			want:      internal.ToPointer(RankTierTactical),
		},
		{
			name:      "nil from factory clears the property",
			enumValue: nil,
			want:      nil,
		},
		{
			name:      "parse error propagates",
			enumValue: nil,
			enumErr:   errParse,
			wantErr:   errParse,
		},
		{
			name:      "unexpected type from factory is rejected",
			enumValue: "strategic",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			account := NewAccount()
			node := mocking.NewMockParseNode()
			node.On("GetEnumValue").Return(tt.enumValue, tt.enumErr)

			err := account.GetFieldDeserializers()[rankTierKey](node)

			switch {
			case tt.wantErr != nil:
				require.ErrorIs(t, err, tt.wantErr)
				return
			case tt.expectErr:
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			got, err := account.GetRankTier()
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAccountGetFieldDeserializersNilReceiver(t *testing.T) {
	var account *Account

	deserializers := account.GetFieldDeserializers()
	require.Len(t, deserializers, totalFields)

	node := mocking.NewMockParseNode()
	node.On("GetStringValue").Return(internal.ToPointer("x"), nil)

	// The map is built from method values on a nil receiver, so invoking one
	// surfaces the nil-model sentinel rather than panicking.
	assert.ErrorIs(t, deserializers[nameKey](node), snerrors.ErrNilModel)
}

// TestAccountZeroValueReceiverNoPanic verifies that an Account built as a struct
// literal (rather than via NewAccount) no longer panics. With *core.BaseModel
// embedded (instead of the core.BackedModel interface), the zero-value has a nil
// *BaseModel pointer that conversion.IsNil catches, returning snerrors.ErrNilModel.
func TestAccountZeroValueReceiverNoPanic(t *testing.T) {
	t.Run("getter", func(t *testing.T) {
		_, err := (&Account{}).GetName()
		assert.ErrorIs(t, err, snerrors.ErrNilStore)
	})

	t.Run("setter", func(t *testing.T) {
		err := (&Account{}).SetName(internal.ToPointer("x"))
		assert.ErrorIs(t, err, snerrors.ErrNilStore)
	})

	t.Run("serialize", func(t *testing.T) {
		err := (&Account{}).Serialize(mocking.NewMockSerializationWriter())
		assert.ErrorIs(t, err, snerrors.ErrNilStore)
	})
}
