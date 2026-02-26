package snippets

import (
	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
)

func _() {
    var client *servicenowsdkgo.ServiceNowClient

	// [START fluent_table]
	client.Now2().TableV2("xSDK_SN_TABLEx")
	// [END fluent_table]

	// [START fluent_attachment]
	client.Now2().Attachment2()
	// [END fluent_attachment]

	// [START fluent_batch]
	client.Now2().Batch()
	// [END fluent_batch]
}
