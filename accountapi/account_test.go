package accountapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
			err := account.setName(&name)
			assert.NoError(t, err)

			retrievedName, err := account.GetName()
			assert.NoError(t, err)
			assert.Equal(t, &name, retrievedName)
		})
	}
}

func TestAccountGettersSetters(t *testing.T) {
	account := NewAccount()
	val := "test-value"

	tests := []struct {
		name   string
		setter func(*string) error
		getter func() (*string, error)
	}{
		{"BannerImageLight", account.setBannerImageLight, account.GetBannerImageLight},
		{"Country", account.setCountry, account.GetCountry},
		{"Parent", account.setParent, account.GetParent},
		{"Notes", account.setNotes, account.GetNotes},
		{"StockSymbol", account.setStockSymbol, account.GetStockSymbol},
		{"Discount", account.setDiscount, account.GetDiscount},
		{"ActiveEscalation", account.setActiveEscalation, account.GetActiveEscalation},
		{"SysUpdatedOn", account.setSysUpdatedOn, account.GetSysUpdatedOn},
		{"AppleIcon", account.setAppleIcon, account.GetAppleIcon},
		{"Number", account.setNumber, account.GetNumber},
		{"SysUpdatedBy", account.setSysUpdatedBy, account.GetSysUpdatedBy},
		{"FiscalYear", account.setFiscalYear, account.GetFiscalYear},
		{"SysCreatedOn", account.setSysCreatedOn, account.GetSysCreatedOn},
		{"Contact", account.setContact, account.GetContact},
		{"StockPrice", account.setStockPrice, account.GetStockPrice},
		{"State", account.setState, account.GetState},
		{"BannerImage", account.setBannerImage, account.GetBannerImage},
		{"SysCreatedBy", account.setSysCreatedBy, account.GetSysCreatedBy},
		{"Longitude", account.setLongitude, account.GetLongitude},
		{"Zip", account.setZip, account.GetZip},
		{"Profits", account.setProfits, account.GetProfits},
		{"Phone", account.setPhone, account.GetPhone},
		{"FaxPhone", account.setFaxPhone, account.GetFaxPhone},
		{"Name", account.setName, account.GetName},
		{"BannerText", account.setBannerText, account.GetBannerText},
		{"AccountCode", account.setAccountCode, account.GetAccountCode},
		{"Primary", account.setPrimary, account.GetPrimary},
		{"City", account.setCity, account.GetCity},
		{"Latitude", account.setLatitude, account.GetLatitude},
		{"SysClassName", account.setSysClassName, account.GetSysClassName},
		{"Manufacturer", account.setManufacturer, account.GetManufacturer},
		{"AccountParent", account.setAccountParent, account.GetAccountParent},
		{"SysID", account.setSysID, account.GetSysID},
		{"MarketCap", account.setMarketCap, account.GetMarketCap},
		{"NumEmployees", account.setNumEmployees, account.GetNumEmployees},
		{"RankTier", account.setRankTier, account.GetRankTier},
		{"Street", account.setStreet, account.GetStreet},
		{"Vendor", account.setVendor, account.GetVendor},
		{"LatLongError", account.setLatLongError, account.GetLatLongError},
		{"Theme", account.setTheme, account.GetTheme},
		{"VendorType", account.setVendorType, account.GetVendorType},
		{"Website", account.setWebsite, account.GetWebsite},
		{"RevenuePerYear", account.setRevenuePerYear, account.GetRevenuePerYear},
		{"PubliclyTraded", account.setPubliclyTraded, account.GetPubliclyTraded},
		{"SysModCount", account.setSysModCount, account.GetSysModCount},
		{"SysTags", account.setSysTags, account.GetSysTags},
		{"Partner", account.setPartner, account.GetPartner},
		{"RegistrationCode", account.setRegistrationCode, account.GetRegistrationCode},
		{"VendorManager", account.setVendorManager, account.GetVendorManager},
		{"AccountPath", account.setAccountPath, account.GetAccountPath},
		{"PrimaryContact", account.setPrimaryContact, account.GetPrimaryContact},
		{"Customer", account.setCustomer, account.GetCustomer},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setter(&val)
			assert.NoError(t, err)
			got, err := tt.getter()
			assert.NoError(t, err)
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
		_ = account.setName(&name)

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
		assert.NoError(t, err)
		assert.NotNil(t, instance)
		assert.IsType(t, &AccountModel{}, instance)
	})
}
