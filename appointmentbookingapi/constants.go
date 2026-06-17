package appointmentbookingapi

const (
	// Shared keys
	activeKey            = "active"
	activeStringKey      = "active_string"
	catalogIdKey         = "catalogId"
	locationKey          = "location"
	openedForKey         = "openedFor"
	serviceConfigRuleKey = "service_cofig_rule"
	taskIdKey            = "taskId"
	taskTableKey         = "taskTable"

	// AppointmentRequest specific keys
	actualEndDateKey   = "actualEndDate"
	actualStartDateKey = "actualStartDate"
	endDateUTCKey      = "endDateUTC"
	rescheduleKey      = "reschedule"
	startDateUTCKey    = "startDateUTC"
	timezoneKey        = "timezone"
	validateRequestKey = "validate_request"

	// AvailabilityRequest specific keys
	endDateKey              = "end_date"
	fetchDaysSlotKey        = "fetch_days_slot"
	fullDayKey              = "full_day"
	getNextAvailableSlotKey = "get_next_available_slot"
	limitKey                = "limit"
	otherInputsKey          = "otherInputs"
	startDateKey            = "start_date"
	useReadReplicaKey       = "use_read_replica"
	viewKey                 = "view"

	// ConfigurationResult specific keys
	advancedCalendarViewPortalKey = "advanced_calendar_view_portal"
	autoAcceptanceKey             = "auto_acceptance"
	localeLanguageKey             = "locale_language"
	serviceConfigKey              = "service_config"
	translationsKey               = "translations"
	userDateFormatOptionsKey      = "userDateFormatOptions"
	useRRKey                      = "useRR"
	userTimeFormatKey             = "userTimeFormat"
	userTimeFormatOptionsKey      = "userTimeFormatOptions"
	viewScaleKey                  = "view_scale"

	// ServiceConfig specific keys
	appointmentBookingConfigKey    = "appointment_booking_config"
	appointmentDurationKey         = "appointment_duration"
	appointmentsPerBookableSlotKey = "appointments_per_bookable_slot"
	bookableDaysKey                = "bookable_days"
	cancelByTimeKey                = "cancel_by_time"
	defaultTimezoneKey             = "default_timezone"
	enableAdvancedConfigKey        = "enable_advanced_config"
	fieldMappingKey                = "field_mapping"
	futureBookableMaxDaysKey       = "future_bookable_max_days"
	leadTimeKey                    = "lead_time"
	mandatoryKey                   = "mandatory"
	useSlotEndTimeAsKey            = "use_slot_end_time_as"
	workDurationKey                = "work_duration"

	// FieldMapping specific keys
	contactKey            = "contact"
	contactRPVariableKey  = "contactRPVariable"
	locationRPVariableKey = "locationRPVariable"

	// RPVariable specific keys
	displayNameKey = "displayName"
	labelKey       = "label"
	nameKey        = "name"

	// UserDateFormatOptions specific keys
	dayKey     = "day"
	monthKey   = "month"
	weekKey    = "week"
	weekdayKey = "weekday"

	// UserTimeFormat specific keys
	typeKey  = "type"
	valueKey = "value"

	// UserTimeFormatOptions specific keys
	hourKey      = "hour"
	hourCycleKey = "hourCycle"
	minuteKey    = "minute"

	// AppointmentResult keys
	dataKey    = "data"
	messageKey = "message"
	reasonKey  = "reason"
	successKey = "success"

	// AvailabilityResult keys
	availabilityKey         = "availability"
	hasMoreKey              = "has_more"
	nextAvailableSlotKey    = "next_available_slot"
	noApptAvailableKey      = "no_appt_available"
	timeZoneDisplayValueKey = "time_zone_display_value"

	// CalendarResponse keys
	rangeEndKey   = "range_end"
	rangeStartKey = "range_start"

	// ExecuteRuleConditionsResult keys
	dedicatedCapacityKey     = "dedicatedCapacity"
	futureMaxBookableDaysKey = "futureMaxBookableDays"
	ruleIdKey                = "ruleId"
	ruleNameKey              = "ruleName"
)
