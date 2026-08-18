package accountapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"

	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// TODO: figure out dates

const (
	bannerImageLightKey = "banner_image_light"
	countryKey          = "country"
	parentKey           = "parent"
	notesKey            = "notes"
	stockSymbolKey      = "stock_symbol"
	discountKey         = "discount"
	activeEscalationKey = "active_escalation"
	sysUpdatedOnKey     = "sys_updated_on"
	appleIconKey        = "apple_icon"
	numberKey           = "number"
	sysUpdatedByKey     = "sys_updated_by"
	fiscalYearKey       = "fiscal_year"
	sysCreatedOnKey     = "sys_created_on"
	contactKey          = "contact"
	stockPriceKey       = "stock_price"
	stateKey            = "state"
	bannerImageKey      = "banner_image"
	sysCreatedByKey     = "sys_created_by"
	longitudeKey        = "longitude"
	zipKey              = "zip"
	profitsKey          = "profits"
	phoneKey            = "phone"
	faxPhoneKey         = "fax_phone"
	nameKey             = "name"
	bannerTextKey       = "banner_text"
	accountCodeKey      = "account_code"
	primaryKey          = "primary"
	cityKey             = "city"
	latitudeKey         = "latitude"
	sysClassNameKey     = "sys_class_name"
	manufacturerKey     = "manufacturer"
	accountParentKey    = "account_parent"
	sysIDKey            = "sys_id"
	marketCapKey        = "market_cap"
	numEmployeesKey     = "num_employees"
	rankTierKey         = "rank_tier"
	streetKey           = "street"
	vendorKey           = "vendor"
	latLongErrorKey     = "lat_long_error"
	themeKey            = "theme"
	vendorTypeKey       = "vendor_type"
	websiteKey          = "website"
	revenuePerYearKey   = "revenue_per_year"
	publiclyTradedKey   = "publicly_traded"
	sysModCountKey      = "sys_mod_count"
	sysTagsKey          = "sys_tags"
	partnerKey          = "partner"
	registrationCodeKey = "registration_code"
	vendorManagerKey    = "vendor_manager"
	accountPathKey      = "account_path"
	primaryContactKey   = "primary_contact"
	customerKey         = "customer"
)

// Account represents an account object in ServiceNow.
type Account struct {
	*core.BaseModel
}

// NewAccount creates a new instance of AccountModel
func NewAccount() *Account {
	return &Account{
		BaseModel: core.NewBaseModel(),
	}
}

// CreateAccountFromDiscriminatorValue is a factory for creating an Account model.
func CreateAccountFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewAccount(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *Account) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}

	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(bannerImageLightKey, m.GetBannerImageLight),
		internalSerialization.SerializeStringFunc(countryKey, m.GetCountry),
		internalSerialization.SerializeStringFunc(parentKey, m.GetParent),
		internalSerialization.SerializeStringFunc(notesKey, m.GetNotes),
		internalSerialization.SerializeStringFunc(stockSymbolKey, m.GetStockSymbol),
		internalSerialization.SerializeInt64Func(discountKey, m.GetDiscount),
		internalSerialization.SerializeStringFunc(activeEscalationKey, m.GetActiveEscalation),
		internalSerialization.SerializeStringFunc(sysUpdatedOnKey, m.GetSysUpdatedOn),
		internalSerialization.SerializeStringFunc(appleIconKey, m.GetAppleIcon),
		internalSerialization.SerializeStringFunc(numberKey, m.GetNumber),
		internalSerialization.SerializeStringFunc(sysUpdatedByKey, m.GetSysUpdatedBy),
		internalSerialization.SerializeStringFunc(fiscalYearKey, m.GetFiscalYear),
		internalSerialization.SerializeStringFunc(sysCreatedOnKey, m.GetSysCreatedOn),
		internalSerialization.SerializeStringFunc(contactKey, m.GetContact),
		internalSerialization.SerializeStringFunc(stockPriceKey, m.GetStockPrice),
		internalSerialization.SerializeStringFunc(stateKey, m.GetState),
		internalSerialization.SerializeStringFunc(bannerImageKey, m.GetBannerImage),
		internalSerialization.SerializeStringFunc(sysCreatedByKey, m.GetSysCreatedBy),
		internalSerialization.SerializeFloat64Func(longitudeKey, m.GetLongitude),
		internalSerialization.SerializeStringFunc(zipKey, m.GetZip),
		internalSerialization.SerializeInt64Func(profitsKey, m.GetProfits),
		internalSerialization.SerializeStringFunc(phoneKey, m.GetPhone),
		internalSerialization.SerializeStringFunc(faxPhoneKey, m.GetFaxPhone),
		internalSerialization.SerializeStringFunc(nameKey, m.GetName),
		internalSerialization.SerializeStringFunc(bannerTextKey, m.GetBannerText),
		internalSerialization.SerializeStringFunc(accountCodeKey, m.GetAccountCode),
		internalSerialization.SerializeBoolFunc(primaryKey, m.GetPrimary),
		internalSerialization.SerializeStringFunc(cityKey, m.GetCity),
		internalSerialization.SerializeFloat64Func(latitudeKey, m.GetLatitude),
		internalSerialization.SerializeStringFunc(sysClassNameKey, m.GetSysClassName),
		internalSerialization.SerializeBoolFunc(manufacturerKey, m.GetManufacturer),
		internalSerialization.SerializeStringFunc(accountParentKey, m.GetAccountParent),
		internalSerialization.SerializeStringFunc(sysIDKey, m.GetSysID),
		internalSerialization.SerializeFloat64Func(marketCapKey, m.GetMarketCap),
		internalSerialization.SerializeInt64Func(numEmployeesKey, m.GetNumEmployees),
		internalSerialization.SerializeEnumFunc(rankTierKey, m.GetRankTier),
		internalSerialization.SerializeStringFunc(streetKey, m.GetStreet),
		internalSerialization.SerializeBoolFunc(vendorKey, m.GetVendor),
		internalSerialization.SerializeStringFunc(latLongErrorKey, m.GetLatLongError),
		internalSerialization.SerializeStringFunc(themeKey, m.GetTheme),
		internalSerialization.SerializeStringFunc(vendorTypeKey, m.GetVendorType),
		internalSerialization.SerializeStringFunc(websiteKey, m.GetWebsite),
		internalSerialization.SerializeInt64Func(revenuePerYearKey, m.GetRevenuePerYear),
		internalSerialization.SerializeBoolFunc(publiclyTradedKey, m.GetPubliclyTraded),
		internalSerialization.SerializeStringFunc(sysModCountKey, m.GetSysModCount),
		internalSerialization.SerializeStringFunc(sysTagsKey, m.GetSysTags),
		internalSerialization.SerializeBoolFunc(partnerKey, m.GetPartner),
		internalSerialization.SerializeStringFunc(registrationCodeKey, m.GetRegistrationCode),
		internalSerialization.SerializeStringFunc(vendorManagerKey, m.GetVendorManager),
		internalSerialization.SerializeStringFunc(accountPathKey, m.GetAccountPath),
		internalSerialization.SerializeStringFunc(primaryContactKey, m.GetPrimaryContact),
		internalSerialization.SerializeBoolFunc(customerKey, m.GetCustomer),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *Account) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		bannerImageLightKey: internalSerialization.DeserializeStringFunc(m.SetBannerImageLight),
		countryKey:          internalSerialization.DeserializeStringFunc(m.SetCountry),
		parentKey:           internalSerialization.DeserializeStringFunc(m.SetParent),
		notesKey:            internalSerialization.DeserializeStringFunc(m.SetNotes),
		stockSymbolKey:      internalSerialization.DeserializeStringFunc(m.SetStockSymbol),
		discountKey:         internalSerialization.DeserializeInt64Func(m.SetDiscount),
		activeEscalationKey: internalSerialization.DeserializeStringFunc(m.SetActiveEscalation),
		sysUpdatedOnKey:     internalSerialization.DeserializeStringFunc(m.SetSysUpdatedOn),
		appleIconKey:        internalSerialization.DeserializeStringFunc(m.SetAppleIcon),
		numberKey:           internalSerialization.DeserializeStringFunc(m.SetNumber),
		sysUpdatedByKey:     internalSerialization.DeserializeStringFunc(m.SetSysUpdatedBy),
		fiscalYearKey:       internalSerialization.DeserializeStringFunc(m.SetFiscalYear),
		sysCreatedOnKey:     internalSerialization.DeserializeStringFunc(m.SetSysCreatedOn),
		contactKey:          internalSerialization.DeserializeStringFunc(m.SetContact),
		stockPriceKey:       internalSerialization.DeserializeStringFunc(m.SetStockPrice),
		stateKey:            internalSerialization.DeserializeStringFunc(m.SetState),
		bannerImageKey:      internalSerialization.DeserializeStringFunc(m.SetBannerImage),
		sysCreatedByKey:     internalSerialization.DeserializeStringFunc(m.SetSysCreatedBy),
		longitudeKey:        internalSerialization.DeserializeFloat64Func(m.SetLongitude),
		zipKey:              internalSerialization.DeserializeStringFunc(m.SetZip),
		profitsKey:          internalSerialization.DeserializeInt64Func(m.SetProfits),
		phoneKey:            internalSerialization.DeserializeStringFunc(m.SetPhone),
		faxPhoneKey:         internalSerialization.DeserializeStringFunc(m.SetFaxPhone),
		nameKey:             internalSerialization.DeserializeStringFunc(m.SetName),
		bannerTextKey:       internalSerialization.DeserializeStringFunc(m.SetBannerText),
		accountCodeKey:      internalSerialization.DeserializeStringFunc(m.SetAccountCode),
		primaryKey:          internalSerialization.DeserializeBoolFunc(m.SetPrimary),
		cityKey:             internalSerialization.DeserializeStringFunc(m.SetCity),
		latitudeKey:         internalSerialization.DeserializeFloat64Func(m.SetLatitude),
		sysClassNameKey:     internalSerialization.DeserializeStringFunc(m.SetSysClassName),
		manufacturerKey:     internalSerialization.DeserializeBoolFunc(m.SetManufacturer),
		accountParentKey:    internalSerialization.DeserializeStringFunc(m.SetAccountParent),
		sysIDKey:            internalSerialization.DeserializeStringFunc(m.SetSysID),
		marketCapKey:        internalSerialization.DeserializeFloat64Func(m.SetMarketCap),
		numEmployeesKey:     internalSerialization.DeserializeInt64Func(m.SetNumEmployees),
		rankTierKey:         internalSerialization.DeserializeEnumFunc(ParseRankTier, m.SetRankTier),
		streetKey:           internalSerialization.DeserializeStringFunc(m.SetStreet),
		vendorKey:           internalSerialization.DeserializeBoolFunc(m.SetVendor),
		latLongErrorKey:     internalSerialization.DeserializeStringFunc(m.SetLatLongError),
		themeKey:            internalSerialization.DeserializeStringFunc(m.SetTheme),
		vendorTypeKey:       internalSerialization.DeserializeStringFunc(m.SetVendorType),
		websiteKey:          internalSerialization.DeserializeStringFunc(m.SetWebsite),
		revenuePerYearKey:   internalSerialization.DeserializeInt64Func(m.SetRevenuePerYear),
		publiclyTradedKey:   internalSerialization.DeserializeBoolFunc(m.SetPubliclyTraded),
		sysModCountKey:      internalSerialization.DeserializeStringFunc(m.SetSysModCount),
		sysTagsKey:          internalSerialization.DeserializeStringFunc(m.SetSysTags),
		partnerKey:          internalSerialization.DeserializeBoolFunc(m.SetPartner),
		registrationCodeKey: internalSerialization.DeserializeStringFunc(m.SetRegistrationCode),
		vendorManagerKey:    internalSerialization.DeserializeStringFunc(m.SetVendorManager),
		accountPathKey:      internalSerialization.DeserializeStringFunc(m.SetAccountPath),
		primaryContactKey:   internalSerialization.DeserializeStringFunc(m.SetPrimaryContact),
		customerKey:         internalSerialization.DeserializeBoolFunc(m.SetCustomer),
	}
}

// GetBannerImageLight returns the banner image light value.
func (m *Account) GetBannerImageLight() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, bannerImageLightKey)
}

// SetBannerImageLight sets the banner image light value.
func (m *Account) SetBannerImageLight(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, bannerImageLightKey, val)
}

// GetCountry returns the country value.
func (m *Account) GetCountry() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, countryKey)
}

// SetCountry sets the country value.
func (m *Account) SetCountry(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, countryKey, val)
}

// GetParent returns the parent value.
func (m *Account) GetParent() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, parentKey)
}

// SetParent sets the parent value.
func (m *Account) SetParent(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, parentKey, val)
}

// GetNotes returns the notes value.
func (m *Account) GetNotes() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, notesKey)
}

// SetNotes sets the notes value.
func (m *Account) SetNotes(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, notesKey, val)
}

// GetStockSymbol returns the stock symbol value.
func (m *Account) GetStockSymbol() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, stockSymbolKey)
}

// SetStockSymbol sets the stock symbol value.
func (m *Account) SetStockSymbol(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, stockSymbolKey, val)
}

// GetDiscount returns the discount value.
func (m *Account) GetDiscount() (*int64, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *int64](m, discountKey)
}

// SetDiscount sets the discount value.
func (m *Account) SetDiscount(val *int64) error {
	return store.DefaultBackedModelMutatorFunc(m, discountKey, val)
}

// GetActiveEscalation returns the active escalation value.
func (m *Account) GetActiveEscalation() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, activeEscalationKey)
}

// SetActiveEscalation sets the active escalation value.
func (m *Account) SetActiveEscalation(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, activeEscalationKey, val)
}

// GetSysUpdatedOn returns the sys updated on value.
func (m *Account) GetSysUpdatedOn() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, sysUpdatedOnKey)
}

// SetSysUpdatedOn sets the sys updated on value.
func (m *Account) SetSysUpdatedOn(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysUpdatedOnKey, val)
}

// GetAppleIcon returns the apple icon value.
func (m *Account) GetAppleIcon() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, appleIconKey)
}

// SetAppleIcon sets the apple icon value.
func (m *Account) SetAppleIcon(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, appleIconKey, val)
}

// GetNumber returns the number value.
func (m *Account) GetNumber() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, numberKey)
}

// SetNumber sets the number value.
func (m *Account) SetNumber(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, numberKey, val)
}

// GetSysUpdatedBy returns the sys updated by value.
func (m *Account) GetSysUpdatedBy() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, sysUpdatedByKey)
}

// SetSysUpdatedBy sets the sys updated by value.
func (m *Account) SetSysUpdatedBy(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysUpdatedByKey, val)
}

// GetFiscalYear returns the fiscal year value.
func (m *Account) GetFiscalYear() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, fiscalYearKey)
}

// SetFiscalYear sets the fiscal year value.
func (m *Account) SetFiscalYear(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, fiscalYearKey, val)
}

// GetSysCreatedOn returns the sys created on value.
func (m *Account) GetSysCreatedOn() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, sysCreatedOnKey)
}

// SetSysCreatedOn sets the sys created on value.
func (m *Account) SetSysCreatedOn(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysCreatedOnKey, val)
}

// GetContact returns the contact value.
func (m *Account) GetContact() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, contactKey)
}

// SetContact sets the contact value.
func (m *Account) SetContact(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, contactKey, val)
}

// GetStockPrice returns the stock price value.
func (m *Account) GetStockPrice() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, stockPriceKey)
}

// SetStockPrice sets the stock price value.
func (m *Account) SetStockPrice(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, stockPriceKey, val)
}

// GetState returns the state value.
func (m *Account) GetState() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, stateKey)
}

// SetState sets the state value.
func (m *Account) SetState(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, stateKey, val)
}

// GetBannerImage returns the banner image value.
func (m *Account) GetBannerImage() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, bannerImageKey)
}

// SetBannerImage sets the banner image value.
func (m *Account) SetBannerImage(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, bannerImageKey, val)
}

// GetSysCreatedBy returns the sys created by value.
func (m *Account) GetSysCreatedBy() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, sysCreatedByKey)
}

// SetSysCreatedBy sets the sys created by value.
func (m *Account) SetSysCreatedBy(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysCreatedByKey, val)
}

// GetLongitude returns the longitude value.
func (m *Account) GetLongitude() (*float64, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *float64](m, longitudeKey)
}

// SetLongitude sets the longitude value.
func (m *Account) SetLongitude(val *float64) error {
	return store.DefaultBackedModelMutatorFunc(m, longitudeKey, val)
}

// GetZip returns the zip value.
func (m *Account) GetZip() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, zipKey)
}

// SetZip sets the zip value.
func (m *Account) SetZip(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, zipKey, val)
}

// GetProfits returns the profits value.
func (m *Account) GetProfits() (*int64, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *int64](m, profitsKey)
}

// SetProfits sets the profits value.
func (m *Account) SetProfits(val *int64) error {
	return store.DefaultBackedModelMutatorFunc(m, profitsKey, val)
}

// GetPhone returns the phone value.
func (m *Account) GetPhone() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, phoneKey)
}

// SetPhone sets the phone value.
func (m *Account) SetPhone(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, phoneKey, val)
}

// GetFaxPhone returns the fax phone value.
func (m *Account) GetFaxPhone() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, faxPhoneKey)
}

// SetFaxPhone sets the fax phone value.s
func (m *Account) SetFaxPhone(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, faxPhoneKey, val)
}

// GetName returns the name value.
func (m *Account) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, nameKey)
}

// SetName sets the name value.
func (m *Account) SetName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nameKey, val)
}

// GetBannerText returns the banner text value.
func (m *Account) GetBannerText() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, bannerTextKey)
}

// SetBannerText sets the banner text value.
func (m *Account) SetBannerText(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, bannerTextKey, val)
}

// GetAccountCode returns the account code value.
func (m *Account) GetAccountCode() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, accountCodeKey)
}

// SetAccountCode sets the account code value.
func (m *Account) SetAccountCode(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, accountCodeKey, val)
}

// GetPrimary returns the primary value.
func (m *Account) GetPrimary() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *bool](m, primaryKey)
}

// SetPrimary sets the primary value.
func (m *Account) SetPrimary(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, primaryKey, val)
}

// GetCity returns the city value.
func (m *Account) GetCity() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, cityKey)
}

// SetCity sets the city value.
func (m *Account) SetCity(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, cityKey, val)
}

// GetLatitude returns the latitude value.
func (m *Account) GetLatitude() (*float64, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *float64](m, latitudeKey)
}

// SetLatitude sets the latitude value.
func (m *Account) SetLatitude(val *float64) error {
	return store.DefaultBackedModelMutatorFunc(m, latitudeKey, val)
}

// GetSysClassName returns the sys class name value.
func (m *Account) GetSysClassName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, sysClassNameKey)
}

// SetSysClassName sets the sys class name value.
func (m *Account) SetSysClassName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysClassNameKey, val)
}

// GetManufacturer returns the manufacturer value.
func (m *Account) GetManufacturer() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *bool](m, manufacturerKey)
}

// SetManufacturer sets the manufacturer value.
func (m *Account) SetManufacturer(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, manufacturerKey, val)
}

// GetAccountParent returns the account parent value.
func (m *Account) GetAccountParent() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, accountParentKey)
}

// SetAccountParent sets the account parent value.
func (m *Account) SetAccountParent(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, accountParentKey, val)
}

// GetSysID returns the sys id value.
func (m *Account) GetSysID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, sysIDKey)
}

// SetSysID sets the sys id value.
func (m *Account) SetSysID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysIDKey, val)
}

// GetMarketCap returns the market cap value.
func (m *Account) GetMarketCap() (*float64, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *float64](m, marketCapKey)
}

// SetMarketCap sets the market cap value.
func (m *Account) SetMarketCap(val *float64) error {
	return store.DefaultBackedModelMutatorFunc(m, marketCapKey, val)
}

// GetNumEmployees returns the num employees value.
func (m *Account) GetNumEmployees() (*int64, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *int64](m, numEmployeesKey)
}

// SetNumEmployees sets the num employees value.
func (m *Account) SetNumEmployees(val *int64) error {
	return store.DefaultBackedModelMutatorFunc(m, numEmployeesKey, val)
}

// GetRankTier returns the rank tier value.
func (m *Account) GetRankTier() (*RankTier, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *RankTier](m, rankTierKey)
}

// SetRankTier sets the rank tier value.
func (m *Account) SetRankTier(val *RankTier) error {
	return store.DefaultBackedModelMutatorFunc(m, rankTierKey, val)
}

// GetStreet returns the street value.
func (m *Account) GetStreet() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, streetKey)
}

// SetStreet sets the street value.
func (m *Account) SetStreet(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, streetKey, val)
}

// GetVendor returns the vendor value.
func (m *Account) GetVendor() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *bool](m, vendorKey)
}

// SetVendor sets the vendor value.
func (m *Account) SetVendor(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, vendorKey, val)
}

// GetLatLongError returns the lat long error value.
func (m *Account) GetLatLongError() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, latLongErrorKey)
}

// SetLatLongError sets the lat long error value.
func (m *Account) SetLatLongError(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, latLongErrorKey, val)
}

// GetTheme returns the theme value.
func (m *Account) GetTheme() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, themeKey)
}

// SetTheme sets the theme value.
func (m *Account) SetTheme(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, themeKey, val)
}

// GetVendorType returns the vendor type value.
func (m *Account) GetVendorType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, vendorTypeKey)
}

// SetVendorType sets the vendor type value.
func (m *Account) SetVendorType(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, vendorTypeKey, val)
}

// GetWebsite returns the website value.
func (m *Account) GetWebsite() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, websiteKey)
}

// SetWebsite sets the website value.
func (m *Account) SetWebsite(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, websiteKey, val)
}

// GetRevenuePerYear returns the revenue per year value.
func (m *Account) GetRevenuePerYear() (*int64, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *int64](m, revenuePerYearKey)
}

// SetRevenuePerYear sets the revenue per year value.
func (m *Account) SetRevenuePerYear(val *int64) error {
	return store.DefaultBackedModelMutatorFunc(m, revenuePerYearKey, val)
}

// GetPubliclyTraded returns the publicly traded value.
func (m *Account) GetPubliclyTraded() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *bool](m, publiclyTradedKey)
}

// SetPubliclyTraded sets the publicly traded value.
func (m *Account) SetPubliclyTraded(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, publiclyTradedKey, val)
}

// GetSysModCount returns the sys mod count value.
func (m *Account) GetSysModCount() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, sysModCountKey)
}

// SetSysModCount sets the sys mod count value.
func (m *Account) SetSysModCount(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysModCountKey, val)
}

// GetSysTags returns the sys tags value.
func (m *Account) GetSysTags() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, sysTagsKey)
}

// SetSysTags sets the sys tags value.
func (m *Account) SetSysTags(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysTagsKey, val)
}

// GetPartner returns the partner value.
func (m *Account) GetPartner() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *bool](m, partnerKey)
}

// SetPartner sets the partner value.
func (m *Account) SetPartner(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, partnerKey, val)
}

// GetRegistrationCode returns the registration code value.
func (m *Account) GetRegistrationCode() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, registrationCodeKey)
}

// SetRegistrationCode sets the registration code value.
func (m *Account) SetRegistrationCode(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, registrationCodeKey, val)
}

// GetVendorManager returns the vendor manager value.
func (m *Account) GetVendorManager() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, vendorManagerKey)
}

// SetVendorManager sets the vendor manager value.
func (m *Account) SetVendorManager(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, vendorManagerKey, val)
}

// GetAccountPath returns the account path value.
func (m *Account) GetAccountPath() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, accountPathKey)
}

// SetAccountPath sets the account path value.
func (m *Account) SetAccountPath(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, accountPathKey, val)
}

// GetPrimaryContact returns the primary contact value.
func (m *Account) GetPrimaryContact() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *string](m, primaryContactKey)
}

// SetPrimaryContact sets the primary contact value.
func (m *Account) SetPrimaryContact(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, primaryContactKey, val)
}

// GetCustomer returns the customer value.
func (m *Account) GetCustomer() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*Account, *bool](m, customerKey)
}

// SetCustomer sets the customer value.
func (m *Account) SetCustomer(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, customerKey, val)
}
