//go:build snippets

package snippets

import (
	"context"
	"log"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
	actsubapi "github.com/michaeldcanady/servicenow-sdk-go/actsubapi"
	appointmentbookingapi "github.com/michaeldcanady/servicenow-sdk-go/appointmentbookingapi"
	appserviceapi "github.com/michaeldcanady/servicenow-sdk-go/appserviceapi"
	caseapi "github.com/michaeldcanady/servicenow-sdk-go/caseapi"
	cdmapplicationsapi "github.com/michaeldcanady/servicenow-sdk-go/cdmapplicationsapi"
	cdmeditorapi "github.com/michaeldcanady/servicenow-sdk-go/cdmeditorapi"
	cmdbinstanceapi "github.com/michaeldcanady/servicenow-sdk-go/cmdbinstanceapi"
)

func ref_list_accounts() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_list_accounts]
	response, err := client.Now().Account().Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_list_accounts]
	_ = response
}

func ref_get_account() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_get_account]
	response, err := client.Now().Account().ByID("{accountID}").Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_get_account]
	_ = response
}

func ref_get_aggregates() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_get_aggregates]
	response, err := client.Now().Stats("incident").Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_get_aggregates]
	_ = response
}

func ref_list_activities() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_list_activities]
	response, err := client.Now().ActSub().Activities().Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_list_activities]
	_ = response
}

func ref_list_contexts() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_list_contexts]
	response, err := client.Now().ActSub().Contexts().Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_list_contexts]
	_ = response
}

func ref_list_subscription_objects() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_list_subscription_objects]
	response, err := client.Now().ActSub().SubObjects().Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_list_subscription_objects]
	_ = response
}

func ref_get_facet_instances() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_get_facet_instances]
	response, err := client.Now().ActSub().Facets().ByContext("{activityContext}").ByInstance("{contextInstance}").Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_get_facet_instances]
	_ = response
}

func ref_get_followings() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_get_followings]
	response, err := client.Now().ActSub().Followings().ByFollower("{follower}").Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_get_followings]
	_ = response
}

func ref_get_preference() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_get_preference]
	response, err := client.Now().ActSub().Preferences().ByProfileId("{profileID}").Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_get_preference]
	_ = response
}

func ref_create_preference() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_create_preference]
	// Build the request body
	var body *actsubapi.ActivitySubscriptionModel
	response, err := client.Now().ActSub().Preferences().Post(ctx, body, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_create_preference]
	_ = response
}

func ref_get_subscription() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_get_subscription]
	response, err := client.Now().ActSub().Subscriptions().BySubscriberId("{subscriberID}").Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_get_subscription]
	_ = response
}

func ref_check_subscription() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_check_subscription]
	response, err := client.Now().ActSub().Subscriptions().ByObjectId("{objectID}").IsSubscribed().Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_check_subscription]
	_ = response
}

func ref_subscribe() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_subscribe]
	// Build the request body
	var body *actsubapi.ActivitySubscriptionModel
	response, err := client.Now().ActSub().Subscriptions().ByObjectId("{objectID}").Subscribe().Post(ctx, body, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_subscribe]
	_ = response
}

func ref_unsubscribe() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_unsubscribe]
	if err := client.Now().ActSub().Subscriptions().ByObjectId("{objectID}").Unsubscribe().Delete(ctx, nil); err != nil {
		log.Fatal(err)
	}
	// [END ref_unsubscribe]
}

func ref_get_subscribers() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_get_subscribers]
	response, err := client.Now().ActSub().Subscribers().BySubObject("{subObject}").Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_get_subscribers]
	_ = response
}

func ref_get_user_stream() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_get_user_stream]
	response, err := client.Now().ActSub().UserStream().ByProfileId("{profileID}").Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_get_user_stream]
	_ = response
}

func ref_update_user_stream() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_update_user_stream]
	// Build the request body
	var body *actsubapi.ActivitySubscriptionModel
	response, err := client.Now().ActSub().UserStream().ByProfileId("{profileID}").Put(ctx, body, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_update_user_stream]
	_ = response
}

func ref_check_availability() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_check_availability]
	// Build the request body
	var body appointmentbookingapi.AvailabilityRequest
	response, err := client.AppointmentBooking().Availability().Post(ctx, body, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_check_availability]
	_ = response
}

func ref_book_appointment() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_book_appointment]
	// Build the request body
	var body appointmentbookingapi.AppointmentRequest
	response, err := client.AppointmentBooking().Appointment().Post(ctx, body, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_book_appointment]
	_ = response
}

func ref_get_calendar() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_get_calendar]
	response, err := client.AppointmentBooking().Calendar().Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_get_calendar]
	_ = response
}

func ref_get_configuration() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_get_configuration]
	response, err := client.AppointmentBooking().Configuration().Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_get_configuration]
	_ = response
}

func ref_execute_rule_conditions() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_execute_rule_conditions]
	// Build the request body
	var body appointmentbookingapi.ExecuteRuleConditionsRequest
	response, err := client.AppointmentBooking().ExecuteRuleConditions().Post(ctx, body, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_execute_rule_conditions]
	_ = response
}

func ref_request_user_window() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_request_user_window]
	// Build the request body
	body := map[string]any{ /* window inputs */ }
	response, err := client.AppointmentBooking().UserWindow().Post(ctx, body, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_request_user_window]
	_ = response
}

func ref_create_application_service() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_create_application_service]
	// Build the request body
	var body *appserviceapi.CreateServiceRequest
	response, err := client.Now().Cmdb().AppService().Create().Post(ctx, body, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_create_application_service]
	_ = response
}

func ref_find_service() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_find_service]
	response, err := client.Now().Cmdb().AppService().Csdm().FindService().Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_find_service]
	_ = response
}

func ref_register_service() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_register_service]
	// Build the request body
	var body *appserviceapi.RegisterServiceRequest
	response, err := client.Now().Cmdb().AppService().Csdm().RegisterService().Post(ctx, body, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_register_service]
	_ = response
}

func ref_populate_service() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_populate_service]
	// Build the request body
	var body *appserviceapi.PopulateServiceRequest
	response, err := client.Now().Cmdb().AppService().Csdm().ByID("{sysID}").PopulateService().Put(ctx, body, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_populate_service]
	_ = response
}

func ref_update_service_details() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_update_service_details]
	// Build the request body
	var body *appserviceapi.ServiceDetailsRequest
	response, err := client.Now().Cmdb().AppService().Csdm().ByID("{sysID}").ServiceDetails().Put(ctx, body, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_update_service_details]
	_ = response
}

func ref_list_cases() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_list_cases]
	response, err := client.CustomerService().Case().Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_list_cases]
	_ = response
}

func ref_create_case() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_create_case]
	// Build the request body
	var body caseapi.CaseResult
	response, err := client.CustomerService().Case().Post(ctx, body, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_create_case]
	_ = response
}

func ref_get_case() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_get_case]
	response, err := client.CustomerService().Case().ByID("{caseSysID}").Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_get_case]
	_ = response
}

func ref_update_case() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_update_case]
	// Build the request body
	var body caseapi.CaseResult
	response, err := client.CustomerService().Case().ByID("{caseSysID}").Put(ctx, body, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_update_case]
	_ = response
}

func ref_list_case_activities() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_list_case_activities]
	response, err := client.CustomerService().Case().ByID("{caseSysID}").Activities().Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_list_case_activities]
	_ = response
}

func ref_get_case_field_values() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_get_case_field_values]
	response, err := client.CustomerService().Case().FieldValues("{fieldName}").Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_get_case_field_values]
	_ = response
}

func ref_list_configuration_items() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_list_configuration_items]
	response, err := client.Now().Cmdb().Instance().ByClass("cmdb_ci_linux_server").Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_list_configuration_items]
	_ = response
}

func ref_create_configuration_item() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_create_configuration_item]
	// Build the request body
	var body cmdbinstanceapi.CmdbInstance
	response, err := client.Now().Cmdb().Instance().ByClass("cmdb_ci_linux_server").Post(ctx, body, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_create_configuration_item]
	_ = response
}

func ref_get_configuration_item() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_get_configuration_item]
	response, err := client.Now().Cmdb().Instance().ByClass("cmdb_ci_linux_server").ByID("{sysID}").Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_get_configuration_item]
	_ = response
}

func ref_update_configuration_item() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_update_configuration_item]
	// Build the request body
	var body cmdbinstanceapi.CmdbInstance
	response, err := client.Now().Cmdb().Instance().ByClass("cmdb_ci_linux_server").ByID("{sysID}").Put(ctx, body, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_update_configuration_item]
	_ = response
}

func ref_patch_configuration_item() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_patch_configuration_item]
	// Build the request body
	var body cmdbinstanceapi.CmdbInstance
	response, err := client.Now().Cmdb().Instance().ByClass("cmdb_ci_linux_server").ByID("{sysID}").Patch(ctx, body, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_patch_configuration_item]
	_ = response
}

func ref_create_relation() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_create_relation]
	// Build the request body
	var body cmdbinstanceapi.CmdbInstance
	response, err := client.Now().Cmdb().Instance().ByClass("cmdb_ci_linux_server").ByID("{sysID}").Relation().Post(ctx, body, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_create_relation]
	_ = response
}

func ref_delete_relation() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_delete_relation]
	if err := client.Now().Cmdb().Instance().ByClass("cmdb_ci_linux_server").ByID("{sysID}").Relation().ByID("{relSysID}").Delete(ctx, nil); err != nil {
		log.Fatal(err)
	}
	// [END ref_delete_relation]
}

func ref_explore_documents() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_explore_documents]
	response, err := client.Now().Documents().Explore().Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_explore_documents]
	_ = response
}

func ref_create_document() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_create_document]
	response, err := client.Now().Documents().Create().Post(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_create_document]
	_ = response
}

func ref_create_or_link_document() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_create_or_link_document]
	response, err := client.Now().Documents().CreateDocument().Post(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_create_or_link_document]
	_ = response
}

func ref_delete_document() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_delete_document]
	if err := client.Now().Documents().Delete().Delete(ctx, nil); err != nil {
		log.Fatal(err)
	}
	// [END ref_delete_document]
}

func ref_get_document_content() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_get_document_content]
	response, err := client.Now().Documents().Content("{documentSysID}").Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_get_document_content]
	_ = response
}

func ref_list_document_versions() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_list_document_versions]
	response, err := client.Now().Documents().Versions("{documentSysID}").Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_list_document_versions]
	_ = response
}

func ref_get_version_state() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_get_version_state]
	response, err := client.Now().Documents().VersionState("{versionSysID}").Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_get_version_state]
	_ = response
}

func ref_sync_down_document() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_sync_down_document]
	response, err := client.Now().Documents().SyncDown("{documentSysID}").Post(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_sync_down_document]
	_ = response
}

func ref_attach_document() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_attach_document]
	response, err := client.Now().Documents().Attach("{providerID}").Post(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_attach_document]
	_ = response
}

func ref_execute_version_action() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_execute_version_action]
	if err := client.Now().Documents().Action("checkout").Document("{documentSysID}").Version("{versionSysID}").Patch(ctx, nil); err != nil {
		log.Fatal(err)
	}
	// [END ref_execute_version_action]
}

func ref_update_deployables() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_update_deployables]
	// Build the request body
	var body *cdmapplicationsapi.DeployableUpdateRequest
	response, err := client.Cdm().Applications().Deployables().Put(ctx, body, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_update_deployables]
	_ = response
}

func ref_delete_deployables() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_delete_deployables]
	if err := client.Cdm().Applications().Deployables().Delete(ctx, nil); err != nil {
		log.Fatal(err)
	}
	// [END ref_delete_deployables]
}

func ref_list_exports() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_list_exports]
	response, err := client.Cdm().Applications().Deployables().Exports().Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_list_exports]
	_ = response
}

func ref_get_export_status() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_get_export_status]
	response, err := client.Cdm().Applications().Deployables().Exports().ByID("{exportID}").Status().Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_get_export_status]
	_ = response
}

func ref_get_export_content() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_get_export_content]
	response, err := client.Cdm().Applications().Deployables().Exports().ByID("{exportID}").Content().Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_get_export_content]
	_ = response
}

func ref_update_shared_components() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_update_shared_components]
	// Build the request body
	var body *cdmapplicationsapi.SharedComponentUpdateRequest
	response, err := client.Cdm().Applications().SharedComponents().Put(ctx, body, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_update_shared_components]
	_ = response
}

func ref_delete_shared_components() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_delete_shared_components]
	if err := client.Cdm().Applications().SharedComponents().Delete(ctx, nil); err != nil {
		log.Fatal(err)
	}
	// [END ref_delete_shared_components]
}

func ref_list_shared_library_applications() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_list_shared_library_applications]
	response, err := client.Cdm().Applications().SharedLibraries().Components().Applications().Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_list_shared_library_applications]
	_ = response
}

func ref_upload_components() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_upload_components]
	// Build the request body
	var body *cdmapplicationsapi.ComponentUploadRequest
	response, err := client.Cdm().Applications().Uploads().Components().Post(ctx, body, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_upload_components]
	_ = response
}

func ref_upload_component_vars() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_upload_component_vars]
	// Build the request body
	var body *cdmapplicationsapi.ComponentVarsUploadRequest
	response, err := client.Cdm().Applications().Uploads().Components().Vars().Post(ctx, body, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_upload_component_vars]
	_ = response
}

func ref_upload_collections() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_upload_collections]
	// Build the request body
	var body *cdmapplicationsapi.CollectionUploadRequest
	response, err := client.Cdm().Applications().Uploads().Collections().Post(ctx, body, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_upload_collections]
	_ = response
}

func ref_upload_collection_file() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_upload_collection_file]
	// Build the request body
	var body *cdmapplicationsapi.Media
	response, err := client.Cdm().Applications().Uploads().Collections().File().Post(ctx, body, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_upload_collection_file]
	_ = response
}

func ref_upload_deployable_file() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_upload_deployable_file]
	// Build the request body
	var body *cdmapplicationsapi.Media
	response, err := client.Cdm().Applications().Uploads().Deployables().File().Post(ctx, body, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_upload_deployable_file]
	_ = response
}

func ref_get_upload_status() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_get_upload_status]
	response, err := client.Cdm().Applications().UploadStatus().ByID("{uploadID}").Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_get_upload_status]
	_ = response
}

func ref_list_changesets() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_list_changesets]
	response, err := client.Cdm().Changesets().Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_list_changesets]
	_ = response
}

func ref_delete_changesets() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_delete_changesets]
	if err := client.Cdm().Changesets().Delete(ctx, nil); err != nil {
		log.Fatal(err)
	}
	// [END ref_delete_changesets]
}

func ref_get_changeset_activity() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_get_changeset_activity]
	response, err := client.Cdm().Changesets().Activity().Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_get_changeset_activity]
	_ = response
}

func ref_get_commit_status() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_get_commit_status]
	response, err := client.Cdm().Changesets().CommitStatus().ByID("{commitID}").Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_get_commit_status]
	_ = response
}

func ref_list_impacted_deployables() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_list_impacted_deployables]
	response, err := client.Cdm().Changesets().ImpactedDeployables().Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_list_impacted_deployables]
	_ = response
}

func ref_list_impacted_shared_components() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_list_impacted_shared_components]
	response, err := client.Cdm().Changesets().ImpactedSharedComponents().Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_list_impacted_shared_components]
	_ = response
}

func ref_get_changeset_impacted_deployables() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_get_changeset_impacted_deployables]
	response, err := client.Cdm().Changesets().ByID("{changesetID}").ImpactedDeployables().Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_get_changeset_impacted_deployables]
	_ = response
}

func ref_list_nodes() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_list_nodes]
	response, err := client.Cdm().Editor().Nodes().Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_list_nodes]
	_ = response
}

func ref_create_node() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_create_node]
	// Build the request body
	var body cdmeditorapi.NodeCreateRequest
	response, err := client.Cdm().Editor().Nodes().Post(ctx, body, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_create_node]
	_ = response
}

func ref_update_node() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_update_node]
	// Build the request body
	var body cdmeditorapi.NodeUpdateRequest
	response, err := client.Cdm().Editor().Nodes().ByID("{nodeSysID}").Put(ctx, body, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_update_node]
	_ = response
}

func ref_delete_node() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_delete_node]
	response, err := client.Cdm().Editor().Nodes().ByID("{nodeSysID}").Delete(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_delete_node]
	_ = response
}

func ref_validate_configuration() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_validate_configuration]
	response, err := client.Cdm().Editor().Validation().Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_validate_configuration]
	_ = response
}

func ref_create_policy_mapping() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_create_policy_mapping]
	response, err := client.Cdm().Policies().Mappings().Post(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END ref_create_policy_mapping]
	_ = response
}

func ref_delete_policy_mapping() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_delete_policy_mapping]
	if err := client.Cdm().Policies().Mappings().Delete(ctx, nil); err != nil {
		log.Fatal(err)
	}
	// [END ref_delete_policy_mapping]
}

func _() {
	ref_list_accounts()
	ref_get_account()
	ref_get_aggregates()
	ref_list_activities()
	ref_list_contexts()
	ref_list_subscription_objects()
	ref_get_facet_instances()
	ref_get_followings()
	ref_get_preference()
	ref_create_preference()
	ref_get_subscription()
	ref_check_subscription()
	ref_subscribe()
	ref_unsubscribe()
	ref_get_subscribers()
	ref_get_user_stream()
	ref_update_user_stream()
	ref_check_availability()
	ref_book_appointment()
	ref_get_calendar()
	ref_get_configuration()
	ref_execute_rule_conditions()
	ref_request_user_window()
	ref_create_application_service()
	ref_find_service()
	ref_register_service()
	ref_populate_service()
	ref_update_service_details()
	ref_list_cases()
	ref_create_case()
	ref_get_case()
	ref_update_case()
	ref_list_case_activities()
	ref_get_case_field_values()
	ref_list_configuration_items()
	ref_create_configuration_item()
	ref_get_configuration_item()
	ref_update_configuration_item()
	ref_patch_configuration_item()
	ref_create_relation()
	ref_delete_relation()
	ref_explore_documents()
	ref_create_document()
	ref_create_or_link_document()
	ref_delete_document()
	ref_get_document_content()
	ref_list_document_versions()
	ref_get_version_state()
	ref_sync_down_document()
	ref_attach_document()
	ref_execute_version_action()
	ref_update_deployables()
	ref_delete_deployables()
	ref_list_exports()
	ref_get_export_status()
	ref_get_export_content()
	ref_update_shared_components()
	ref_delete_shared_components()
	ref_list_shared_library_applications()
	ref_upload_components()
	ref_upload_component_vars()
	ref_upload_collections()
	ref_upload_collection_file()
	ref_upload_deployable_file()
	ref_get_upload_status()
	ref_list_changesets()
	ref_delete_changesets()
	ref_get_changeset_activity()
	ref_get_commit_status()
	ref_list_impacted_deployables()
	ref_list_impacted_shared_components()
	ref_get_changeset_impacted_deployables()
	ref_list_nodes()
	ref_create_node()
	ref_update_node()
	ref_delete_node()
	ref_validate_configuration()
	ref_create_policy_mapping()
	ref_delete_policy_mapping()
}
