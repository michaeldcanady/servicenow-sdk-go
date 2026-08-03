package accountapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestAccountSerialization(t *testing.T) {
	tests := []struct {
		name string
		val  string
	}{
		{"ValidName", "Boxeo EMEA"},
		{"EmptyName", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			account := NewAccount()
			name := tt.val
			err := account.SetName(&name)
			require.NoError(t, err)

			retrievedName, err := account.GetName()
			require.NoError(t, err)
			assert.Equal(t, &name, retrievedName)
		})
	}
}

func TestAccountGettersSetters(t *testing.T) {
	account := NewAccount()
	val := "test-value"

	tests := []struct {
		name   string
		Setter func(*string) error
		getter func() (*string, error)
	}{
		{"BannerImageLight", account.SetBannerImageLight, account.GetBannerImageLight},
		{"Country", account.SetCountry, account.GetCountry},
		{"Parent", account.SetParent, account.GetParent},
		{"Notes", account.SetNotes, account.GetNotes},
		{"StockSymbol", account.SetStockSymbol, account.GetStockSymbol},
		{"Discount", account.SetDiscount, account.GetDiscount},
		{"ActiveEscalation", account.SetActiveEscalation, account.GetActiveEscalation},
		{"SysUpdatedOn", account.SetSysUpdatedOn, account.GetSysUpdatedOn},
		{"AppleIcon", account.SetAppleIcon, account.GetAppleIcon},
		{"Number", account.SetNumber, account.GetNumber},
		{"SysUpdatedBy", account.SetSysUpdatedBy, account.GetSysUpdatedBy},
		{"FiscalYear", account.SetFiscalYear, account.GetFiscalYear},
		{"SysCreatedOn", account.SetSysCreatedOn, account.GetSysCreatedOn},
		{"Contact", account.SetContact, account.GetContact},
		{"StockPrice", account.SetStockPrice, account.GetStockPrice},
		{"State", account.SetState, account.GetState},
		{"BannerImage", account.SetBannerImage, account.GetBannerImage},
		{"SysCreatedBy", account.SetSysCreatedBy, account.GetSysCreatedBy},
		{"Longitude", account.SetLongitude, account.GetLongitude},
		{"Zip", account.SetZip, account.GetZip},
		{"Profits", account.SetProfits, account.GetProfits},
		{"Phone", account.SetPhone, account.GetPhone},
		{"FaxPhone", account.SetFaxPhone, account.GetFaxPhone},
		{"Name", account.SetName, account.GetName},
		{"BannerText", account.SetBannerText, account.GetBannerText},
		{"AccountCode", account.SetAccountCode, account.GetAccountCode},
		{"Primary", account.SetPrimary, account.GetPrimary},
		{"City", account.SetCity, account.GetCity},
		{"Latitude", account.SetLatitude, account.GetLatitude},
		{"SysClassName", account.SetSysClassName, account.GetSysClassName},
		{"Manufacturer", account.SetManufacturer, account.GetManufacturer},
		{"AccountParent", account.SetAccountParent, account.GetAccountParent},
		{"SysID", account.SetSysID, account.GetSysID},
		{"MarketCap", account.SetMarketCap, account.GetMarketCap},
		{"NumEmployees", account.SetNumEmployees, account.GetNumEmployees},
		{"RankTier", account.SetRankTier, account.GetRankTier},
		{"Street", account.SetStreet, account.GetStreet},
		{"Vendor", account.SetVendor, account.GetVendor},
		{"LatLongError", account.SetLatLongError, account.GetLatLongError},
		{"Theme", account.SetTheme, account.GetTheme},
		{"VendorType", account.SetVendorType, account.GetVendorType},
		{"Website", account.SetWebsite, account.GetWebsite},
		{"RevenuePerYear", account.SetRevenuePerYear, account.GetRevenuePerYear},
		{"PubliclyTraded", account.SetPubliclyTraded, account.GetPubliclyTraded},
		{"SysModCount", account.SetSysModCount, account.GetSysModCount},
		{"SysTags", account.SetSysTags, account.GetSysTags},
		{"Partner", account.SetPartner, account.GetPartner},
		{"RegistrationCode", account.SetRegistrationCode, account.GetRegistrationCode},
		{"VendorManager", account.SetVendorManager, account.GetVendorManager},
		{"AccountPath", account.SetAccountPath, account.GetAccountPath},
		{"PrimaryContact", account.SetPrimaryContact, account.GetPrimaryContact},
		{"Customer", account.SetCustomer, account.GetCustomer},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.Setter(&val)
			require.NoError(t, err)
			got, err := tt.getter()
			require.NoError(t, err)
			assert.Equal(t, &val, got)
		})
	}
}

func TestAccountModel_GetFieldDeserializers(t *testing.T) {
	t.Run("ValidDeserializers", func(t *testing.T) {
		account := NewAccount()
		deserializers := account.GetFieldDeserializers()
		assert.NotNil(t, deserializers)
		assert.NotEmpty(t, deserializers)
	})
}

func TestAccountModel_Serialize(t *testing.T) {
	t.Run("SuccessfulSerialization", func(t *testing.T) {
		account := NewAccount()
		name := "test-account"
		_ = account.SetName(&name)

		writer := &mocking.MockSerializationWriter{}
		writer.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
		writer.On("WriteObjectValue", mock.Anything, mock.Anything, mock.Anything).Return(nil)
		writer.On("WriteAdditionalData", mock.Anything).Return(nil)

		err := account.Serialize(writer)
		assert.NoError(t, err)
	})
}

func TestCreateAccountFromDiscriminatorValue(t *testing.T) {
	t.Run("SuccessfulCreation", func(t *testing.T) {
		instance, err := CreateAccountFromDiscriminatorValue(nil)
		require.NoError(t, err)
		assert.NotNil(t, instance)
		assert.IsType(t, &Account{}, instance)
	})
}
