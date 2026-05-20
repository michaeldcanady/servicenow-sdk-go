package cdmchangesetapi

const (
	// Shared keys
	sysIdKey  = "sys_id"
	numberKey = "number"
	stateKey  = "state"
	typeKey   = "type"
	linkKey   = "link"
	valueKey  = "value"
	nodeKey   = "node"
	nameKey   = "name"

	// Audit keys
	sysCreatedByKey = "sys_created_by"
	sysCreatedOnKey = "sys_created_on"
	sysUpdatedByKey = "sys_updated_by"
	sysUpdatedOnKey = "sys_updated_on"

	// Changeset keys
	autoValidateKey              = "auto_validate"
	cdmApplicationKey            = "cdm_application"
	committedAtKey               = "committed_at"
	committedByKey               = "committed_by"
	descriptionKey               = "description"
	lastConflictDetectionTimeKey = "last_conflict_detection_time"
	publishOptionKey             = "publish_option"
	titleKey                     = "title"

	// Activity keys
	changesetIdKey = "changeset_id"
	conflictKey    = "conflict"
	namePathKey    = "name_path"
	newNameKey     = "new_name"
	oldNameKey     = "old_name"
	newValueKey    = "new_value"
	oldValueKey    = "old_value"
	secureKey      = "secure"

	// Impacted Shared Component keys
	cdmSharedLibraryKey = "cdm_shared__library"
	nodeMainKey         = "node_main"
	versionCounterKey   = "version_counter"

	// Impacted Deployable keys (Query-based)
	cdiCountKey               = "cdi_count"
	cdiUsageKey               = "cdi_usage"
	cdmAppKey                 = "cdm_app"
	cdmCiKey                  = "cdm_ci"
	environmentTypeKey        = "environment_type"
	snapshotVersionCounterKey = "snapshot_version_counter"

	// Impacted Deployable keys (Path-based)
	conflictTypeKey   = "conflict_type"
	effectiveFromKey  = "effective_from"
	effectiveToKey    = "effective_to"
	levelKey          = "level"
	linkedToKey       = "linked_to"
	mainIdKey         = "main_id"
	mainIdEncodedKey  = "main_id_encoded"
	nodeClassifierKey = "node_classifier"
	statusKey         = "status"
	secureValueKey    = "secure_value"
)
