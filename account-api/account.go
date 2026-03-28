package accountapi

import (
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

const (
	account_codeKey       = "account_code"
	account_parentKey     = "account_parent"
	account_pathKey       = "account_path"
	active_escalationKey  = "active_escalation"
	apple_iconKey         = "apple_icon"
	banner_imageKey       = "banner_image"
	banner_image_lightKey = "banner_image_light"
	banner_textKey        = "banner_text"
	cityKey               = "city"
	contactKey            = "contact"
	countryKey            = "country"
	customerKey           = "customer"
	discountKey           = "discount"
	fax_phoneKey          = "fax_phone"
	fiscal_yearKey        = "fiscal_year"
	lat_long_errorKey     = "lat_long_error"
	latitudeKey           = "latitude"
	longitudeKey          = "longitude"
	manufacturerKey       = "manufacturer"
	market_capKey         = "market_cap"
	nameKey               = "name"
	notesKey              = "notes"
	num_employeesKey      = "num_employees"
	numberKey             = "number"
	parentKey             = "parent"
	partnerKey            = "partner"
	phoneKey              = "phone"
	primaryKey            = "primary"
	primary_contactKey    = "primary_contact"
	profitsKey            = "profits"
	publicly_tradedKey    = "publicly_traded"
	rank_tierKey          = "rank_tier"
	registration_codeKey  = "registration_code"
	revenue_per_yearKey   = "revenue_per_year"
	stateKey              = "state"
	stock_priceKey        = "stock_price"
	stock_symbolKey       = "stock_symbol"
	streetKey             = "street"
	sys_class_nameKey     = "sys_class_name"
	sys_created_byKey     = "sys_created_by"
	sys_created_onKey     = "sys_created_on"
	sys_idKey             = "sys_id"
	sys_mod_countKey      = "sys_mod_count"
	sys_updated_byKey     = "sys_updated_by"
	sys_updated_onKey     = "sys_updated_on"
	themeKey              = "theme"
	vendorKey             = "vendor"
	vendor_managerKey     = "vendor_manager"
	vendor_typeKey        = "vendor_type"
	websiteKey            = "website"
	zipKey                = "zip"
)

// Account defines the interface for the Account model.
type Account interface {

	// GetAccountCode gets the account_code property value.
	GetAccountCode() (*string, error)
	// SetAccountCode sets the account_code property value.
	SetAccountCode(*string) error

	// GetAccountParent gets the account_parent property value.
	GetAccountParent() (*string, error)
	// SetAccountParent sets the account_parent property value.
	SetAccountParent(*string) error

	// GetAccountPath gets the account_path property value.
	GetAccountPath() (*string, error)
	// SetAccountPath sets the account_path property value.
	SetAccountPath(*string) error

	// GetActiveEscalation gets the active_escalation property value.
	GetActiveEscalation() (*string, error)
	// SetActiveEscalation sets the active_escalation property value.
	SetActiveEscalation(*string) error

	// GetAppleIcon gets the apple_icon property value.
	GetAppleIcon() (*string, error)
	// SetAppleIcon sets the apple_icon property value.
	SetAppleIcon(*string) error

	// GetBannerImage gets the banner_image property value.
	GetBannerImage() (*string, error)
	// SetBannerImage sets the banner_image property value.
	SetBannerImage(*string) error

	// GetBannerImageLight gets the banner_image_light property value.
	GetBannerImageLight() (*string, error)
	// SetBannerImageLight sets the banner_image_light property value.
	SetBannerImageLight(*string) error

	// GetBannerText gets the banner_text property value.
	GetBannerText() (*string, error)
	// SetBannerText sets the banner_text property value.
	SetBannerText(*string) error

	// GetCity gets the city property value.
	GetCity() (*string, error)
	// SetCity sets the city property value.
	SetCity(*string) error

	// GetContact gets the contact property value.
	GetContact() (*string, error)
	// SetContact sets the contact property value.
	SetContact(*string) error

	// GetCountry gets the country property value.
	GetCountry() (*string, error)
	// SetCountry sets the country property value.
	SetCountry(*string) error

	// GetCustomer gets the customer property value.
	GetCustomer() (*bool, error)
	// SetCustomer sets the customer property value.
	SetCustomer(*bool) error

	// GetDiscount gets the discount property value.
	GetDiscount() (*float64, error)
	// SetDiscount sets the discount property value.
	SetDiscount(*float64) error

	// GetFaxPhone gets the fax_phone property value.
	GetFaxPhone() (*string, error)
	// SetFaxPhone sets the fax_phone property value.
	SetFaxPhone(*string) error

	// GetFiscalYear gets the fiscal_year property value.
	GetFiscalYear() (*string, error)
	// SetFiscalYear sets the fiscal_year property value.
	SetFiscalYear(*string) error

	// GetLatLongError gets the lat_long_error property value.
	GetLatLongError() (*string, error)
	// SetLatLongError sets the lat_long_error property value.
	SetLatLongError(*string) error

	// GetLatitude gets the latitude property value.
	GetLatitude() (*float64, error)
	// SetLatitude sets the latitude property value.
	SetLatitude(*float64) error

	// GetLongitude gets the longitude property value.
	GetLongitude() (*float64, error)
	// SetLongitude sets the longitude property value.
	SetLongitude(*float64) error

	// GetManufacturer gets the manufacturer property value.
	GetManufacturer() (*bool, error)
	// SetManufacturer sets the manufacturer property value.
	SetManufacturer(*bool) error

	// GetMarketCap gets the market_cap property value.
	GetMarketCap() (*float64, error)
	// SetMarketCap sets the market_cap property value.
	SetMarketCap(*float64) error

	// GetName gets the name property value.
	GetName() (*string, error)
	// SetName sets the name property value.
	SetName(*string) error

	// GetNotes gets the notes property value.
	GetNotes() (*string, error)
	// SetNotes sets the notes property value.
	SetNotes(*string) error

	// GetNumEmployees gets the num_employees property value.
	GetNumEmployees() (*int64, error)
	// SetNumEmployees sets the num_employees property value.
	SetNumEmployees(*int64) error

	// GetNumber gets the number property value.
	GetNumber() (*string, error)
	// SetNumber sets the number property value.
	SetNumber(*string) error

	// GetParent gets the parent property value.
	GetParent() (*string, error)
	// SetParent sets the parent property value.
	SetParent(*string) error

	// GetPartner gets the partner property value.
	GetPartner() (*bool, error)
	// SetPartner sets the partner property value.
	SetPartner(*bool) error

	// GetPhone gets the phone property value.
	GetPhone() (*string, error)
	// SetPhone sets the phone property value.
	SetPhone(*string) error

	// GetPrimary gets the primary property value.
	GetPrimary() (*bool, error)
	// SetPrimary sets the primary property value.
	SetPrimary(*bool) error

	// GetPrimaryContact gets the primary_contact property value.
	GetPrimaryContact() (*string, error)
	// SetPrimaryContact sets the primary_contact property value.
	SetPrimaryContact(*string) error

	// GetProfits gets the profits property value.
	GetProfits() (*float64, error)
	// SetProfits sets the profits property value.
	SetProfits(*float64) error

	// GetPubliclyTraded gets the publicly_traded property value.
	GetPubliclyTraded() (*bool, error)
	// SetPubliclyTraded sets the publicly_traded property value.
	SetPubliclyTraded(*bool) error

	// GetRankTier gets the rank_tier property value.
	GetRankTier() (*string, error)
	// SetRankTier sets the rank_tier property value.
	SetRankTier(*string) error

	// GetRegistrationCode gets the registration_code property value.
	GetRegistrationCode() (*string, error)
	// SetRegistrationCode sets the registration_code property value.
	SetRegistrationCode(*string) error

	// GetRevenuePerYear gets the revenue_per_year property value.
	GetRevenuePerYear() (*float64, error)
	// SetRevenuePerYear sets the revenue_per_year property value.
	SetRevenuePerYear(*float64) error

	// GetState gets the state property value.
	GetState() (*string, error)
	// SetState sets the state property value.
	SetState(*string) error

	// GetStockPrice gets the stock_price property value.
	GetStockPrice() (*string, error)
	// SetStockPrice sets the stock_price property value.
	SetStockPrice(*string) error

	// GetStockSymbol gets the stock_symbol property value.
	GetStockSymbol() (*string, error)
	// SetStockSymbol sets the stock_symbol property value.
	SetStockSymbol(*string) error

	// GetStreet gets the street property value.
	GetStreet() (*string, error)
	// SetStreet sets the street property value.
	SetStreet(*string) error

	// GetSysClassName gets the sys_class_name property value.
	GetSysClassName() (*string, error)
	// SetSysClassName sets the sys_class_name property value.
	SetSysClassName(*string) error

	// GetSysCreatedBy gets the sys_created_by property value.
	GetSysCreatedBy() (*string, error)
	// SetSysCreatedBy sets the sys_created_by property value.
	SetSysCreatedBy(*string) error

	// GetSysCreatedOn gets the sys_created_on property value.
	GetSysCreatedOn() (*time.Time, error)
	// SetSysCreatedOn sets the sys_created_on property value.
	SetSysCreatedOn(*time.Time) error

	// GetSysId gets the sys_id property value.
	GetSysId() (*string, error)
	// SetSysId sets the sys_id property value.
	SetSysId(*string) error

	// GetSysModCount gets the sys_mod_count property value.
	GetSysModCount() (*int64, error)
	// SetSysModCount sets the sys_mod_count property value.
	SetSysModCount(*int64) error

	// GetSysUpdatedBy gets the sys_updated_by property value.
	GetSysUpdatedBy() (*string, error)
	// SetSysUpdatedBy sets the sys_updated_by property value.
	SetSysUpdatedBy(*string) error

	// GetSysUpdatedOn gets the sys_updated_on property value.
	GetSysUpdatedOn() (*time.Time, error)
	// SetSysUpdatedOn sets the sys_updated_on property value.
	SetSysUpdatedOn(*time.Time) error

	// GetTheme gets the theme property value.
	GetTheme() (*string, error)
	// SetTheme sets the theme property value.
	SetTheme(*string) error

	// GetVendor gets the vendor property value.
	GetVendor() (*bool, error)
	// SetVendor sets the vendor property value.
	SetVendor(*bool) error

	// GetVendorManager gets the vendor_manager property value.
	GetVendorManager() (*string, error)
	// SetVendorManager sets the vendor_manager property value.
	SetVendorManager(*string) error

	// GetVendorType gets the vendor_type property value.
	GetVendorType() (*string, error)
	// SetVendorType sets the vendor_type property value.
	SetVendorType(*string) error

	// GetWebsite gets the website property value.
	GetWebsite() (*string, error)
	// SetWebsite sets the website property value.
	SetWebsite(*string) error

	// GetZip gets the zip property value.
	GetZip() (*string, error)
	// SetZip sets the zip property value.
	SetZip(*string) error

	serialization.Parsable
	kiotaStore.BackedModel
}

// AccountModel is the concrete implementation of the Account interface.
type AccountModel struct {
	newInternal.Model
}

// NewAccount creates a new instance of AccountModel with a backing store.
func NewAccount() *AccountModel {
	return &AccountModel{
		newInternal.NewBaseModel(),
	}
}

// CreateAccountFromDiscriminatorValue is a factory function for creating AccountModel instances during deserialization.
func CreateAccountFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewAccount(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *AccountModel) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}

	return internalSerialization.Serialize(writer,

		internalSerialization.SerializeStringFunc(account_codeKey)(m.GetAccountCode),

		internalSerialization.SerializeStringFunc(account_parentKey)(m.GetAccountParent),

		internalSerialization.SerializeStringFunc(account_pathKey)(m.GetAccountPath),

		internalSerialization.SerializeStringFunc(active_escalationKey)(m.GetActiveEscalation),

		internalSerialization.SerializeStringFunc(apple_iconKey)(m.GetAppleIcon),

		internalSerialization.SerializeStringFunc(banner_imageKey)(m.GetBannerImage),

		internalSerialization.SerializeStringFunc(banner_image_lightKey)(m.GetBannerImageLight),

		internalSerialization.SerializeStringFunc(banner_textKey)(m.GetBannerText),

		internalSerialization.SerializeStringFunc(cityKey)(m.GetCity),

		internalSerialization.SerializeStringFunc(contactKey)(m.GetContact),

		internalSerialization.SerializeStringFunc(countryKey)(m.GetCountry),

		internalSerialization.SerializeStringToBoolFunc(customerKey)(m.GetCustomer),

		internalSerialization.SerializeStringToFloat64Func(discountKey)(m.GetDiscount),

		internalSerialization.SerializeStringFunc(fax_phoneKey)(m.GetFaxPhone),

		internalSerialization.SerializeStringFunc(fiscal_yearKey)(m.GetFiscalYear),

		internalSerialization.SerializeStringFunc(lat_long_errorKey)(m.GetLatLongError),

		internalSerialization.SerializeStringToFloat64Func(latitudeKey)(m.GetLatitude),

		internalSerialization.SerializeStringToFloat64Func(longitudeKey)(m.GetLongitude),

		internalSerialization.SerializeStringToBoolFunc(manufacturerKey)(m.GetManufacturer),

		internalSerialization.SerializeStringToFloat64Func(market_capKey)(m.GetMarketCap),

		internalSerialization.SerializeStringFunc(nameKey)(m.GetName),

		internalSerialization.SerializeStringFunc(notesKey)(m.GetNotes),

		internalSerialization.SerializeStringToInt64Func(num_employeesKey)(m.GetNumEmployees),

		internalSerialization.SerializeStringFunc(numberKey)(m.GetNumber),

		internalSerialization.SerializeStringFunc(parentKey)(m.GetParent),

		internalSerialization.SerializeStringToBoolFunc(partnerKey)(m.GetPartner),

		internalSerialization.SerializeStringFunc(phoneKey)(m.GetPhone),

		internalSerialization.SerializeStringToBoolFunc(primaryKey)(m.GetPrimary),

		internalSerialization.SerializeStringFunc(primary_contactKey)(m.GetPrimaryContact),

		internalSerialization.SerializeStringToFloat64Func(profitsKey)(m.GetProfits),

		internalSerialization.SerializeStringToBoolFunc(publicly_tradedKey)(m.GetPubliclyTraded),

		internalSerialization.SerializeStringFunc(rank_tierKey)(m.GetRankTier),

		internalSerialization.SerializeStringFunc(registration_codeKey)(m.GetRegistrationCode),

		internalSerialization.SerializeStringToFloat64Func(revenue_per_yearKey)(m.GetRevenuePerYear),

		internalSerialization.SerializeStringFunc(stateKey)(m.GetState),

		internalSerialization.SerializeStringFunc(stock_priceKey)(m.GetStockPrice),

		internalSerialization.SerializeStringFunc(stock_symbolKey)(m.GetStockSymbol),

		internalSerialization.SerializeStringFunc(streetKey)(m.GetStreet),

		internalSerialization.SerializeStringFunc(sys_class_nameKey)(m.GetSysClassName),

		internalSerialization.SerializeStringFunc(sys_created_byKey)(m.GetSysCreatedBy),

		internalSerialization.SerializeStringToTimeFunc(sys_created_onKey, "2006-01-02 15:04:05")(m.GetSysCreatedOn),

		internalSerialization.SerializeStringFunc(sys_idKey)(m.GetSysId),

		internalSerialization.SerializeStringToInt64Func(sys_mod_countKey)(m.GetSysModCount),

		internalSerialization.SerializeStringFunc(sys_updated_byKey)(m.GetSysUpdatedBy),

		internalSerialization.SerializeStringToTimeFunc(sys_updated_onKey, "2006-01-02 15:04:05")(m.GetSysUpdatedOn),

		internalSerialization.SerializeStringFunc(themeKey)(m.GetTheme),

		internalSerialization.SerializeStringToBoolFunc(vendorKey)(m.GetVendor),

		internalSerialization.SerializeStringFunc(vendor_managerKey)(m.GetVendorManager),

		internalSerialization.SerializeStringFunc(vendor_typeKey)(m.GetVendorType),

		internalSerialization.SerializeStringFunc(websiteKey)(m.GetWebsite),

		internalSerialization.SerializeStringFunc(zipKey)(m.GetZip),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *AccountModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{

		account_codeKey: internalSerialization.DeserializeStringFunc()(m.SetAccountCode),

		account_parentKey: internalSerialization.DeserializeStringFunc()(m.SetAccountParent),

		account_pathKey: internalSerialization.DeserializeStringFunc()(m.SetAccountPath),

		active_escalationKey: internalSerialization.DeserializeStringFunc()(m.SetActiveEscalation),

		apple_iconKey: internalSerialization.DeserializeStringFunc()(m.SetAppleIcon),

		banner_imageKey: internalSerialization.DeserializeStringFunc()(m.SetBannerImage),

		banner_image_lightKey: internalSerialization.DeserializeStringFunc()(m.SetBannerImageLight),

		banner_textKey: internalSerialization.DeserializeStringFunc()(m.SetBannerText),

		cityKey: internalSerialization.DeserializeStringFunc()(m.SetCity),

		contactKey: internalSerialization.DeserializeStringFunc()(m.SetContact),

		countryKey: internalSerialization.DeserializeStringFunc()(m.SetCountry),

		customerKey: internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToBoolPtr)(m.SetCustomer),

		discountKey: internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToFloat64Ptr)(m.SetDiscount),

		fax_phoneKey: internalSerialization.DeserializeStringFunc()(m.SetFaxPhone),

		fiscal_yearKey: internalSerialization.DeserializeStringFunc()(m.SetFiscalYear),

		lat_long_errorKey: internalSerialization.DeserializeStringFunc()(m.SetLatLongError),

		latitudeKey: internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToFloat64Ptr)(m.SetLatitude),

		longitudeKey: internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToFloat64Ptr)(m.SetLongitude),

		manufacturerKey: internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToBoolPtr)(m.SetManufacturer),

		market_capKey: internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToFloat64Ptr)(m.SetMarketCap),

		nameKey: internalSerialization.DeserializeStringFunc()(m.SetName),

		notesKey: internalSerialization.DeserializeStringFunc()(m.SetNotes),

		num_employeesKey: internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToInt64Ptr)(m.SetNumEmployees),

		numberKey: internalSerialization.DeserializeStringFunc()(m.SetNumber),

		parentKey: internalSerialization.DeserializeStringFunc()(m.SetParent),

		partnerKey: internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToBoolPtr)(m.SetPartner),

		phoneKey: internalSerialization.DeserializeStringFunc()(m.SetPhone),

		primaryKey: internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToBoolPtr)(m.SetPrimary),

		primary_contactKey: internalSerialization.DeserializeStringFunc()(m.SetPrimaryContact),

		profitsKey: internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToFloat64Ptr)(m.SetProfits),

		publicly_tradedKey: internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToBoolPtr)(m.SetPubliclyTraded),

		rank_tierKey: internalSerialization.DeserializeStringFunc()(m.SetRankTier),

		registration_codeKey: internalSerialization.DeserializeStringFunc()(m.SetRegistrationCode),

		revenue_per_yearKey: internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToFloat64Ptr)(m.SetRevenuePerYear),

		stateKey: internalSerialization.DeserializeStringFunc()(m.SetState),

		stock_priceKey: internalSerialization.DeserializeStringFunc()(m.SetStockPrice),

		stock_symbolKey: internalSerialization.DeserializeStringFunc()(m.SetStockSymbol),

		streetKey: internalSerialization.DeserializeStringFunc()(m.SetStreet),

		sys_class_nameKey: internalSerialization.DeserializeStringFunc()(m.SetSysClassName),

		sys_created_byKey: internalSerialization.DeserializeStringFunc()(m.SetSysCreatedBy),

		sys_created_onKey: internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToTimePtr("2006-01-02 15:04:05"))(m.SetSysCreatedOn),

		sys_idKey: internalSerialization.DeserializeStringFunc()(m.SetSysId),

		sys_mod_countKey: internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToInt64Ptr)(m.SetSysModCount),

		sys_updated_byKey: internalSerialization.DeserializeStringFunc()(m.SetSysUpdatedBy),

		sys_updated_onKey: internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToTimePtr("2006-01-02 15:04:05"))(m.SetSysUpdatedOn),

		themeKey: internalSerialization.DeserializeStringFunc()(m.SetTheme),

		vendorKey: internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToBoolPtr)(m.SetVendor),

		vendor_managerKey: internalSerialization.DeserializeStringFunc()(m.SetVendorManager),

		vendor_typeKey: internalSerialization.DeserializeStringFunc()(m.SetVendorType),

		websiteKey: internalSerialization.DeserializeStringFunc()(m.SetWebsite),

		zipKey: internalSerialization.DeserializeStringFunc()(m.SetZip),
	}
}

// GetAccountCode returns the account_code property value.
func (m *AccountModel) GetAccountCode() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, account_codeKey)
}

// SetAccountCode sets the account_code property value.
func (m *AccountModel) SetAccountCode(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, account_codeKey, val)
}

// GetAccountParent returns the account_parent property value.
func (m *AccountModel) GetAccountParent() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, account_parentKey)
}

// SetAccountParent sets the account_parent property value.
func (m *AccountModel) SetAccountParent(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, account_parentKey, val)
}

// GetAccountPath returns the account_path property value.
func (m *AccountModel) GetAccountPath() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, account_pathKey)
}

// SetAccountPath sets the account_path property value.
func (m *AccountModel) SetAccountPath(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, account_pathKey, val)
}

// GetActiveEscalation returns the active_escalation property value.
func (m *AccountModel) GetActiveEscalation() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, active_escalationKey)
}

// SetActiveEscalation sets the active_escalation property value.
func (m *AccountModel) SetActiveEscalation(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, active_escalationKey, val)
}

// GetAppleIcon returns the apple_icon property value.
func (m *AccountModel) GetAppleIcon() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, apple_iconKey)
}

// SetAppleIcon sets the apple_icon property value.
func (m *AccountModel) SetAppleIcon(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, apple_iconKey, val)
}

// GetBannerImage returns the banner_image property value.
func (m *AccountModel) GetBannerImage() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, banner_imageKey)
}

// SetBannerImage sets the banner_image property value.
func (m *AccountModel) SetBannerImage(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, banner_imageKey, val)
}

// GetBannerImageLight returns the banner_image_light property value.
func (m *AccountModel) GetBannerImageLight() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, banner_image_lightKey)
}

// SetBannerImageLight sets the banner_image_light property value.
func (m *AccountModel) SetBannerImageLight(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, banner_image_lightKey, val)
}

// GetBannerText returns the banner_text property value.
func (m *AccountModel) GetBannerText() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, banner_textKey)
}

// SetBannerText sets the banner_text property value.
func (m *AccountModel) SetBannerText(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, banner_textKey, val)
}

// GetCity returns the city property value.
func (m *AccountModel) GetCity() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, cityKey)
}

// SetCity sets the city property value.
func (m *AccountModel) SetCity(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, cityKey, val)
}

// GetContact returns the contact property value.
func (m *AccountModel) GetContact() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, contactKey)
}

// SetContact sets the contact property value.
func (m *AccountModel) SetContact(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, contactKey, val)
}

// GetCountry returns the country property value.
func (m *AccountModel) GetCountry() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, countryKey)
}

// SetCountry sets the country property value.
func (m *AccountModel) SetCountry(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, countryKey, val)
}

// GetCustomer returns the customer property value.
func (m *AccountModel) GetCustomer() (*bool, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](backingStore, customerKey)
}

// SetCustomer sets the customer property value.
func (m *AccountModel) SetCustomer(val *bool) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, customerKey, val)
}

// GetDiscount returns the discount property value.
func (m *AccountModel) GetDiscount() (*float64, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *float64](backingStore, discountKey)
}

// SetDiscount sets the discount property value.
func (m *AccountModel) SetDiscount(val *float64) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, discountKey, val)
}

// GetFaxPhone returns the fax_phone property value.
func (m *AccountModel) GetFaxPhone() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, fax_phoneKey)
}

// SetFaxPhone sets the fax_phone property value.
func (m *AccountModel) SetFaxPhone(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, fax_phoneKey, val)
}

// GetFiscalYear returns the fiscal_year property value.
func (m *AccountModel) GetFiscalYear() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, fiscal_yearKey)
}

// SetFiscalYear sets the fiscal_year property value.
func (m *AccountModel) SetFiscalYear(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, fiscal_yearKey, val)
}

// GetLatLongError returns the lat_long_error property value.
func (m *AccountModel) GetLatLongError() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, lat_long_errorKey)
}

// SetLatLongError sets the lat_long_error property value.
func (m *AccountModel) SetLatLongError(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, lat_long_errorKey, val)
}

// GetLatitude returns the latitude property value.
func (m *AccountModel) GetLatitude() (*float64, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *float64](backingStore, latitudeKey)
}

// SetLatitude sets the latitude property value.
func (m *AccountModel) SetLatitude(val *float64) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, latitudeKey, val)
}

// GetLongitude returns the longitude property value.
func (m *AccountModel) GetLongitude() (*float64, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *float64](backingStore, longitudeKey)
}

// SetLongitude sets the longitude property value.
func (m *AccountModel) SetLongitude(val *float64) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, longitudeKey, val)
}

// GetManufacturer returns the manufacturer property value.
func (m *AccountModel) GetManufacturer() (*bool, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](backingStore, manufacturerKey)
}

// SetManufacturer sets the manufacturer property value.
func (m *AccountModel) SetManufacturer(val *bool) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, manufacturerKey, val)
}

// GetMarketCap returns the market_cap property value.
func (m *AccountModel) GetMarketCap() (*float64, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *float64](backingStore, market_capKey)
}

// SetMarketCap sets the market_cap property value.
func (m *AccountModel) SetMarketCap(val *float64) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, market_capKey, val)
}

// GetName returns the name property value.
func (m *AccountModel) GetName() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, nameKey)
}

// SetName sets the name property value.
func (m *AccountModel) SetName(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, nameKey, val)
}

// GetNotes returns the notes property value.
func (m *AccountModel) GetNotes() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, notesKey)
}

// SetNotes sets the notes property value.
func (m *AccountModel) SetNotes(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, notesKey, val)
}

// GetNumEmployees returns the num_employees property value.
func (m *AccountModel) GetNumEmployees() (*int64, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *int64](backingStore, num_employeesKey)
}

// SetNumEmployees sets the num_employees property value.
func (m *AccountModel) SetNumEmployees(val *int64) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, num_employeesKey, val)
}

// GetNumber returns the number property value.
func (m *AccountModel) GetNumber() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, numberKey)
}

// SetNumber sets the number property value.
func (m *AccountModel) SetNumber(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, numberKey, val)
}

// GetParent returns the parent property value.
func (m *AccountModel) GetParent() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, parentKey)
}

// SetParent sets the parent property value.
func (m *AccountModel) SetParent(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, parentKey, val)
}

// GetPartner returns the partner property value.
func (m *AccountModel) GetPartner() (*bool, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](backingStore, partnerKey)
}

// SetPartner sets the partner property value.
func (m *AccountModel) SetPartner(val *bool) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, partnerKey, val)
}

// GetPhone returns the phone property value.
func (m *AccountModel) GetPhone() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, phoneKey)
}

// SetPhone sets the phone property value.
func (m *AccountModel) SetPhone(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, phoneKey, val)
}

// GetPrimary returns the primary property value.
func (m *AccountModel) GetPrimary() (*bool, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](backingStore, primaryKey)
}

// SetPrimary sets the primary property value.
func (m *AccountModel) SetPrimary(val *bool) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, primaryKey, val)
}

// GetPrimaryContact returns the primary_contact property value.
func (m *AccountModel) GetPrimaryContact() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, primary_contactKey)
}

// SetPrimaryContact sets the primary_contact property value.
func (m *AccountModel) SetPrimaryContact(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, primary_contactKey, val)
}

// GetProfits returns the profits property value.
func (m *AccountModel) GetProfits() (*float64, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *float64](backingStore, profitsKey)
}

// SetProfits sets the profits property value.
func (m *AccountModel) SetProfits(val *float64) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, profitsKey, val)
}

// GetPubliclyTraded returns the publicly_traded property value.
func (m *AccountModel) GetPubliclyTraded() (*bool, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](backingStore, publicly_tradedKey)
}

// SetPubliclyTraded sets the publicly_traded property value.
func (m *AccountModel) SetPubliclyTraded(val *bool) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, publicly_tradedKey, val)
}

// GetRankTier returns the rank_tier property value.
func (m *AccountModel) GetRankTier() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, rank_tierKey)
}

// SetRankTier sets the rank_tier property value.
func (m *AccountModel) SetRankTier(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, rank_tierKey, val)
}

// GetRegistrationCode returns the registration_code property value.
func (m *AccountModel) GetRegistrationCode() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, registration_codeKey)
}

// SetRegistrationCode sets the registration_code property value.
func (m *AccountModel) SetRegistrationCode(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, registration_codeKey, val)
}

// GetRevenuePerYear returns the revenue_per_year property value.
func (m *AccountModel) GetRevenuePerYear() (*float64, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *float64](backingStore, revenue_per_yearKey)
}

// SetRevenuePerYear sets the revenue_per_year property value.
func (m *AccountModel) SetRevenuePerYear(val *float64) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, revenue_per_yearKey, val)
}

// GetState returns the state property value.
func (m *AccountModel) GetState() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, stateKey)
}

// SetState sets the state property value.
func (m *AccountModel) SetState(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, stateKey, val)
}

// GetStockPrice returns the stock_price property value.
func (m *AccountModel) GetStockPrice() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, stock_priceKey)
}

// SetStockPrice sets the stock_price property value.
func (m *AccountModel) SetStockPrice(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, stock_priceKey, val)
}

// GetStockSymbol returns the stock_symbol property value.
func (m *AccountModel) GetStockSymbol() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, stock_symbolKey)
}

// SetStockSymbol sets the stock_symbol property value.
func (m *AccountModel) SetStockSymbol(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, stock_symbolKey, val)
}

// GetStreet returns the street property value.
func (m *AccountModel) GetStreet() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, streetKey)
}

// SetStreet sets the street property value.
func (m *AccountModel) SetStreet(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, streetKey, val)
}

// GetSysClassName returns the sys_class_name property value.
func (m *AccountModel) GetSysClassName() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, sys_class_nameKey)
}

// SetSysClassName sets the sys_class_name property value.
func (m *AccountModel) SetSysClassName(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sys_class_nameKey, val)
}

// GetSysCreatedBy returns the sys_created_by property value.
func (m *AccountModel) GetSysCreatedBy() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, sys_created_byKey)
}

// SetSysCreatedBy sets the sys_created_by property value.
func (m *AccountModel) SetSysCreatedBy(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sys_created_byKey, val)
}

// GetSysCreatedOn returns the sys_created_on property value.
func (m *AccountModel) GetSysCreatedOn() (*time.Time, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *time.Time](backingStore, sys_created_onKey)
}

// SetSysCreatedOn sets the sys_created_on property value.
func (m *AccountModel) SetSysCreatedOn(val *time.Time) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sys_created_onKey, val)
}

// GetSysId returns the sys_id property value.
func (m *AccountModel) GetSysId() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, sys_idKey)
}

// SetSysId sets the sys_id property value.
func (m *AccountModel) SetSysId(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sys_idKey, val)
}

// GetSysModCount returns the sys_mod_count property value.
func (m *AccountModel) GetSysModCount() (*int64, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *int64](backingStore, sys_mod_countKey)
}

// SetSysModCount sets the sys_mod_count property value.
func (m *AccountModel) SetSysModCount(val *int64) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sys_mod_countKey, val)
}

// GetSysUpdatedBy returns the sys_updated_by property value.
func (m *AccountModel) GetSysUpdatedBy() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, sys_updated_byKey)
}

// SetSysUpdatedBy sets the sys_updated_by property value.
func (m *AccountModel) SetSysUpdatedBy(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sys_updated_byKey, val)
}

// GetSysUpdatedOn returns the sys_updated_on property value.
func (m *AccountModel) GetSysUpdatedOn() (*time.Time, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *time.Time](backingStore, sys_updated_onKey)
}

// SetSysUpdatedOn sets the sys_updated_on property value.
func (m *AccountModel) SetSysUpdatedOn(val *time.Time) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sys_updated_onKey, val)
}

// GetTheme returns the theme property value.
func (m *AccountModel) GetTheme() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, themeKey)
}

// SetTheme sets the theme property value.
func (m *AccountModel) SetTheme(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, themeKey, val)
}

// GetVendor returns the vendor property value.
func (m *AccountModel) GetVendor() (*bool, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](backingStore, vendorKey)
}

// SetVendor sets the vendor property value.
func (m *AccountModel) SetVendor(val *bool) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, vendorKey, val)
}

// GetVendorManager returns the vendor_manager property value.
func (m *AccountModel) GetVendorManager() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, vendor_managerKey)
}

// SetVendorManager sets the vendor_manager property value.
func (m *AccountModel) SetVendorManager(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, vendor_managerKey, val)
}

// GetVendorType returns the vendor_type property value.
func (m *AccountModel) GetVendorType() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, vendor_typeKey)
}

// SetVendorType sets the vendor_type property value.
func (m *AccountModel) SetVendorType(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, vendor_typeKey, val)
}

// GetWebsite returns the website property value.
func (m *AccountModel) GetWebsite() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, websiteKey)
}

// SetWebsite sets the website property value.
func (m *AccountModel) SetWebsite(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, websiteKey, val)
}

// GetZip returns the zip property value.
func (m *AccountModel) GetZip() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, zipKey)
}

// SetZip sets the zip property value.
func (m *AccountModel) SetZip(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, zipKey, val)
}
