// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package steps

import (
	"context"
	"fmt"
	"strings"

	"github.com/cucumber/godog"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/tableapi"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/tests/integration/v2/support"
)

type tableSteps struct{}

func (s *tableSteps) iCreateANewIncidentWithDescription(ctx context.Context, desc string) (context.Context, error) {
	w := support.WorldFrom(ctx)

	record := tableapi.NewTableRecord()
	_ = record.SetValue("short_description", desc)

	resp, err := w.Client.Now().Table("incident").Post(ctx, record, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	result, err := resp.GetResult()
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	sysID, err := result.GetSysID()
	if err != nil || sysID == nil || *sysID == "" {
		w.Err = fmt.Errorf("missing or invalid sys_id in created record")
		return support.WithWorld(ctx, w), nil
	}

	w.LastSysID = *sysID
	w.TrackCreation(*sysID)
	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *tableSteps) iRequestTheIncidentBySysID(ctx context.Context, field string) (context.Context, error) {
	w := support.WorldFrom(ctx)
	sysID := w.LastSysID
	if sysID == "" {
		if support.IsOffline() {
			sysID = "mock_sys_id_1"
		} else {
			return ctx, fmt.Errorf("no record sys_id available — create a record first")
		}
	}

	resp, err := w.Client.Now().Table("incident").ByID(sysID).Get(ctx, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *tableSteps) iRequestTheIncidentWithSysID(ctx context.Context, sysID string) (context.Context, error) {
	w := support.WorldFrom(ctx)

	resp, err := w.Client.Now().Table("incident").ByID(sysID).Get(ctx, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *tableSteps) iRequestAllIncidentsFromTable(ctx context.Context, table string) (context.Context, error) {
	w := support.WorldFrom(ctx)

	resp, err := w.Client.Now().Table(table).Get(ctx, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *tableSteps) iRequestAllIncidentsFromTableWithQuery(ctx context.Context, table, query string) (context.Context, error) {
	w := support.WorldFrom(ctx)

	cfg := &tableapi.TableRequestBuilderGetRequestConfiguration{
		QueryParameters: &tableapi.TableRequestBuilderGetQueryParameters{
			Query: &query,
		},
	}

	resp, err := w.Client.Now().Table(table).Get(ctx, cfg)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *tableSteps) iRequestAllIncidentsFromTableWithDisplayValue(ctx context.Context, table, displayValue string) (context.Context, error) {
	w := support.WorldFrom(ctx)

	var dv tableapi.DisplayValue
	switch displayValue {
	case "true":
		dv = tableapi.DisplayValueTrue
	case "false":
		dv = tableapi.DisplayValueFalse
	case "all":
		dv = tableapi.DisplayValueAll
	default:
		dv = tableapi.DisplayValueUnknown
	}

	cfg := &tableapi.TableRequestBuilderGetRequestConfiguration{
		QueryParameters: &tableapi.TableRequestBuilderGetQueryParameters{
			DisplayValue: &dv,
		},
	}

	resp, err := w.Client.Now().Table(table).Get(ctx, cfg)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *tableSteps) iRequestAllIncidentsFromTableWithFields(ctx context.Context, table, fields string) (context.Context, error) {
	w := support.WorldFrom(ctx)

	cfg := &tableapi.TableRequestBuilderGetRequestConfiguration{
		QueryParameters: &tableapi.TableRequestBuilderGetQueryParameters{
			Fields: []string{fields},
		},
	}

	resp, err := w.Client.Now().Table(table).Get(ctx, cfg)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *tableSteps) iRequestAllIncidentsFromTableWithLimitAndOffset(ctx context.Context, table string, limit, offset int) (context.Context, error) {
	w := support.WorldFrom(ctx)

	l := int32(limit)
	o := int32(offset)
	cfg := &tableapi.TableRequestBuilderGetRequestConfiguration{
		QueryParameters: &tableapi.TableRequestBuilderGetQueryParameters{
			Limit:  &l,
			Offset: &o,
		},
	}

	resp, err := w.Client.Now().Table(table).Get(ctx, cfg)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *tableSteps) iUpdateIncidentWithPUTDescription(ctx context.Context, sysID, desc string) (context.Context, error) {
	w := support.WorldFrom(ctx)

	record := tableapi.NewTableRecord()
	_ = record.SetValue("short_description", desc)

	resp, err := w.Client.Now().Table("incident").ByID(sysID).Put(ctx, record, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *tableSteps) iUpdateLastIncidentWithPUTDescription(ctx context.Context, desc string) (context.Context, error) {
	w := support.WorldFrom(ctx)
	sysID := w.LastSysID
	if sysID == "" {
		if support.IsOffline() {
			sysID = "mock_sys_id_1"
		} else {
			return ctx, fmt.Errorf("no record sys_id available — create a record first")
		}
	}

	record := tableapi.NewTableRecord()
	_ = record.SetValue("short_description", desc)

	resp, err := w.Client.Now().Table("incident").ByID(sysID).Put(ctx, record, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *tableSteps) iPatchIncidentDescriptionTo(ctx context.Context, sysID, desc string) (context.Context, error) {
	w := support.WorldFrom(ctx)

	record := tableapi.NewTableRecord()
	_ = record.SetValue("short_description", desc)

	resp, err := w.Client.Now().Table("incident").ByID(sysID).Patch(ctx, record, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *tableSteps) iPatchLastIncidentDescriptionTo(ctx context.Context, desc string) (context.Context, error) {
	w := support.WorldFrom(ctx)
	sysID := w.LastSysID
	if sysID == "" {
		if support.IsOffline() {
			sysID = "mock_sys_id_1"
		} else {
			return ctx, fmt.Errorf("no record sys_id available — create a record first")
		}
	}

	record := tableapi.NewTableRecord()
	_ = record.SetValue("short_description", desc)

	resp, err := w.Client.Now().Table("incident").ByID(sysID).Patch(ctx, record, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *tableSteps) iDeleteTheIncident(ctx context.Context, sysID string) (context.Context, error) {
	w := support.WorldFrom(ctx)

	err := w.Client.Now().Table("incident").ByID(sysID).Delete(ctx, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	return support.WithWorld(ctx, w), nil
}

func (s *tableSteps) iDeleteLastIncident(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	sysID := w.LastSysID
	if sysID == "" {
		if support.IsOffline() {
			sysID = "mock_sys_id_1"
		} else {
			return ctx, fmt.Errorf("no record sys_id available — create a record first")
		}
	}

	err := w.Client.Now().Table("incident").ByID(sysID).Delete(ctx, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	return support.WithWorld(ctx, w), nil
}

func (s *tableSteps) iSendAHEADRequestForTheLastIncident(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	sysID := w.LastSysID
	if sysID == "" {
		if support.IsOffline() {
			sysID = "mock_sys_id_1"
		} else {
			return ctx, fmt.Errorf("no record sys_id available — create a record first")
		}
	}

	resp, err := w.Client.Now().Table("incident").Head(ctx, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *tableSteps) theCreatedRecordShouldHaveValid(ctx context.Context, field string) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.Response == nil {
		return ctx, fmt.Errorf("response is nil, cannot check field %q", field)
	}
	return ctx, nil
}

func (s *tableSteps) theResultsShouldContainAtLeastRecords(ctx context.Context, minCount int) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.Response == nil {
		return ctx, fmt.Errorf("response is nil, expected at least %d records", minCount)
	}
	colResp, ok := w.Response.(core.ServiceNowCollectionResponse[*tableapi.TableRecord])
	if !ok {
		return ctx, fmt.Errorf("response is not a collection response")
	}
	results, err := colResp.GetResult()
	if err != nil {
		return ctx, fmt.Errorf("failed to get results: %v", err)
	}
	if len(results) < minCount {
		return ctx, fmt.Errorf("expected at least %d records, got %d", minCount, len(results))
	}
	return ctx, nil
}

func (s *tableSteps) theRetrievedRecordShouldHaveFieldContaining(ctx context.Context, field, expectedSubstring string) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.Response == nil {
		return ctx, fmt.Errorf("response is nil, cannot check field %q", field)
	}
	itemResp, ok := w.Response.(core.ServiceNowItemResponse[*tableapi.TableRecord])
	if !ok {
		return ctx, fmt.Errorf("response is not an item response")
	}
	result, err := itemResp.GetResult()
	if err != nil {
		return ctx, fmt.Errorf("failed to get result: %v", err)
	}
	elem, err := result.Get(field)
	if err != nil || elem == nil {
		return ctx, fmt.Errorf("field %q is nil or missing", field)
	}
	val, err := elem.GetValue()
	if err != nil {
		return ctx, fmt.Errorf("failed to get value for field %q: %v", field, err)
	}
	valStr := fmt.Sprintf("%v", val)
	if !strings.Contains(valStr, expectedSubstring) {
		return ctx, fmt.Errorf("field %q value %q does not contain %q", field, valStr, expectedSubstring)
	}
	return ctx, nil
}

// InitializeTableScenario registers all table step definitions.
func InitializeTableScenario(sc *godog.ScenarioContext) {
	s := &tableSteps{}

	RegisterSharedSteps(sc)

	sc.Step(`^I create a new incident with description "([^"]*)"$`, s.iCreateANewIncidentWithDescription)
	sc.Step(`^I request the incident by its "([^"]*)"$`, s.iRequestTheIncidentBySysID)
	sc.Step(`^I request the incident with sys_id "([^"]*)"$`, s.iRequestTheIncidentWithSysID)
	sc.Step(`^I request all incidents from the "([^"]*)" table$`, s.iRequestAllIncidentsFromTable)
	sc.Step(`^I request all incidents from the "([^"]*)" table with query "([^"]*)"$`, s.iRequestAllIncidentsFromTableWithQuery)
	sc.Step(`^I request all incidents from the "([^"]*)" table with display_value "([^"]*)"$`, s.iRequestAllIncidentsFromTableWithDisplayValue)
	sc.Step(`^I request all incidents from the "([^"]*)" table with fields "([^"]*)"$`, s.iRequestAllIncidentsFromTableWithFields)
	sc.Step(`^I request all incidents from the "([^"]*)" table with limit (\d+) and offset (\d+)$`, s.iRequestAllIncidentsFromTableWithLimitAndOffset)
	sc.Step(`^I update the incident "([^"]*)" with PUT description "([^"]*)"$`, s.iUpdateIncidentWithPUTDescription)
	sc.Step(`^I update the last incident with PUT description "([^"]*)"$`, s.iUpdateLastIncidentWithPUTDescription)
	sc.Step(`^I patch the incident "([^"]*)" description to "([^"]*)"$`, s.iPatchIncidentDescriptionTo)
	sc.Step(`^I patch the last incident description to "([^"]*)"$`, s.iPatchLastIncidentDescriptionTo)
	sc.Step(`^I delete the incident "([^"]*)"$`, s.iDeleteTheIncident)
	sc.Step(`^I delete the last incident$`, s.iDeleteLastIncident)
	sc.Step(`^I send a HEAD request for the last incident$`, s.iSendAHEADRequestForTheLastIncident)
	sc.Step(`^the created record should have a valid "([^"]*)"$`, s.theCreatedRecordShouldHaveValid)
	sc.Step(`^the results should contain at least (\d+) record$`, s.theResultsShouldContainAtLeastRecords)
	sc.Step(`^the retrieved record should have field "([^"]*)" containing "([^"]*)"$`, s.theRetrievedRecordShouldHaveFieldContaining)

	sc.Before(BeforeScenario)
	sc.After(AfterScenario)
}
