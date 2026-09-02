// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package steps

import (
	"context"

	"github.com/cucumber/godog"
	actsubapi "github.com/michaeldcanady/servicenow-sdk-go/v2/activitysubscriptionsapi"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/tests/integration/v2/support"
)

type activitySubscriptionsSteps struct{}

func (s *activitySubscriptionsSteps) iRetrieveActivitiesForStream(ctx context.Context, streamID string) (context.Context, error) {
	w := support.WorldFrom(ctx)

	contextID := "mock_context_id"
	contextInstanceID := streamID

	cfg := &actsubapi.ActivitiesRequestBuilderGetRequestConfiguration{
		QueryParameters: &actsubapi.ActivitiesRequestBuilderGetQueryParameters{
			Context:         &contextID,
			ContextInstance: &contextInstanceID,
		},
	}

	resp, err := w.Client.Now().ActSub().Activities().Get(ctx, cfg)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *activitySubscriptionsSteps) iRetrieveFacetsForContextAndInstance(ctx context.Context, contextID, instanceID string) (context.Context, error) {
	w := support.WorldFrom(ctx)

	resp, err := w.Client.Now().ActSub().Facets().ByContext(contextID).ByInstance(instanceID).Get(ctx, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *activitySubscriptionsSteps) iRetrieveActivitiesForEmptyStream(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)

	contextID := ""
	contextInstanceID := ""

	cfg := &actsubapi.ActivitiesRequestBuilderGetRequestConfiguration{
		QueryParameters: &actsubapi.ActivitiesRequestBuilderGetQueryParameters{
			Context:         &contextID,
			ContextInstance: &contextInstanceID,
		},
	}

	resp, err := w.Client.Now().ActSub().Activities().Get(ctx, cfg)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

// InitializeActivitySubscriptionsScenario registers all activity subscriptions step definitions.
func InitializeActivitySubscriptionsScenario(sc *godog.ScenarioContext) {
	s := &activitySubscriptionsSteps{}

	RegisterSharedSteps(sc)

	sc.Step(`^I retrieve activities for stream "([^"]*)"$`, s.iRetrieveActivitiesForStream)
	sc.Step(`^I retrieve facets for context "([^"]*)" and instance "([^"]*)"$`, s.iRetrieveFacetsForContextAndInstance)
	sc.Step(`^I retrieve activities for empty stream$`, s.iRetrieveActivitiesForEmptyStream)

	sc.Before(BeforeScenario)
	sc.After(AfterScenario)
}
