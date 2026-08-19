package steps

import (
	"context"

	"github.com/cucumber/godog"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/tests/integration/v2/support"
)

type documentsSteps struct{}

func (s *documentsSteps) iExploreAvailableDocuments(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)

	resp, err := w.Client.Now().Documents().Explore().Get(ctx, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *documentsSteps) iCreateADocument(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)

	resp, err := w.Client.Now().Documents().Create().Post(ctx, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *documentsSteps) iDeleteADocument(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)

	err := w.Client.Now().Documents().Delete().Delete(ctx, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	return support.WithWorld(ctx, w), nil
}

func (s *documentsSteps) iGetVersionsForDocument(ctx context.Context, docID string) (context.Context, error) {
	w := support.WorldFrom(ctx)

	resp, err := w.Client.Now().Documents().Versions(docID).Get(ctx, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *documentsSteps) iGetContentForDocument(ctx context.Context, docID string) (context.Context, error) {
	w := support.WorldFrom(ctx)

	resp, err := w.Client.Now().Documents().Content(docID).Get(ctx, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

// InitializeDocumentsScenario registers all documents step definitions.
func InitializeDocumentsScenario(sc *godog.ScenarioContext) {
	s := &documentsSteps{}

	RegisterSharedSteps(sc)

	sc.Step(`^I explore available documents$`, s.iExploreAvailableDocuments)
	sc.Step(`^I create a document$`, s.iCreateADocument)
	sc.Step(`^I delete a document$`, s.iDeleteADocument)
	sc.Step(`^I get versions for document "([^"]*)"$`, s.iGetVersionsForDocument)
	sc.Step(`^I get content for document "([^"]*)"$`, s.iGetContentForDocument)

	sc.Before(BeforeScenario)
	sc.After(AfterScenario)
}
