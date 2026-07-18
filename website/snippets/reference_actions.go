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

func refListAccounts() {
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

func refGetAccount() {
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

func refGetAggregates() {
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

func refListActivities() {
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

func refListContexts() {
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

func refListSubscriptionObjects() {
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

func refGetFacetInstances() {
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

func refGetFollowings() {
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

func refGetPreference() {
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

func refCreatePreference() {
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

func refGetSubscription() {
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

func refCheckSubscription() {
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

func refSubscribe() {
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

func refUnsubscribe() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_unsubscribe]
	if err := client.Now().ActSub().Subscriptions().ByObjectId("{objectID}").Unsubscribe().Delete(ctx, nil); err != nil {
		log.Fatal(err)
	}
	// [END ref_unsubscribe]
}

func refGetSubscribers() {
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

func refGetUserStream() {
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

func refUpdateUserStream() {
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

func refCheckAvailability() {
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

func refBookAppointment() {
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

func refGetCalendar() {
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

func refGetConfiguration() {
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

func refExecuteRuleConditions() {
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

func refRequestUserWindow() {
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

func refCreateApplicationService() {
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

func refFindService() {
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

func refRegisterService() {
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

func refPopulateService() {
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

func refUpdateServiceDetails() {
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

func refListCases() {
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

func refCreateCase() {
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

func refGetCase() {
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

func refUpdateCase() {
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

func refListCaseActivities() {
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

func refGetCaseFieldValues() {
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

func refListConfigurationItems() {
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

func refCreateConfigurationItem() {
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

func refGetConfigurationItem() {
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

func refUpdateConfigurationItem() {
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

func refPatchConfigurationItem() {
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

func refCreateRelation() {
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

func refDeleteRelation() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_delete_relation]
	if err := client.Now().Cmdb().Instance().ByClass("cmdb_ci_linux_server").ByID("{sysID}").Relation().ByID("{relSysID}").Delete(ctx, nil); err != nil {
		log.Fatal(err)
	}
	// [END ref_delete_relation]
}

func refExploreDocuments() {
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

func refCreateDocument() {
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

func refCreateOrLinkDocument() {
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

func refDeleteDocument() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_delete_document]
	if err := client.Now().Documents().Delete().Delete(ctx, nil); err != nil {
		log.Fatal(err)
	}
	// [END ref_delete_document]
}

func refGetDocumentContent() {
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

func refListDocumentVersions() {
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

func refGetVersionState() {
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

func refSyncDownDocument() {
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

func refAttachDocument() {
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

func refExecuteVersionAction() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_execute_version_action]
	if err := client.Now().Documents().Action("checkout").Document("{documentSysID}").Version("{versionSysID}").Patch(ctx, nil); err != nil {
		log.Fatal(err)
	}
	// [END ref_execute_version_action]
}

func refUpdateDeployables() {
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

func refDeleteDeployables() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_delete_deployables]
	if err := client.Cdm().Applications().Deployables().Delete(ctx, nil); err != nil {
		log.Fatal(err)
	}
	// [END ref_delete_deployables]
}

func refListExports() {
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

func refGetExportStatus() {
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

func refGetExportContent() {
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

func refUpdateSharedComponents() {
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

func refDeleteSharedComponents() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_delete_shared_components]
	if err := client.Cdm().Applications().SharedComponents().Delete(ctx, nil); err != nil {
		log.Fatal(err)
	}
	// [END ref_delete_shared_components]
}

func refListSharedLibraryApplications() {
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

func refUploadComponents() {
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

func refUploadComponentVars() {
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

func refUploadCollections() {
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

func refUploadCollectionFile() {
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

func refUploadDeployableFile() {
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

func refGetUploadStatus() {
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

func refListChangesets() {
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

func refDeleteChangesets() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_delete_changesets]
	if err := client.Cdm().Changesets().Delete(ctx, nil); err != nil {
		log.Fatal(err)
	}
	// [END ref_delete_changesets]
}

func refGetChangesetActivity() {
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

func refGetCommitStatus() {
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

func refListImpactedDeployables() {
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

func refListImpactedSharedComponents() {
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

func refGetChangesetImpactedDeployables() {
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

func refListNodes() {
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

func refCreateNode() {
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

func refUpdateNode() {
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

func refDeleteNode() {
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

func refValidateConfiguration() {
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

func refCreatePolicyMapping() {
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

func refDeletePolicyMapping() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()
	// [START ref_delete_policy_mapping]
	if err := client.Cdm().Policies().Mappings().Delete(ctx, nil); err != nil {
		log.Fatal(err)
	}
	// [END ref_delete_policy_mapping]
}

func _() {
	refListAccounts()
	refGetAccount()
	refGetAggregates()
	refListActivities()
	refListContexts()
	refListSubscriptionObjects()
	refGetFacetInstances()
	refGetFollowings()
	refGetPreference()
	refCreatePreference()
	refGetSubscription()
	refCheckSubscription()
	refSubscribe()
	refUnsubscribe()
	refGetSubscribers()
	refGetUserStream()
	refUpdateUserStream()
	refCheckAvailability()
	refBookAppointment()
	refGetCalendar()
	refGetConfiguration()
	refExecuteRuleConditions()
	refRequestUserWindow()
	refCreateApplicationService()
	refFindService()
	refRegisterService()
	refPopulateService()
	refUpdateServiceDetails()
	refListCases()
	refCreateCase()
	refGetCase()
	refUpdateCase()
	refListCaseActivities()
	refGetCaseFieldValues()
	refListConfigurationItems()
	refCreateConfigurationItem()
	refGetConfigurationItem()
	refUpdateConfigurationItem()
	refPatchConfigurationItem()
	refCreateRelation()
	refDeleteRelation()
	refExploreDocuments()
	refCreateDocument()
	refCreateOrLinkDocument()
	refDeleteDocument()
	refGetDocumentContent()
	refListDocumentVersions()
	refGetVersionState()
	refSyncDownDocument()
	refAttachDocument()
	refExecuteVersionAction()
	refUpdateDeployables()
	refDeleteDeployables()
	refListExports()
	refGetExportStatus()
	refGetExportContent()
	refUpdateSharedComponents()
	refDeleteSharedComponents()
	refListSharedLibraryApplications()
	refUploadComponents()
	refUploadComponentVars()
	refUploadCollections()
	refUploadCollectionFile()
	refUploadDeployableFile()
	refGetUploadStatus()
	refListChangesets()
	refDeleteChangesets()
	refGetChangesetActivity()
	refGetCommitStatus()
	refListImpactedDeployables()
	refListImpactedSharedComponents()
	refGetChangesetImpactedDeployables()
	refListNodes()
	refCreateNode()
	refUpdateNode()
	refDeleteNode()
	refValidateConfiguration()
	refCreatePolicyMapping()
	refDeletePolicyMapping()
}
