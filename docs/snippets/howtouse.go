package snippets

import (
	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
)

func _() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	// [START fluent_table]
	client.Now().Table("xSDK_SN_TABLEx")
	// [END fluent_table]

	// [START fluent_attachment]
	client.Now().Attachment()
	// [END fluent_attachment]

	// [START fluent_batch]
	client.Now().Batch()
	// [END fluent_batch]
}
