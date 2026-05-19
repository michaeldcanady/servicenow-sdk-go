package accountapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

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
type Account interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetBannerImageLight() (*string, error)
	setBannerImageLight(*string) error
	GetCountry() (*string, error)
	setCountry(*string) error
	GetParent() (*string, error)
	setParent(*string) error
	GetNotes() (*string, error)
	setNotes(*string) error
	GetStockSymbol() (*string, error)
	setStockSymbol(*string) error
	GetDiscount() (*string, error)
	setDiscount(*string) error
	GetActiveEscalation() (*string, error)
	setActiveEscalation(*string) error
	GetSysUpdatedOn() (*string, error)
	setSysUpdatedOn(*string) error
	GetAppleIcon() (*string, error)
	setAppleIcon(*string) error
	GetNumber() (*string, error)
	setNumber(*string) error
	GetSysUpdatedBy() (*string, error)
	setSysUpdatedBy(*string) error
	GetFiscalYear() (*string, error)
	setFiscalYear(*string) error
	GetSysCreatedOn() (*string, error)
	setSysCreatedOn(*string) error
	GetContact() (*string, error)
	setContact(*string) error
	GetStockPrice() (*string, error)
	setStockPrice(*string) error
	GetState() (*string, error)
	setState(*string) error
	GetBannerImage() (*string, error)
	setBannerImage(*string) error
	GetSysCreatedBy() (*string, error)
	setSysCreatedBy(*string) error
	GetLongitude() (*string, error)
	setLongitude(*string) error
	GetZip() (*string, error)
	setZip(*string) error
	GetProfits() (*string, error)
	setProfits(*string) error
	GetPhone() (*string, error)
	setPhone(*string) error
	GetFaxPhone() (*string, error)
	setFaxPhone(*string) error
	GetName() (*string, error)
	setName(*string) error
	GetBannerText() (*string, error)
	setBannerText(*string) error
	GetAccountCode() (*string, error)
	setAccountCode(*string) error
	GetPrimary() (*string, error)
	setPrimary(*string) error
	GetCity() (*string, error)
	setCity(*string) error
	GetLatitude() (*string, error)
	setLatitude(*string) error
	GetSysClassName() (*string, error)
	setSysClassName(*string) error
	GetManufacturer() (*string, error)
	setManufacturer(*string) error
	GetAccountParent() (*string, error)
	setAccountParent(*string) error
	GetSysID() (*string, error)
	setSysID(*string) error
	GetMarketCap() (*string, error)
	setMarketCap(*string) error
	GetNumEmployees() (*string, error)
	setNumEmployees(*string) error
	GetRankTier() (*string, error)
	setRankTier(*string) error
	GetStreet() (*string, error)
	setStreet(*string) error
	GetVendor() (*string, error)
	setVendor(*string) error
	GetLatLongError() (*string, error)
	setLatLongError(*string) error
	GetTheme() (*string, error)
	setTheme(*string) error
	GetVendorType() (*string, error)
	setVendorType(*string) error
	GetWebsite() (*string, error)
	setWebsite(*string) error
	GetRevenuePerYear() (*string, error)
	setRevenuePerYear(*string) error
	GetPubliclyTraded() (*string, error)
	setPubliclyTraded(*string) error
	GetSysModCount() (*string, error)
	setSysModCount(*string) error
	GetSysTags() (*string, error)
	setSysTags(*string) error
	GetPartner() (*string, error)
	setPartner(*string) error
	GetRegistrationCode() (*string, error)
	setRegistrationCode(*string) error
	GetVendorManager() (*string, error)
	setVendorManager(*string) error
	GetAccountPath() (*string, error)
	setAccountPath(*string) error
	GetPrimaryContact() (*string, error)
	setPrimaryContact(*string) error
	GetCustomer() (*string, error)
	setCustomer(*string) error
}

// AccountModel implementation of Account
type AccountModel struct {
	newInternal.BaseModel
}

// NewAccount creates a new instance of AccountModel
func NewAccount() *AccountModel {
	return &AccountModel{
		BaseModel: *newInternal.NewBaseModel(),
	}
}

// CreateAccountFromDiscriminatorValue is a factory for creating an Account model.
func CreateAccountFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewAccount(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *AccountModel) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}

	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(bannerImageLightKey)(m.GetBannerImageLight),
		internalSerialization.SerializeStringFunc(countryKey)(m.GetCountry),
		internalSerialization.SerializeStringFunc(parentKey)(m.GetParent),
		internalSerialization.SerializeStringFunc(notesKey)(m.GetNotes),
		internalSerialization.SerializeStringFunc(stockSymbolKey)(m.GetStockSymbol),
		internalSerialization.SerializeStringFunc(discountKey)(m.GetDiscount),
		internalSerialization.SerializeStringFunc(activeEscalationKey)(m.GetActiveEscalation),
		internalSerialization.SerializeStringFunc(sysUpdatedOnKey)(m.GetSysUpdatedOn),
		internalSerialization.SerializeStringFunc(appleIconKey)(m.GetAppleIcon),
		internalSerialization.SerializeStringFunc(numberKey)(m.GetNumber),
		internalSerialization.SerializeStringFunc(sysUpdatedByKey)(m.GetSysUpdatedBy),
		internalSerialization.SerializeStringFunc(fiscalYearKey)(m.GetFiscalYear),
		internalSerialization.SerializeStringFunc(sysCreatedOnKey)(m.GetSysCreatedOn),
		internalSerialization.SerializeStringFunc(contactKey)(m.GetContact),
		internalSerialization.SerializeStringFunc(stockPriceKey)(m.GetStockPrice),
		internalSerialization.SerializeStringFunc(stateKey)(m.GetState),
		internalSerialization.SerializeStringFunc(bannerImageKey)(m.GetBannerImage),
		internalSerialization.SerializeStringFunc(sysCreatedByKey)(m.GetSysCreatedBy),
		internalSerialization.SerializeStringFunc(longitudeKey)(m.GetLongitude),
		internalSerialization.SerializeStringFunc(zipKey)(m.GetZip),
		internalSerialization.SerializeStringFunc(profitsKey)(m.GetProfits),
		internalSerialization.SerializeStringFunc(phoneKey)(m.GetPhone),
		internalSerialization.SerializeStringFunc(faxPhoneKey)(m.GetFaxPhone),
		internalSerialization.SerializeStringFunc(nameKey)(m.GetName),
		internalSerialization.SerializeStringFunc(bannerTextKey)(m.GetBannerText),
		internalSerialization.SerializeStringFunc(accountCodeKey)(m.GetAccountCode),
		internalSerialization.SerializeStringFunc(primaryKey)(m.GetPrimary),
		internalSerialization.SerializeStringFunc(cityKey)(m.GetCity),
		internalSerialization.SerializeStringFunc(latitudeKey)(m.GetLatitude),
		internalSerialization.SerializeStringFunc(sysClassNameKey)(m.GetSysClassName),
		internalSerialization.SerializeStringFunc(manufacturerKey)(m.GetManufacturer),
		internalSerialization.SerializeStringFunc(accountParentKey)(m.GetAccountParent),
		internalSerialization.SerializeStringFunc(sysIDKey)(m.GetSysID),
		internalSerialization.SerializeStringFunc(marketCapKey)(m.GetMarketCap),
		internalSerialization.SerializeStringFunc(numEmployeesKey)(m.GetNumEmployees),
		internalSerialization.SerializeStringFunc(rankTierKey)(m.GetRankTier),
		internalSerialization.SerializeStringFunc(streetKey)(m.GetStreet),
		internalSerialization.SerializeStringFunc(vendorKey)(m.GetVendor),
		internalSerialization.SerializeStringFunc(latLongErrorKey)(m.GetLatLongError),
		internalSerialization.SerializeStringFunc(themeKey)(m.GetTheme),
		internalSerialization.SerializeStringFunc(vendorTypeKey)(m.GetVendorType),
		internalSerialization.SerializeStringFunc(websiteKey)(m.GetWebsite),
		internalSerialization.SerializeStringFunc(revenuePerYearKey)(m.GetRevenuePerYear),
		internalSerialization.SerializeStringFunc(publiclyTradedKey)(m.GetPubliclyTraded),
		internalSerialization.SerializeStringFunc(sysModCountKey)(m.GetSysModCount),
		internalSerialization.SerializeStringFunc(sysTagsKey)(m.GetSysTags),
		internalSerialization.SerializeStringFunc(partnerKey)(m.GetPartner),
		internalSerialization.SerializeStringFunc(registrationCodeKey)(m.GetRegistrationCode),
		internalSerialization.SerializeStringFunc(vendorManagerKey)(m.GetVendorManager),
		internalSerialization.SerializeStringFunc(accountPathKey)(m.GetAccountPath),
		internalSerialization.SerializeStringFunc(primaryContactKey)(m.GetPrimaryContact),
		internalSerialization.SerializeStringFunc(customerKey)(m.GetCustomer),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *AccountModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		bannerImageLightKey: internalSerialization.DeserializeStringFunc()(m.setBannerImageLight),
		countryKey:          internalSerialization.DeserializeStringFunc()(m.setCountry),
		parentKey:           internalSerialization.DeserializeStringFunc()(m.setParent),
		notesKey:            internalSerialization.DeserializeStringFunc()(m.setNotes),
		stockSymbolKey:      internalSerialization.DeserializeStringFunc()(m.setStockSymbol),
		discountKey:         internalSerialization.DeserializeStringFunc()(m.setDiscount),
		activeEscalationKey: internalSerialization.DeserializeStringFunc()(m.setActiveEscalation),
		sysUpdatedOnKey:     internalSerialization.DeserializeStringFunc()(m.setSysUpdatedOn),
		appleIconKey:        internalSerialization.DeserializeStringFunc()(m.setAppleIcon),
		numberKey:           internalSerialization.DeserializeStringFunc()(m.setNumber),
		sysUpdatedByKey:     internalSerialization.DeserializeStringFunc()(m.setSysUpdatedBy),
		fiscalYearKey:       internalSerialization.DeserializeStringFunc()(m.setFiscalYear),
		sysCreatedOnKey:     internalSerialization.DeserializeStringFunc()(m.setSysCreatedOn),
		contactKey:          internalSerialization.DeserializeStringFunc()(m.setContact),
		stockPriceKey:       internalSerialization.DeserializeStringFunc()(m.setStockPrice),
		stateKey:            internalSerialization.DeserializeStringFunc()(m.setState),
		bannerImageKey:      internalSerialization.DeserializeStringFunc()(m.setBannerImage),
		sysCreatedByKey:     internalSerialization.DeserializeStringFunc()(m.setSysCreatedBy),
		longitudeKey:        internalSerialization.DeserializeStringFunc()(m.setLongitude),
		zipKey:              internalSerialization.DeserializeStringFunc()(m.setZip),
		profitsKey:          internalSerialization.DeserializeStringFunc()(m.setProfits),
		phoneKey:            internalSerialization.DeserializeStringFunc()(m.setPhone),
		faxPhoneKey:         internalSerialization.DeserializeStringFunc()(m.setFaxPhone),
		nameKey:             internalSerialization.DeserializeStringFunc()(m.setName),
		bannerTextKey:       internalSerialization.DeserializeStringFunc()(m.setBannerText),
		accountCodeKey:      internalSerialization.DeserializeStringFunc()(m.setAccountCode),
		primaryKey:          internalSerialization.DeserializeStringFunc()(m.setPrimary),
		cityKey:             internalSerialization.DeserializeStringFunc()(m.setCity),
		latitudeKey:         internalSerialization.DeserializeStringFunc()(m.setLatitude),
		sysClassNameKey:     internalSerialization.DeserializeStringFunc()(m.setSysClassName),
		manufacturerKey:     internalSerialization.DeserializeStringFunc()(m.setManufacturer),
		accountParentKey:    internalSerialization.DeserializeStringFunc()(m.setAccountParent),
		sysIDKey:            internalSerialization.DeserializeStringFunc()(m.setSysID),
		marketCapKey:        internalSerialization.DeserializeStringFunc()(m.setMarketCap),
		numEmployeesKey:     internalSerialization.DeserializeStringFunc()(m.setNumEmployees),
		rankTierKey:         internalSerialization.DeserializeStringFunc()(m.setRankTier),
		streetKey:           internalSerialization.DeserializeStringFunc()(m.setStreet),
		vendorKey:           internalSerialization.DeserializeStringFunc()(m.setVendor),
		latLongErrorKey:     internalSerialization.DeserializeStringFunc()(m.setLatLongError),
		themeKey:            internalSerialization.DeserializeStringFunc()(m.setTheme),
		vendorTypeKey:       internalSerialization.DeserializeStringFunc()(m.setVendorType),
		websiteKey:          internalSerialization.DeserializeStringFunc()(m.setWebsite),
		revenuePerYearKey:   internalSerialization.DeserializeStringFunc()(m.setRevenuePerYear),
		publiclyTradedKey:   internalSerialization.DeserializeStringFunc()(m.setPubliclyTraded),
		sysModCountKey:      internalSerialization.DeserializeStringFunc()(m.setSysModCount),
		sysTagsKey:          internalSerialization.DeserializeStringFunc()(m.setSysTags),
		partnerKey:          internalSerialization.DeserializeStringFunc()(m.setPartner),
		registrationCodeKey: internalSerialization.DeserializeStringFunc()(m.setRegistrationCode),
		vendorManagerKey:    internalSerialization.DeserializeStringFunc()(m.setVendorManager),
		accountPathKey:      internalSerialization.DeserializeStringFunc()(m.setAccountPath),
		primaryContactKey:   internalSerialization.DeserializeStringFunc()(m.setPrimaryContact),
		customerKey:         internalSerialization.DeserializeStringFunc()(m.setCustomer),
	}
}

// Getters and Setters...
func (m *AccountModel) GetBannerImageLight() (*string, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), bannerImageLightKey) }
func (m *AccountModel) setBannerImageLight(val *string) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), bannerImageLightKey, val) }
func (m *AccountModel) GetCountry() (*string, error)          { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), countryKey) }
func (m *AccountModel) setCountry(val *string) error          { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), countryKey, val) }
func (m *AccountModel) GetParent() (*string, error)           { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), parentKey) }
func (m *AccountModel) setParent(val *string) error           { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), parentKey, val) }
func (m *AccountModel) GetNotes() (*string, error)            { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), notesKey) }
func (m *AccountModel) setNotes(val *string) error            { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), notesKey, val) }
func (m *AccountModel) GetStockSymbol() (*string, error)      { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), stockSymbolKey) }
func (m *AccountModel) setStockSymbol(val *string) error      { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), stockSymbolKey, val) }
func (m *AccountModel) GetDiscount() (*string, error)         { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), discountKey) }
func (m *AccountModel) setDiscount(val *string) error         { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), discountKey, val) }
func (m *AccountModel) GetActiveEscalation() (*string, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), activeEscalationKey) }
func (m *AccountModel) setActiveEscalation(val *string) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), activeEscalationKey, val) }
func (m *AccountModel) GetSysUpdatedOn() (*string, error)     { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysUpdatedOnKey) }
func (m *AccountModel) setSysUpdatedOn(val *string) error     { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sysUpdatedOnKey, val) }
func (m *AccountModel) GetAppleIcon() (*string, error)        { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), appleIconKey) }
func (m *AccountModel) setAppleIcon(val *string) error        { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), appleIconKey, val) }
func (m *AccountModel) GetNumber() (*string, error)           { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), numberKey) }
func (m *AccountModel) setNumber(val *string) error           { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), numberKey, val) }
func (m *AccountModel) GetSysUpdatedBy() (*string, error)     { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysUpdatedByKey) }
func (m *AccountModel) setSysUpdatedBy(val *string) error     { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sysUpdatedByKey, val) }
func (m *AccountModel) GetFiscalYear() (*string, error)       { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), fiscalYearKey) }
func (m *AccountModel) setFiscalYear(val *string) error       { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), fiscalYearKey, val) }
func (m *AccountModel) GetSysCreatedOn() (*string, error)     { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysCreatedOnKey) }
func (m *AccountModel) setSysCreatedOn(val *string) error     { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sysCreatedOnKey, val) }
func (m *AccountModel) GetContact() (*string, error)          { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), contactKey) }
func (m *AccountModel) setContact(val *string) error          { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), contactKey, val) }
func (m *AccountModel) GetStockPrice() (*string, error)       { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), stockPriceKey) }
func (m *AccountModel) setStockPrice(val *string) error       { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), stockPriceKey, val) }
func (m *AccountModel) GetState() (*string, error)            { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), stateKey) }
func (m *AccountModel) setState(val *string) error            { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), stateKey, val) }
func (m *AccountModel) GetBannerImage() (*string, error)      { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), bannerImageKey) }
func (m *AccountModel) setBannerImage(val *string) error      { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), bannerImageKey, val) }
func (m *AccountModel) GetSysCreatedBy() (*string, error)     { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysCreatedByKey) }
func (m *AccountModel) setSysCreatedBy(val *string) error     { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sysCreatedByKey, val) }
func (m *AccountModel) GetLongitude() (*string, error)        { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), longitudeKey) }
func (m *AccountModel) setLongitude(val *string) error        { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), longitudeKey, val) }
func (m *AccountModel) GetZip() (*string, error)              { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), zipKey) }
func (m *AccountModel) setZip(val *string) error              { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), zipKey, val) }
func (m *AccountModel) GetProfits() (*string, error)          { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), profitsKey) }
func (m *AccountModel) setProfits(val *string) error          { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), profitsKey, val) }
func (m *AccountModel) GetPhone() (*string, error)            { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), phoneKey) }
func (m *AccountModel) setPhone(val *string) error            { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), phoneKey, val) }
func (m *AccountModel) GetFaxPhone() (*string, error)         { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), faxPhoneKey) }
func (m *AccountModel) setFaxPhone(val *string) error         { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), faxPhoneKey, val) }
func (m *AccountModel) GetName() (*string, error)             { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), nameKey) }
func (m *AccountModel) setName(val *string) error             { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), nameKey, val) }
func (m *AccountModel) GetBannerText() (*string, error)       { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), bannerTextKey) }
func (m *AccountModel) setBannerText(val *string) error       { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), bannerTextKey, val) }
func (m *AccountModel) GetAccountCode() (*string, error)      { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), accountCodeKey) }
func (m *AccountModel) setAccountCode(val *string) error      { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), accountCodeKey, val) }
func (m *AccountModel) GetPrimary() (*string, error)          { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), primaryKey) }
func (m *AccountModel) setPrimary(val *string) error          { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), primaryKey, val) }
func (m *AccountModel) GetCity() (*string, error)             { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), cityKey) }
func (m *AccountModel) setCity(val *string) error             { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), cityKey, val) }
func (m *AccountModel) GetLatitude() (*string, error)         { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), latitudeKey) }
func (m *AccountModel) setLatitude(val *string) error         { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), latitudeKey, val) }
func (m *AccountModel) GetSysClassName() (*string, error)     { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysClassNameKey) }
func (m *AccountModel) setSysClassName(val *string) error     { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sysClassNameKey, val) }
func (m *AccountModel) GetManufacturer() (*string, error)     { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), manufacturerKey) }
func (m *AccountModel) setManufacturer(val *string) error     { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), manufacturerKey, val) }
func (m *AccountModel) GetAccountParent() (*string, error)    { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), accountParentKey) }
func (m *AccountModel) setAccountParent(val *string) error    { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), accountParentKey, val) }
func (m *AccountModel) GetSysID() (*string, error)            { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysIDKey) }
func (m *AccountModel) setSysID(val *string) error            { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sysIDKey, val) }
func (m *AccountModel) GetMarketCap() (*string, error)        { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), marketCapKey) }
func (m *AccountModel) setMarketCap(val *string) error        { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), marketCapKey, val) }
func (m *AccountModel) GetNumEmployees() (*string, error)     { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), numEmployeesKey) }
func (m *AccountModel) setNumEmployees(val *string) error     { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), numEmployeesKey, val) }
func (m *AccountModel) GetRankTier() (*string, error)         { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), rankTierKey) }
func (m *AccountModel) setRankTier(val *string) error         { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), rankTierKey, val) }
func (m *AccountModel) GetStreet() (*string, error)           { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), streetKey) }
func (m *AccountModel) setStreet(val *string) error           { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), streetKey, val) }
func (m *AccountModel) GetVendor() (*string, error)           { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), vendorKey) }
func (m *AccountModel) setVendor(val *string) error           { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), vendorKey, val) }
func (m *AccountModel) GetLatLongError() (*string, error)     { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), latLongErrorKey) }
func (m *AccountModel) setLatLongError(val *string) error     { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), latLongErrorKey, val) }
func (m *AccountModel) GetTheme() (*string, error)            { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), themeKey) }
func (m *AccountModel) setTheme(val *string) error            { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), themeKey, val) }
func (m *AccountModel) GetVendorType() (*string, error)       { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), vendorTypeKey) }
func (m *AccountModel) setVendorType(val *string) error       { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), vendorTypeKey, val) }
func (m *AccountModel) GetWebsite() (*string, error)          { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), websiteKey) }
func (m *AccountModel) setWebsite(val *string) error          { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), websiteKey, val) }
func (m *AccountModel) GetRevenuePerYear() (*string, error)   { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), revenuePerYearKey) }
func (m *AccountModel) setRevenuePerYear(val *string) error   { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), revenuePerYearKey, val) }
func (m *AccountModel) GetPubliclyTraded() (*string, error)   { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), publiclyTradedKey) }
func (m *AccountModel) setPubliclyTraded(val *string) error   { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), publiclyTradedKey, val) }
func (m *AccountModel) GetSysModCount() (*string, error)      { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysModCountKey) }
func (m *AccountModel) setSysModCount(val *string) error      { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sysModCountKey, val) }
func (m *AccountModel) GetSysTags() (*string, error)          { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysTagsKey) }
func (m *AccountModel) setSysTags(val *string) error          { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sysTagsKey, val) }
func (m *AccountModel) GetPartner() (*string, error)          { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), partnerKey) }
func (m *AccountModel) setPartner(val *string) error          { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), partnerKey, val) }
func (m *AccountModel) GetRegistrationCode() (*string, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), registrationCodeKey) }
func (m *AccountModel) setRegistrationCode(val *string) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), registrationCodeKey, val) }
func (m *AccountModel) GetVendorManager() (*string, error)    { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), vendorManagerKey) }
func (m *AccountModel) setVendorManager(val *string) error    { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), vendorManagerKey, val) }
func (m *AccountModel) GetAccountPath() (*string, error)      { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), accountPathKey) }
func (m *AccountModel) setAccountPath(val *string) error      { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), accountPathKey, val) }
func (m *AccountModel) GetPrimaryContact() (*string, error)   { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), primaryContactKey) }
func (m *AccountModel) setPrimaryContact(val *string) error   { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), primaryContactKey, val) }
func (m *AccountModel) GetCustomer() (*string, error)         { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), customerKey) }
func (m *AccountModel) setCustomer(val *string) error         { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), customerKey, val) }
