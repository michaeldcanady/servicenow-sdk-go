package tableapi

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

const keyName = "key1"

func TestTableEntry(t *testing.T) {
	responseJSON := `{
			"parent": "",
			"made_sla": "true",
			"watch_list": "",
			"upon_reject": "cancel",
			"sys_updated_on": "2016-01-19 04:52:04",
			"approval_history": "",
			"number": "PRB0000050",
			"sys_updated_by": "glide.maint",
			"opened_by": {
			  "link": "https://instance.servicenow.com/api/now/table/sys_user/glide.maint",
			  "value": "glide.maint"
			},
			"user_input": "",
			"sys_created_on": "2016-01-19 04:51:19",
			"sys_domain": {
			  "link": "https://instance.servicenow.com/api/now/table/sys_user_group/global",
			  "value": "global"
			},
			"state": "4",
			"sys_created_by": "glide.maint",
			"knowledge": "false",
			"order": "",
			"closed_at": "2016-01-19 04:52:04",
			"cmdb_ci": {
			  "link": "https://instance.servicenow.com/api/now/table/cmdb_ci/55b35562c0a8010e01cff22378e0aea9",
			  "value": "55b35562c0a8010e01cff22378e0aea9"
			},
			"delivery_plan": "",
			"impact": "3",
			"active": "false",
			"work_notes_list": "",
			"business_service": "",
			"priority": "4",
			"sys_domain_path": "/",
			"time_worked": "",
			"expected_start": "",
			"rejection_goto": "",
			"opened_at": "2016-01-19 04:49:47",
			"business_duration": "1970-01-01 00:00:00",
			"group_list": "",
			"work_end": "",
			"approval_set": "",
			"wf_activity": "",
			"work_notes": "",
			"short_description": "Switch occasionally drops connections",
			"correlation_display": "",
			"delivery_task": "",
			"work_start": "",
			"assignment_group": "",
			"additional_assignee_list": "",
			"description": "Switch occasionally drops connections",
			"calendar_duration": "1970-01-01 00:02:17",
			"close_notes": "updated firmware",
			"sys_class_name": "problem",
			"closed_by": "",
			"follow_up": "",
			"sys_id": "04ce72c9c0a8016600b5b7f75ac67b5b",
			"contact_type": "phone",
			"urgency": "3",
			"company": "",
			"reassignment_count": "",
			"activity_due": "",
			"assigned_to": "",
			"comments": "",
			"approval": "not requested",
			"sla_due": "",
			"comments_and_work_notes": "",
			"due_date": "",
			"sys_mod_count": "1",
			"sys_tags": "",
			"escalation": "0",
			"upon_approval": "proceed",
			"correlation_id": "",
			"location": ""
		  }`

	expected := TableEntry{
		"parent":           "",
		"made_sla":         "true",
		"watch_list":       "",
		"upon_reject":      "cancel",
		"sys_updated_on":   "2016-01-19 04:52:04",
		"approval_history": "",
		"number":           "PRB0000050",
		"sys_updated_by":   "glide.maint",
		"opened_by": map[string]interface{}{
			"link":  "https://instance.servicenow.com/api/now/table/sys_user/glide.maint",
			"value": "glide.maint",
		},
		"user_input":     "",
		"sys_created_on": "2016-01-19 04:51:19",
		"sys_domain": map[string]interface{}{
			"link":  "https://instance.servicenow.com/api/now/table/sys_user_group/global",
			"value": "global",
		},
		"state":          "4",
		"sys_created_by": "glide.maint",
		"knowledge":      "false",
		"order":          "",
		"closed_at":      "2016-01-19 04:52:04",
		"cmdb_ci": map[string]interface{}{
			"link":  "https://instance.servicenow.com/api/now/table/cmdb_ci/55b35562c0a8010e01cff22378e0aea9",
			"value": "55b35562c0a8010e01cff22378e0aea9",
		},
		"delivery_plan":            "",
		"impact":                   "3",
		"active":                   "false",
		"work_notes_list":          "",
		"business_service":         "",
		"priority":                 "4",
		"sys_domain_path":          "/",
		"time_worked":              "",
		"expected_start":           "",
		"rejection_goto":           "",
		"opened_at":                "2016-01-19 04:49:47",
		"business_duration":        "1970-01-01 00:00:00",
		"group_list":               "",
		"work_end":                 "",
		"approval_set":             "",
		"wf_activity":              "",
		"work_notes":               "",
		"short_description":        "Switch occasionally drops connections",
		"correlation_display":      "",
		"delivery_task":            "",
		"work_start":               "",
		"assignment_group":         "",
		"additional_assignee_list": "",
		"description":              "Switch occasionally drops connections",
		"calendar_duration":        "1970-01-01 00:02:17",
		"close_notes":              "updated firmware",
		"sys_class_name":           "problem",
		"closed_by":                "",
		"follow_up":                "",
		"sys_id":                   "04ce72c9c0a8016600b5b7f75ac67b5b",
		"contact_type":             "phone",
		"urgency":                  "3",
		"company":                  "",
		"reassignment_count":       "",
		"activity_due":             "",
		"assigned_to":              "",
		"comments":                 "",
		"approval":                 "not requested",
		"sla_due":                  "",
		"comments_and_work_notes":  "",
		"due_date":                 "",
		"sys_mod_count":            "1",
		"sys_tags":                 "",
		"escalation":               "0",
		"upon_approval":            "proceed",
		"correlation_id":           "",
		"location":                 "",
	}

	var entry TableEntry

	err := json.Unmarshal([]byte(responseJSON), &entry)
	if err != nil {
		t.Errorf("Expected nil got '%s'", err)
	}

	assert.Equal(t, expected, entry)
}

func TestTableEntryValueValidKey(t *testing.T) {
	entry := TableEntry{
		keyName: "value1",
	}

	value := entry.Value(keyName)

	assert.NotNil(t, value)
	assert.Equal(t, value, &TableValue{value: "value1"})

	entry = TableEntry{
		keyName: map[string]interface{}{
			"link":  "https://instance.servicenow.com/api/now/table/cmdb_ci/55b35562c0a8010e01cff22378e0aea9",
			"value": "55b35562c0a8010e01cff22378e0aea9",
		},
	}

	value = entry.Value(keyName)

	assert.NotNil(t, value)
	assert.Equal(t, value, &TableValue{value: "55b35562c0a8010e01cff22378e0aea9"})
}

func TestTableEntryValueMissingKey(t *testing.T) {
	entry := TableEntry{
		"key2": "value1",
	}

	value := entry.Value(keyName)

	assert.Nil(t, value)
}

func TestTableEntryKeys(t *testing.T) {
	entry := TableEntry{
		keyName: "value2",
		"key2":  "value1",
	}

	keys := entry.Keys()

	assert.Contains(t, keys, keyName)
	assert.Contains(t, keys, "key2")
}
