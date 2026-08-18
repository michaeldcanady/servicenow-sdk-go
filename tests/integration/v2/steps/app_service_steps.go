package steps

import (
	"context"
	"fmt"

	"github.com/cucumber/godog"
	appserviceapi "github.com/michaeldcanady/servicenow-sdk-go/appserviceapi"
	"github.com/michaeldcanady/servicenow-sdk-go/tests/integration/v2/support"
)

type appServiceSteps struct{}

func (s *appServiceSteps) iFindAServiceWithQuery(ctx context.Context, query string) (context.Context, error) {
	w := support.WorldFrom(ctx)

	name := query
	cfg := &appserviceapi.FindServiceRequestConfiguration{
		QueryParameters: &appserviceapi.FindServiceQueryParameters{
			Name: &name,
		},
	}

	resp, err := w.Client.Now().Cmdb().AppService().Csdm().FindService().Get(ctx, cfg)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *appServiceSteps) iCreateAServiceWithName(ctx context.Context, name string) (context.Context, error) {
	w := support.WorldFrom(ctx)

	body := appserviceapi.NewCreateServiceRequest()
	_ = body.SetName(&name)

	resp, err := w.Client.Now().Cmdb().AppService().Create().Post(ctx, body, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	result, err := resp.GetResult()
	if err == nil {
		sysID, err := result.GetSysID()
		if err == nil && sysID != nil && *sysID != "" {
			w.LastSysID = *sysID
		}
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *appServiceSteps) iGetContentOfTheLastService(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	sysID := w.LastSysID
	if sysID == "" {
		if support.IsOffline() {
			sysID = "mock_app_service_sys_id_1"
		} else {
			return ctx, fmt.Errorf("no app service sys_id available — create a service first")
		}
	}

	resp, err := w.Client.Now().Cmdb().AppService().ByID(sysID).GetContent().Get(ctx, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *appServiceSteps) iGetContentOfServiceWithSysID(ctx context.Context, sysID string) (context.Context, error) {
	w := support.WorldFrom(ctx)

	resp, err := w.Client.Now().Cmdb().AppService().ByID(sysID).GetContent().Get(ctx, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

// InitializeAppServiceScenario registers all app service step definitions.
func InitializeAppServiceScenario(sc *godog.ScenarioContext) {
	s := &appServiceSteps{}

	RegisterSharedSteps(sc)

	sc.Step(`^I find a service with query "([^"]*)"$`, s.iFindAServiceWithQuery)
	sc.Step(`^I create a service with name "([^"]*)"$`, s.iCreateAServiceWithName)
	sc.Step(`^I get content of the last service$`, s.iGetContentOfTheLastService)
	sc.Step(`^I get content of service "([^"]*)"$`, s.iGetContentOfServiceWithSysID)

	sc.Before(BeforeScenario)
	sc.After(AfterScenario)
}
