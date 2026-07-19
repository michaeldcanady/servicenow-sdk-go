//go:build snippets

package snippets

import (
	"context"
	"log"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
	aggregationapi "github.com/michaeldcanady/servicenow-sdk-go/aggregationapi"
	appointmentbookingapi "github.com/michaeldcanady/servicenow-sdk-go/appointmentbookingapi"
	appserviceapi "github.com/michaeldcanady/servicenow-sdk-go/appserviceapi"
	batchapi "github.com/michaeldcanady/servicenow-sdk-go/batchapi"
	caseapi "github.com/michaeldcanady/servicenow-sdk-go/caseapi"
	cdmapplicationsapi "github.com/michaeldcanady/servicenow-sdk-go/cdmapplicationsapi"
	cdmeditorapi "github.com/michaeldcanady/servicenow-sdk-go/cdmeditorapi"
	cmdbinstanceapi "github.com/michaeldcanady/servicenow-sdk-go/cmdbinstanceapi"
)

func _() {
	moduleAccount()
	moduleActSub()
	moduleAppointmentBooking()
	moduleAppService()
	moduleAttachment()
	moduleBatch()
	moduleCase()
	moduleCdmApplications()
	moduleCdmChangesets()
	moduleCdmEditor()
	moduleCmdbInstance()
	moduleDocuments()
	modulePolicy()
	moduleStats()
	moduleTables()
}

func moduleAccount() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	// [START module_account]
	// List accounts
	accounts, err := client.Now().Account().Get(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Get a single account by its account ID
	account, err := client.Now().Account().ByID("{accountID}").Get(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END module_account]
	_ = accounts
	_ = account
}

func moduleActSub() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	// [START module_actsub]
	actSub := client.Now().ActSub()

	// List activities
	activities, err := actSub.Activities().Get(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Check whether the current user is subscribed to an object
	subscribed, err := actSub.Subscriptions().
		ByObjectId("{objectID}").
		IsSubscribed().
		Get(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END module_actsub]
	_ = activities
	_ = subscribed
}

func moduleAppointmentBooking() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	var catalogID, startDate, endDate string

	// [START module_appointment_booking]
	booking := client.AppointmentBooking()

	// Check availability
	availabilityRequest := appointmentbookingapi.NewAvailabilityRequest()
	if err := availabilityRequest.SetCatalogId(&catalogID); err != nil {
		log.Fatal(err)
	}
	if err := availabilityRequest.SetStartDate(&startDate); err != nil {
		log.Fatal(err)
	}
	if err := availabilityRequest.SetEndDate(&endDate); err != nil {
		log.Fatal(err)
	}
	availability, err := booking.Availability().Post(context.Background(), availabilityRequest, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Book an appointment
	appointmentRequest := appointmentbookingapi.NewAppointmentRequest()
	if err := appointmentRequest.SetCatalogId(&catalogID); err != nil {
		log.Fatal(err)
	}
	appointment, err := booking.Appointment().Post(context.Background(), appointmentRequest, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END module_appointment_booking]
	_ = availability
	_ = appointment
}

func moduleAppService() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	var registerRequest *appserviceapi.RegisterServiceRequest

	// [START module_app_service]
	appService := client.Now().Cmdb().AppService()

	// Register an existing application service
	resp, err := appService.Csdm().RegisterService().Post(context.Background(), registerRequest, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Find a service
	found, err := appService.Csdm().FindService().Get(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END module_app_service]
	_ = resp
	_ = found
}

func moduleAttachment() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	// [START module_attachment]
	// Access the attachment API
	attachmentAPI := client.Now().Attachment()

	// Fetch metadata for all attachments (optionally filtered)
	result, err := attachmentAPI.Get(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END module_attachment]
	_ = result
}

func moduleBatch() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	// [START module_batch]
	// 1. Prepare individual requests using ToRequestInformation methods
	request1, err := client.Now().Table("incident").ToGetRequestInformation(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	request2, err := client.Now().Table("sys_user").ToGetRequestInformation(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// 2. Create a batch request model and add the requests
	body := batchapi.NewBatchRequestModel()

	// 3. Execute the batch
	response, err := client.Now().Batch().Post(context.Background(), body, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END module_batch]
	_ = request1
	_ = request2
	_ = response
}

func moduleCase() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	var newCase *caseapi.CaseResultModel
	var caseBody *caseapi.CaseResultModel

	// [START module_case]
	cases := client.CustomerService().Case()

	// List cases
	list, err := cases.Get(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Create a case
	created, err := cases.Post(context.Background(), newCase, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Get and update a specific case
	item, err := cases.ByID("{caseSysID}").Get(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	updated, err := cases.ByID("{caseSysID}").Put(context.Background(), caseBody, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END module_case]
	_ = list
	_ = created
	_ = item
	_ = updated
}

func moduleCdmApplications() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	var uploadRequest *cdmapplicationsapi.ComponentUploadRequest

	// [START module_cdm_applications]
	applications := client.Cdm().Applications()

	// Upload component configuration data
	status, err := applications.Uploads().Components().Post(context.Background(), uploadRequest, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Check an upload's status
	uploadStatus, err := applications.UploadStatus().ByID("{uploadID}").Get(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Export a deployable's configuration and fetch the content
	export, err := applications.Deployables().Exports().Get(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	content, err := applications.Deployables().Exports().ByID("{exportID}").Content().Get(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END module_cdm_applications]
	_ = status
	_ = uploadStatus
	_ = export
	_ = content
}

func moduleCdmChangesets() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	// [START module_cdm_changesets]
	changesets := client.Cdm().Changesets()

	// List changesets
	list, err := changesets.Get(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Check the status of a commit
	status, err := changesets.CommitStatus().ByID("{commitID}").Get(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Impacted deployables for a specific changeset
	impact, err := changesets.ByID("{changesetSysID}").ImpactedDeployables().Get(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END module_cdm_changesets]
	_ = list
	_ = status
	_ = impact
}

func moduleCdmEditor() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	var createRequest cdmeditorapi.NodeCreateRequest
	var updateRequest cdmeditorapi.NodeUpdateRequest

	// [START module_cdm_editor]
	editor := client.Cdm().Editor()

	// List nodes
	nodes, err := editor.Nodes().Get(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Create and update a node
	created, err := editor.Nodes().Post(context.Background(), createRequest, nil)
	if err != nil {
		log.Fatal(err)
	}
	updated, err := editor.Nodes().ByID("{nodeID}").Put(context.Background(), updateRequest, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Validate configuration data
	validation, err := editor.Validation().Get(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END module_cdm_editor]
	_ = nodes
	_ = created
	_ = updated
	_ = validation
}

func moduleCmdbInstance() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	var ciBody *cmdbinstanceapi.CmdbInstanceModel

	// [START module_cmdb_instance]
	instance := client.Now().Cmdb().Instance()

	// Query CIs of a class
	servers, err := instance.ByClass("cmdb_ci_linux_server").Get(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Get / update a specific CI
	ci, err := instance.ByClass("cmdb_ci_linux_server").ByID("{sysID}").Get(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	updated, err := instance.ByClass("cmdb_ci_linux_server").ByID("{sysID}").Patch(context.Background(), ciBody, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END module_cmdb_instance]
	_ = servers
	_ = ci
	_ = updated
}

func moduleDocuments() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	// [START module_documents]
	documents := client.Now().Documents()

	// Search documents and folders
	results, err := documents.Explore().Get(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Download a document's content
	content, err := documents.Content("{documentSysID}").Get(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// List a document's versions
	versions, err := documents.Versions("{documentSysID}").Get(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END module_documents]
	_ = results
	_ = content
	_ = versions
}

func modulePolicy() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	// [START module_policy]
	mappings := client.Cdm().Policies().Mappings()

	// Create a policy mapping
	created, err := mappings.Post(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Resolve policy inputs
	resolved := mappings.Inputs().Resolved()

	// Delete a mapping
	err = mappings.Delete(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END module_policy]
	_ = created
	_ = resolved
}

func moduleStats() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	// [START module_stats]
	config := &aggregationapi.StatsRequestBuilderGetRequestConfiguration{
		QueryParameters: &aggregationapi.StatsRequestBuilderGetQueryParameters{
			Count:     true,
			SumFields: []string{"reassignment_count"},
			Query:     "active=true",
		},
	}

	response, err := client.Now().Stats("incident").Get(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}

	result, err := response.GetResult()
	if err != nil {
		log.Fatal(err)
	}

	stats, err := result.GetStats()
	if err != nil {
		log.Fatal(err)
	}

	count, _ := stats.GetCount()
	// [END module_stats]
	_ = count
}

func moduleTables() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	// [START module_tables]
	// Access the incident table
	incidentTable := client.Now().Table("incident")

	// Fetch a list of records
	result, err := incidentTable.Get(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END module_tables]
	_ = result
}
